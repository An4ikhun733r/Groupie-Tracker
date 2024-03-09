function validateInputsCreation() {
    var minValue = document.getElementById("creation-min").value;
    var maxValue = document.getElementById("creation-max").value;

    if (parseInt(maxValue) >= 1958 && parseInt(maxValue) <= 2015){
        if (parseInt(minValue) > parseInt(maxValue)) {
        document.getElementById("creation-min").value = maxValue;
        }
    }
}

function validateInputsAlbum() {
    var minValue = document.getElementById("album-min").value;
    var maxValue = document.getElementById("album-max").value;

    if (parseInt(maxValue) >= 1963 && parseInt(maxValue) <= 2018){
        if (parseInt(minValue) > parseInt(maxValue)) {
        document.getElementById("album-min").value = maxValue;
        }
    }
}

function validateInputsMembers() {
    var minValue = document.getElementById("members-min").value;
    var maxValue = document.getElementById("members-max").value;

    if (parseInt(maxValue) >= 1 && parseInt(maxValue) <= 8){
        if (parseInt(minValue) > parseInt(maxValue)) {
        document.getElementById("members-min").value = maxValue;
        }
    }
}