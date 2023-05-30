package form_handler

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (f form) buildVariablesJsFields() (out_vars string) {

	out_vars += fmt.Sprintf(`const form = module.querySelector('#%v-form');`, f.Name)
	out_vars += fmt.Sprintf(`const fieldset = module.querySelectorAll("#%v-form fieldset");`, f.Name)

	for _, field := range f.Fields {
		if field.NotRenderHtml {
			out_vars += buildFieldJsVar(&field)
		}
	}
	return
}

func buildFieldJsVar(field *model.Field) (out string) {
	if field != nil {
		query_type := `querySelector`
		switch field.Input.Name() {
		case "radio":
			query_type = `querySelectorAll`
		case "checkbox":
			query_type = `querySelectorAll`
		}

		out = fmt.Sprintf(`let input_%v = form.%v("[name='%v']");`, field.Name, query_type, field.Name)
	}
	return
}
