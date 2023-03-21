import request from "./request"

// 登录返回token
interface LoginData {
   username: string
   password: string
}
interface LoginRes {
   code: number
   message: string
   data: {
      token: string
      tokenHead: string
   }
}
export const loginApi = (data: LoginData): Promise<LoginRes> =>
   request.post("/user/login", data)
