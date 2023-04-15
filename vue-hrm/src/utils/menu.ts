import request from "@/utils/request"
import type { PromiseResponse } from "@/utils/common"

//
export interface Menu {
   id: number
   name: string
   parent_id: number
   icon: string
   url: string
   updated_at: string
}

export interface MenuListResponse {
   total: number
   menus: Menu[]
}

export const menuListApi = (): PromiseResponse<MenuListResponse> => request.get("/menu/list")

// Partial让Menu中的所有成员变成可选字段
// type MenuAddRequest = Partial<Menu>

// 从Menu中挑选出几个成员做为新的类型
type MenuAddRequest = Pick<Menu, "name" | "parent_id" | "icon" | "url">

// 组合other成为新类型
// type MenuAddRequest = Pick<Menu, "name" | "parent_id" | "icon" | "url"> & { other: string}

// 另一种写法, 忽略id和updated_at, 保留其它成员做为新的类型
// type MenuAddRequest = Omit<Menu, "id" | "updated_at">

export const menuAddApi = (data: MenuAddRequest): PromiseResponse<Menu> => request.post("/menu/add")
