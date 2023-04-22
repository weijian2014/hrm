import request from "@/utils/request"
import type { PromiseResponse } from "@/utils/common"

export interface PostListResponse {
   total: number
   posts: Post[]
}

export const postListApi = (): PromiseResponse<PostListResponse> => request.get("/post/list")

export const postAddApi = (data: Post): PromiseResponse<Post> => request.post("/post/add", data)

export const postUpdateApi = (data: Post): PromiseResponse<string> => request.put("/post/update", data)

export const postDeleteApi = (id: number): PromiseResponse<string> => request.delete("/post/" + id)
