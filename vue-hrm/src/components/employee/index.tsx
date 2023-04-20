import { getAgeByBirthday } from "@/utils/common"
import type { TableColumnCtx, UploadFile } from "element-plus"
import moment from "moment"

export function useSettings() {
   // 表格配置
   const table = reactive<TableSettings>({
      border: true,
      fit: true,
      highlight_current_row: true,
      height: 500,
      empty_text: "暂无数据",
      table_layout: "auto",
      row_key: "id",
   })

   // 列配置
   const columns = reactive<TableColumn[]>([
      {
         prop: "id",
         label: "序号",
         visible: false,
         disable: true,
         sortable: false,
         align: "center",
         keyPosition: "",
         valuePosition: "",
         width: 30,
      },
      {
         prop: "name",
         label: "姓名",
         visible: true,
         disable: true,
         sortable: true,
         align: "center",
         keyPosition: "A2",
         valuePosition: "B2",
         width: 50,
      },
      {
         prop: "gender",
         label: "性别",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "C2",
         valuePosition: "D2",
         width: 30,
      },
      {
         prop: "birthday",
         label: "生日",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "E2",
         valuePosition: "F2",
         width: 70,
         formatter: dateFormatter,
      },
      {
         prop: "height",
         label: "身高(cm)",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "A3",
         valuePosition: "B3",
         width: 50,
      },
      {
         prop: "weight",
         label: "体重(kg)",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "C3",
         valuePosition: "D3",
         width: 50,
      },
      {
         prop: "degree",
         label: "学历",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "E3",
         valuePosition: "F3",
         width: 70,
      },
      {
         prop: "identifier",
         label: "身份证号",
         visible: true,
         disable: false,
         sortable: false,
         align: "center",
         keyPosition: "A4",
         valuePosition: "B4",
         width: 125,
      },
      {
         prop: "phone",
         label: "手机号码",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "C4",
         valuePosition: "D4",
         width: 90,
      },
      {
         prop: "political_status",
         label: "政治面貌",
         visible: false,
         disable: false,
         sortable: true,
         align: "left",
         keyPosition: "E4",
         valuePosition: "F4",
         width: 80,
      },
      {
         prop: "social_security",
         label: "社保",
         visible: false,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "A5",
         valuePosition: "B5",
         width: 30,
      },
      {
         prop: "current_address",
         label: "现住址",
         visible: false,
         disable: false,
         sortable: false,
         align: "left",
         keyPosition: "C5",
         valuePosition: "D5",
         width: 220,
      },
      {
         prop: "first_work_time",
         label: "首次工作日期",
         visible: false,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "A6",
         valuePosition: "B6",
         width: 90,
         formatter: dateFormatter,
      },
      {
         prop: "former_employer",
         label: "原单位",
         visible: false,
         disable: false,
         sortable: true,
         align: "left",
         keyPosition: "C6",
         valuePosition: "D6",
         width: 160,
      },
      {
         prop: "post",
         label: "岗位",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "A7",
         valuePosition: "B7",
         width: 40,
      },
      {
         prop: "salary",
         label: "工资(¥)",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         keyPosition: "C7",
         valuePosition: "D7",
         width: 55,
      },
      {
         prop: "security_card",
         label: "保安证",
         visible: false,
         disable: false,
         sortable: false,
         align: "center",
         keyPosition: "E7",
         valuePosition: "F7",
         width: 70,
      },
      {
         prop: "comments",
         label: "备注",
         visible: false,
         disable: false,
         sortable: false,
         align: "left",
         keyPosition: "",
         valuePosition: "A8",
         width: 300,
      },
   ])

   // 分页配置
   const pagination = reactive<PaginationSettings>({
      layout: "->, total, sizes, prev, pager, next",
      prev_text: "上一页",
      next_text: "下一页",
      page_sizes: [10, 20, 30, 40, 50, 100],
      default_page_size: 10,
      default_current_page: 1,
      hide_on_single_page: false,
   })

   // 需要显示列的lable
   const checkList = computed<string[]>(() => {
      let visibles: string[] = []
      columns.forEach((item) => {
         if (item.visible) {
            visibles.push(item.label)
         }
      })
      return visibles
   })

   return {
      table,
      columns,
      pagination,
      checkList,
   }
}

export function useData() {
   // el-table组件对象, 自动关联到el-table组件
   const tableRef = ref()

   // 表格数据
   const tableData = ref<Employee[]>([] as Employee[])

   // 是否禁用表头的修改和删除按钮
   const isButtonDisabled = ref(true)

   // 搜索框的值
   const searchInputValue = ref("")

   // AddOrEdit组件属性
   const isAddOrEditShow = ref(false)
   const addOrEditTitle = ref("修改员工")
   const addOrEditData = ref<Employee>({} as Employee)

   // 表格选中的行
   const selections = ref<Employee[]>([] as Employee[])

   return {
      tableRef,
      tableData,
      isButtonDisabled,
      searchInputValue,
      isAddOrEditShow,
      addOrEditTitle,
      addOrEditData,
      selections,
   }
}

export const excelKeyPosition = computed(() => {
   let ekp = new Map()
   useSettings().columns.forEach((item) => {
      if (item.prop != "id" && item.prop != "comments") {
         ekp.set(item.label, item.keyPosition)
      }
   })
   return ekp
})

export const excelValuePosition = computed(() => {
   let evp = new Map()
   useSettings().columns.forEach((item) => {
      if (item.prop != "id") {
         evp.set(item.prop, item.valuePosition)
      }
   })
   return evp
})

const excelHeader = computed(() => {
   let header: string[] = []
   useSettings().columns.forEach((item) => {
      header.push(item.label)
   })
   return header
})

const cexelWidth = computed(() => {
   let widthArr: number[] = []
   useSettings().columns.forEach((item) => {
      widthArr.push(item.width)
   })
   return widthArr
})

export function convertForExport(rows: Employee[]): {
   excelHeaders: string[]
   excelBodys: any[]
   excelColumnsWidth: any[]
} {
   const excelData: any = []

   rows.map((row) => {
      let rowData = [
         row.id,
         row.name,
         row.gender,
         moment(row.birthday).format("YYYY-MM-DD"),
         row.height,
         row.weight,
         row.degree,
         row.identifier,
         row.phone,
         row.political_status,
         row.social_security,
         row.current_address,
         moment(row.first_work_time).format("YYYY-MM-DD"),
         row.former_employer,
         row.post,
         row.salary,
         row.security_card,
         row.comments,
      ]

      // 根据单元格内容的长度计算列宽
      // rowData.forEach((v, i) => {
      //    maxLenght[i] = Math.max(v.toString().length, maxLenght[i])
      // })

      excelData.push(rowData)
   })
   return { excelHeaders: excelHeader.value, excelBodys: excelData, excelColumnsWidth: cexelWidth.value }
}

export function dateFormatter(row: any, column: TableColumnCtx<any>, cellValue: any, index: number) {
   var date = row[column.property]
   if (date == undefined) {
      return "未知"
   }

   const dateStr = moment(date).format("YYYY-MM-DD")
   const age = getAgeByBirthday(dateStr)
   if (column.property === "birthday") {
      // 生日
      return dateStr + " (" + age + "岁)"
   } else {
      // 首次工作
      return dateStr + " (" + age + "年)"
   }
}

export function readFile(file: UploadFile) {
   return new Promise(function (resolve, reject) {
      const reader = new FileReader()
      reader.onload = (e) => {
         resolve(e.target?.result)
      }
      reader.readAsBinaryString(file.raw as Blob)
   })
}
