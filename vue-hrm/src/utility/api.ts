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

export const refreshToken = (): PromiseResponse<string> =>
   request({
      method: "POST",
      url: "/user/refresh",
      params: {
         refresh_token: useUserStore().tokenInfo.token,
      },
   })

//
export const getUserInfo = (): PromiseResponse<string> => request.get("/user/info")

// employee
interface EmployeeListResponseData {
   total: number
   rows: Employee[]
}
export const employeeListApi = (): PromiseResponse<EmployeeListResponseData> => request.get("/employee/list")
