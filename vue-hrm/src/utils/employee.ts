import request from "@/utils/request"
import type { PromiseResponse } from "@/utils/common"

interface EmployeeListResponseData {
   total: number
   rows: Employee[]
}
export const employeeListApi = (): PromiseResponse<EmployeeListResponseData> => request.get("/employee/list")
