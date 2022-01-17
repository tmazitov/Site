import { createLocalstorageItem, setNone } from './jwt';
import {auth} from '../client/client';
import router from '../router.js'

export function signInAction(username, password, errField) {
    auth
        .post('/user/entry', {
            'username': username,
            'password': password,
        })
        .then((response) => {
            let token = response.data
            console.log(token)
            createLocalstorageItem(token)
            window.location.href = "/profile"
        })
        .catch((error) => {
            if (error.response.status === 400) {
                errField.innerHTML = "Invalid username or password"
            } else if (error.response.status === 401) {
                errField.innerHTML = "Invalid username or password"
            } else if (error.response.status === 500) {
                errField.innerHTML = "Server error"
            }
        });
}

export function signUpAction(username, password, email, errField) {
    auth
        .post('/user/new', {
            'username': username,
            'password': password,
            'email': email
        })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "/profile"
        })
        .catch((error) => {
            if (error.response.status === 400){
                errField.innerHTML = "Invalid credentials"
            }
            else if (error.response.status === 500) {
                errField.innerHTML = "Server error"
            }
        });
}

export function signOutAction() {
    setNone()
    auth
        .get('/user/exit')
        .then(() => {
            window.location.href = "/"
        })
}

export async function refreshTokens() {
    const response = await auth
        .put('/user/refresh')
        .catch(() => {
            router.push('auth')
        })
    let token = response.data
    createLocalstorageItem(token)
}