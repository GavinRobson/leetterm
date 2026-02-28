package log

import (
	"fmt"
	"reflect"
	"strings"
)

// Struct pretty-prints any struct pointer or struct value as:
// fieldName: value
// fieldName2: value2
func Struct(v any) string {
	if v == nil {
		return "<nil>"
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return "<nil>"
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", v)
	}

	rt := rv.Type()
	var b strings.Builder

	for i := 0; i < rv.NumField(); i++ {
		sf := rt.Field(i)

		// skip unexported fields
		if sf.PkgPath != "" {
			continue
		}

		name := sf.Name

		// prefer json tag name, if present
		if tag := sf.Tag.Get("json"); tag != "" {
			tagName := strings.Split(tag, ",")[0]
			if tagName != "" && tagName != "-" {
				name = tagName
			}
		}

		fv := rv.Field(i).Interface()
		fmt.Fprintf(&b, "%s: %v\n", name, fv)
	}

	out := b.String()
	return strings.TrimRight(out, "\n")
}
