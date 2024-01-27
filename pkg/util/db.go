package util

import "fmt"

func CalculateOffset(page, perPage int) (result int) {
	result = (page - 1) * perPage
	return
}

func TransformSortClause(column, sort string) (result string) {
	if sort == "latest" {
		return fmt.Sprintf("%s DESC", column)
	}

	return column
}
