import request from "@/utils/request"
import type { ApiResponse, PromiseResponse } from "@/utils/common"

// Id        uint64    `json:"id" description:"菜单ID"`
// Name      string    `json:"name" description:"菜单名"`
// Url       string    `json:"url" description:"菜单链接"`
// ParentId  uint64    `json:"parent_id" description:"菜单的父级菜单ID"`
// UpdatedAt time.Time `json:"updated_at" description:"更新时间"`

//
export interface MenuListResponse {
   id: number
   name: string
   url: string
   icon: string
   parent_id: number
   updated_at: string
}

export const menuListApi = (): PromiseResponse<MenuListResponse[]> => request.post("/menu/list")
