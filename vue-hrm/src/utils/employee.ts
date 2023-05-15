import request from "@/utils/request"
import type { PromiseResponse } from "@/utils/common"

interface EmployeeListResponseData {
   total: number
   employees: Employee[]
}
export const employeeListApi = (): PromiseResponse<EmployeeListResponseData> => request.get("/employee/list")

interface SearchRequest {
   key: string
}
export const employeeSearchApi = (data: SearchRequest): PromiseResponse<EmployeeListResponseData> =>
   request.post("/employee/search", data)

export const employeeUpdateApi = (data: Employee): PromiseResponse<string> => request.put("/employee/update", data)

export const employeeAddApi = (data: Employee): PromiseResponse<string> => request.post("/employee/add", data)

export const employeeDeleteApi = (id: number): PromiseResponse<string> => request.delete("/employee/" + id)
