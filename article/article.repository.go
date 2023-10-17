package article

import (
	"explore-gofiber/database"
)

func GetArticleListData() []IArticleListItem {
	// // Method 1: get all columns
	// articles := []Article{}
	// database.MySQL.Find(&articles)

	// Method 2: get specific columns
	articles := []IArticleListItem{}
	database.MySQL.Model(&Article{}).Find(&articles)

	return articles
}
