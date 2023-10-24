package tag

type FindAllArgs struct {
	OrderBy string `validate:"omitempty,oneof=id title createdAt updatedAt image status" json:"orderBy"`
	Order   string `validate:"omitempty,oneof=asc desc" json:"order"`
}
