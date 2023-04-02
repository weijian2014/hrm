import request from "./request"
import type { ApiResponse } from "./common"

type PromiseResponse<T> = Promise<ApiResponse<T>>

//
interface LoginRequest {
   username: string
   password: string
}

export const loginApi = (data: LoginRequest): PromiseResponse<TokenInfo> => request.post("/user/login", data)

//
export const getUserInfo = (): PromiseResponse<string> => request.get("/user/info")

// employee
interface EmployeeListResponseData {
   total: number
   rows: Employee[]
}
export const employeeListApi = (): PromiseResponse<EmployeeListResponseData> => request.get("/employee/list")
