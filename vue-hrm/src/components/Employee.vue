<template>
   <el-table
      ref="multipleTableRef"
      :data="tableData"
      :border="tableSettings?.border"
      :fit="tableSettings?.fit"
      :height="tableSettings?.height"
      :table-layout="tableSettings?.table_layout"
      :empty-text="tableSettings?.empty_text"
      :highlight-current-row="tableSettings?.highlight_current_row"
      :row-key="tableSettings?.row_key"
      default-sort="{ prop: 'name', order: 'descending' }"
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
</template>

<script lang="ts" setup>
import type { TableColumnCtx, ElTable } from "element-plus"
import { ref, computed } from "vue"
import Axios from "axios"

interface Employee {
   id: number
   name: string
   gender: string
   age: number
   work_time: number
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

interface TableSettings {
   border: boolean
   fit: boolean
   height: number
   table_layout: string
   empty_text: string
   highlight_current_row: boolean
   row_key: string
   table_columns: TableColumnSettings[]
}

const multipleTableRef = ref<InstanceType<typeof ElTable>>()

const formatter = (row: Employee, column: TableColumnCtx<Employee>) => {
   return row.current_address
}

const tableData = ref<Employee[]>()
Axios.get("./public/testData.json", {
   params: {},
})
   .then(function (response) {
      console.log(response)
      tableData.value = response.data.rows
   })
   .catch(function (error) {
      console.log(error)
   })
   .then(function () {
      // always executed
   })

const tableSettings = ref<TableSettings>()
Axios.get("./public/tableSettings.json", {
   params: {},
})
   .then(function (response) {
      console.log(response)
      tableSettings.value = response.data
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
   multipleTableRef.value!.toggleRowSelection(row)
}

const handleEdit = (index: number, row: Employee) => {
   console.log(index, row)
}
const handleDelete = (index: number, row: Employee) => {
   console.log(index, row)
}
</script>
<style lang="scss" scoped></style>
