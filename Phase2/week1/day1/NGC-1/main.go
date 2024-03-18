package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type Avengers struct {
	Name  string `required:"true" minLen:"5" maxLen:"15"`
	Age   int    `required:"true" min:"5" max:"10"`
	Email string `required:"true" minLen:"5" maxLen:"15" regex:"true"`
}

func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(s).Field(i).Interface()
		if field.Tag.Get("required") == "true" {
			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
		if field.Tag.Get("minLen") == "5" {
			minLen, _ := strconv.Atoi(field.Tag.Get("minLen"))
			if len(value.(string)) < minLen {
				return fmt.Errorf("%s min len is 5", field.Name)
			}
		}
		if field.Tag.Get("maxLen") == "15" {
			maxLen, _ := strconv.Atoi(field.Tag.Get("maxLen"))
			if len(value.(string)) > maxLen {
				return fmt.Errorf("%s max len is 10", field.Name)
			}
		}
		if field.Tag.Get("min") == "5" {
			min, _ := strconv.Atoi(field.Tag.Get("min"))
			if value.(int) < min {
				return fmt.Errorf("%s min num is 5", field.Name)
			}
		}
		if field.Tag.Get("max") == "10" {
			max, _ := strconv.Atoi(field.Tag.Get("max"))
			if value.(int) > max {
				return fmt.Errorf("%s max num is 10", field.Name)
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

func main() {
	avengers := Avengers{
		Name:  "budiw",
		Age:   5,
		Email: "budi@email.com",
	}

	err := ValidateStruct(avengers)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("struct is valid")
	}
}
