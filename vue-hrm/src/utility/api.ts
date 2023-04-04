import request from "./request"
import type { ApiResponse } from "./common"
import { useUserStore } from "@/store/user"

type PromiseResponse<T> = Promise<ApiResponse<T>>

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

// employee
interface EmployeeListResponseData {
   total: number
   rows: Employee[]
}
export const employeeListApi = (): PromiseResponse<EmployeeListResponseData> => request.get("/employee/list")
