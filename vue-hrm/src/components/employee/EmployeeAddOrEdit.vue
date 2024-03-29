<script setup lang="ts">
import { employeeUpdateApi, employeeAddApi } from "@/utils/employee"
import { excelKeyPosition, excelValuePosition, readFile } from "./index"
import { rules } from "./validation"
import { toIsoString } from "@/utils/common"
import { newDefaultEmployee } from "@/utils/common"
import type { UploadFile, UploadFiles, UploadUserFile } from "element-plus"
import * as XLSXS from "xlsx-js-style"

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps<{
   isShow: boolean
   title: string
   formData: Employee
   postData: Post[]
}>()

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "cancel"])

const state = reactive<{
   isEscapeClose: boolean
   isShowClose: boolean
   isClickModalToClose: boolean
   dialogVisible: boolean
   rawFormData: Employee
   isImportDisabled: boolean
   // form参数
   labelWidth: string
   labelPosition: "top" | "right" | "left"
   isInline: boolean
   isHideRequiredAsterisk: boolean
   requireAsteriskPosition: "left" | "right"
   isInlineMessage: boolean
   isStatusIcon: boolean
   isScrollToError: boolean
}>({
   isEscapeClose: false, // 是否按ESC关闭
   isShowClose: false, // 是否显示右上角的关闭
   isClickModalToClose: false, // 是否可以通过点击 modal 关闭 Dialog (对话框以外的任意位置)
   dialogVisible: false, // 是否显示对话框
   rawFormData: {} as Employee,
   isImportDisabled: false, // 是否禁用导入按钮
   // form参数
   labelWidth: "90px", // 标签的长度
   labelPosition: "right", // 表单对齐方式
   isInline: false, // 行内表单模式
   isHideRequiredAsterisk: true, // 是否隐藏必填字段标签旁边的红色星号
   requireAsteriskPosition: "left", // 星号的位置
   isInlineMessage: true, // 是否以行内形式展示校验信息
   isStatusIcon: false, // 是否在输入框中显示校验结果反馈图标
   isScrollToError: true, // 当校验失败时，滚动到第一个错误表单项
})

// 解构
const {
   isEscapeClose,
   isShowClose,
   isClickModalToClose,
   dialogVisible,
   rawFormData,
   isImportDisabled,
   labelWidth,
   labelPosition,
   isInline,
   isHideRequiredAsterisk,
   requireAsteriskPosition,
   isInlineMessage,
   isStatusIcon,
   isScrollToError,
} = toRefs(state)

// watch写法上支持一个或者多个监听源, 这些监听源必须只能是getter/effect函数, ref数据, reactive对象或者数组类型
watch(
   () => props.formData,
   (newValue) => {
      // 当props.formData变化为newValue时, 拷贝一份原始数据,
      // 不要直接使用props.formData, 否则会影响父组件的数据
      rawFormData.value = { ...props.formData }
   }
)

// 当props.isShow变化时会传递给dialogVisible, 而dialogVisible被绑定给了el-dialog, 从而达到外部控制显示隐藏el-dialog的目的
watch(
   () => props.isShow,
   (newValue) => {
      if (props.title === "修改员工") {
         isImportDisabled.value = true
      } else {
         isImportDisabled.value = false
      }

      dialogVisible.value = newValue as boolean
      // 清除校验信息
      formRef.value?.clearValidate()
   }
)

// 下载员工信息模板
const handleDownload = () => {
   let a = document.createElement("a")
   a.href = "../static/template/employee.xlsx"
   a.download = "员工信息模板" + ".xlsx"
   a.style.display = "none"
   document.body.appendChild(a)
   a.click()
   document.body.removeChild(a)
}

// el-form组件对象, 自动关联到el-form组件
// const formRef = ref<FormInstance>()
const formRef = ref()
const isLoading = ref(false)
const handleSave = async () => {
   isLoading.value = true

   // 表格输入规则校验, 通过进入than, 不通过则进入catch
   await formRef.value
      ?.validate()
      .then(async () => {
         console.log("表单校验成功")

         // 转换成Golang一致的时间格式 -- ISO 8601 格式
         rawFormData.value.birthday = toIsoString(new Date(rawFormData.value.birthday))
         rawFormData.value.first_work_time = toIsoString(new Date(rawFormData.value.first_work_time))
         if (rawFormData.value.updated_at) {
            rawFormData.value.updated_at = toIsoString(new Date(rawFormData.value.updated_at))
         }

         // 转换Golang一致的类型
         rawFormData.value.phone = rawFormData.value.phone.toString()
         rawFormData.value.security_card = rawFormData.value.security_card.toString()
         props.postData.forEach((item) => {
            if (item.name === rawFormData.value.post) {
               rawFormData.value.post_id = item.id
            }
         })

         console.log("handleSave", rawFormData.value)
         if (props.title === "修改员工") {
            await employeeUpdateApi(rawFormData.value)
               .then((res) => {
                  if (res.code === 200) {
                     console.log(res)
                     emits("save", res.message)
                  }
               })
               .catch((res) => {
                  console.log(res)
                  ElMessage.error(res.message)
                  return new Promise(() => {})
               })
               .finally(() => {
                  isLoading.value = false
               })
         } else {
            await employeeAddApi(rawFormData.value)
               .then((res) => {
                  if (res.code === 200) {
                     console.log(res)
                     emits("save", res.message)
                  }
               })
               .catch((res) => {
                  console.log(res)
                  ElMessage.error(res.message)
                  return new Promise(() => {})
               })
               .finally(() => {
                  isLoading.value = false
               })
         }
      })
      .catch((err: any) => {
         isLoading.value = false
         console.log("表单校验失败", err)
         return new Promise(() => {})
      })
}

