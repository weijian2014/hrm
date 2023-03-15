import axios from "axios"
import Employee from "../class/Employee"
import { Table } from "../class/Table"

export default {
   getTableData(): any {
      ;async (req: any) => {
         const response = await axios
            .get("./public/table_data.json", {
               params: { req },
            })
            .then(function (response) {
               return {
                  total: response.data.total as number,
                  rows: response.data.rows as Employee[],
               }
            })
            .catch(function (error) {
               console.log(error)
            })
            .then(function () {
               // always executed
            })
      }
   },
   getTableSettings(): any {
      ;async (req: any) => {
         const response = await axios
            .get("./public/table_settings.json", {
               params: { req },
            })
            .then(function (response) {
               return {
                  pagination: response.data.pagination as Table.PaginationSettings,
                  table_settings: response.data.table_settings as Table.TableSettings,
                  table_columns: response.data.table_columns as Table.TableColumn,
               }
            })
            .catch(function (error) {
               console.log(error)
            })
            .then(function () {
               // always executed
            })
      }
   },
}
