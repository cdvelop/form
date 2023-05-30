package form_handler

import (
	"fmt"
	"strings"
)

func (f form) RegisteredHtmlForm() {
	fieldsets := f.buildHtmlFieldsetInputs()

	form_store.built_html_form[f.Name] = fmt.Sprintf(`<form class="form-distributed-fields" id="%v-form" name="%v-form" autocomplete="off">
	%v
	</form>`, f.Name, f.Name, fieldsets)
}

func (f *form) buildHtmlFieldsetInputs() string {
	var allFieldSet []string

	for index, field := range f.RenderFields() {

		if field.Input.Name() != "hidden" {

			allFieldSet = append(allFieldSet, f.fieldsetBuild(field, index))
		} else {
			id := f.buildID(field.Name, index)

			inpuTag := field.Input.Build(id, field.Name, field.SkipCompletionAllowed)
			allFieldSet = append(allFieldSet, inpuTag)
		}

	}
	return strings.Join(allFieldSet, "\n")
}
