package models

// Admin model
type Admin struct {
	BaseModel

	Role     string `json:"role"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (Admin) TableName() string {
	return "Admin"
}
