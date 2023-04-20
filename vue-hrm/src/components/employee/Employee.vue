<script lang="ts" setup>
import AddOrEdit from "@/components/employee/AddOrEdit.vue"
import { employeeListApi, employeeSearchApi, employeeDeleteApi } from "@/utils/employee"
import type { CheckboxValueType } from "element-plus"
import { useSettings, useData, convertForExport } from "./index"
import moment from "moment"
import * as XLSX from "xlsx"

const { table, columns, pagination, checkList } = useSettings()

const {
   tableRef,
   tableData,
   isButtonDisabled,
   searchInputValue,
   isAddOrEditShow,
   addOrEditTitle,
   addOrEditData,
   selections,
} = useData()

const refresh = async () => {
   await employeeListApi()
      .then((res) => {
         if (res.code === 200) {
            console.log(res)
            // totalRows.value = res.data.total
            tableData.value = res.data.rows
         }
      })
      .catch((res) => {
         console.log(res)
         return new Promise(() => {})
      })
}

refresh()

const deleteAndRefresh = async (ids: number[]) => {
   for (let i = 0; i < ids.length; ++i) {
      await employeeDeleteApi(ids[i])
         .then((res) => {
            if (res.code === 200) {
               console.log(res)
            }
         })
         .catch((res) => {
            console.log(res)
            isButtonDisabled.value = false
            return new Promise(() => {})
         })
   }

   await refresh()
}

// 可显示的列
const visibleColumns = computed(() => {
   return columns.filter((column: { visible: boolean }) => column.visible)
})

// 查询输入变化时
watch(
   () => searchInputValue.value,
   async (newValue) => {
      if (newValue && newValue.length != 0) {
         await employeeSearchApi({ key: newValue })
            .then((res) => {
               if (res.code === 200) {
                  tableData.value = res.data.rows
               }
            })
            .catch((err) => {
               console.log(err)
               return new Promise(() => {})
            })
      } else {
         await refresh()
      }
   }
)

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

// 勾选表格的行时
const handleSelectChange = (rows: Employee[]) => {
   // rows为所有选中的行
   console.log("handleSelectChange", rows)
   selections.value = rows
   isButtonDisabled.value = rows.length === 0
}

const handleRowClick = (row: Employee) => {
   // console.log("handleRowClick", row)
}

const handleAddOrImport = () => {
   console.log("handleAddOrImport")
   addOrEditTitle.value = "新增或导入员工"
   addOrEditData.value = {} as Employee
   isAddOrEditShow.value = true
}

const handleExport = () => {
   if (selections.value.length === 0) {
      return
   }

   // excelData包含表头和行数据
   const { excelHeaders, excelBodys, excelColumnsWidth } = convertForExport(selections.value)
   console.log("handleExport", excelHeaders, excelBodys, excelColumnsWidth)

   // 创建工作表, skipHeader=true, 因为excelData中已经包含表头
   const workSheet = XLSX.utils.json_to_sheet([excelHeaders, ...excelBodys], { skipHeader: true })

   // 设置列宽度
   let colsWidth: XLSX.ColInfo[] = []
   excelColumnsWidth.forEach((w) => {
      let col = { wpx: w }
      colsWidth.push(col)
   })
   workSheet["!cols"] = colsWidth

   // 创建工作簿
   const workBook = XLSX.utils.book_new()
   // 将工作表放入工作簿中
   XLSX.utils.book_append_sheet(workBook, workSheet, "员工信息")
   let excelFileName = "员工信息_" + moment(new Date()).format("YYYY-MM-DD") + ".xlsx"
   // 生成文件并下载
   XLSX.writeFile(workBook, excelFileName)
   tableRef.value.clearSelection()
}

const handleEdit = (index: number, row: Employee | undefined) => {
   console.log("handleEdit", index, row)
   if (row) {
      addOrEditData.value = row
      console.log("handleEdit", addOrEditData.value)
      addOrEditTitle.value = "修改员工"
      isAddOrEditShow.value = true
   }
}

const handleDelete = async (index: number, row: Employee) => {
   console.log("handleDelete", index, row)
   await ElMessageBox.confirm(
      h("p", null, [
         h("span", null, "确认删除姓名 "),
         h("i", { style: "color: red" }, row.name),
         h("span", null, " 的记录?"),
      ]),
      "员工删除确认",
      {
         confirmButtonText: "确认",
         cancelButtonText: "取消",
         type: "warning",
      }
   ).catch(() => {
      ElMessage.info("操作已取消")
      return new Promise(() => {})
   })

   const ids = [row.id]
   deleteAndRefresh(ids)
   ElMessage.success("删除成功")
}

