<script lang="tsx" setup>
import type { UploadFile, UploadUserFile } from "element-plus"
import { excelPosition } from "./index"
import { newDefaultEmployee } from "@/utils/common"
import { employeeAddApi } from "@/utils/employee"
import moment from "moment"
import * as XLSX from "xlsx"
import { useData } from "./index"

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps<{
   isShow: boolean
   title: string
   fileList: UploadUserFile[]
}>()

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "cancel"])

const state = reactive<{
   isEscapeClose: boolean
   isShowClose: boolean
   isClickModalToClose: boolean
   isDialogVisible: boolean
   isMultiple: boolean
   isAutoUpload: boolean
   acceptFileFormat: string
   importFiles: UploadUserFile[]
}>({
   isEscapeClose: false, // 是否按ESC关闭
   isShowClose: true, // 是否显示右上角的关闭
   isClickModalToClose: false, // 是否可以通过点击 modal 关闭 Dialog (对话框以外的任意位置)
   isDialogVisible: false, // 是否显示对话框
   isMultiple: true, // 是否支持多选文件
   isAutoUpload: false, // 是否自动上传
   acceptFileFormat: ".xls, .xlsx", // 接受上传的文件类型
   importFiles: [],
})

// 解构
const {
   isEscapeClose,
   isShowClose,
   isClickModalToClose,
   isDialogVisible,
   isMultiple,
   isAutoUpload,
   acceptFileFormat,
   importFiles,
} = toRefs(state)

const { isAddOrEditShow, addOrEditTitle, addOrEditData } = useData()

// 当props.isShow变化时会传递给dialogVisible, 而dialogVisible被绑定给了el-dialog, 从而达到外部控制显示隐藏el-dialog的目的
watch(
   () => props.isShow,
   (newValue) => {
      isDialogVisible.value = newValue as boolean
   }
)

watch(
   () => props.fileList,
   (newValue) => {
      importFiles.value = newValue
   }
)

function readFile(file: UploadFile) {
   return new Promise(function (resolve, reject) {
      const reader = new FileReader()
      reader.onload = (e) => {
         resolve(e.target?.result)
      }
      reader.readAsBinaryString(file.raw as Blob)
   })
}

const handleChange = async (uploadFile: UploadFile) => {
   await readFile(uploadFile)
      .then(async (rawData) => {
         // 解析二进制格式数据
         let workbook = XLSX.read(rawData, { type: "binary", cellDates: true })
         // 获取第一个Sheet
         let worksheet = workbook.Sheets[workbook.SheetNames[0]]

         // 转成json对象
         // let json = XLSX.utils.sheet_to_json(worksheet, { defval: "", header: 1, blankrows: false })

         let employee = newDefaultEmployee()

         excelPosition.forEach((v, k) => {
            const cell = worksheet[v]
            if (!cell) {
               ElMessage.warning("无法读取" + uploadFile.name + "的" + v + "单元格")
               return
            }
            const val = cell.v
            switch (k) {
               case "name":
                  employee.name = val
                  break
               case "gender":
                  employee.gender = val
                  break
               case "birthday":
                  employee.birthday = val
                  break
               case "height":
                  employee.height = val
                  break
               case "weight":
                  employee.weight = val
                  break
               case "degree":
                  employee.degree = val
                  break
               case "identifier":
                  employee.identifier = val
                  break
               case "phone":
                  employee.phone = val
                  break
               case "political_status":
                  employee.political_status = val
                  break
               case "social_security":
                  employee.social_security = val
                  break
               case "current_address":
                  employee.current_address = val
                  break
               case "first_work_time":
                  employee.first_work_time = val
                  break
               case "former_employer":
                  employee.former_employer = val
                  break
               case "post":
                  employee.post = val
                  break
               case "salary":
                  employee.salary = val
                  break
               case "security_card":
                  employee.security_card = val
                  break
               case "comments":
                  employee.comments = val
                  break
            }
         })

         console.log(employee)

         addOrEditData.value = employee
         addOrEditTitle.value = "从表格中导入"
         isAddOrEditShow.value = true
         // await employeeAddApi(employee)
         //    .then((res) => {
         //       if (res.code === 200) {
         //          ElMessage.success(uploadFile.name + "导入成功, 姓名:" + employee.name + ", 性别:" + employee.gender)
         //       }
         //    })
         //    .catch((res) => {
         //       console.log(res)
         //       ElMessage.warning(uploadFile.name + "导入失败, 姓名:" + employee.name + ", 性别:" + employee.gender)
         //       return new Promise(() => {})
         //    })
      })
      .catch((error) => {
         ElMessage.error(uploadFile.name + "读取失败, " + error)
         return new Promise(() => {})
      })
}

// AddVue组件发送的保存事件
const handleSave = async (message: string) => {
   // todo 更新表格数据
   isAddOrEditShow.value = false
   console.log("handleSave", message, addOrEditData.value)
}

// AddVue组件发送的消息事件
const handleCancel = (message: string) => {
   isAddOrEditShow.value = false
   console.log("handleCancel", message, addOrEditData.value)
   ElMessage.info(message)
}
</script>

<template>
   <!-- 数据上传的功能 -->
   <el-dialog
      width="30%"
      draggable
      v-model="isDialogVisible"
      :title="title"
      :close-on-press-escape="isEscapeClose"
      :show-close="isShowClose"
      :close-on-click-modal="isClickModalToClose">
      <el-upload
         v-model.file-list="importFiles"
         :multiple="isMultiple"
         :auto-upload="isAutoUpload"
         :accept="acceptFileFormat"
         :on-change="handleChange">
         <el-button type="primary">导入</el-button>
         <template #tip>
            <div class="el-upload__tip">选择需要导入的员工表格, 支持xls和xlsx格式</div>
         </template>
      </el-upload>
      <AddOrEdit
         :isShow="isAddOrEditShow"
         :title="addOrEditTitle"
         :formData="addOrEditData"
         @save="handleSave"
         @cancel="handleCancel">
      </AddOrEdit>
   </el-dialog>
</template>

<style lang="scss" scoped></style>
