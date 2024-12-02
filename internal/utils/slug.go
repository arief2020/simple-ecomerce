package utils

import "strings"

func CreateSlug(input string) string {
	return strings.ReplaceAll(strings.ToLower(input), " ", "-")
}