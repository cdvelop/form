function checkFullObject() {
    console.log("chequeando y validando objeto");
    // UPDATE MODE
    if (main_target_selected.TAG != undefined) {
        MAIN_SAVE_FUNCTION = "UPDATE";
    }

    if (form.checkValidity()) {
        main_target_selected.VALID = true;

        console.log("OBJETO VALIDO: ", DataForm(form))

    } else {
        main_target_selected.VALID = false;
        console.log("FALTAN CAMPOS AL OBJETO: ", DataForm(form))
    }
};







function InputRightInForm(input) {
    checkFullObject();
    InputRight(input)
};