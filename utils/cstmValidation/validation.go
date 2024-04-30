package cstmvalidation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

func NoWhiteSpace(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if strings.Contains(value, " ") {
		fmt.Println(value)
		return false
	}
	return true
}

func NoDot(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	pattern := `\.`
	re := regexp.MustCompile(pattern)
	if re.MatchString(value) {
		return false
	} else {
		return true
	}
}
