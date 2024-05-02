package utils

import (
	"reflect"

	"github.com/go-playground/validator"
)

func ValidateReflector(test interface{}, err error) map[string]string {
	fieldErrors := make(map[string]string)
	for _, newErr := range err.(validator.ValidationErrors) {
		t := reflect.TypeOf(test)
		getTags, _ := t.FieldByName(newErr.Field())
		jsonTag := getTags.Tag.Get("json")
		fieldErrors[jsonTag] = ValidationMessage(getTags, newErr)
	}
	return fieldErrors
}
