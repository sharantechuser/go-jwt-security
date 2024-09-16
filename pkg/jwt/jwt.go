package jwt

import (
	"time"

	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/sharantechuser/go-jwt-security/pkg/config"
)

var jwtSecret = []byte("LRu9fN2X6gQ29hpJ1Lc+/vur2IAQiEBzML+Mbd/9bR8=") // Replace with your own secret

type Claims struct {
	Username string `json:"username"`
	jwtlib.StandardClaims
}

// GenerateToken generates a JWT token
func GenerateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwtlib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    config.AppConfig.AppName,
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates a JWT token
func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(tokenStr, &Claims{}, func(token *jwtlib.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
