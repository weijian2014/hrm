import request from "@/utils/request"
import type { PromiseResponse } from "@/utils/common"

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

export const refreshToken = (data: RefreshRequest): PromiseResponse<TokenInfo> => request.post("/user/refresh", data)

//
export const getUserInfo = (): PromiseResponse<string> => request.get("/user/info")
