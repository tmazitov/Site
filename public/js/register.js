function addUser() {

    // Block for error message
    let errBlock = document.getElementById("register_message")

    // Get username and check it
    let name = document.getElementById("username").value
    if (name.length < 4) {
        errBlock.innerHTML = "Username is too short"
        return
    }

    // Get username and check it
    let pass = document.getElementById("password").value
    if (pass.length < 8) {
        errBlock.innerHTML = "Password is too short"
        return
    }

    // Init request
    var xhr = new XMLHttpRequest();

    // Set params
    var params =
        'username=' + encodeURIComponent(name) +
        '&password=' + encodeURIComponent(pass);

    // Request configuration
    xhr.open("POST", '/user/new?' + params, true);

    // Error handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            console.log(xhr.status + ": " + xhr.statusText)
            if (xhr.status === 200) {
                console.log(xhr.responseText)
            } else if (xhr.status == 403) {
                errBlock.innerHTML = "User with this username alredy exists"
            }
        }
    }

    // Send request
    xhr.send();
}