interface PaginationSettings {
   layout: string // 分页器布局
   prev_text: string // 替代图标显示的上一页文字
   next_text: string // 替代图标显示的下一页文字
   page_sizes: number[] // 每一页展示多少条数据选项
   default_page_size: number //  默认每一页需要展示多少条数据
   default_current_page: number // 默认当前第几页
   hide_on_single_page: boolean // 只有一页时是否隐藏
   background: boolean // 显示当前页背景颜色
}

interface TableSettings {
   border: boolean
   fit: boolean
   height: Number
   table_layout: "auto" | "fixed" | undefined
   empty_text: string
   highlight_current_row: boolean
   row_key: string
}

interface TableColumn {
   prop: string
   label: string
   visible: boolean
   disable: boolean
   sortable: boolean
   align: string
   keyPosition: string
   valuePosition: string
   width: number
   formatter?: any | undefined
}
