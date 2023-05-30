package form_handler

import (
	_ "embed"
	"strings"
)

//go:embed jsmodule/basic_form.js
var basic_form string

func (f form) AttachFuntion() string {
	sb := strings.Builder{}

	// 1 - cargar js basico formulario
	sb.WriteString(basic_form)

	// 2 - cargar variables campos
	sb.WriteString(f.buildVariablesJsFields())

	// 3 - cargar funciones input registrados
	f.loadJsModuleFunctions(&sb)

	// 4 - cargar switch cuando se selecciona otro objetivo
	f.buildSwitchWhenAnotherTargetSelected(&sb)

	// 5 - cargar switch cuando valor de input cambia
	f.buildSwitchWhenInputValueChange(&sb)

	return sb.String()
}

func (f form) loadJsModuleFunctions(sb *strings.Builder) {
	var resgistered = map[string]struct{}{}

	for _, field := range f.Fields {
		if js, ok := form_store.js_functions[field.Input.Name()]; ok {
			if _, repeated := resgistered[field.Input.Name()]; !repeated {

				sb.WriteString(js.AttachJsFunctions())

				resgistered[field.Input.Name()] = struct{}{}

			}
		}
	}

}