// 批量删除
const handleBatchDelete = () => {
   console.log("handleBatchDelete", selections)
   if (selections.value.length === 0) {
      isButtonDisabled.value = true
      return
   }

   isButtonDisabled.value = true
   let names: string = ""
   let ids: number[] = []
   selections.value.forEach((item, i) => {
      names += item.name
      ids.push(item.id)
      if (i != selections.value.length - 1) {
         names += ", "
      }
   })

   ElMessageBox.confirm(
      h("p", null, [
         h("span", null, "确认删除姓名 "),
         h("i", { style: "color: red" }, names),
         h("span", null, " 共" + selections.value.length + "条记录?"),
      ]),
      "员工删除确认",
      {
         confirmButtonText: "确认",
         cancelButtonText: "取消",
         type: "warning",
      }
   )
      .then(() => {
         deleteAndRefresh(ids)

         selections.value = [] as Employee[]
         isButtonDisabled.value = true
         ElMessage.success("删除成功")
      })
      .catch(() => {
         ElMessage.info("操作已取消")
         isButtonDisabled.value = false
         return new Promise(() => {})
      })
}

// 刷新
const handleRefresh = async () => {
   console.log("handleRefresh")
   await refresh()
}

////// Add组件

// AddVue组件发送的保存事件
const handleSave = async (message: string) => {
   // todo 更新表格数据
   isAddOrEditShow.value = false
   console.log("handleSave", message, addOrEditData.value)

   // 从数据库中读取最新的数据
   await employeeListApi()
      .then((res) => {
         if (res.code === 200) {
            console.log(res)
            // totalRows.value = res.data.total
            tableData.value = res.data.rows
         }
      })
      .catch((res) => {
         console.log(res)
         return new Promise(() => {})
      })

   ElMessage.success(message)
}

// AddVue组件发送的消息事件
const handleCancel = (message: string) => {
   isAddOrEditShow.value = false
   console.log("handleCancel", message, addOrEditData.value)
   ElMessage.info(message)
}
</script>

<template>
   <div>
      <!-- 表头工具 -->
      <el-row class="mb-4">
         <el-col :span="7">
            <el-button type="primary" @click="handleAddOrImport">
               <IEpPlus />
               <span style="vertical-align: middle">新增或导入</span>
            </el-button>
            <el-button type="primary" :disabled="isButtonDisabled">
               <IEpDownload />
               <span style="vertical-align: middle" @click="handleExport">导出</span>
            </el-button>
            <el-button type="primary" :disabled="isButtonDisabled" @click="handleBatchDelete">
               <IEpDelete />
               <span style="vertical-align: middle">批量删除</span>
            </el-button>
         </el-col>
         <el-col :span="7">
            <el-input v-model="searchInputValue" placeholder="" clearable>
               <template #suffix>
                  <el-icon><IEpSearch /></el-icon>
               </template>
               <template #prepend>模糊搜索</template>
            </el-input>
         </el-col>
         <el-col class="ml-4" :span="1">
            <el-button type="primary">
               <IEpRefresh />
               <span style="vertical-align: middle" @click="handleRefresh">刷新</span>
            </el-button>
         </el-col>
         <el-col class="ml-8" :span="1">
            <el-popover placement="right-end" title="列筛选" :width="300" trigger="click">
               <template #reference>
                  <el-button type="primary" style="vertical-align: middle"><IEpPlus />列筛选</el-button>
               </template>
               <el-checkbox-group v-model="checkList" @change="handleCheckedChange" :min="6" :max="17">
                  <el-checkbox
                     v-for="item in columns"
                     :label="item.label"
                     :checked="item.visible"
                     :disabled="item.disable"
                     v-show="item.prop != 'id'"
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
         :default-sort="{ prop: 'id', order: 'descending' }"
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
   <AddOrEdit
      :isShow="isAddOrEditShow"
      :title="addOrEditTitle"
      :formData="addOrEditData"
      @save="handleSave"
      @cancel="handleCancel">
   </AddOrEdit>
</template>

<style lang="scss" scoped></style>
