package form

func (form) JsGlobal() string {
	return `function DataForm(form) {
		let formData = new FormData(form);
		let object = new Object();
		formData.forEach(function (value, key) {
	
			if (!Reflect.has(object, key)) {
				object[key] = value;
				return;
			}
			if (!Array.isArray(object[key])) {
				object[key] = [object[key]]
			}
			object[key].push(value)
		});
	
		console.log("FORM OBJECT: ", object);
	
		return object
	};

	function InputRight(input) {
		let fd = input.closest('fieldset');
		if (fd != null) {
			fd.classList.add("foka");
			fd.classList.remove("ferr");
		}
	};
	
	function InputWrong(input) {
		let fd = input.closest('fieldset');
		if (fd != null) {
			fd.classList.remove("foka");
			fd.classList.add("ferr");
		}
	};
	
	function FocusInput(input) {
		// blur quita el foco del elemento activo 
		if (document.activeElement != document.body) document.activeElement.blur();
		setTimeout(function () {
			input.focus();
		}, 300);
	};
	`
}
