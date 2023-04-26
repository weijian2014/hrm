<script setup lang="ts">
import { changePasswordApi } from "@/utils/user"
import { FormRules } from "element-plus"
import { useUserStore } from "@/store/user"
import { storeToRefs } from "pinia"

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps<{
   isShow: boolean
   title: string
}>()

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "cancel"])

const state = reactive<{
   isEscapeClose: boolean
   isShowClose: boolean
   isClickModalToClose: boolean
   dialogVisible: boolean
   rawFormData: ChangePasswordRequest
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
   rawFormData: {} as ChangePasswordRequest,
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
   labelWidth,
   labelPosition,
   isInline,
   isHideRequiredAsterisk,
   requireAsteriskPosition,
   isInlineMessage,
   isStatusIcon,
   isScrollToError,
} = toRefs(state)

const { tokenInfo } = storeToRefs(useUserStore())

// 当props.isShow变化时会传递给dialogVisible, 而dialogVisible被绑定给了el-dialog, 从而达到外部控制显示隐藏el-dialog的目的
watch(
   () => props.isShow,
   (newValue) => {
      dialogVisible.value = newValue as boolean
      if (dialogVisible.value) {
         rawFormData.value.username = tokenInfo.value.user_name
      }
      // 清除校验信息
      formRef.value?.clearValidate()
   }
)

const validatePassword = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入新密码")
      return
   }

   const valueString = value?.toString()
   if (valueString === rawFormData.value.old_password) {
      callback("新密码不能与旧密码一致")
      return
   } else if (
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[._~!@#$^&*])[A-Za-z0-9._~!@#$^&*]{8,16}$/g.test(valueString) == false
   ) {
      callback("请输入包含英文字母大小写、数字和特殊符号的 8-16 位组合")
      return
   }

   callback()
}

const validatePasswordRepeat = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入确认密码")
      return
   }

   const valueString = value?.toString()
   if (valueString !== rawFormData.value.new_password) {
      callback("确认密码与新密码不一致")
      return
   }

   callback()
}

// 校验规则
const rules = reactive<FormRules>({
   old_password: [{ required: true, min: 8, max: 20, message: "原密码长度在8~20之间", trigger: "blur" }],
   new_password: [
      {
         required: true,
         validator: validatePassword,
         trigger: "blur",
      },
      { min: 8, max: 20, message: "新密码长度在8~20之间", trigger: "blur" },
   ],
   new_password_repeat: [
      {
         required: true,
         validator: validatePasswordRepeat,
         trigger: "blur",
      },
      { min: 8, max: 20, message: "确认密码长度在8~20之间", trigger: "blur" },
   ],
})

const formRef = ref()
const isLoading = ref(false)
const handleSave = async () => {
   console.log("handleSave", rawFormData.value)
   isLoading.value = true
   await formRef.value
      ?.validate()
      .then(async () => {
         await changePasswordApi(rawFormData.value)
            .then((res) => {
               if (res.code === 200) {
                  console.log(res)
                  emits("save", res.data)
                  isLoading.value = false
               }
            })
            .catch((res) => {
               console.log(res)
               isLoading.value = false
               ElMessage.error(props.title + "失败, " + res)
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
   // 清除校验信息
   formRef.value?.clearValidate()
   // 恢复表单数据
   // 向外发送cancel(取消)事件
   emits("cancel", "已取消")
}
</script>

<template>
   <div class="container">
      <!-- 当点击对话框右上角关闭时, 向外部发送canel事件 -->
      <el-dialog
         width="20%"
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
            <el-form-item label="原密码" prop="old_password">
               <el-input v-model.trim="rawFormData.old_password" type="password"> </el-input>
            </el-form-item>
            <el-form-item label="新密码" prop="new_password">
               <el-input v-model.trim="rawFormData.new_password" type="password"> </el-input>
            </el-form-item>
            <el-form-item label="新密码确认" prop="new_password_repeat">
               <el-input v-model.trim="rawFormData.new_password_repeat" type="password"> </el-input>
            </el-form-item>
         </el-form>
         <template #footer>
            <span class="dialog-footer">
               <el-button type="primary" @click="handleSave" :loading="isLoading">保存</el-button>
               <el-button type="danger" @click="handleCancel" :loading="isLoading">取消</el-button>
            </span>
         </template>
      </el-dialog>
   </div>
</template>

<style lang="scss" scoped></style>
