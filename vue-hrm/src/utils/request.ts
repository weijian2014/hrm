import axios from "axios"
import { useUserStore } from "@/store/user"
import { refreshToken } from "@/utils/user"
import router from "@/router/index"

const request = axios.create({
   // 这里的/api是vite.config.ts中的server.proxy配置的
   baseURL: "/api",
   // baseURL: "http://0.0.0.0:8080",
   timeout: 15000,
   headers: {
      "Content-Type": "application/json;charset=utf-8",
   },
})

// 请求拦截器
request.interceptors.request.use(
   (config) => {
      const token = useUserStore().tokenInfo.access_token
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
request.interceptors.response.use(
   (response) => {
      return response.data
   },
   async (err) => {
      if (!err.response) {
         return Promise.reject(err)
      }

      // Token过期
      if (err.response.status === 401) {
         console.log("token过期, 需要刷新")
         const res = await refreshToken({
            refresh_token: useUserStore().tokenInfo.refresh_token,
         })

         if (res.code === 200) {
            console.log("刷新token成功")
            // 保存新生成的Token
            useUserStore().saveToken(JSON.stringify(res.data))

            // 再次发送失败的请求, 并返回结果
            return request(err.config)
         } else {
            // 刷新Token失败, 跳转至登录页
            console.log("刷新Token失败, 请重新登录")
            ElMessage.error("令牌过期, 请重新登录")
            router.push({ name: "login", query: { redirect: router.currentRoute.value.fullPath } })
            return
         }
      }
      return Promise.reject(err)
   }
)

export default request
