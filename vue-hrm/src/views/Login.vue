<template>
   <el-form
      ref="ruleFormRef"
      :model="ruleForm"
      :rules="rules"
      class="demo-ruleForm">
      <el-form-item prop="username">
         <el-input v-model="ruleForm.username" type="text" autocomplete="off" />
      </el-form-item>
      <el-form-item prop="password">
         <el-input
            v-model="ruleForm.password"
            type="password"
            autocomplete="off" />
      </el-form-item>
      <el-form-item>
         <el-button type="primary" @click="loginFn()">登录</el-button>
      </el-form-item>
   </el-form>
</template>

<script setup lang="ts">
import { ref, reactive, toRefs } from "vue"
import type { FormInstance, FormRules } from "element-plus"
import { loginApi, getUserInfo } from "@/utility/api"

// el-form组件对象, 自动关联到el-form组件
const ruleFormRef = ref()

const state = reactive({
   ruleForm: {
      username: "",
      password: "",
   },
})

// 使用toRefs解构后得到的对象也是响应式的
let { ruleForm } = toRefs(state)

const validatePassword = (
   rule: unknown,
   value: string | undefined,
   callback: (msg?: string) => void
) => {
   if (!value) {
      callback("请输入密码")
   } else {
      callback()
   }
}

// 校验规则
const rules = reactive({
   username: [
      {
         required: true,
         message: "请输入用户名",
         trigger: "blur",
      },
      { min: 3, max: 20, message: "用户名长度在3~20之间", trigger: "blur" },
   ],
   password: [
      {
         validator: validatePassword,
         trigger: "blur",
      },
   ],
})

const loginFn = () => {
   // 表格输入规则校验, 通过进入than, 不通过则进入catch
   ruleFormRef.value
      ?.validate()
      .then(() => {
         console.log("输入规则校验通过")
         loginApi({
            username: ruleForm.value.username,
            password: ruleForm.value.password,
         })
            .then((res) => {
               if (res.code === 200) {
                  console.log(res)
                  // 先存储token, 可以存到cookie中, https://github.com/js-cookie/js-cookie
                  localStorage.setItem(
                     "token",
                     res.data.token_header + " " + res.data.token
                  )
                  getUserInfo().then((res) => {})
               }
            })
            .catch((res) => {
               console.log(res)
            })
      })
      .catch(() => {
         console.log("输入规则校验不通过")
      })
}
</script>
<style lang="scss" scoped></style>
