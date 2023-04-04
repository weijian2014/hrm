export interface ApiResponse<T> {
   code: number
   message: string
   data: T
}

export type PromiseResponse<T> = Promise<ApiResponse<T>>
