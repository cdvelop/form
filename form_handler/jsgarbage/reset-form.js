SUBSCRIBE_ACTION_BUTTON_CANCEL.push(EnableForm);

function EnableForm() {
    DisableFormModule(false, '` + css_handler.CssColorSettings()["--col-sec"] + `');
};

function resetform() {
    console.log("RESETEANDO FORMULARIO");
    form.reset();
    resetColorFieldset();
    MAIN_SAVE_FUNCTION = undefined;
    main_target_selected.VALID = false;
};

function FormModuleDisable() {
    DisableFormModule(true, '` + css_handler.CssColorSettings()["--col-ter"] + `');
};

SUBSCRIBE_ACTION_BUTTON_DELETE.push(FormModuleDisable);

function DisableFormModule(t, c) {
    for (let i = 0; i < fieldset.length; i++) {
        fieldset[i].disabled = t;
        fieldset[i].querySelector('legend').style.backgroundColor = c;
    }
};

function resetColorFieldset() {
    for (let i = 0; i < fieldset.length; i++) {
        // quitar clase foka y ferr
        fieldset[i].classList.remove("foka");
        fieldset[i].classList.remove("ferr");
    }
};