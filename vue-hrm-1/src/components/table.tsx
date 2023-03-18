interface TableColumnSettings {
   prop: string
   label: string
   visible: boolean
   sortable: boolean
   align: string
}

interface PaginationSettings {
   layout: string // 分页器布局
   prev_text: string
   next_text: string
   page_sizes: number[] // 每一页展示多少条数据
   default_page_size: number // 代表的是每一页需要展示多少条数据, 与page-size属性相同
   default_current_page: number // 当前第几页, 与current-page属性相同
   hide_on_single_page: boolean
}

type TableColumnSettingsList = TableColumnSettings[]

export function useColumns() {
   const columns: TableColumnSettingsList = [
      {
         prop: "id",
         label: "序号",
         visible: false,
         sortable: false,
         align: "center",
      },
      {
         prop: "name",
         label: "姓名",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "gender",
         label: "性别",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "age",
         label: "年龄",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "first_work_time",
         label: "工作时间",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "salary",
         label: "工资",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "post",
         label: "岗位",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "social_security",
         label: "社保",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "phone",
         label: "电话",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "former_employer",
         label: "原单位",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "height",
         label: "身高",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "weight",
         label: "体重",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "degree",
         label: "学历",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "political_status",
         label: "政治面貌",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "identifier",
         label: "身份证",
         visible: true,
         sortable: false,
         align: "center",
      },
      {
         prop: "security_card",
         label: "保安证",
         visible: true,
         sortable: false,
         align: "center",
      },
      {
         prop: "current_address",
         label: "现住址",
         visible: false,
         sortable: false,
         align: "center",
      },
      {
         prop: "comments",
         label: "备注",
         visible: false,
         sortable: false,
         align: "center",
      },
   ]

   const paginationSettings: PaginationSettings = {}
   return {
      columns,
   }
}
