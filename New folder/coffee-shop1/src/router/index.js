import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/home.vue'
import Admin from '@/views/admin.vue'
import Overview from '@/views/overview.vue'
import Products from '@/views/products.vue'
import Productssell from '@/views/productssell.vue'
import User from '@/views/user.vue'
import Orders from '@/views/orders.vue'
import Comment from '@/views/comment.vue'
import {store} from '@/store/Store'


Vue.use(Router)

const router = new Router({
  mode:"history",
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/?error=true',
      name: 'home-error',
      component: Home,
    },
    {
      path: "/admin",
      name: "admin",
      component: Admin,
      meta: { requiresAuth: true },
      children:[
        {
          path: "overview",
          name: "overview",
          component: Overview
        },
        {
          path: "products",
          name: "products",
          component: Products
        },
        {
          path: "productssell",
          name: "productssell",
          component: Productssell
        },
        {
          path: "user",
          name: "user",
          component: User
        },
        {
          path: "orders",
          name: "orders",
          component: Orders
        },
        {
          path: "comment",
          name: "comment",
          component: Comment
        }
      ]
    },
    {
      path:"/checkout",
      name:"checkout",
      component:()=>
      import("@/views/checkout.vue")
    },
    {
      path: '/carabeli',
      name: 'carabeli',
      component:()=>
      import("@/views/carabeli.vue")
    },
    {
      path: "/profile",
      name: "profile",
      component:()=>
      import("@/views/profil.vue")
    },
  ]
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.getters.loggedIn) {
      next({
        name: 'home',
      })
    }else if(store.getters.loggedAdmin){
      next()
    }
    else {
      store.dispatch("CekAdmin")
    }
  }else{
    next()
  }
})
export default router