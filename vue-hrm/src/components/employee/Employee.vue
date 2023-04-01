<script lang="ts" setup>
import AddVue from "@/components/employee/Add.vue"
import { employeeListApi } from "@/utility/api"

const state = reactive<{
   isButtonDisabled: boolean
   zoomValue: boolean
   inputValue: string
   tableRef: {}
   selections: []
   isShow: boolean
   title: string
   rowData: Employee
   tableSettings: TableSettings
   tableColumns: TableColumn[]
   paginationSettings: PaginationSettings
   totalRows: number
   tableData: Employee[]
}>({
   // 表格
   isButtonDisabled: true, // 是否禁用表头的修改和删除按钮
   zoomValue: true, // 开关控制的表格样式
   inputValue: "", // 搜索框的值
   tableRef: {}, // el-table组件对象, 自动关联到el-table组件
   selections: [], // 表格选中的行
   // Add组件属性
   isShow: false,
   title: "查看",
   rowData: {} as Employee,
   // 表格数据
   tableSettings: {
      border: true,
      fit: true,
      highlight_current_row: true,
      height: 500,
      empty_text: "N/A",
      table_layout: "auto",
      row_key: "id",
   },
   tableColumns: [
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
         prop: "birthday",
         label: "生日",
         visible: true,
         sortable: true,
         align: "center",
      },
      {
         prop: "first_work_time",
         label: "参加工作时间",
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
         visible: false,
         sortable: false,
         align: "center",
      },
      {
         prop: "security_card",
         label: "保安证",
         visible: false,
         sortable: false,
         align: "center",
      },
      {
         prop: "current_address",
         label: "现住址",
         visible: false,
         sortable: false,
         align: "left",
      },
      {
         prop: "comments",
         label: "备注",
         visible: false,
         sortable: false,
         align: "left",
      },
   ],
   paginationSettings: {
      layout: "->, total, sizes, prev, pager, next",
      prev_text: "上一页",
      next_text: "下一页",
      page_sizes: [10, 20, 30, 40, 50, 100],
      default_page_size: 10,
      default_current_page: 1,
      hide_on_single_page: false,
   },
   totalRows: 0,
   tableData: [],
})

// 解构
const {
   isButtonDisabled,
   inputValue,
   zoomValue,
   tableRef,
   selections,
   isShow,
   title,
   rowData,
   tableSettings,
   tableColumns,
   totalRows,
   tableData,
} = toRefs(state)

employeeListApi()
   .then((res) => {
      if (res.code === 200) {
         console.log(res)
         totalRows.value = res.data.total
         tableData.value = res.data.rows
      }
   })
   .catch((res) => {
      console.log(res)
   })

////// 表头工具栏
const handleDropdownClick = () => {}
const zoomChange = (val: any) => {
   console.log("zoomChange", val)
}
const tableLayout = computed(() => {
   return zoomValue.value ? "auto" : "fixed"
})

////// 表格
const handleSelectChange = (rows: Employee[]) => {
   console.log("handleSelectChange", rows)
   isButtonDisabled.value = rows.length === 0
}

const handleRowClick = (row: Employee) => {
   // console.log("handleRowClick", row)
}

// 向AddVue组件传值
const handleAdd = () => {
   console.log("新增")
   title.value = "新增"
   rowData.value = {} as Employee
   isShow.value = true
}

const handleEdit = (index: number, row: Employee | undefined) => {
   console.log(index, row)
   if (row) {
      rowData.value = row
      console.log("handleEdit", rowData.value)
      title.value = "修改"
      isShow.value = true
   }
}

const handleDelete = (index: number, row: Employee) => {
   console.log("handleDelete", index, row)
   tableData.value.splice(index, 1)
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

////// 数据
const columns = computed(() => {
   return tableColumns.value.filter((column) => column.visible)
})
</script>

<template>
   <div class="container" style="margin: auto">
      <!-- 表头工具 -->
      <el-row class="my-2">
         <el-col :span="12">
            <el-button type="primary">
               <el-icon style="vertical-align: middle">
                  <Plus />
               </el-icon>
               <span style="vertical-align: middle">新增</span>
            </el-button>
            <el-button type="primary">
               <el-icon style="vertical-align: middle">
                  <Upload />
               </el-icon>
               <span style="vertical-align: middle">导入</span>
            </el-button>
            <el-button type="primary" :disabled="isButtonDisabled">
               <el-icon style="vertical-align: middle">
                  <Download />
               </el-icon>
               <span style="vertical-align: middle">导出</span>
            </el-button>
            <el-button type="primary" :disabled="isButtonDisabled">
               <el-icon style="vertical-align: middle">
                  <Delete />
               </el-icon>
               <span style="vertical-align: middle">删除</span>
            </el-button>
         </el-col>
         <el-col :span="6">
            <el-input
               v-model="inputValue"
               class="w-88"
               placeholder=""
               clearable>
               <template #prepend>模糊搜索</template>
               <template #prefix>
                  <el-icon class="el-input__icon"><search /></el-icon>
               </template>
            </el-input>
         </el-col>
         <el-col :span="6">
            <el-dropdown
               split-button
               :hide-on-click="false"
               class="mr-3"
               type="primary"
               @click="handleDropdownClick">
               列
               <template #dropdown>
                  <el-dropdown-menu>
                     <el-dropdown-item v-for="(column, index) in tableColumns"
                        ><el-checkbox
                           :key="column.prop"
                           :checked="column.visible"
                           >{{ column.label }}</el-checkbox
                        ></el-dropdown-item
                     >
                  </el-dropdown-menu>
               </template>
            </el-dropdown>
            <el-switch
               v-model="zoomValue"
               size="large"
               inline-prompt
               active-text="自动"
               inactive-text="固定"
               @change="zoomChange" />
         </el-col>
      </el-row>
      <!-- 表格 -->
      <el-table
         ref="tableRef"
         style="width: 100%"
         :data="tableData"
         :border="tableSettings.border"
         :fit="tableSettings.fit"
         :height="tableSettings.height"
         :table-layout="tableLayout"
         :empty-text="tableSettings.empty_text"
         :highlight-current-row="tableSettings.highlight_current_row"
         :row-key="tableSettings.row_key"
         :default-sort="{ prop: 'name', order: 'descending' }"
         @selection-change="handleSelectChange"
         @row-click="handleRowClick">
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
               <el-button
                  size="small"
                  type="primary"
                  @click="handleEdit(scope.$index, scope.row)"
                  >修改</el-button
               >
               <el-button
                  size="small"
                  type="danger"
                  @click="handleDelete(scope.$index, scope.row)"
                  >删除</el-button
               >
            </template>
         </el-table-column>
      </el-table>
   </div>
   <AddVue
      :isShow="isShow"
      :title="title"
      :formData="rowData"
      @save="handleSave"
      @cancel="handleCancel">
   </AddVue>
</template>

<style lang="scss" scoped></style>
