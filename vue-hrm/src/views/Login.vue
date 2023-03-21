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
import { useRouter } from "vue-router"
import { loginApi, getUserInfo } from "@/utility/api"
import { useMainStore } from "@/store/index"
import { storeToRefs } from "pinia"

// 获取当前项目的路由对象
let router = useRouter()

// 使用pinia的main存储
const main = useMainStore()
// 对pinia中main存储的对象进行解构, 使之成为响应式对象
const { data, roles, menus } = storeToRefs(main)

// el-form组件对象, 自动关联到el-form组件
const ruleFormRef = ref()

const state = reactive({
   ruleForm: {
      username: "admin",
      password: "123456",
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
                  // 先存储token
                  localStorage.setItem(
                     "token",
                     res.data.token_header + " " + res.data.token
                  )
                  // 再获取用户信息
                  getUserInfo()
                     .then((res) => {
                        if (res.code === 200) {
                           console.log(res)
                           data.value = res.data
                           router.push("/home")
                        }
                     })
                     .catch((res) => {
                        console.log(res)
                     })
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
