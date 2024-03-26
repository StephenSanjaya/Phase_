package helper

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

func ValidateUser(s interface{}) error {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(s).Field(i).Interface()
		if field.Tag.Get("required") == "true" {
			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
		if field.Tag.Get("minLen") == "8" {
			minLen, _ := strconv.Atoi(field.Tag.Get("minLen"))
			if len(value.(string)) < minLen {
				return fmt.Errorf("%s min len is 8", field.Name)
			}
		}
		if field.Tag.Get("minLen") == "6" {
			minLen, _ := strconv.Atoi(field.Tag.Get("minLen"))
			if len(value.(string)) < minLen {
				return fmt.Errorf("%s min len is 6", field.Name)
			}
		}
		if field.Tag.Get("maxLen") == "15" {
			maxLen, _ := strconv.Atoi(field.Tag.Get("maxLen"))
			if len(value.(string)) > maxLen {
				return fmt.Errorf("%s max len is 15", field.Name)
			}
		}
		if field.Tag.Get("min") == "17" {
			min, _ := strconv.Atoi(field.Tag.Get("min"))
			if value.(int) < min {
				return fmt.Errorf("%s min num is 17", field.Name)
			}
		}
		if field.Tag.Get("regex") == "true" {
			regex, _ := regexp.Compile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
			isMatch := regex.MatchString(value.(string))
			if !isMatch {
				return fmt.Errorf("email wrong format")
			}
		}
	}
	return nil
}
