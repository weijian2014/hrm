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
      },
      {
         prop: "name",
         label: "姓名",
         visible: true,
         disable: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "gender",
         label: "性别",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "birthday",
         label: "生日",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
         formatter: dateFormatter,
      },
      {
         prop: "height",
         label: "身高",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "weight",
         label: "体重",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "degree",
         label: "学历",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "identifier",
         label: "身份证号",
         visible: true,
         disable: false,
         sortable: false,
         align: "center",
      },
      {
         prop: "phone",
         label: "手机号码",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "political_status",
         label: "政治面貌",
         visible: false,
         disable: false,
         sortable: true,
         align: "left",
      },
      {
         prop: "social_security",
         label: "社保",
         visible: false,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "current_address",
         label: "现住址",
         visible: false,
         disable: false,
         sortable: false,
         align: "left",
      },
      {
         prop: "first_work_time",
         label: "首次工作日期",
         visible: false,
         disable: false,
         sortable: true,
         align: "center",
         formatter: dateFormatter,
      },
      {
         prop: "former_employer",
         label: "原单位",
         visible: false,
         disable: false,
         sortable: true,
         align: "left",
      },
      {
         prop: "post",
         label: "岗位",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "salary",
         label: "工资",
         visible: true,
         disable: false,
         sortable: true,
         align: "center",
      },
      {
         prop: "security_card",
         label: "保安证",
         visible: false,
         disable: false,
         sortable: false,
         align: "center",
      },
      {
         prop: "comments",
         label: "备注",
         visible: false,
         disable: false,
         sortable: false,
         align: "left",
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
   const seachInputValue = ref("")

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
      seachInputValue,
      isAddOrEditShow,
      addOrEditTitle,
      addOrEditData,
      selections,
   }
}

export const excelPosition = new Map([
   ["name", "B2"],
   ["gender", "D2"],
   ["birthday", "F2"],
   ["height", "B3"],
   ["weight", "D3"],
   ["degree", "F3"],
   ["identifier", "B4"],
   ["phone", "D4"],
   ["political_status", "F4"],
   ["social_security", "B5"],
   ["current_address", "D5"],
   ["first_work_time", "B6"],
   ["former_employer", "D6"],
   ["post", "B7"],
   ["salary", "D7"],
   ["security_card", "F7"],
   ["comments", "A8"],
])

export const excelFeilds = new Map([
   ["人员情况调查", "A1"],
   ["姓名", "A2"],
   ["性别", "C2"],
   ["生日", "E2"],
   ["身高(cm)", "A3"],
   ["体重(kg)", "C3"],
   ["学历", "E3"],
   ["身份证号", "A4"],
   ["手机号码", "C4"],
   ["政治面貌", "E4"],
   ["社保", "A5"],
   ["现住址", "C5"],
   ["首次工作日期", "A6"],
   ["原单位", "C6"],
   ["岗位", "A7"],
   ["工资(¥)", "C7"],
   ["保安证", "E7"],
])

// const propToLableMap = computed(() => {
//    let pl = new Map<string, string>()
//    useSettings().columns.forEach((item) => {
//       pl.set(item.prop, item.label)
//    })
//    return pl
// })

// const lableToPropMap = computed(() => {
//    let lp = new Map<string, string>()
//    useSettings().columns.forEach((item) => {
//       lp.set(item.label, item.prop)
//    })
//    return lp
// })

const excelHeader = computed(() => {
   let header: string[] = []
   useSettings().columns.forEach((item) => {
      header.push(item.label)
   })
   return header
})

export function convertForExport(rows: Employee[]): {
   excelHeaders: string[]
   excelBodys: any[]
   excelColumnsWidth: any[]
} {
   const excelData: any = []
   const maxLenght: number[] = [30, 50, 30, 70, 40, 40, 70, 125, 90, 80, 30, 220, 90, 160, 40, 55, 70, 300]

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
   return { excelHeaders: excelHeader.value, excelBodys: excelData, excelColumnsWidth: maxLenght }
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