const handleCancel = () => {
   // 清除校验信息
   formRef.value?.clearValidate()
   // 恢复表单数据
   rawFormData.value = { ...props.formData }
   // 向外发送cancel(取消)事件
   emits("cancel", "操作已取消")
}

//// upload组件
const fileList = ref<UploadUserFile[]>([] as UploadUserFile[])
const handleChange = async (file: UploadFile, files: UploadFiles) => {
   isLoading.value = true
   console.log("handleChange", file, files)
   await readFile(file)
      .then((rawData) => {
         // 解析二进制格式数据
         let workbook = XLSXS.read(rawData, { type: "binary", cellDates: true })
         // 获取第一个Sheet
         let worksheet = workbook.Sheets[workbook.SheetNames[0]]

         // 检查表格格式
         if (!isExcelValid(worksheet)) {
            console.log("无法读取表格, 请使用模板导入")
            ElMessage.warning("无法读取表格, 请使用模板导入")
            isLoading.value = false
            return
         }

         // 转成json对象
         // let rows = XLSX.utils.sheet_to_json(worksheet, { defval: "", header: 1, blankrows: false })
         // console.log("sheet_to_json", rows)

         let employee = newDefaultEmployee()

         for (let ep of excelValuePosition.value) {
            const k = ep[0]
            const v = ep[1]
            const cell = worksheet[v]
            if (!cell) {
               ElMessage.warning("无法读取" + file.name + "的" + v + "单元格")
               isLoading.value = false
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
         }

         console.log("读取表格完成", employee)
         // 清除校验信息
         formRef.value?.clearValidate()
         rawFormData.value = employee
      })
      .catch((error) => {
         ElMessage.error(file.name + "读取失败, " + error)
         return new Promise(() => {})
      })
      .finally(() => {
         isLoading.value = false
      })
}

const isExcelValid = (worksheet: XLSXS.WorkSheet): boolean => {
   let isOk = true
   // 这里不要用forEach
   for (let item of excelKeyPosition.value) {
      const k = item[0]
      const v = item[1]
      const cell = worksheet[v]
      console.log(cell.v, k, v)
      if (!cell || cell.v != k) {
         isOk = false
         break
      }
   }

   return isOk
}

const genderOptions = [
   {
      value: "男",
      label: "男",
   },
   {
      value: "女",
      label: "女",
   },
]

const socialSecurityOptions = [
   {
      value: "有",
      label: "有",
   },
   {
      value: "无",
      label: "无",
   },
]

const politicalStatusOptions = [
   {
      value: "群众",
      label: "群众",
   },
   {
      value: "共青团员",
      label: "共青团员",
   },
   {
      value: "党员",
      label: "党员",
   },
   {
      value: "无党派人士",
      label: "无党派人士",
   },
   {
      value: "民主党派成员",
      label: "民主党派成员",
   },
]

const degreeOptions = [
   {
      value: "小学",
      label: "小学",
   },
   {
      value: "初中",
      label: "初中",
   },
   {
      value: "高中",
      label: "高中",
   },
   {
      value: "中专",
      label: "中专",
   },
   {
      value: "大专",
      label: "大专",
   },
   {
      value: "专升本",
      label: "专升本",
   },
   {
      value: "本科",
      label: "本科",
   },
   {
      value: "研究生",
      label: "研究生",
   },
   {
      value: "博士",
      label: "博士",
   },
   {
      value: "博士后",
      label: "博士后",
   },
]

const postOptions = computed(() => {
   let p: any[] = []
   props.postData.forEach((item) => {
      const o = {
         value: item.name,
         label: item.name,
      }

      p.push(o)
   })
   return p
})
</script>

<template>
   <!-- 当点击对话框右上角关闭时, 向外部发送canel事件 -->
   <el-dialog
      width="55%"
      draggable
      v-model="dialogVisible"
      :title="title"
      :close-on-press-escape="isEscapeClose"
      :show-close="isShowClose"
      :close-on-click-modal="isClickModalToClose">
      <el-form
         ref="formRef"
         :model="rawFormData"
         :rules="rules"
         :label-width="labelWidth"
         :hide-required-asterisk="isHideRequiredAsterisk"
         :require-asterisk-position="requireAsteriskPosition"
         :label-position="labelPosition"
         :inline="isInline"
         :inline-message="isInlineMessage"
         :status-icon="isStatusIcon"
         :scroll-to-error="isScrollToError">
         <el-row>
            <el-col :span="8">
               <el-form-item label="姓名" prop="name">
                  <el-input v-model="rawFormData.name"> </el-input>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="性别" prop="gender">
                  <el-select v-model="rawFormData.gender" placeholder="">
                     <el-option
                        v-for="item in genderOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value" />
                  </el-select>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="生日" prop="birthday">
                  <el-date-picker v-model="rawFormData.birthday" type="date" placeholder="" />
               </el-form-item>
            </el-col>
         </el-row>
         <el-row>
            <el-col :span="8">
               <el-form-item label="身高(cm)" prop="height">
                  <el-input v-model.number="rawFormData.height" placeholder=""> </el-input
               ></el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="体重(kg)" prop="weight">
                  <el-input v-model.number="rawFormData.weight" placeholder=""> </el-input
               ></el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="学历" prop="degree">
                  <el-select v-model="rawFormData.degree" placeholder="">
                     <el-option
                        v-for="item in degreeOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value" />
                  </el-select>
               </el-form-item>
            </el-col>
         </el-row>
         <el-row>
            <el-col :span="8">
               <el-form-item label="身份证号" prop="identifier">
                  <el-input v-model="rawFormData.identifier" placeholder=""> </el-input>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="手机号码" prop="phone">
                  <el-input v-model.number="rawFormData.phone" placeholder=""> </el-input>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="政治面貌" prop="political_status">
                  <el-select v-model="rawFormData.political_status" placeholder="">
                     <el-option
                        v-for="item in politicalStatusOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value" />
                  </el-select>
               </el-form-item>
            </el-col>
         </el-row>
         <el-row>
            <el-col :span="8">
               <el-form-item label="社保" prop="social_security">
                  <el-select v-model="rawFormData.social_security" placeholder="">
                     <el-option
                        v-for="item in socialSecurityOptions"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value" />
                  </el-select>
               </el-form-item>
            </el-col>
            <el-col :span="16">
               <el-form-item label="现住址" prop="current_address">
                  <el-input v-model="rawFormData.current_address"> </el-input>
               </el-form-item>
            </el-col>
         </el-row>
         <el-row>
            <el-col :span="8">
               <el-form-item label="首次工作" prop="first_work_time">
                  <el-date-picker v-model="rawFormData.first_work_time" type="date" placeholder="" />
               </el-form-item>
            </el-col>
            <el-col :span="16">
               <el-form-item label="原单位" prop="former_employer">
                  <el-input v-model="rawFormData.former_employer"> </el-input>
               </el-form-item>
            </el-col>
         </el-row>
         <el-row>
            <el-col :span="8">
               <el-form-item label="岗位" prop="post">
                  <el-select v-model="rawFormData.post" placeholder="">
                     <el-option v-for="item in postOptions" :key="item.value" :label="item.label" :value="item.value" />
                  </el-select>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="工资(¥)" prop="salary">
                  <el-input v-model.number="rawFormData.salary" placeholder=""> </el-input>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="保安证" prop="security_card">
                  <el-input v-model.number="rawFormData.security_card" placeholder=""> </el-input>
               </el-form-item>
            </el-col>
         </el-row>
         <el-row>
            <el-col :span="24">
               <el-form-item label="备注" prop="comments">
                  <el-input
                     v-model="rawFormData.comments"
                     maxlength="8888"
                     :rows="3"
                     placeholder="填写了解的情况"
                     show-word-limit
                     type="textarea" />
               </el-form-item>
            </el-col>
         </el-row>
      </el-form>
      <el-upload
         v-if="!isImportDisabled"
         v-model.file-list="fileList"
         drag
         accept=".xls, .xlsx"
         :multiple="false"
         :show-file-list="false"
         :auto-upload="false"
         :disabled="isImportDisabled"
         :on-change="handleChange">
         <!-- <el-icon :size="26" class="el-icon--upload"><IEpUploadFilled /></el-icon> -->
         <h1 class="el-upload__text">支持.xls, .xlsx格式, 将表格(员工信息模板)拖拽到此处, 或者<em>单击上传</em></h1>
         <!-- <template #tip>
            <div class="el-upload__tip">支持.xls, .xlsx格式(请使用模板导入)</div>
         </template> -->
      </el-upload>
      <template #footer>
         <span class="dialog-footer">
            <el-button type="primary" :loading="isLoading" @click="handleDownload">下载模板</el-button>
            <el-button type="primary" :loading="isLoading" @click="handleSave">保存</el-button>
            <el-button type="danger" :loading="isLoading" @click="handleCancel">取消</el-button>
         </span>
      </template>
   </el-dialog>
</template>

<style lang="scss" scoped>
// ::v-deep .el-upload {
//    width: 100%;
//    height: 56px;
//    text-align: center;
//    line-height: 56px;

//    .h1 {
//       text-align: center;
//       margin: 0 auto;
//    }
// }

// ::v-deep .el-upload .el-upload-dragger {
//    width: 100%;
//    height: 30%;
// }
</style>
