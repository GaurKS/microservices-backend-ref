package services

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateId(strlen int) string {
	// const chars = "abcdefghijklmnopqrstuvwxyzABC0123456789"
	chars := os.Getenv("CHARS")
	result := make([]byte, strlen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(email string, name string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("error parsing jwt-claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
