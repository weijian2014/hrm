import { ref } from "vue"
import axios from "axios"
import Employee from "../class/Employee"
import { Table } from "../class/Table"

export default {
   setup() {
      // 这时可以定义为非响应式的变量
      const total = ref(0)
      const rows = ref<Employee[]>()
      const pagination = ref<Table.PaginationSettings>()
      const table_settings = ref<Table.TableSettings>()
      const table_columns = ref<Table.TableColumn>()

      const getTableData = (req: any) => {
         axios
            .get("./public/table_data.json", {
               params: req,
            })
            .then((response) => {
               console.log(response)
               total.value = response.data.total as number
               rows.value = response.data.rows as Employee[]
            })
            .catch((err) => {
               console.log(err)
            })
      }

      const getTableSettings = (req: any) => {
         axios
            .get("./public/table_settings.json", {
               params: req,
            })
            .then((response) => {
               console.log(response)
               pagination.value = response.data.pagination as Table.PaginationSettings
               table_settings.value = response.data.table_settings as Table.TableSettings
               table_columns.value = response.data.table_columns as Table.TableColumn
            })
            .catch((err) => {
               console.log(err)
            })
      }

      return {
         total,
         rows,
         pagination,
         table_settings,
         table_columns,
         getTableData,
         getTableSettings,
      }
   },
}
