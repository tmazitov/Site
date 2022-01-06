import MainPage from './pages/MainPage.vue'

import VueRouter from 'vue-router'

const routes = [
    {
      path: '/',
      name: 'Home',
      component: MainPage
    },
    {
      path: '/profile',
      name: 'Profile',
      component: () => import('./pages/ProfilePage.vue')
    },
    {
        path: '/about',
        name: 'About',
        component: () => import('./pages/AboutPage.vue')
    },
    {
        path: '/auth',
        name: 'Auth',
        component: () => import('./pages/AuthPage.vue')
    },
]



const router = new VueRouter({
    routes : routes,
    mode : 'history'
})

export default router