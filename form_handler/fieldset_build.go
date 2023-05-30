package form_handler

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (f form) fieldsetBuild(field model.Field, index int) string {

	id := f.buildID(field.Name, index)

	tabIndex := fmt.Sprintf(` tabindex="%v"`, index)

	labelFor := fmt.Sprintf(` for="%v"`, id)

	class := ` class="` + cssClass(&field) + `"`

	t := `
	<fieldset data-name="` + field.Name + `"` + tabIndex + class + `>
	<legend><label` + labelFor + `>` + field.Legend + `</label></legend>
	` + field.Input.Build(id, field.Name, field.SkipCompletionAllowed) + `
    </fieldset>
	`

	return t
}
