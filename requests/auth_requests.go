package requests

type RegisterRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	UserName   string `json:"username" validate:"required,min=4"`
	Password   string `json:"password" validate:"required,min=6"`
	Department string `json:"department" validate:"required,min=3"`
}

type LoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}