


function actionAfterModifyDom() {
    if (!MULTI_CREATE_MODE_ACTIVATED) {
        // console.log("MODO CREACION NORMAL");
        // fin loop reiniciamos
        ActionButtonCANCEL();
        // foco en elemento agregado
        list.querySelectorAll("li")[0].click();
    } else {
        // console.log("MODO CREACION MULTIPLE ACTIVADO LOOP");
        resetform();
        ActionButtonADD(true);

        DEFAULT_INPUT_FOCUS_FUNCTION();
    };
};