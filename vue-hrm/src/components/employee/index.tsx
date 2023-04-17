import { getAgeByBirthday } from "@/utils/common"
import type { TableColumnCtx } from "element-plus"
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
   const tableRef = {}

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

export default class Validator {
   /**
    * @description 根据身份证号前17位计算出最后一位
    * @param { string } n 身份证号前17位
    * @returns { string } 最后一位
    */
   private static computeLastNum(n: string): string {
      const computeNums = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2]
      const remainderToLast = ["1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"]
      let total = 0
      n.split("").forEach((n, i) => {
         total += Number(n) * computeNums[i]
      })
      return remainderToLast[total % 11]
   }

   /**
    * @description 检验身份证号格式是否正确
    * @param { string } n 身份证号
    * @returns { boolean }
    */
   static isIdentifierValid(n: string): boolean {
      // 检验长度
      if (n.length !== 18) return false
      // 地区是否存在(需要有地区的库，可以查看开源项目)
      // if(!Area.validateAreaCodeExist(n.substr(0, 6))) return false
      // 校验最后一位
      return this.computeLastNum(n.substr(0, 17)) === n.substr(17, 1)
   }

   /**
    * @description 校验手机号格式是否正确
    * @param { string } phone 手机号
    * @returns { boolean }
    */
   static isPhoneValid(phone: string): boolean {
      const phoneRegExp: RegExp =
         /^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$/
      return phoneRegExp.test(phone)
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
