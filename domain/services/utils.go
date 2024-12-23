package services

import (
	"errors"
	"reflect"
	"slices"
)

func sanitizeSortOrder(sortOrder string) string {
	if sortOrder != "asc" && sortOrder != "desc" {
		return "desc"
	}
	return sortOrder
}

func sanitizeSortBy(sortBy string, s any) (string, error) {

	fields, err := getAllFields(s)
	if err != nil {
		return "", err
	}

	if !slices.Contains(fields, sortBy) {
		return "", errors.New("The field that you want to sort by is not exist")
	}

	return sortBy, nil
}

func getAllFields(s any) ([]string, error) {

	typ := reflect.TypeOf(s)

	if typ.Kind() != reflect.Pointer {
		if typ.Elem().Kind() != reflect.Struct {
			return nil, errors.New("The argument that you passed is not a struct")
		}
		return nil, errors.New("The argument that you passed is not a pointer")
	}

	fields := getAllStructFields(s)

	return fields, nil
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

		if field.Anonymous {
			keys = append(keys, getAllStructFields(val.Field(i).Interface())...)
		}

		key := field.Tag.Get("db")
		keys = append(keys, key)
	}

	return keys
}
