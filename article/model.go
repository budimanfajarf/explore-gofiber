package article

// Article model
type Article struct {
	id      uint
	title   string
	content string
	image   string
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Article) TableName() string {
	return "Article"
}
