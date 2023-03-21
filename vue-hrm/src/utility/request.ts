import axios from "axios"

const instance = axios.create({
   // 这里的/api是vite.config.ts中的server.proxy配置的
   baseURL: "/api",
   // baseURL: "http://0.0.0.0:8080",
   timeout: 15000,
   headers: {
      "Content-Type": "application/json;charset=utf-8",
   },
})

// 请求拦截器
instance.interceptors.request.use(
   (config) => {
      return config
   },
   (err) => {
      return Promise.reject(err)
   }
)

// 响应拦截器
instance.interceptors.response.use(
   (result) => {
      return result.data
   },
   (err) => {
      return Promise.reject(err)
   }
)

export default instance
