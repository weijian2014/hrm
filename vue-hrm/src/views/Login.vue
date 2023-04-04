<script setup lang="ts">
import { useRouter, useRoute } from "vue-router"
import { loginApi, getUserInfo } from "@/utility/api"
import { useUserStore } from "@/store/user"
import { storeToRefs } from "pinia"
import { User, Key } from "@element-plus/icons-vue"
import type { FormInstance, FormRules } from "element-plus"

// 获取当前项目的路由对象
const router = useRouter()
const route = useRoute()

// 使用pinia的main存储
const store = useUserStore()
// 对pinia中main存储的对象进行解构, 使之成为响应式对象
// const { token, saveToken } = storeToRefs(store)

// el-form组件对象, 自动关联到el-form组件
// const formRef = ref<FormInstance>()
const formRef = ref()

const state = reactive({
   form: {
      username: "admin",
      password: "123456",
   },
   isLoading: false,
})

// 使用toRefs解构后得到的对象也是响应式的
let { form, isLoading } = toRefs(state)

const validatePassword = (rule: unknown, value: string | undefined, callback: (msg?: string) => void) => {
   if (!value) {
      callback("请输入密码")
   } else {
      callback()
   }
}

// 校验规则
const rules = reactive<FormRules>({
   username: [
      {
         required: true,
         message: "请输入用户名",
         trigger: "blur", // 失去焦点时
      },
      { min: 3, max: 20, message: "用户名长度在3~20之间", trigger: "blur" },
   ],
   password: [
      {
         required: true,
         validator: validatePassword,
         trigger: "blur",
      },
      { min: 6, max: 20, message: "密码长度在3~20之间", trigger: "blur" },
   ],
})

const loginFn = () => {
   isLoading.value = true
   // 表格输入规则校验, 通过进入than, 不通过则进入catch
   formRef.value
      ?.validate()
      .then(() => {
         console.log("表单校验成功")
         loginApi({
            username: form.value.username,
            password: form.value.password,
         })
            .then((res) => {
               if (res.code === 200) {
                  console.log("登录成功", res)
                  // 先存储token
                  store.saveToken(JSON.stringify(res.data))
                  // 再获取用户信息
                  getUserInfo()
                     .then((res) => {
                        if (res.code === 200) {
                           console.log("获取用户信息成功", res)
                           store.data = res.data
                           ElMessage.success("登录成功")
                           // 有重定向地址(路由拦截设置的)时跳转到该地址, 否则跳转到首页
                           router.push((route.query.redirect as string) || "/")
                           return
                        }
                     })
                     .catch((res) => {
                        console.log("获取用户信息失败", res)
                        ElMessage.error(res.message)
                     })
               } else {
                  // 登录失败
                  console.log("登录失败", res)
                  ElMessage.error(res.message)
               }
               isLoading.value = false
            })
            .catch((res) => {
               isLoading.value = false
               console.log("登录错误", res)
               ElMessage.error(res.data.message)
            })
      })
      .catch((err: any) => {
         isLoading.value = false
         console.log("表单校验失败", err)
      })
}
</script>

<template>
   <div class="login">
      <el-form ref="formRef" label-position="left" size="large" :model="form" :rules="rules" class="demo-form">
         <h1>欢迎使用HRM管理系统</h1>
         <el-form-item prop="username" label="账号">
            <el-input v-model="form.username" type="text" autocomplete="off" :prefix-icon="User" />
         </el-form-item>
         <el-form-item prop="password" label="密码">
            <el-input v-model="form.password" type="password" autocomplete="off" :prefix-icon="Key" />
         </el-form-item>
         <el-form-item>
            <el-button type="primary" :loading="isLoading" @click="loginFn()">登录</el-button>
         </el-form-item>
      </el-form>
   </div>
</template>
<style lang="scss" scoped>
.login {
   background-color: #ccc;
   height: 100vh;
   display: flex;
   justify-content: center;
   align-items: center;
   text-align: center;

   .el-form {
      width: 300px;
      background-color: #fff;
      padding: 30px;
      border-radius: 15px;
      .el-form-item {
         margin-top: 20px;
      }

      .el-button {
         width: 100%;
         margin-top: 5px;
      }
   }
}
</style>
