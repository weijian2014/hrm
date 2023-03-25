interface PaginationSettings {
   layout: string = "" // 分页器布局
   prev_text: string = ""
   next_text: string = ""
   page_sizes: number[] = [10, 20, 30, 40, 50, 100] // 每一页展示多少条数据
   default_page_size: number = 10 // 代表的是每一页需要展示多少条数据, 与page-size属性相同
   default_current_page: number = 1 // 当前第几页, 与current-page属性相同
   hide_on_single_page: boolean = false
}

interface TableSettings {
   border: boolean = true
   fit: boolean = true
   height: number = 500
   table_layout: string = "auto" as const
   empty_text: string = "N/A"
   highlight_current_row: boolean = true
   row_key: string = "id"
}

interface TableColumn {
   prop: string = ""
   label: string = ""
   visible: boolean = true
   sortable: boolean = true
   align: string = "center"
}
