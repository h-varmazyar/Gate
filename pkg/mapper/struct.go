package mapper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func Struct(from, to interface{}) {
	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Struct &&
		(fromValue.Kind() != reflect.Ptr && fromValue.Elem().Kind() != reflect.Struct) {
		panic(fmt.Errorf("from is not struct or pointer of struct"))
	}
	fromType := fromValue.Type()
	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
		fromValue = fromValue.Elem()
	}

	toValue := reflect.ValueOf(to)
	if toValue.Kind() != reflect.Ptr || toValue.Type().Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("to is not pointer of struct"))
	}
	if toValue.IsZero() {
		panic(fmt.Errorf("to is nil"))
	}
	toType := toValue.Type().Elem()
	toValue = toValue.Elem()

	for i := 0; i < fromType.NumField(); i++ {
		fromTypeField := fromType.Field(i)
		fromFieldTags := getTagNames(fromTypeField)
		fromField := fromValue.FieldByName(fromTypeField.Name)
		fromFieldType := fromTypeField.Type
		if fromFieldType.Kind() == reflect.Ptr {
			fromFieldType = fromFieldType.Elem()
		}
		if fromTypeField.Anonymous {
			Struct(fromField.Interface(), to)
		}

		for i := 0; i < toType.NumField(); i++ {
			toTypeField := toType.Field(i)
			toFieldTags := getTagNames(toTypeField)
			toField := toValue.FieldByName(toTypeField.Name)
			toFieldType := toTypeField.Type
			if toFieldType.Kind() == reflect.Ptr {
				toFieldType = toFieldType.Elem()
			}
			if !toField.CanSet() {
				continue
			}

			if compareTags(fromFieldTags, toFieldTags) {
				if fromField.IsZero() {
					toField.Set(reflect.Zero(toField.Type()))
					continue
				}

				if fromFieldType.Kind() == toFieldType.Kind() && toFieldType.Kind() == reflect.Struct {
					val := reflect.New(toFieldType)
					if val.Kind() != reflect.Ptr {
						val = val.Addr()
					}
					Struct(fromField.Interface(), val.Interface())
					if toField.Kind() == reflect.Ptr {
						toField.Set(val)
					} else {
						toField.Set(val.Elem())
					}
					continue
				}

				if fromFieldType.Kind() == toFieldType.Kind() && toFieldType.Kind() == reflect.Slice {
					Slice(fromField.Interface(), toField.Addr().Interface())
					continue
				}

				if fromFieldType.Kind() == toFieldType.Kind() && toFieldType.Kind() == reflect.Map {
					fromFieldTypeElem := fromFieldType.Elem()
					if fromFieldTypeElem.Kind() == reflect.Ptr {
						fromFieldTypeElem = fromFieldTypeElem.Elem()
					}
					toFieldTypeElem := toFieldType.Elem()
					if toFieldTypeElem.Kind() == reflect.Ptr {
						toFieldTypeElem = toFieldTypeElem.Elem()
					}
					if fromFieldTypeElem.Kind() == toFieldTypeElem.Kind() && toFieldTypeElem.Kind() == reflect.Struct {
						dic := reflect.MakeMap(toFieldType)
						for _, key := range fromField.MapKeys() {
							item := reflect.New(toFieldTypeElem)
							Struct(fromField.MapIndex(key).Interface(), item.Interface())
							dic.SetMapIndex(key, item)
						}
						toField.Set(dic)
						continue
					}
				}

				if fromField.Type() == toField.Type() {
					toField.Set(fromField)
				} else {
					switch val := fromField.Interface().(type) {
					case time.Time:
						if toField.Kind() == reflect.Int64 {
							toField.Set(reflect.ValueOf(val.Unix()))
							continue
						}
					case *time.Time:
						if toField.Kind() == reflect.Int64 {
							toField.Set(reflect.ValueOf(val.Unix()))
							continue
						}
					}
					if fromField.Kind() == reflect.Int64 {
						switch toField.Interface().(type) {
						case time.Time:
							toField.Set(reflect.ValueOf(time.Unix(fromField.Int(), 0)))
							continue
						case *time.Time:
							toField.Set(reflect.ValueOf(time.Unix(fromField.Int(), 0)))
							continue
						}
					}
					bin, _ := json.Marshal(fromField.Interface())
					_ = json.Unmarshal(bin, toField.Addr().Interface())
				}
			}
		}
	}
}
