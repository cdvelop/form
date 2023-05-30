subscribe_main_target_change.push(check_FORM_UPDATE_MODE);

function check_FORM_UPDATE_MODE(form_target) {
    console.log("NOTIFICACION check_FORM_UPDATE_MODE: ", form_target);
    if (MAIN_SAVE_FUNCTION === "UPDATE" && main_target_selected.TAG != undefined) {

        if (!OBSERVED_FORM.UPDATE_MODE_ON) {

            // console.log("MODO UPDATE:ACTIVADO ? ", MAIN_SAVE_FUNCTION);
            BTN_SAVE.disabled = false;
            BTN_DELETE.disabled = true;
            BTN_ADD.disabled = true;
            MSJ.innerHTML = '<H4 class="ok">Modo Edicion Activado</H4>';
            OBSERVED_FORM.UPDATE_MODE_ON = true
        }
    };
};

function Update() {
    let updateData = new Object();
    let formData = DataForm(form);
    let originalData = main_target_selected.TAG.firstElementChild.querySelectorAll("data");

    for (let i = 0; i < originalData.length; i++) {
        let input_name = originalData[i].dataset.name;
        let input_value = formData[input_name];
        if (input_value == undefined) {
            input_value = "";
        }
        // console.log("VALUE: ", originalData[i].dataset.value, "Orig NAME:", input_name, "Form value :",input_value);
        if (input_value != originalData[i].dataset.value) {
            updateData[input_name] = input_value;
        }
    };

    if (Object.keys(updateData).length != 0) {
        updateData["id_" + module.id] = main_target_selected.TAG.dataset.id
        let object = new Object();
        object["Object"] = module.id;
        object["Data"] = updateData
        // console.log("UPDATE DATA RECOLECTADA:", object);
        let dataObject = JSON.stringify(object);
        SOCKSEND("update", dataObject, module.id, module.id);
    };
};