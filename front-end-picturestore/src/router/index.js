import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Signup from '@/components/Signup'
import Login from '@/components/Login'
import VueResource from 'vue-resource'
import VueFlashMessage from 'vue-flash-message'

Vue.use(Router)
Vue.use(VueResource)
// Vue.use(VueFlashMessage)

Vue.use(VueFlashMessage, {
  messageOptions: {
    timeout: 1000
  }
})

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    }
  ]
})
