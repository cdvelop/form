
	
	let DEFAULT_INPUT_FOCUS_FUNCTION = FocusOnFirstInput;

	const main_form_state = { "UPDATE_MODE_ON": false };

	let SUBSCRIBE_STATUS_FORM = [test1];

	function test1(value) {
		console.log("TEST FUNCION 1 SUSCRITA ESTADO FORMULARIO ES: ", value)
	};

	const main_form_handler = {
		set: function (target, property, value) {
			target[property] = value;

			console.log("CAMBIO ESTADO FORMULARIO: ", main_form_state)

			SUBSCRIBE_STATUS_FORM.forEach(function (func) {
				func(target);
			});

			return true;
		}
	};

	const OBSERVED_FORM = new Proxy(main_form_state, main_form_handler);

	function test2(value) {
		console.log("TEST FUNCION 2 SUSCRITA DESPUES ESTADO FORMULARIO ES: ", value)
	};