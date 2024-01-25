package dtos

type AccountVerificationDto struct {
	Token string `json:"token" form:"token" binding:"required"`
}
