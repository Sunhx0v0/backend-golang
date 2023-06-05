// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
// 这里是引用element-ui的代码
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// 以下是引用axios的代码（专门存在axios.js里了）
import axios from './axios'
// 把axios.js文件中的axios实例加载到Vue的实例属性prototype中
Vue.prototype.$AXIOS = axios
// 引入element
Vue.use(Element)
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
