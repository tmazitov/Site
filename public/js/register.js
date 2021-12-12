function getRndInteger(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function addUser() {

    let name = document.getElementById("username").value
    let pass = document.getElementById("password").value

    var xhr = new XMLHttpRequest();

    var params =
        'username=' + encodeURIComponent(name) +
        '&password=' + encodeURIComponent(pass);

    xhr.open("POST", '/user/new?' + params, true);

    xhr.onreadystatechange = function(oEvent) {
        if (xhr.readyState === 4) {
            console.log(xhr.status + ": " + xhr.statusText)
            if (xhr.status === 200) {
                console.log(xhr.responseText)
            } else if (xhr.status == 403) {
                let block = document.getElementById("register_message")
                block.innerHTML += "User with this username alredy exists"
            }
        }
    }

    xhr.send();



}