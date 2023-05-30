package form_handler

import "fmt"

func (f form) buildID(field_name string, index int) string {
	return fmt.Sprintf("form.%v.%v.%v", f.Name, field_name, index)
}
