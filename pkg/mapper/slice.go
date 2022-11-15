package mapper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Slice(from, to interface{}) {
	fromValue := reflect.ValueOf(from)
	if fromValue.Kind() != reflect.Slice &&
		(fromValue.Kind() != reflect.Ptr && fromValue.Elem().Kind() != reflect.Slice) {
		panic(fmt.Errorf("from is not slice or pointer of slice"))
	}
	fromType := fromValue.Type()
	if fromType.Kind() == reflect.Ptr {
		fromType = fromType.Elem()
		fromValue = fromValue.Elem()
	}

	toValue := reflect.ValueOf(to)
	if toValue.Kind() != reflect.Ptr || toValue.Elem().Kind() != reflect.Slice {
		panic(fmt.Errorf("to is not pointer of slice"))
	}
	toType := toValue.Elem().Type()
	toValue = toValue.Elem()

	fromTypeElem := fromType.Elem()
	if fromTypeElem.Kind() == reflect.Ptr {
		fromTypeElem = fromTypeElem.Elem()
	}
	toTypeElem := toType.Elem()
	if toTypeElem.Kind() == reflect.Ptr {
		toTypeElem = toTypeElem.Elem()
	}
	toValue.Set(reflect.MakeSlice(toType, 0, toValue.Cap()))

	if fromTypeElem.Kind() == toTypeElem.Kind() && toTypeElem.Kind() == reflect.Struct {
		for i := 0; i < fromValue.Len(); i++ {
			fromItem := fromValue.Index(i)
			isExists := false
			if fromID, ok := getIDField(fromValue.Index(i)); ok {
				for j := 0; j < toValue.Len(); j++ {
					toItem := toValue.Index(j)
					if toID, ok := getIDField(toItem); ok {
						if fmt.Sprint(fromID.Interface()) == fmt.Sprint(toID.Interface()) {
							Struct(fromItem.Interface(), toItem.Interface())
							isExists = true
							break
						}
					}
				}
			}
			if !isExists {
				toItem := reflect.New(toTypeElem)
				Struct(fromItem.Interface(), toItem.Interface())
				if toValue.Type().Elem().Kind() == reflect.Ptr {
					toValue.Set(reflect.Append(toValue, toItem))
				} else {
					toValue.Set(reflect.Append(toValue, toItem.Elem()))
				}
			}
		}
		return
	}

	if fromTypeElem.Kind() == toTypeElem.Kind() && toTypeElem.Kind() == reflect.Slice {
		for i := 0; i < fromValue.Len(); i++ {
			item := reflect.New(toTypeElem)
			Slice(fromValue.Index(i).Interface(), item.Interface())
			toValue.Set(reflect.Append(toValue, item.Elem()))
		}
		return
	}

	bin, _ := json.Marshal(fromValue.Interface())
	_ = json.Unmarshal(bin, toValue.Addr().Interface())
}

func getIDField(value reflect.Value) (reflect.Value, bool) {
	if value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	t := value.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			if v, ok := getIDField(value.FieldByName(field.Name)); ok {
				return v, true
			}
		}
		if strings.ToLower(field.Name) == "id" {
			return value.FieldByName(field.Name), true
		}
	}
	return reflect.ValueOf(nil), false
}
