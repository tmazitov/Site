import Vue from 'vue'
import VueCookies from 'vue-cookies'

import MainPage from './components/MainPage.vue'
import AuthPage from './components/AuthPage.vue'
import AboutPage from './components/AboutPage.vue'
import ProfilePage from './components/ProfilePage.vue'

const routes = {
    "/": MainPage,
    "/auth": AuthPage,
    "/profile": ProfilePage,
    "/about": AboutPage,
}

new Vue({
    el: '#app',
    data: {
        currentRoute: window.location.pathname
    },
    computed: {
        ViewComponent() {
            return routes[this.currentRoute] //|| NotFound
        }
    },
    render(h) { return h(this.ViewComponent) }
})

Vue.use(VueCookies)