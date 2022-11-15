package mapper

import (
	"reflect"
	"strings"
)

func compareTags(from, to []string) bool {
	for _, i := range from {
		for _, j := range to {
			if flatTag(i) == flatTag(j) {
				return true
			}
		}
	}
	return false
}
func flatTag(tag string) string {
	return strings.ToLower(strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(tag, " ", ""),
			"_", ""),
		"-", ""))
}
func getTagNames(field reflect.StructField) []string {
	tags := []string{
		field.Name,
	}
	if tag := getJsonTagName(field); tag != "" {
		tags = append(tags, tag)
	}
	if tag := getBsonTagName(field); tag != "" {
		tags = append(tags, tag)
	}
	if tag := getCqlTagName(field); tag != "" {
		tags = append(tags, tag)
	}
	if tag := getProtobufTagName(field); tag != "" {
		tags = append(tags, tag)
	}
	if tag := getGormTagName(field); tag != "" {
		tags = append(tags, tag)
	}
	if tag := getMapTagName(field); tag != "" {
		tags = append(tags, tag)
	}
	return tags
}
func getJsonTagName(field reflect.StructField) string {
	return strings.Split(field.Tag.Get("json"), ",")[0]
}
func getBsonTagName(field reflect.StructField) string {
	return strings.Split(field.Tag.Get("bson"), ",")[0]
}
func getCqlTagName(field reflect.StructField) string {
	return field.Tag.Get("cql")
}
func getProtobufTagName(field reflect.StructField) string {
	for _, item := range strings.Split(field.Tag.Get("protobuf"), ",") {
		if strings.HasPrefix(item, "name=") {
			return strings.ReplaceAll(item, "name=", "")
		}
	}
	return ""
}
func getGormTagName(field reflect.StructField) string {
	for _, item := range strings.Split(field.Tag.Get("gorm"), ";") {
		if strings.HasPrefix(item, "column:") {
			return strings.ReplaceAll(item, "column:", "")
		}
	}
	return ""
}
func getMapTagName(field reflect.StructField) string {
	return field.Tag.Get("map")
}
