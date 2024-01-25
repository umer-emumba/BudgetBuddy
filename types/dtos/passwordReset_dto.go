package dtos

type PasswordResetDto struct {
	Token    string `json:"token" form:"token" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}
