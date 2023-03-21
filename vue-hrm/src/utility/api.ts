import request from "./request"
import type { ApiResponse } from "./common"

type PromiseResponse<T> = Promise<ApiResponse<T>>

//
interface LoginRequest {
   username: string
   password: string
}
interface LoginResponseData {
   token: string
   token_header: string
}
export const loginApi = (
   data: LoginRequest
): PromiseResponse<LoginResponseData> => request.post("/user/login", data)

//
interface InfoResponseData {}
export const getUserInfo = (): PromiseResponse<InfoResponseData> =>
   request.get("/user/info")
