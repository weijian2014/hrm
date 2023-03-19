import axios from "axios"

const instance = axios.create({
   baseURL: "http://0.0.0.0:8080/api/v1",
   timeout: 15000,
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
