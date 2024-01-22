package types

import "github.com/golang-jwt/jwt/v5"

type TokenType string

const (
	EmailVerification TokenType = "email_verification"
	Refresh           TokenType = "refresh"
	Access            TokenType = "access"
	PasswordReset     TokenType = "password_reset"
)

type UserType string

const (
	User UserType = "user"
)

type JwtToken struct {
	Id int
	TokenType
	UserType
	jwt.RegisteredClaims
}
