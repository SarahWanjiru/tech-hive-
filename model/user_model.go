package model

type UserModel struct {
 	Username string `json:"username" validate:"required,email" example:"user@example.com"`
 	Password string `json:"password" validate:"required" example:"password123"`
 }

type UserRegistrationModel struct {
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,max=150"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,oneof=customer admin"`
}
