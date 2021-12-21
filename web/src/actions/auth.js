import axios from 'axios';
import { createLocalstorageItem, setNone } from './jwt';

export function signInAction(username, password, errField) {
    axios
        .post('http://localhost:8000/user/entry?' +
            'username=' + username +
            '&password=' + password, {}, { withCredentials: true })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "http://localhost:8080/profile"
        })
        .catch((error) => {
            if (error.response.status === 401) {
                errField.innerHTML = "Username or password is not valid"
            }
        });
}

export function signUpAction(username, password, email, errField) {
    axios
        .post('http://localhost:8000/user/new?' +
            'username=' + username +
            '&password=' + password +
            '&email=' + email)
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "http://localhost:8080/profile"
        })
        .catch((error) => {
            if (error.response.status === 403) {
                errField.innerHTML = error.response.data
            }
        });
}

export function signOutAction() {
    setNone()
    window.location.href = "http://localhost:8080"
}

export function refreshTokens() {
    axios
        .put('http://localhost:8000/user/entry', {}, {
            withCredentials: true,
        })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            document.location.reload();
        })
        .catch(() => {
            window.location.href = "http://localhost:8080/auth"
        });


}