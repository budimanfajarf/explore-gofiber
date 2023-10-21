package auth

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserResponse todo
type UserResponse struct {
	ID       uint   `json:"id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// AccessResponse todo
type AccessResponse struct {
	Token string `json:"token"`
}

// AuthResponse todo
type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
