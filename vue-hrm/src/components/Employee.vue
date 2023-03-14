<template>
   <div class="container">
      <el-row>
         <el-col> </el-col>
         <el-col> </el-col>
      </el-row>
      <el-table
         ref="tableRef"
         :data="tableData"
         :border="tableSettings?.border"
         :fit="tableSettings?.fit"
         :height="tableSettings?.height"
         :table-layout="tableSettings?.table_layout"
         :empty-text="tableSettings?.empty_text"
         :highlight-current-row="tableSettings?.highlight_current_row"
         :row-key="tableSettings?.row_key"
         :default-sort="{ prop: 'name', order: 'descending' }"
         @row-click="rowClick"
         style="width: 100%">
         <el-table-column fixed type="selection" />
         <el-table-column
            v-for="(column, index) in tableColumns"
            :key="column.prop"
            :prop="column.prop"
            :label="column.label"
            :sortable="column.sortable"
            :align="column.align"
            :index="index"></el-table-column>
         <el-table-column width="140" label="操作" align="center">
            <template #default="scope">
               <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
               <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
            </template>
         </el-table-column>
      </el-table>
      <el-pagination
         ref="paginationRef"
         background
         :layout="paginationSettings?.layout"
         :page-sizes="paginationSettings?.page_sizes"
         :prev-text="paginationSettings?.prev_text"
         :next-text="paginationSettings?.next_text"
         :default-current-page="paginationSettings?.default_current_page"
         :default-page-size="paginationSettings?.default_page_size"
         :hide-on-single-page="false"
         :total="tableTotal"
         :page-size="pageSize"
         :current-page="currentPage"
         @size-change="handleSizeChange"
         @current-change="handleCurrentChange"
         @prev-click="handlePrevClick"
         @next-click="handleNextClick" />
   </div>
</template>

<script lang="ts" setup>
import { TableColumnCtx, ElTable, ElPagination } from "element-plus"
import { ref, computed } from "vue"
import Axios from "axios"

interface Employee {
   id: number
   name: string
   gender: string
   age: number
   work_time: string
   salary: number
   post: string
   social_security: string
   phone: string
   former_employer: string
   height: number
   weight: number
   diploma: string
   political_status: string
   identifier: string
   security_card: string
   current_address: string
   comments: string
}

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

interface TableSettings {
   paginationSettings: PaginationSettings
   border: boolean
   fit: boolean
   height: number
   table_layout: string
   empty_text: string
   highlight_current_row: boolean
   row_key: string
   table_columns: TableColumnSettings[]
}

const tableRef = ref<InstanceType<typeof ElTable>>()
const paginationRef = ref<InstanceType<typeof ElPagination>>()

const formatter = (row: Employee, column: TableColumnCtx<Employee>) => {
   return row.current_address
}

const tableData = ref<Employee[]>()
const tableTotal = ref<number>()
Axios.get("./public/table_data.json", {
   params: {},
})
   .then(function (response) {
      // console.log(response)
      tableData.value = response.data.rows
      tableTotal.value = response.data.total
      console.log(tableTotal)
   })
   .catch(function (error) {
      console.log(error)
   })
   .then(function () {
      // always executed
   })

const tableSettings = ref<TableSettings>()
const paginationSettings = ref<PaginationSettings>()
Axios.get("./public/table_settings.json", {
   params: {},
})
   .then(function (response) {
      tableSettings.value = response.data
      paginationSettings.value = response.data.pagination
      console.log(paginationSettings)
   })
   .catch(function (error) {
      console.log(error)
   })
   .then(function () {
      // always executed
   })

let tableColumns = ref<TableColumnSettings[]>()
tableColumns = computed(() => {
   return tableSettings.value?.table_columns.filter((column) => column.visible)
})

const rowClick = (row: Employee) => {
   tableRef.value!.toggleRowSelection(row)
}

const handleEdit = (index: number, row: Employee) => {
   console.log(index, row)
}
const handleDelete = (index: number, row: Employee) => {
   console.log(index, row)
}

const pageSize = ref(10)
const currentPage = ref(1)
const handleSizeChange = (value: number) => {
   console.log(value)
   pageSize.value = value
}
const handleCurrentChange = (value: number) => {
   console.log(value)
   currentPage.value = value
}
const handlePrevClick = (value: number) => {
   console.log(value)
   currentPage.value = value
}
const handleNextClick = (value: number) => {
   console.log(value)
   currentPage.value = value
}
</script>
<style lang="scss" scoped></style>
