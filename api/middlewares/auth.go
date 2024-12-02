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

func (m *AuthMiddlewares) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header is required")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Unauthorized(c, "Invalid token format")
			return
		}

		tokenString := parts[1]
		jwtValidator, err := m.newJWTValidator()
		if err != nil {
			response.InternalServerError(c, err, "Failed to create JWT validator")
			return
		}

		claims, err := jwtValidator.ValidateToken(context.Background(), tokenString)
		if err != nil {
			response.Unauthorized(c, "Invalid token")
			return
		}

		customClaims := claims.(*validator.ValidatedClaims).CustomClaims.(*CustomClaims)

		c.Set("user", customClaims.Email)
		c.Set("roles", customClaims.Roles)

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
