import request from "@/utils/request"
import type { ApiResponse, PromiseResponse } from "@/utils/common"

//
interface LoginRequest {
   username: string
   password: string
}

export const loginApi = (data: LoginRequest): PromiseResponse<TokenInfo> => request.post("/user/login", data)

export const logoutApi = (): PromiseResponse<string> => request.post("/user/logout")

interface RefreshRequest {
   refresh_token: string
}

let refreshPromise: Promise<any>
let isRefreshing = false
export const refreshToken = (data: RefreshRequest): PromiseResponse<TokenInfo> => {
   // 防止重复刷新
   if (isRefreshing) {
      return refreshPromise
   }

   isRefreshing = true
   refreshPromise = request.post("/user/refresh", data).finally(() => {
      // 无论成功还是失败都会走入finally
      isRefreshing = false
   })

   return refreshPromise
}

//
export const getUserInfo = (): PromiseResponse<string> => request.get("/user/info")
