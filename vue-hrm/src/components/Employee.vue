<template>
   <div class="container" style="margin: auto">
      <el-row class="m-1 pt-2 pb-2">
         <el-col :span="12">
            <el-button :icon="Plus" type="primary" @click="handleAdd">新增</el-button>
            <el-button :icon="Upload" type="primary">导入</el-button>
            <el-button :icon="Download" :disabled="isDisabled" type="primary">导出</el-button>
            <el-button :icon="Delete" :disabled="isDisabled" type="danger">删除</el-button>
         </el-col>
         <el-col :span="6">
            <el-input v-model="input" class="w-50" placeholder="" clearable>
               <template #prepend>模糊搜索</template>
               <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
               </template>
            </el-input>
         </el-col>
         <el-col :span="6" class="pl-2">
            <el-dropdown split-button :hide-on-click="false" type="primary" @click="handleDropdownClick">
               列
               <template #dropdown>
                  <el-dropdown-menu>
                     <el-dropdown-item v-for="(column, index) in tableColumns"
                        ><el-checkbox :key="column.prop" :checked="column.visible">{{
                           column.label
                        }}</el-checkbox></el-dropdown-item
                     >
                  </el-dropdown-menu>
               </template>
            </el-dropdown>
         </el-col>
      </el-row>
      <el-table
         class=""
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
         @selection-change="handleSelectChange"
         @row-click="handleRowClick"
         style="width: 100%">
         <el-table-column type="selection" />
         <el-table-column
            v-for="(column, index) in columns"
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
         layout="->, total, sizes, prev, pager, next"
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
   <AddVue :isShow="isShow" :id="id" :employee="employee" @save="handleSave" @cancel="handleCancel"> </AddVue>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue"
import { ElTable, ElPagination, ElMessage } from "element-plus"
import { Upload, Download, Delete, Plus, Search } from "@element-plus/icons-vue"
import Employee from "@/class/Employee"
import AddVue from "./Add.vue"
import type Table from "@/class/Table"
import Axios from "axios"

const tableRef = ref<InstanceType<typeof ElTable>>()
const paginationRef = ref<InstanceType<typeof ElPagination>>()

// 在ref中使用数组或者对象, 返回的是reactive类型
// 基本类型一般使用ref, 数组或者对象一般使用reactive
const tableData = ref<Employee[]>()
const tableTotal = ref<number>()
const tablePaginationSettings = ref<Table.PaginationSettings>()
const tableSettings = ref<Table.TableSettings>()
const tableColumns = ref<Table.TableColumn[]>()

const columns = computed(() => {
   return tableColumns.value?.filter((column) => column.visible)
})

Axios.get("./public/table_data.json", {
   params: {},
})
   .then(function (response) {
      tableData.value = response.data.rows
      tableTotal.value = Number(response.data.total)
      console.log("getTableData", tableData.value, tableTotal.value)
   })
   .catch(function (error) {
      console.log(error)
   })
   .then(function () {
      // always executed
   })

Axios.get("./public/table_settings.json", {
   params: {},
})
   .then(function (response) {
      tablePaginationSettings.value = response.data.pagination
      tableSettings.value = response.data.table_settings
      tableColumns.value = response.data.table_columns
      console.log("getTableSettings", tablePaginationSettings.value, tableSettings.value, tableColumns.value)
   })
   .catch(function (error) {
      console.log(error)
   })
   .then(function () {
      // always executed
   })

////////////////////////////////////////////////////////
// AddVue组件的属性
const isShow = ref(false)
const id = ref(0)
const employee = ref<Employee>()

// 向AddVue组件传值 -- 只传isShow
const handleAdd = () => {
   console.log("新增")
   employee.value = new Employee()
   isShow.value = true
}

// 向AddVue组件传值
const handleEdit = (index: number, row: Employee | undefined) => {
   console.log(index, row)
   id.value = index
   employee.value = row
   isShow.value = true
}

// AddVue组件发送的保存事件
const handleSave = (message: string) => {
   isShow.value = false
   console.log("handleSave", message)
   ElMessage.success(message)
}

// AddVue组件发送的消息事件
const handleCancel = (message: string) => {
   isShow.value = false
   console.log("handleCancel", employee.value)
   employee.value = new Employee()
   ElMessage.info(message)
}

//表格操作//////////////////////////////////////////////////////
const selections = ref<Employee[]>()
let isDisabled = ref<boolean>(true)
const handleSelectChange = (rows: Employee[]) => {
   console.log("handleSelectChange", rows)
   selections.value = rows
   isDisabled.value = selections.value.length == 0 ? true : false
}
const handleRowClick = (row: Employee) => {
   console.log("handleRowClick", row)
}
const handleDelete = (index: number, row: Employee) => {
   console.log(index, row)
}
const pageSize = ref(tablePaginationSettings.value?.default_page_size)
const currentPage = ref(tablePaginationSettings.value?.default_current_page)
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

// 搜索框
const input = ref("")

// 列
const handleDropdownClick = () => {}

const checkedPoliticalStatus = computed(() => {
   let isChecked = false
   tableColumns.value?.forEach((column) => {
      if (column.prop == "political_status") {
         isChecked = column.visible
      }
   })
   return isChecked
})

const checkedFormerEmployer = computed(() => {
   let isChecked = false
   tableColumns.value?.forEach((column) => {
      if (column.prop == "former_employer") {
         isChecked = column.visible
      }
   })
   return isChecked
})

const checkedIdentifier = computed(() => {
   let isChecked = false
   tableColumns.value?.forEach((column) => {
      if (column.prop == "identifier") {
         isChecked = column.visible
      }
   })
   return isChecked
})

const checkedSecurityCard = computed(() => {
   let isChecked = false
   tableColumns.value?.forEach((column) => {
      if (column.prop == "security_card") {
         isChecked = column.visible
      }
   })
   return isChecked
})

const checkedCurrentAddress = computed(() => {
   let isChecked = false
   tableColumns.value?.forEach((column) => {
      if (column.prop == "current_address") {
         isChecked = column.visible
      }
   })
   return isChecked
})

const checkedComments = computed(() => {
   let isChecked = false
   tableColumns.value?.forEach((column) => {
      if (column.prop == "comments") {
         isChecked = column.visible
      }
   })
   return isChecked
})
</script>
<style lang="scss" scoped></style>
