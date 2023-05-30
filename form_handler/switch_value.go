package form_handler

import (
	"bytes"
	_ "embed"
	"log"
	"strings"
	"text/template"
	// "github.com/cdvelop/yoyi/pkg/ui/js_handler"
)

//go:embed switch_value.js
var switch_value string

func init() {
	// js_handler.RemoveJSCommentFromTemplateVariable(&switch_value, "SwitchOptions")
}

func (f *form) buildSwitchWhenInputValueChange(sb *strings.Builder) {

	f.buildSwitchOptionValueChange()

	sb.WriteString(f.buildSwitchChangeValue())

}

func (f *form) buildSwitchOptionValueChange() {
	var out string

	var resgistered = map[string]struct{}{}

	for _, field := range f.Fields {
		if js, ok := form_store.js_value_changes[field.Input.Name()]; ok {

			if _, repeated := resgistered[field.Input.Name()]; !repeated {

				out += `case "` + field.Input.Name() + `":` + "\n"

				out += js.InputValueChanges() + "\n"

				out += "break;\n"

				resgistered[field.Input.Name()] = struct{}{}
			}
		}
	}

	f.config.SwitchOptions = out
}

func (f form) buildSwitchChangeValue() string {

	t, err := template.New("").Parse(switch_value)
	if err != nil {
		log.Println(err)
		return ""
	}

	var html bytes.Buffer
	err = t.Execute(&html, f.config)
	if err != nil {
		log.Println(err)
		return ""
	}

	return html.String()

}
