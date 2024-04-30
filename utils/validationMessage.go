package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

func ValidationMessage(reflectTags reflect.StructField, currentError validator.FieldError) string {
	validateValue := strings.Split(reflectTags.Tag.Get("validate"), ",")
	var fieldErrors string
	validation := make(map[string]string)
	validation["required"] = "This field cannot be empty"
	validation["email"] = "This field must be in email format"
	validation["email"] = "This field must be in email format"
	validation["min"] = "This field must be at least __ characters long."
	validation["max"] = "This field must be at max __ characters long."
	validation["noWhiteSpace"] = "This field must not containing any whitespace."
	validation["noDot"] = "This field must not containing any Dot."

	for _, validateType := range validateValue {
		fmt.Println(currentError.Tag())
		if currentError.ActualTag() == strings.Split(validateType, "=")[0] {
			value := strings.Split(validateType, "=")
			if len(value) > 1 {
				// fieldErrors = append(fieldErrors, strings.ReplaceAll(validation[value[0]], "__", value[1]))
				fieldErrors = strings.ReplaceAll(validation[value[0]], "__", value[1])
			} else {
				fieldErrors = validation[validateType]
			}
		}
	}
	return fieldErrors
}
