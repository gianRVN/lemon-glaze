package utils

import (
	"strconv"
	"strings"
)

func GetIDFromPath(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/api/categories/")
	return strconv.Atoi(idStr)
}