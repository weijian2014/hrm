<script setup lang="ts">
import { employeeUpdateApi } from "@/utils/employee"
import Validator from "./index"
import type { FormRules } from "element-plus"

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps<{
   isShow: boolean
   title: string
   formData: Employee
}>()

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "cancel"])

const state = reactive<{
   isEscapeClose: boolean
   isShowClose: boolean
   isClickModalToClose: boolean
   dialogVisible: boolean
   rawFormData: Employee
}>({
   isEscapeClose: false, // 是否按ESC关闭
   isShowClose: false, // 是否显示右上角的关闭
   isClickModalToClose: false, // 是否可以通过点击 modal 关闭 Dialog (对话框以外的任意位置)
   dialogVisible: false, // 是否显示对话框
   rawFormData: {} as Employee,
})

// 标签的长度
const labelWidth = ref("110px")
// 表单对齐方式, 'left' | 'right' | 'top'
const labelPosition = ref("right")
// 行内表单模式
const isInline = ref(false)
// 是否隐藏必填字段标签旁边的红色星号
const isHideRequiredAsterisk = ref(false)
// 星号的位置, 'left' | 'right'
const requireAsteriskPosition = ref("left")
// 是否以行内形式展示校验信息
const isInlineMessage = ref(true)
// 是否在输入框中显示校验结果反馈图标
const isStatusIcon = ref(false)
// 当校验失败时，滚动到第一个错误表单项
const isScrollToError = ref(true)

// 解构
const { isEscapeClose, isShowClose, isClickModalToClose, dialogVisible, rawFormData } = toRefs(state)

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
      dialogVisible.value = newValue as boolean
      formRef.value?.clearValidate()
   }
)

// el-form组件对象, 自动关联到el-form组件
// const formRef = ref<FormInstance>()
const formRef = ref()
const isLoading = ref(false)
const handleSave = async () => {
   console.log("handleSave", rawFormData.value)
   isLoading.value = true

   // 表格输入规则校验, 通过进入than, 不通过则进入catch
   await formRef.value
      ?.validate()
      .then(async () => {
         console.log("表单校验成功")

         await employeeUpdateApi(rawFormData.value)
            .then((res) => {
               if (res.code === 200) {
                  console.log(res)
                  // 向外发送save(保存)事件
                  emits("save", "保存成功")
                  isLoading.value = false
               }
            })
            .catch((res) => {
               console.log(res)
               isLoading.value = false
               return new Promise(() => {})
            })
      })
      .catch((err: any) => {
         isLoading.value = false
         console.log("表单校验失败", err)
         return new Promise(() => {})
      })
}

const handleCancel = () => {
   // 向外发送cancel(取消)事件
   emits("cancel", "已取消")
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
      value: "无党派人士",
      label: "无党派人士",
   },
   {
      value: "民主党派成员",
      label: "民主党派成员",
   },
   {
      value: "共青团员",
      label: "共青团员",
   },
   {
      value: "党员",
      label: "党员",
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

const postOptions = [
   {
      value: "前台",
      label: "前台",
   },
   {
      value: "巡逻",
      label: "巡逻",
   },
   {
      value: "消防",
      label: "消防",
   },
   {
      value: "教练",
      label: "教练",
   },
   {
      value: "保安",
      label: "保安",
   },
   {
      value: "银保",
      label: "银保",
   },
]

const isHeightValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入身高, 单位厘米")
      return
   }

   let h = Number(value)
   if (h < 120 || h > 230) {
      callback("身高在120~230厘米之间")
      return
   }

   callback()
}

const isWeightValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入体重, 单位公斤")
      return
   }

   let w = Number(value)
   if (w < 40 || w > 130) {
      callback("体重在40~130公斤之间")
      return
   }

   callback()
}

const isIdentifierValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入身份证号")
      return
   }

   const identifierString = value.toString()
   if (identifierString.length != 11) {
      callback("身份证号必须是18位")
      return
   }

   if (!Validator.isIdentifierValid(value?.toString())) {
      callback("身份证号输入有误")
      return
   }

   callback()
}

const isPhoneValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入手机号")
      return
   }

   const phoneString = value.toString()
   if (phoneString.length != 11) {
      callback("手机号是11位数字")
      return
   }

   if (!Validator.isPhoneValid(phoneString)) {
      callback("手机号输入有误")
      return
   }

   callback()
}

