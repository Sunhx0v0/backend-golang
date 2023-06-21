// 引入axios
import axios from 'axios'
// 引入Vue
import Vue from 'vue'

// 创建axios实例，添加api路径
const http = axios.create({
  baseURL: 'http://localhost:8080/mongo'
})

// 全局进行响应的拦截(axios内的响应拦截方法)
// 作用：用vue弹框显示json的消息
http.interceptors.response.use(res => {
  return res
}, err => {
  // 如果拦截到错误的操作,使用VUE将错误信息进行弹出展示
  // 获取错误信息console.log(err.response.data.message)
  Vue.prototype.$message({
    type: 'error',
    message: err.response.data.message
  })
  return Promise.reject(err)
})

// 使用axios的interceptors拦截器，将http调用时拦截
http.interceptors.request.use(function (config) {
  // 将token值传入请求头，"bearer + 空格"是代码规范，看到Bearer(持票人)大家就明白是对token的验证
  config.headers.Authorization = 'bearer ' + localStorage.token
  return config
}, function (error) {
  // 错误处理
  return Promise.reject(error)
})

// 将文件变量导出，导出到main.js
export default http
