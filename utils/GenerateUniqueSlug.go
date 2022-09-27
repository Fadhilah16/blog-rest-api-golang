package utils

import (
	"blog-api-golang/services"
	"strconv"

	"github.com/gosimple/slug"
)

func GenerateUniqueSlug(title string) string {
	titleSlug := slug.Make(title)
	tempSlug := titleSlug

	i := 0
	for true {
		exists := services.ExistsBlogBySlug(tempSlug)

		if exists {
			tempSlug = titleSlug + "-" + strconv.Itoa(i)
			i = i + 1

		} else {
			break
		}
	}
	return tempSlug

}
