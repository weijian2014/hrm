import request from "@/utils/request"
import type { PromiseResponse } from "@/utils/common"

export interface PostListResponse {
   total: number
   posts: Post[]
}

export const postListApi = (): PromiseResponse<PostListResponse> => request.get("/post/list")

export const menuAddApi = (data: Post): PromiseResponse<Post> => request.post("/post/add")
