// 引入axios
import axios from 'axios'

// 创建axios实例，添加api路径
const http = axios.create({
  baseURL: 'http://localhost:8080/mongo'
})

// 将文件变量导出，导出到main.js
export default http
