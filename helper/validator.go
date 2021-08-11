package helper

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type (
	// ValidationErrors define validation error
	ValidationErrors struct {
		Validation []ErrorsField `json:"validation"`
	}
	// ErrorsField define field and reason error
	ErrorsField struct {
		Field  string `json:"field"`
		Reason string `json:"reason"`
	}
	// fieldError struct {
	// 	err validator.FieldError
	// }
)

// BindingFieldsTags to use the names which have been specified for JSON representations of structs, rather than normal Go field names:
func BindingFieldsTags() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// IsValidField check form request is valid field
func IsValidField(err error) ValidationErrors {
	var field []ErrorsField
	for _, fieldErr := range err.(validator.ValidationErrors) {
		field = append(field, ErrorsField{
			Field:  strings.ToLower(fieldErr.Field()),
			Reason: fieldErr.ActualTag(),
		})
	}
	return ValidationErrors{
		Validation: field,
	}
}
