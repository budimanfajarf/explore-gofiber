package utils

import (
	"explore-gofiber/config"
	"explore-gofiber/constant"
	"regexp"
	"strings"
)

func RemoveUnnecessarySlashesFromURL(url string) string {
	re := regexp.MustCompile(`([^:]/)\/+`)
	return re.ReplaceAllString(url, "$1")
}

func GetImageURL(path, imageName string) string {
	if imageName == "" {
		return ""
	}

	// check if first word is http or not
	if strings.HasPrefix(imageName, "http") {
		return imageName
	}

	env := config.Env
	// log.Println(env.StorageUrl)

	return RemoveUnnecessarySlashesFromURL(
		strings.Join([]string{env.StorageUrl, path, imageName}, "/"),
	)
}

func GetArticleImageURL(imageName string) string {
	return GetImageURL(constant.ImagePathArticle, imageName)
}