const isSecurityCardValid = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   const securityCardString = value?.toString()
   if (securityCardString && (securityCardString.length < 6 || securityCardString.length > 10)) {
      callback("保安证是6~10位数字")
      return
   }

   callback()
}

// 校验规则
const rules = reactive<FormRules>({
   name: [
      {
         required: true,
         message: "请输入姓名",
         trigger: "blur", // 失去焦点时
      },
   ],
   gender: [
      {
         required: true,
         message: "请选择性别",
         trigger: "blur",
      },
   ],
   birthday: [
      {
         required: true,
         message: "请选择生日",
         trigger: "blur",
      },
   ],
   height: [
      {
         required: true,
         validator: isHeightValid,
         trigger: "blur",
      },
   ],
   weight: [
      {
         required: true,
         validator: isWeightValid,
         trigger: "blur",
      },
   ],
   political_status: [
      {
         required: true,
         message: "请选择政治面貌",
         trigger: "blur",
      },
   ],
   degree: [
      {
         required: true,
         message: "请选择学历",
         trigger: "blur",
      },
   ],
   social_security: [
      {
         required: true,
         message: "请选择社保",
         trigger: "blur",
      },
   ],
   identifier: [
      {
         required: true,
         validator: isIdentifierValid,
         trigger: "blur",
      },
   ],
   phone: [
      { type: "number", message: "手机号是数字", trigger: "change" },
      {
         required: true,
         validator: isPhoneValid,
         trigger: "blur",
      },
   ],
   first_work_time: [
      {
         required: true,
         message: "请选择参加工作时间",
         trigger: "blur",
      },
   ],
   post: [
      {
         required: true,
         message: "请输入岗位",
         trigger: "blur",
      },
   ],
   salary: [
      { required: true, message: "请输入工资", trigger: "blur" },
      { type: "number", message: "工资是数字, 单位人民币元", trigger: "change" },
   ],
   security_card: [
      { type: "number", message: "保安证是数字", trigger: "change" },
      {
         required: false,
         validator: isSecurityCardValid,
         trigger: "blur",
      },
   ],
})
</script>

<template>
   <!-- 当点击对话框右上角关闭时, 向外部发送canel事件 -->
   <el-dialog
      width="50%"
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
         <!-- 第1行 -->
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
         <!-- 第2行 -->
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
         </el-row>
         <!-- 第3行 -->
         <el-row>
            <el-col :span="24">
               <el-form-item label="现住址" prop="current_address">
                  <el-input v-model="rawFormData.current_address"> </el-input>
               </el-form-item>
            </el-col>
         </el-row>
         <!-- 第4行 -->
         <el-row>
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
         </el-row>
         <!-- 第5行 -->
         <el-row>
            <el-col :span="8">
               <el-form-item label="身份证" prop="identifier">
                  <el-input v-model="rawFormData.identifier" placeholder=""> </el-input>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="电话" prop="phone">
                  <el-input v-model.number="rawFormData.phone" placeholder=""> </el-input>
               </el-form-item>
            </el-col>
         </el-row>
         <!-- 第6行 -->
         <el-row>
            <el-col :span="8">
               <el-form-item label="原单位" prop="former_employer">
                  <el-input v-model="rawFormData.former_employer"> </el-input>
               </el-form-item>
            </el-col>
            <el-col :span="8">
               <el-form-item label="参加工作时间" prop="first_work_time">
                  <el-date-picker v-model="rawFormData.first_work_time" type="date" placeholder="" />
               </el-form-item>
            </el-col>
         </el-row>
         <!-- 第7行 -->
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
                     maxlength="800"
                     placeholder="填写了解的情况"
                     show-word-limit
                     type="textarea" />
               </el-form-item>
            </el-col>
         </el-row>
      </el-form>
      <template #footer>
         <span class="dialog-footer">
            <el-button type="primary" :loading="isLoading" @click="handleSave">保存</el-button>
            <el-button type="danger" :loading="isLoading" @click="handleCancel">取消</el-button>
         </span>
      </template>
   </el-dialog>
</template>

<style lang="scss" scoped></style>
