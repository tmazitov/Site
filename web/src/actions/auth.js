import { createLocalstorageItem, setNone } from './jwt';
import axios from 'axios';

export function signInAction(username, password, errField) {
    axios
        .post('/api/user/entry', {
            'username': username,
            'password': password,
        }, { withCredentials: true })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "/profile"
        })
        .catch((error) => {
            if (error.response.status === 401) {
                errField.innerHTML = "Username or password is not valid"
            }
        });
}

export function signUpAction(username, password, email, errField) {
    axios
        .post('/api/user/new', {
            'username': username,
            'password': password,
            'email': email
        }, { withCredentials: true })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "/profile"
        })
        .catch((error) => {
            if (error.response.status === 403) {
                errField.innerHTML = error.response.data
            }
        });
}

export function signOutAction() {
    setNone()
    axios
        .get('/api/user/exit', { withCredentials: true })
        .then(() => {
            window.location.href = "/"
        })
}

export function refreshTokens() {
    axios
        .put('/api/user/refresh', {}, {
            withCredentials: true,
        })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            document.location.reload();
        })
        .catch(() => {
            window.location.href = "/auth"
        });


}