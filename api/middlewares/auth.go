package middlewares

import (
	"context"
	"fmt"

	"log"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/config"
	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

type (
	AuthMiddlewares struct {
		auth0Conf config.Auth0Config
	}

	CustomClaims struct {
		Email string   `json:"email"`
		Roles []string `json:"user_roles"`
	}
)

func NewAuthMiddlewares(auth0Conf config.Auth0Config) *AuthMiddlewares {
	return &AuthMiddlewares{
		auth0Conf: auth0Conf,
	}
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	if c.Email == "" {
		return fmt.Errorf("email is required")
	}

	return nil
}

func (m *AuthMiddlewares) AuthMiddleware(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, roles, err := m.getUserInfo(c)
		if err != nil {
			if !required {
				c.Next()
				return
			}

			switch true {
			case errors.Is(err, errors.ErrUnauthorized):
				response.Unauthorized(c, "Unauthorized")
			default:
				response.InternalServerError(c, err, "Failed to get user info")
			}
			return
		}

		c.Set("user", email)
		c.Set("roles", roles)

		c.Next()
	}
}

func (m *AuthMiddlewares) AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")
		if !exists || !slices.Contains(roles.([]string), "admin") {
			response.Forbidden(c, "Forbidden")
			return
		}
		c.Next()
	}
}

func (m *AuthMiddlewares) getUserInfo(c *gin.Context) (string, []string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil, errors.NewUnauthorizedError("Authorization header is required")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", nil, errors.NewUnauthorizedError("Invalid token format")
	}

	tokenString := parts[1]
	jwtValidator, err := m.newJWTValidator()
	if err != nil {
		return "", nil, fmt.Errorf("failed to create JWT validator: %w", err)
	}

	claims, err := jwtValidator.ValidateToken(context.Background(), tokenString)
	if err != nil {
		return "", nil, errors.NewUnauthorizedError("Invalid token")
	}

	customClaims := claims.(*validator.ValidatedClaims).CustomClaims.(*CustomClaims)
	return customClaims.Email, customClaims.Roles, nil
}

func (m *AuthMiddlewares) newJWTValidator() (*validator.Validator, error) {
	issuerURL, err := url.Parse("https://" + m.auth0Conf.Domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	return validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{m.auth0Conf.Audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
}
