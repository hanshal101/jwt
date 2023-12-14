package token

import (
	"fmt"
	"hanshal101/jwt/models"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	jWTPrivateToken = "SecretTokenService"
	ip              = "192.168.0.107"
)

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, origin string) (bool, *models.JwtClaims) {
	claims := &models.JwtClaims{}
	token, _ := getTokenfromString(tokenString, claims)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func getTokenfromString(tokenString string, claims *models.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : access denied")
		}
		return []byte(jWTPrivateToken), nil
	})
}

func GetClaims(tokenString string) models.JwtClaims {
	claims := &models.JwtClaims{}
	_, err := getTokenfromString(tokenString, claims)
	if err != nil {
		return *claims
	}
	return *claims
}
