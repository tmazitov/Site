import Vue from 'vue'

import VueRouter from 'vue-router'
import { vfmPlugin } from 'vue-final-modal'
import Unicon from 'vue-unicons/dist/vue-unicons-vue2.umd'
import { uniTimes } from 'vue-unicons/dist/icons'
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';

import App from './App.vue'

import router from './router'

const sweetAlert = {
    confirmButtonColor: '#5f9ea0',
  };

Vue.use(VueSweetalert2, sweetAlert);
Vue.use(VueRouter)
Vue.use(vfmPlugin)
Unicon.add([uniTimes])
Vue.use(Unicon)

new Vue({
    render : h => h(App),
    router
}).$mount('#app')

