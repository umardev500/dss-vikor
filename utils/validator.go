package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ParseData struct {
	Field   string  `json:"field"`
	Tag     string  `json:"tag"`
	Param   *string `json:"param,omitempty"`
	Message string  `json:"message"`
}

func parseTag(tag, field, param string) ParseData {
	var response ParseData
	var paramValue *string
	if param != "" {
		paramValue = &param
	}

	switch tag {
	case "required":
		msg := fmt.Sprintf("%s is required", field)
		response = ParseData{
			Field:   field,
			Tag:     tag,
			Param:   paramValue,
			Message: msg,
		}
	case "gte":
		msg := fmt.Sprintf("%s must be greater than or equal to %s", field, param)
		response = ParseData{
			Field:   field,
			Tag:     tag,
			Param:   paramValue,
			Message: msg,
		}
	}

	return response
}

func ValidateStruct(v *validator.Validate, s interface{}) ([]ParseData, error) {
	var responseData []ParseData

	if err := v.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			tag := err.Tag()
			param := err.Param()

			// Get field by name
			sf, ok := reflect.TypeOf(s).FieldByName(field)
			if !ok {
				data := parseTag(tag, field, param)
				responseData = append(responseData, data)
				continue
			}

			// Lookup json tag
			fJsonName, ok := sf.Tag.Lookup("json")
			if !ok {
				data := parseTag(tag, field, param)
				responseData = append(responseData, data)
				continue
			}

			// append json name if found
			data := parseTag(tag, fJsonName, param)
			responseData = append(responseData, data)
		}

		return responseData, err
	}

	return nil, nil
}
