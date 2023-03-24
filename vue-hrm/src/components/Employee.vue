<template>
   <div class="container" style="margin: auto">
      <el-row class="my-2">
         <!-- 表头工具 -->
         <el-col :span="12">
            <el-button :icon="Plus" type="primary" @click="handleAdd"
               >新增</el-button
            >
            <el-button :icon="Upload" type="primary">导入</el-button>
            <el-button :icon="Download" :disabled="isDisabled" type="primary"
               >导出</el-button
            >
            <el-button :icon="Delete" :disabled="isDisabled" type="danger"
               >删除</el-button
            >
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
         :border="tableSettings?.table_settings.border"
         :fit="tableSettings?.table_settings.fit"
         :height="tableSettings?.table_settings.height"
         :table-layout="tableLayout"
         :empty-text="tableSettings?.table_settings.empty_text"
         :highlight-current-row="
            tableSettings?.table_settings.highlight_current_row
         "
         :row-key="tableSettings?.table_settings.row_key"
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
            :min-width="column.minWidth"
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
      :id="id"
      :title="title"
      :employee="employee"
      @save="handleSave"
      @cancel="handleCancel">
   </AddVue>
</template>

<script lang="ts" setup>
import { ref, reactive, computed } from "vue"
import Employee from "@/class/Employee"
import AddVue from "@/components/Add.vue"
import { Upload, Download, Delete, Plus, Search } from "@element-plus/icons-vue"

////// 表头工具栏
const inputValue = ref("") // 搜索框
const zoomValue = ref(true)
const handleDropdownClick = () => {}
const zoomChange = (val: any) => {
   console.log("zoomChange", val)
}
const tableLayout = computed(() => {
   return zoomValue.value ? "auto" : "fixed"
})

////// 表格
// el-table组件对象, 自动关联到el-table组件
const tableRef = ref()
const selections = ref<Employee[]>()
const isDisabled = ref(true)
const handleSelectChange = (rows: Employee[]) => {
   console.log("handleSelectChange", rows)
   selections.value = rows
   isDisabled.value = selections.value.length == 0 ? true : false
}

const handleRowClick = (row: Employee) => {
   // console.log("handleRowClick", row)
}

// 向AddVue组件传值 -- 只传isShow
const handleAdd = () => {
   console.log("新增")
   title.value = "新增"
   employee.value = new Employee()
   isShow.value = true
}

const handleEdit = (index: number, row: Employee | undefined) => {
   console.log(index, row)
   id.value = index
   title.value = "修改"
   employee.value = row
   isShow.value = true
}

const handleDelete = (index: number, row: Employee) => {
   console.log("handleDelete", index, row)
   tableData.slice(index, 1)
}

////// Add组件
const isShow = ref(false)
const id = ref(0)
const title = ref("查看")
const employee = ref<Employee>()

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
   ElMessage.info(message)
}

////// 数据
const columns = computed(() => {
   return tableColumns.filter((column) => column.visible)
})

const tableSettings = {
   table_settings: {
      border: true,
      fit: true,
      highlight_current_row: true,
      height: 500,
      empty_text: "N/A",
      table_layout: "auto",
      row_key: "id",
   },
   table_columns: [
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
   pagination_settings: {
      layout: "->, total, sizes, prev, pager, next",
      prev_text: "上一页",
      next_text: "下一页",
      page_sizes: [10, 20, 30, 40, 50, 100],
      default_page_size: 10,
      default_current_page: 1,
      hide_on_single_page: false,
   },
}

const tableColumns = [
   {
      prop: "id",
      label: "序号",
      visible: false,
      sortable: false,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "name",
      label: "姓名",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "gender",
      label: "性别",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "birthday",
      label: "生日",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 30,
   },
   {
      prop: "first_work_time",
      label: "参加工作时间",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 30,
   },
   {
      prop: "salary",
      label: "工资",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "post",
      label: "岗位",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "social_security",
      label: "社保",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "phone",
      label: "电话",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 30,
   },
   {
      prop: "former_employer",
      label: "原单位",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "height",
      label: "身高",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "weight",
      label: "体重",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "degree",
      label: "学历",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "political_status",
      label: "政治面貌",
      visible: true,
      sortable: true,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "identifier",
      label: "身份证",
      visible: false,
      sortable: false,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "security_card",
      label: "保安证",
      visible: false,
      sortable: false,
      align: "center",
      minWidth: 18,
   },
   {
      prop: "current_address",
      label: "现住址",
      visible: false,
      sortable: false,
      align: "left",
      minWidth: 18,
   },
   {
      prop: "comments",
      label: "备注",
      visible: false,
      sortable: false,
      align: "left",
      minWidth: 18,
   },
]

const tableData = reactive([
   {
      id: 1,
      name: "李四",
      gender: "女",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 2,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 3,
      name: "张三",
      gender: "女",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 4,
      name: "张三",
      gender: "女",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 5,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 6,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 7,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 8,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 9,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 10,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 11,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 12,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 13,
      name: "张三",
      gender: "女",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 14,
      name: "张三",
      gender: "女",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 15,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 16,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 17,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 18,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 19,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 20,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 21,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 22,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 23,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 24,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
   {
      id: 25,
      name: "张三",
      gender: "男",
      birthday: "1996-11-26",
      first_work_time: "2015-07-01",
      salary: 4500,
      post: "保安",
      social_security: "无",
      phone: "13888888888",
      former_employer: "新盾",
      height: 182,
      weight: 67,
      degree: "大专",
      political_status: "党员",
      identifier: "412345678908172844",
      security_card: "123456789",
      current_address: "广西省桂林市七星区五象街道18号",
      comments: "有此情况需要了解",
   },
])
</script>
