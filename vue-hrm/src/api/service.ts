import service from "./request"

// 获取seller
export function getSeller(params: any) {
   return service.request({
      method: "POST",
      url: "接口一",
      data: params,
   })
}
export function Login(params: any) {
   return service.request({
      method: "POST",
      url: "接口二",
      data: params,
   })
}