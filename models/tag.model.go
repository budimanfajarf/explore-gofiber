package models

// Tag model
type Tag struct {
	BaseModel

	Name string `gorm:"column:name" json:"name"`

	// Tags []Tag `gorm:"many2many:ArticleTag;ForeignKey:id;References:id;JoinForeignKey:articleId;JoinReferences:tagId"`

	// Articles []Article `gorm:"many2many:ArticleTag;ForeignKey:id;References:id;JoinForeignKey:tagId;JoinReferences:articleId"`

}

// func (Tag) TableName() string {
// 	return "Tag"
// }
