package sqlx

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func getUpdateFields(s any) (string, error) {
	fields := []string{}

	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if typ.Kind() != reflect.Pointer {
		if typ.Elem().Kind() != reflect.Struct {
			return "", errors.New("The argument that you passed is not a struct")
		}
		return "", errors.New("The argument that you passed is not a pointer")
	}

	for i := 0; i < val.Elem().NumField(); i++ {
		field := typ.Elem().Field(i)
		value := val.Elem().Field(i)

		if value.IsValid() && !value.IsZero() {
			key := field.Tag.Get("db")
			if field.Type.Kind() == reflect.Slice {
				r, err := regexp.Compile("(.*)(s|sh|ch|x|z|es|ies)$")
				if err != nil {
					return "", err
				}

				arrTyp := key
				if r.MatchString(key) {
					arrTyp = r.FindStringSubmatch(key)[1]
				}

				fields = append(fields, fmt.Sprintf("%s = string_to_array(:%s, ',')::%s[]", key, key, arrTyp))
				continue
			}
			fields = append(fields, fmt.Sprintf("%s = :%s", key, key))
		}
	}

	return strings.Join(fields, ","), nil
}
