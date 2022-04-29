package models

type UserDTO struct {
	Name *string `json:"name" validate:"required"`
	CNP  *string `json:"cnp" validate:"required,min=13,max=13"`
}

type UserDTOList struct {
	UserDTOs []UserDTO `json:"userdtos" validate:"required"`
}
