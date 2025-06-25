package utils

import (
	"reflect"
	"strings"
)

// ValidateStruct validates struct fields using reflection and pointers
func ValidateStruct(s interface{}) []string {
	if s == nil {
		panic("Cannot validate nil struct")
	}
	
	var errors []string
	v := reflect.ValueOf(s)
	
	// Handle pointer to struct
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			panic("Cannot validate nil pointer")
		}
		v = v.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		panic("ValidateStruct expects a struct or pointer to struct")
	}
	
	t := v.Type()
	
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		
		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}
		
		// Get JSON tag for field name
		jsonTag := fieldType.Tag.Get("json")
		fieldName := fieldType.Name
		if jsonTag != "" && jsonTag != "-" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}
		
		// Validate based on field type
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				errors = append(errors, fieldName+" is required")
			}
		case reflect.Ptr:
			if field.IsNil() {
				errors = append(errors, fieldName+" cannot be nil")
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if field.Uint() == 0 {
				errors = append(errors, fieldName+" must be greater than 0")
			}
		}
	}
	
	return errors
}

// ValidateRequired checks if required fields are present using pointers
func ValidateRequired(fields map[string]interface{}) []string {
	var errors []string
	
	for fieldName, value := range fields {
		if value == nil {
			errors = append(errors, fieldName+" is required")
			continue
		}
		
		// Use reflection to check value
		v := reflect.ValueOf(value)
		
		// Handle pointers
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				errors = append(errors, fieldName+" cannot be nil")
				continue
			}
			v = v.Elem()
		}
		
		// Check for empty values
		switch v.Kind() {
		case reflect.String:
			if strings.TrimSpace(v.String()) == "" {
				errors = append(errors, fieldName+" cannot be empty")
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v.Uint() == 0 {
				errors = append(errors, fieldName+" must be greater than 0")
			}
		}
	}
	
	return errors
}