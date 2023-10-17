package article

import (
	"explore-gofiber/database"
	"explore-gofiber/models"
)

func GetArticleListData() []IArticleListItem {
	// // Method 1: get all columns
	// articles := []models.Article{}
	// database.MySQL.Find(&articles)

	// Method 2: get specific columns
	articles := []IArticleListItem{}
	database.MySQL.Model(&models.Article{}).Find(&articles)

	return articles
}
