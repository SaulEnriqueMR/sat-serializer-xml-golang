package helpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func TrimXml(inputXml []byte) ([]byte, error) {
	re := regexp.MustCompile(`(\w+:\w+|[\w-]+)\s*=\s*"([^"]*)"`)
	trimmedXML := re.ReplaceAllStringFunc(string(inputXml), func(match string) string {
		parts := strings.SplitN(match, "=", 2)
		if len(parts) != 2 {
			return match // Return the original match if it doesn't contain "="
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(strings.Trim(parts[1], `"`))
		return fmt.Sprintf(`%s="%s"`, key, value)
	})
	trimmedXML = regexp.MustCompile(`>\s+<`).ReplaceAllString(trimmedXML, "><")
	trimmedXML = regexp.MustCompile(`\s+`).ReplaceAllString(trimmedXML, " ")
	trimmedXML = strings.TrimSpace(trimmedXML)
	return []byte(trimmedXML), nil
}

func TrimStringAttributes(s interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Type()
		switch field.Kind() {
		case reflect.String:
			field.SetString(strings.TrimSpace(field.String()))
		case reflect.Struct:
			TrimStringAttributes(field.Addr().Interface())
		case reflect.Slice:
			if fieldType.Elem().Kind() == reflect.Struct {
				for j := 0; j < field.Len(); j++ {
					TrimStringAttributes(field.Index(j).Addr().Interface())
				}
			}
		case reflect.Ptr:
			if field.Elem().Kind() == reflect.Struct {
				TrimStringAttributes(field.Interface())
			}
		default:
			// Default case: do nothing for other types
		}
	}
}
