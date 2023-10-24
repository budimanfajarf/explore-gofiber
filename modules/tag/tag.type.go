package tag

type FindAllArgs struct {
	OrderBy string `validate:"omitempty,oneof=id title createdAt updatedAt image status" json:"orderBy"`
	Order   string `validate:"omitempty,oneof=asc desc" json:"order"`
}

type ListItem struct {
	ID   uint   `gorm:"column:id;primarykey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
