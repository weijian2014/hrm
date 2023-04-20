interface PaginationSettings {
   layout: string = "" // 分页器布局
   prev_text: string = "" // 替代图标显示的上一页文字
   next_text: string = "" // 替代图标显示的下一页文字
   page_sizes: number[] = [10, 20, 30, 40, 50, 100] // 每一页展示多少条数据选项
   default_page_size: number = 10 //  默认每一页需要展示多少条数据
   default_current_page: number = 1 // 默认当前第几页
   hide_on_single_page: boolean = false // 只有一页时是否隐藏
   background: boolean = true // 显示当前页背景颜色
}

interface TableSettings {
   border: boolean = true
   fit: boolean = true
   height: number = 500
   table_layout: "auto" | "fixed" | undefined
   empty_text: string = "N/A"
   highlight_current_row: boolean = true
   row_key: string = "id"
}

interface TableColumn {
   prop: string = ""
   label: string = ""
   visible: boolean = true
   disable: boolean = false
   sortable: boolean = true
   align: string = "center"
   keyPosition: string
   valuePosition: string
   width: number
   formatter?: any | undefined
}
