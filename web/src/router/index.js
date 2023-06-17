import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import HomeView from '../views/HomeView'
import Cate from '../views/Cate'
import Cateset from '../views/Cateset'
import AboutView from '../views/AboutView'
import Login from '../views/login'
import AdminSet from '../views/AdminSet'
import AdminList from '../views/AdminList'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/about',
      name: 'About',
      component: AboutView
    },
    {
      path: '/cate',
      name: 'Cate',
      component: Cate,
      children: [
        {path: '/cate/home', component: HomeView},
        {path: '/cate/set', component: Cateset},
        {path: '/admins/create', component: AdminSet},
        {path: '/admins/edit/:id', component: AdminSet, props: true},
        {path: '/admins/list', component: AdminList}
      ]
    }
  ]
})
