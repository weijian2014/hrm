import axios from "axios"
import { useUserStore } from "@/store/user"

const store = useUserStore()

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
      let token = store.token.token
      if (token) {
         // 如果headers不为空则等于headers, 否则创建一个headers对象
         config.headers = config.headers || {}
         config.headers.Authorization = token
      }
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
