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

// 将文件变量导出，导出到main.js
export default http
