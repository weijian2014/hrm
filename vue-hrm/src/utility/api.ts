import request from "./request"

// 登录返回token
interface LoginData {
   username: string
   password: string
}
interface LoginRes {
   code: number
   data: {
      token: string
      tokenHead: string
   }
   message: string
}
export const loginApi = (data: LoginData): Promise<LoginRes> =>
   request.post("/acount/login", data)
