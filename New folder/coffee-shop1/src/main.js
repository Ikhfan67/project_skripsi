// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import 'bootstrap-css-only/css/bootstrap.min.css'; 
import 'mdbvue/build/css/mdb.css';
require('vue2-animate/dist/vue2-animate.min.css')


Vue.config.productionTip = false

import '@/assets/wow.min.js';

import '@/assets/style.css';
import '@/assets/style2.css';
// import '@/assets/style3.css';
import '@/assets/book.css';

import '@/assets/modaleffects.css';
import '@/assets/simple-slider.css';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faUserSecret } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faUserSecret)
Vue.component('font-awesome-icon', FontAwesomeIcon)


import VueWow from 'vue-wow'
Vue.use(VueWow)

import * as VueGoogleMaps from 'vue2-google-maps'
Vue.use(VueGoogleMaps, {
  load: {
    key: 'AIzaSyBlYeVH9Bhj9YgXqs57kKXa-DhVU0_E9W4',
    libraries: 'places',
  },
})
import moment from 'moment'
Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('MMMM Do YYYY, h:mm:ss a')
  }
})
Vue.filter('formatDate1', function(value) {
  if (value) {
    return moment(String(value)).format('MMMM Do YYYY')
  }
})
import 'popper.js';
import 'bootstrap';
import '@/assets/app.scss';
import jQuery from 'jquery';
import {store} from './store/Store'


window.$ = window.jQuery = jQuery;


Vue.component('Navbar', require('./components/navbar.vue').default);
Vue.component('add-to-cart', require('./components/addToCart.vue').default);
Vue.component('mini-cart', require('./components/minicart.vue').default);
Vue.component('products-list', require('./sections/products-list.vue').default);
Vue.component('todo-list', require('./views/todoList.vue').default);

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)

import Vue2Filters from 'vue2-filters'
Vue.use(Vue2Filters)

import  VueEditor  from 'vue2-editor'
Vue.use(VueEditor)
import VueCarousel from 'vue-carousel';
Vue.use(VueCarousel);
import { BTable } from 'bootstrap-vue'
Vue.component('b-table', BTable)

import Swal from 'sweetalert2';

window.Swal = Swal;

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000
});

window.Toast = Toast;

import JsonExcel from 'vue-json-excel'
 
Vue.component('downloadExcel', JsonExcel)

new Vue({
  el: '#app',
  store: store,
  router,
  components: { App },
  template: '<App/>'
})