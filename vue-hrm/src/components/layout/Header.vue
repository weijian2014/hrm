<script setup lang="ts">
import { useUserStore } from "@/store/user"
import { storeToRefs } from "pinia"
import { isCollapse } from "./index"
import { useRouter } from "vue-router"
import { logoutApi } from "@/utility/api"

const router = useRouter()
const store = useUserStore()
const { tokenInfo } = storeToRefs(store)

const logoutFn = async () => {
   await ElMessageBox.confirm("确认退出?", "退出询问", {
      confirmButtonText: "确认",
      cancelButtonText: "取消",
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
   <el-header>
      <el-icon @click="isCollapse = !isCollapse"
         ><IEpExpand v-show="isCollapse" /> <IEpFold v-show="!isCollapse" />
      </el-icon>

      <!-- 面包屑 -->
      <el-breadcrumb separator="/">
         <el-breadcrumb-item :to="{ path: '/' }">homepage</el-breadcrumb-item>
         <el-breadcrumb-item><a href="/">promotion management</a></el-breadcrumb-item>
         <el-breadcrumb-item>promotion list</el-breadcrumb-item>
         <el-breadcrumb-item>promotion detail</el-breadcrumb-item>
      </el-breadcrumb>

      <!-- 下拉菜单 -->
      <el-dropdown>
         <span class="el-dropdown-link">
            <el-avatar :size="32" :src="'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'" />
            <el-icon class="el-icon--right">
               <i-ep-arrow-down />
            </el-icon>
         </span>
         <template #dropdown>
            <el-dropdown-menu>
               <el-dropdown-item>{{ tokenInfo.user_name }}</el-dropdown-item>
               <el-dropdown-item divided @click="logoutFn">退出</el-dropdown-item>
            </el-dropdown-menu>
         </template>
      </el-dropdown>
   </el-header>
</template>

<style lang="scss" scoped>
.el-header {
   display: flex; // 让子元素水平显示
   align-items: center;
   background-color: #79bbff;

   .el-icon {
      margin-right: 17px;
   }

   .el-dropdown {
      margin-left: auto;

      .el-dropdown-link {
         display: flex;
         justify-content: center;
         align-items: center;
      }
   }
}
</style>
