package admin

import "explore-gofiber/types"

// Admin model
type Admin struct {
	types.BaseModel

	Role     string `json:"role"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

type Tabler interface {
	TableName() string
}

func (Admin) TableName() string {
	return "Admin"
}
