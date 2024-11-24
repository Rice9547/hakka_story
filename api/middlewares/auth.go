package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
		adminConf config.AdminConfig
		auth0Conf config.Auth0Config
	}

	CustomClaims struct {
		Email string `json:"email"`
	}
)

func NewAuthMiddlewares(adminConf config.AdminConfig, auth0Conf config.Auth0Config) *AuthMiddlewares {
	return &AuthMiddlewares{
		adminConf: adminConf,
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, http.StatusUnauthorized, "Invalid token format")
			return
		}

		tokenString := parts[1]
		jwtValidator, err := m.newJWTValidator()
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Failed to create JWT validator")
			return
		}

		claims, err := jwtValidator.ValidateToken(context.Background(), tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Set("user", claims.(*validator.ValidatedClaims).CustomClaims.(*CustomClaims).Email)

		c.Next()
	}
}

func (m *AuthMiddlewares) AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists || !slices.Contains(m.adminConf.Whitelist, user.(string)) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
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
