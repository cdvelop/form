subscribe_main_target_change.push(switchTargetChange);

function switchTargetChange(e) {

    console.log("switchTargetChange ",e)

    // {{.Reset}}
    
    // let selected = main_target_selected.FIELDATA.firstElementChild.querySelectorAll("data");
    let selected = main_target_selected.FIELDATA;

    for (let i = 0; i < selected.length; i++) {
        let input = form[selected[i].dataset.name];
        //  console.log("SELECCION INPUT FORMULARIO: ",input);
        //  console.log("data set tag: ", selected[i].dataset.name);
        switch (selected[i].dataset.name) {
            // {{.SwitchOptions}}

            default:
                input.value = selected[i].dataset.value;
                InputValidationWithValidity(input);
                break;
        };
    };
    // console.log("Nombre: ", input.name, " Valor: ", input.value, "  INPUT: ", input)
};
