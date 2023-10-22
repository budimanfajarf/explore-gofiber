package models

// Article model
type Article struct {
	BaseModel

	Title    string `gorm:"column:title" json:"title"`
	Content  string `gorm:"column:content" json:"content"`
	Image    string `gorm:"column:image" json:"image"`
	ImageUrl string `gorm:"-" json:"imageUrl"`
	Status   string `gorm:"column:status" json:"status"`

	// type Profile struct {
	// 	gorm.Model
	// 	Name      string
	// 	UserRefer uint
	// }

	// type User struct {
	// 	gorm.Model
	// 	Profiles  []Profile `gorm:"many2many:user_profiles;ForeignKey:Refer;JoinForeignKey:UserReferID;References:UserRefer;JoinReferences:ProfileRefer"`
	// 	Profiles2 []Profile `gorm:"many2many:user_profiles2;ForeignKey:refer;JoinForeignKey:user_refer_id;References:user_refer;JoinReferences:profile_refer"`
	// 	Refer     uint
	// }

	// Tags []Tag `gorm:"many2many:ArticleTag;ForeignKey:id;References:id;JoinForeignKey:articleId;JoinReferences:tagId"`

	Tags []Tag `gorm:"many2many:ArticleTag" json:"tags"`
}

// func (Article) TableName() string {
// 	return "Article"
// }
