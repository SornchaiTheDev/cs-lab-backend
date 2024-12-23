package sqlx

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func getUpdateFields(s any) (string, error) {
	typ := reflect.TypeOf(s)

	if typ.Kind() != reflect.Pointer {
		if typ.Elem().Kind() != reflect.Struct {
			return "", errors.New("The argument that you passed is not a struct")
		}
		return "", errors.New("The argument that you passed is not a pointer")
	}

	fields := getAllStructFields(s)

	return strings.Join(fields, ","), nil
}

func getAllStructFields(s any) []string {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	keys := make([]string, 0)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if value.IsValid() && !value.IsZero() {
			key := field.Tag.Get("db")
			if field.Type.Kind() == reflect.Slice {
				r := regexp.MustCompile("(.*)(s|sh|ch|x|z|es|ies)$")

				arrTyp := key
				if r.MatchString(key) {
					arrTyp = r.FindStringSubmatch(key)[1]
				}

				keys = append(keys, fmt.Sprintf("%s = string_to_array(:%s, ',')::%s[]", key, key, arrTyp))
				continue
			}
			keys = append(keys, fmt.Sprintf("%s = :%s", key, key))
		}

		if field.Anonymous {
			keys = append(keys, getAllStructFields(value.Field(i).Interface())...)
		}
	}

	return keys
}
