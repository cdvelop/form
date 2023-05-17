package form

import (
	"regexp"

	"github.com/cdvelop/model"
)

// parámetro opcional
// "hidden" si se vera oculto o no.
// "Escriba Nombre y dos apellidos" = place_holder.
func (input) Text(option ...string) model.Input {
	t := text{}
	for i, opt := range option {
		switch opt {
		case "hidden":
			t.Hidden = true
		default:
			t.placeHolder = option[i]
		}
	}

	return model.Input{
		Component: model.Component{
			Name:        t.Name(),
			CssGlobal:   nil,
			CssPrivate:  nil,
			JsGlobal:    nil,
			JsPrivate:   nil,
			JsListeners: nil,
		},
		Build:    t,
		Validate: t,
		TestData: t,
	}
}

// const patternText = `^[A-Za-zÑñ 0-9.,() ]{2,100}$`

const patternText = `^[a-zA-ZÑñ]{2,100}[a-zA-ZÑñ0-9()., ]*$`

const titleInfo = "texto, punto,coma, paréntesis o números permitidos max. 100 caracteres"

// texto,punto,coma, paréntesis o números permitidos
type text struct {
	placeHolder string
	Hidden      bool
}

func (text) Name() string {
	return "text"
}

func (t text) HtmlName() string {
	if t.Hidden {
		return "hidden"
	}
	return "text"
}

func (t text) HtmlTAG(id, field_name string, skip_completion_allowed bool) string {
	var info string
	if t.placeHolder == "" {
		info = titleInfo
	}

	a := attribute{
		InputId:               id,
		TagType:               t.HtmlName(),
		PropertieName:         field_name,
		TitleInfo:             info,
		Pattern:               patternText,
		SkipCompletionAllowed: skip_completion_allowed,
		PlaceHolder:           t.placeHolder,
		// Autocomplete:  "off",
	}

	return a.Build()
}

// validación con datos de entrada
func (text) DataField(data_in string, skip_validation bool) (ok bool) {
	if !skip_validation {

		pvalid := regexp.MustCompile(patternText)

		return pvalid.MatchString(data_in)

	} else {
		return true
	}
}

// options: first_name,last_name, phrase
func (text) Good(table_name, field_name string, random bool) (out []string) {

	first_name := []string{"Maria", "Juan", "Marcela", "Luz", "Carmen", "Jose", "Octavio"}

	last_name := []string{"Soto", "Del Rosario", "Del Carmen", "Ñuñez", "Perez", "Cadiz", "Caro"}

	phrase := []string{"Dr. Maria Jose Diaz Cadiz", "son 4 (4 bidones)", "pc dental (1)", "equipo (4)"}

	switch field_name {
	case table_name + "_name":
		return combineStringArray(true, first_name, last_name, last_name)
	case "first_name":
		return first_name

	case "last_name":
		return last_name

	default:
		out = append(out, phrase...)
		out = append(out, first_name...)
		out = append(out, last_name...)
	}

	return
}

func (text) Wrong() (out []string) {

	out = []string{
		"peréz del rozal",
		" estos son \\n los podria",
	}

	out = append(out, wrong_data...)

	return
}
