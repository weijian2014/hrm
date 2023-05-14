<script setup lang="ts">
import { useRouter } from "vue-router"
import { useUserStore } from "@/store/user"
import { storeToRefs } from "pinia"
import { isCollapse } from "./index"
import { logoutApi } from "@/utils/user"
import ChangePassword from "./ChangePassword.vue"

const router = useRouter()
const store = useUserStore()
const { tokenInfo } = storeToRefs(store)

const state = reactive<{
   // 弹窗
   isShow: boolean
   title: string
}>({
   isShow: false,
   title: "修改 " + tokenInfo.value.user_name + " 的密码",
})

const { isShow, title } = toRefs(state)

const changeFn = async () => {
   console.log("changeFn")
   isShow.value = true
}

// 组件发送的保存事件
const handleSave = async (message: string) => {
   isShow.value = false
   console.log("handleSave", message)
   ElMessage.success(message)
   store.saveToken("")
   router.push({ name: "login" })
}

// 组件发送的消息事件
const handleCancel = (message: string) => {
   isShow.value = false
   ElMessage.info(message)
}

// 展开或收起侧边栏时, 改变标题
const collapseTitle = ref("收起侧栏")
watch(
   () => isCollapse.value,
   (newValue) => {
      collapseTitle.value = isCollapse.value ? "展开侧栏" : "收起侧栏"
   }
)

// 修改密码
const handleChangePassword = () => {
   console.log("handleChangePassword")
   isShow.value = true
}

// 退出登录
const handleLogout = async () => {
   await ElMessageBox.confirm("确认退出?", "退出询问", {
      cancelButtonText: "取消",
      confirmButtonText: "确认",
      type: "warning",
   }).catch(() => {
      ElMessage.info("操作已取消")
      return new Promise(() => {})
   })

   await logoutApi().catch(() => {})

   ElMessage.success("已退出")
   store.saveToken("")
   router.push({ name: "login" })
}
</script>

<template>
   <el-menu
      class="el-menu-vertical-demo"
      :collapse="isCollapse"
      router
      unique-opened
      default-active="employee"
      active-text-color="#409EFF"
      background-color="#545c64"
      text-color="#fff">
      <el-menu-item @click="isCollapse = !isCollapse">
         <el-icon @click="isCollapse = !isCollapse"
            ><IEpExpand v-show="isCollapse" /> <IEpFold v-show="!isCollapse" />
         </el-icon>
         <template #title>{{ collapseTitle }}</template>
      </el-menu-item>
      <el-menu-item index="employee">
         <el-icon><IEpUser /></el-icon>
         <template #title>员工管理</template>
      </el-menu-item>
      <el-menu-item index="recruitment">
         <el-icon><IEpCollectionTag /></el-icon>
         <template #title>招聘管理</template>
      </el-menu-item>
      <el-sub-menu index="">
         <template #title>
            <el-icon><IEpKey /></el-icon>
            <span>后台管理</span>
         </template>
         <el-menu-item index="post">
            <template #title>岗位列表</template>
         </el-menu-item>
         <el-menu-item index="menu">
            <template #title>菜单列表</template>
         </el-menu-item>
         <el-menu-item index="role">
            <template #title>角色列表</template>
         </el-menu-item>
         <el-menu-item index="user">
            <template #title>用户列表</template>
         </el-menu-item>
      </el-sub-menu>
      <el-menu-item index="" @click="handleChangePassword">
         <el-icon><IEpSwitch /></el-icon>
         <template #title>修改密码</template>
      </el-menu-item>
      <ChangePassword :is-show="isShow" :title="title" @save="handleSave" @cancel="handleCancel"></ChangePassword>
      <el-menu-item index="" @click="handleLogout">
         <el-icon><IEpClose /></el-icon>
         <template #title>退出</template>
      </el-menu-item>
   </el-menu>
</template>

<style lang="scss" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
   width: 160px;
   min-height: 400px;
}

.el-menu-item.is-active {
   color: #fff !important;
   background: #409eff !important;
}
</style>
