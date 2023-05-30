package form_handler

import "github.com/cdvelop/model"

func cssClass(field *model.Field) (class string) {
	if field != nil {

		switch field.Input.Name() {
		case "text":
			class = "max-height-100"

		case "checkbox":
			class = "min-width-100"

		case "textarea":
			class = "min-width-100"
		}
	}
	return
}
