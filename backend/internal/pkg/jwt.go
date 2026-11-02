package pkg

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("studysmart_jwt_secret_key_change_in_production_2026")

type Claims struct {
	UserID   uuid.UUID `json:"userId"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uuid.UUID, email, fullName, role string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Email:    email,
		FullName: fullName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// AuthMiddleware ensures valid JWT in Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			SendUnauthorized(c, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			SendUnauthorized(c, "invalid authorization format")
			c.Abort()
			return
		}

		claims, err := ParseToken(parts[1])
		if err != nil {
			SendUnauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}

		// Store user details in context
		c.Set("userId", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("fullName", claims.FullName)

		c.Next()
	}
}

// OptionalAuthMiddleware extracts user info if present, but doesn't abort if missing
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				if claims, err := ParseToken(parts[1]); err == nil {
					c.Set("userId", claims.UserID)
					c.Set("email", claims.Email)
					c.Set("role", claims.Role)
					c.Set("fullName", claims.FullName)
				}
			}
		}
		c.Next()
	}
}

