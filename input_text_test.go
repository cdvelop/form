package form_test

import (
	"log"
	"testing"

	"github.com/cdvelop/form"
)

var (
	modelText = form.Input().Text()

	dataText = map[string]struct {
		inputData       string
		skip_validation bool
		result          bool
	}{
		"nombre correcto con punto?":     {"Dr. Maria Jose Diaz Cadiz", false, true},
		"no tilde ":                      {"peréz del rozal", false, false},
		"texto con ñ ":                   {"Ñuñez perez", false, true},
		"texto correcto + 3 caracteres ": {"hola", false, true},
		"texto correcto 3 caracteres ":   {"los", false, true},
		"oración ok ":                    {"hola que tal", false, true},
		"solo Dato numérico permitido?":  {"100", false, false},
		"con caracteres y coma ":         {"los,true, vengadores", false, true},
		"sin data ok":                    {"", false, false},
		"un carácter numérico ":          {"8", false, false},
		"palabra mas numero permitido ":  {"son 4 bidones", false, true},
		"con paréntesis y numero ":       {"son 4 (4 bidones)", false, true},
		"con solo paréntesis ":           {"son (bidones)", false, true},
		"palabras y numero":              {"apellido Actualizado 1", false, true},
		"palabra con slash?":             {" estos son \\n los podria", false, false},
	}
)

func Test_InputText(t *testing.T) {
	for prueba, data := range dataText {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelText.Validate.DataField(data.inputData, data.skip_validation); ok != data.result {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputText(t *testing.T) {
	for _, data := range modelText.TestData.Good("", "", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.Validate.DataField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextFirsNames(t *testing.T) {
	for _, data := range modelText.TestData.Good("patient", "patient_name", true) {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.Validate.DataField(data, false); !ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputText(t *testing.T) {
	for _, data := range modelText.TestData.Wrong() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.Validate.DataField(data, false); ok {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
