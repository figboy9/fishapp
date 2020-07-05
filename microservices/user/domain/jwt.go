package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
	jwt.StandardClaims
}

type ctxKey string

const JwtCtxKey ctxKey = "jwtClaims"

type TokenType string

const (
	IdToken      TokenType = "id_token"
	RefreshToken TokenType = "refresh_token"
)
