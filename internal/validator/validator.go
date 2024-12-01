package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type AppValidator interface {
	Validate(data interface{}) []ErrorResponse
}

type appValidator struct {
	validator *validator.Validate
}

func NewAppValidator() AppValidator {
	return &appValidator{validator: validator.New()}
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v *appValidator) Validate(data interface{}) []ErrorResponse {
	errors := []ErrorResponse{}

	if errs := v.validator.Struct(data); errs != nil {

		for _, err := range errs.(validator.ValidationErrors) {
			fieldName := getJSONFieldName(data, err.Field())
			errors = append(errors, ErrorResponse{
				Field:   fieldName,
				Message: err.Tag(),
			})
		}
	}

	return errors
}

// Thanks ChatGPT
func getJSONFieldName(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem() // Dereference if it's a pointer
	}
	field, found := t.FieldByName(fieldName)
	if !found {
		return fieldName // Fallback to struct field name
	}
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return fieldName // Fallback if no JSON tag or it's ignored
	}
	return jsonTag
}
