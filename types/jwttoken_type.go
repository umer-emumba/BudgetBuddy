package types

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
