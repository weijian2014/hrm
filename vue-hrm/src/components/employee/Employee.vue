<script lang="ts" setup>
import AddVue from "@/components/employee/Add.vue"
import { employeeListApi } from "@/utils/employee"
import type { CheckboxValueType } from "element-plus"
import { useSettings, useData } from "./index"

const { table, columns, pagination, checkList } = useSettings()

const { tableRef, tableData, isButtonDisabled, seachInputValue, isShow, title, rowData, isVisibleColumnsSettings } =
   useData()

employeeListApi()
   .then((res) => {
      if (res.code === 200) {
         console.log(res)
         // totalRows.value = res.data.total
         tableData.value = res.data.rows
      }
   })
   .catch((res) => {
      console.log(res)
   })

// 可显示的列
const visibleColumns = computed(() => {
   return columns.filter((column: { visible: boolean }) => column.visible)
})

////// 表头工具栏
let lastCheckList = checkList.value as CheckboxValueType[]
const handleCheckedChange = (values: CheckboxValueType[]) => {
   // values为所有被选中的列
   // 求出values与lastCheckedLables的差集--取消选中的lable
   const diffA = values.concat(lastCheckList).filter((v) => !values.includes(v))

   // 求出lastCheckedLables与values的差集--新增选中的lable
   var diffB = lastCheckList.concat(values).filter((v) => !lastCheckList.includes(v))

   diffA.forEach((item1) => {
      columns.forEach((item2) => {
         if (item2.label === item1) {
            item2.visible = false
         }
      })
   })

   diffB.forEach((item1) => {
      columns.forEach((item2) => {
         if (item2.label === item1) {
            item2.visible = true
         }
      })
   })

   lastCheckList = values
}

////// 表格
const handleSelectChange = (rows: Employee[]) => {
   console.log("handleSelectChange", rows)
   isButtonDisabled.value = rows.length === 0
}

const handleRowClick = (row: Employee) => {
   // console.log("handleRowClick", row)
}

const handleAdd = () => {
   console.log("新增")
   title.value = "新增"
   rowData.value = {} as Employee
   isShow.value = true
}

const handleEdit = (index: number, row: Employee | undefined) => {
   console.log("handleEdit", index, row)
   if (row) {
      rowData.value = row
      console.log("handleEdit", rowData.value)
      title.value = "修改"
      isShow.value = true
   }
}

const handleDelete = (index: number, row: Employee) => {
   console.log("handleDelete", index, row)
   tableData.splice(index, 1)
}

////// Add组件

// AddVue组件发送的保存事件
const handleSave = (message: string) => {
   // todo 更新表格数据
   isShow.value = false
   console.log("handleSave", message, rowData.value)
   ElMessage.success(message)
}

// AddVue组件发送的消息事件
const handleCancel = (message: string) => {
   isShow.value = false
   console.log("handleCancel", message, rowData.value)
   ElMessage.info(message)
}
</script>

<template>
   <div>
      <!-- 表头工具 -->
      <el-row class="mb-4">
         <el-col :span="12">
            <el-button type="primary" @click="handleAdd">
               <IEpPlus />
               <span style="vertical-align: middle">新增</span>
            </el-button>
            <el-button type="primary">
               <IEpUpload />
               <span style="vertical-align: middle">导入</span>
            </el-button>
            <el-button type="primary" :disabled="isButtonDisabled">
               <IEpDownload />
               <span style="vertical-align: middle">导出</span>
            </el-button>
            <el-button type="primary" :disabled="isButtonDisabled">
               <IEpDelete />
               <span style="vertical-align: middle">删除</span>
            </el-button>
         </el-col>
         <el-col :span="6">
            <el-input v-model="seachInputValue" placeholder="" clearable>
               <template #suffix>
                  <el-icon><IEpSearch /></el-icon>
               </template>
               <template #prepend>模糊搜索</template>
            </el-input>
         </el-col>
         <el-col :offset="1" :span="5">
            <el-popover
               placement="left-end"
               title="列筛选"
               trigger="click"
               width="300"
               :visible="isVisibleColumnsSettings">
               <template #reference>
                  <el-button type="primary" @click="isVisibleColumnsSettings = !isVisibleColumnsSettings">
                     <IEpPlus />
                     <span style="vertical-align: middle">列筛选</span>
                  </el-button>
               </template>
               <el-checkbox-group v-model="checkList" @change="handleCheckedChange">
                  <el-checkbox
                     v-for="item in columns"
                     :label="item.label"
                     :checked="item.visible"
                     :disabled="item.disable"
                     >{{ item.label }}</el-checkbox
                  >
               </el-checkbox-group>
            </el-popover>
         </el-col>
      </el-row>
      <!-- 表格, table-layout为auto时表头无法固定 -->
      <el-table
         ref="tableRef"
         style="width: 100%"
         :data="tableData"
         :border="table.border"
         :fit="table.fit"
         :height="table.height"
         :table-layout="table.table_layout"
         :empty-text="table.empty_text"
         :highlight-current-row="table.highlight_current_row"
         :row-key="table.row_key"
         :default-sort="{ prop: 'name', order: 'descending' }"
         @selection-change="handleSelectChange"
         @row-click="handleRowClick">
         <el-table-column type="selection" align="center" />
         <el-table-column
            v-for="(column, index) in visibleColumns"
            :key="column.prop"
            :prop="column.prop"
            :label="column.label"
            :sortable="column.sortable"
            :align="column.align"
            :formatter="column.formatter"
            :index="index"></el-table-column>
         <el-table-column width="140" label="操作" align="center">
            <template #default="scope">
               <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
               <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
            </template>
         </el-table-column>
      </el-table>
   </div>
   <AddVue :isShow="isShow" :title="title" :formData="rowData" @save="handleSave" @cancel="handleCancel"> </AddVue>
</template>

<style lang="scss" scoped></style>
