package dtos

import "mime/multipart"

type UpdateProfileDTO struct {
	Name  string                `form:"name" binding:"omitempty,min=2,max=50"`
	Image *multipart.FileHeader `form:"image" binding:"omitempty"`
}
