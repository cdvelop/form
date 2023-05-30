package form_handler

import (
	"bytes"
	_ "embed"
	"log"
	"strings"
	"text/template"
	// "github.com/cdvelop/yoyi/pkg/ui/js_handler"
)

//go:embed switch_selected.js
var switch_selected string

func init() {
	// js_handler.RemoveJSCommentFromTemplateVariable(&switch_selected, "Reset", "SwitchOptions")
}

func (f *form) buildSwitchWhenAnotherTargetSelected(sb *strings.Builder) {

	f.buildSwitchOptionSelected()

	sb.WriteString(f.buildSelectedInput())

}

func (f *form) buildSwitchOptionSelected() {
	var out string

	var resgistered = map[string]struct{}{}

	for _, field := range f.Fields {
		if js, ok := form_store.js_target_changes[field.Input.Name()]; ok {

			if _, repeated := resgistered[field.Input.Name()]; !repeated {

				out += `case "` + field.Input.Name() + `":` + "\n"

				out += js.SelectedTargetChanges() + "\n"

				out += "break;\n"

				resgistered[field.Input.Name()] = struct{}{}
			}
		}
	}

	f.config.SwitchOptions = out
}

func (f form) buildSelectedInput() string {

	t, err := template.New("").Parse(switch_selected)
	if err != nil {
		log.Println(err)
		return ""
	}

	if f.config.Reset == "" {
		f.config.Reset = `form.reset();`
	}

	var html bytes.Buffer
	err = t.Execute(&html, f.config)
	if err != nil {
		log.Println(err)
		return ""
	}

	return html.String()

}
