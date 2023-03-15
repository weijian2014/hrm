<template>
   <div class="container">
      <el-row class="m-1 pt-2 pb-2">
         <el-col :span="12">
            <el-button :icon="Plus" type="primary" @click="handleAdd">新增</el-button>
            <el-button :icon="Upload" type="primary">导入</el-button>
            <el-button :icon="Delete" type="danger">删除</el-button>
         </el-col>
         <el-col :span="6">
            <el-input v-model="input" class="w-50" placeholder="" clearable>
               <template #prepend>模糊搜索</template>
               <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
               </template>
            </el-input>
         </el-col>
      </el-row>
      <el-table
         class=""
         ref="tableRef"
         :data="tableRows"
         :border="tableSettings?.border"
         :fit="tableSettings?.fit"
         :height="tableSettings?.height"
         table-layout="auto"
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
         class="pt-2"
         background
         :hide-on-single-page="false"
         :layout="tablePaginationSettings?.layout"
         :page-sizes="tablePaginationSettings?.page_sizes"
         :prev-text="tablePaginationSettings?.prev_text"
         :next-text="tablePaginationSettings?.next_text"
         :default-current-page="currentPage"
         :default-page-size="pageSize"
         v-model:current-page="currentPage"
         v-model:page-size="pageSize"
         v-model:total="tableTotal"
         @size-change="handleSizeChange"
         @current-change="handleCurrentChange"
         @prev-click="handlePrevClick"
         @next-click="handleNextClick" />
   </div>
   <AddVue :isShow="isShow" :id="id" :employee="employee" @save="handleSave" @canel="handleCancel"> </AddVue>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from "vue"
import { ElTable, ElPagination, ElMessage } from "element-plus"
import { Upload, Delete, Plus, Search } from "@element-plus/icons-vue"
import Employee from "./../class/Employee"
import AddVue from "./Add.vue"
import Test from "./../api/test"
import { Table } from "../class/Table"

const tableRef = ref<InstanceType<typeof ElTable>>()
const paginationRef = ref<InstanceType<typeof ElPagination>>()

const tableRows = ref<Employee[]>()
const tableTotal = ref<number>()
const tablePaginationSettings = ref<Table.PaginationSettings>()
const tableSettings = ref<Table.TableSettings>()
const tableColumns = ref<Table.TableColumn[]>()

onMounted(async () => {
   const data = await Test.getTableData()
   tableRows.value = data.rows as Employee[]
   tableTotal.value = data.total as number

   const settings = await Test.getTableSettings()
   tablePaginationSettings.value = settings.pagination
   tableSettings.value = settings.table_settings
   tableColumns.value = settings.table_columns
})

const rowClick = (row: Employee) => {
   tableRef.value!.$emit("select")
}

// AddVue组件的属性
const isShow = ref(false)
const id = ref(0)
const employee = ref<Employee>()

// 向AddVue组件传值 -- 只传isShow
const handleAdd = () => {
   console.log("新增")
   isShow.value = true
}

// 向AddVue组件传值
const handleEdit = (index: number, row: Employee) => {
   console.log(index, row)
   isShow.value = true
   id.value = index
   employee.value = row as Employee
}

// AddVue组件发送的保存事件
const handleSave = (message: string) => {
   isShow.value = false
   employee.value = new Employee()
   console.log(message)
   ElMessage.success(message)
}

// AddVue组件发送的消息事件
const handleCancel = () => {
   isShow.value = false
   employee.value = new Employee()
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

const input = ref("")
</script>
<style lang="scss" scoped></style>
