<script setup lang="ts">
import { isCollapse } from "./index"
</script>

<template>
   <el-aside>
      <el-scrollbar>
         <el-menu router unique-opened :collapse="isCollapse">
            <a href="/" class="logo">
               <img src="@/assets/logo.svg" alt="" />
               <h1>HRM管理系统</h1>
            </a>
            <el-menu-item index="/employee">
               <el-icon><IEpUser /></el-icon><span>员工管理</span>
            </el-menu-item>
            <el-menu-item index="/recruitment">
               <el-icon><IEpCollectionTag /></el-icon><span>招聘管理</span>
            </el-menu-item>

            <el-sub-menu index="/permission">
               <template #title
                  ><el-icon><IEpKey /></el-icon><span>后台管理</span>
               </template>
               <el-menu-item index="post">
                  <el-icon><IEpStamp /></el-icon><span>岗位列表</span>
               </el-menu-item>
               <el-menu-item index="menu">
                  <el-icon><IEpMenu /></el-icon><span>菜单列表</span>
               </el-menu-item>
               <el-menu-item index="role">
                  <el-icon><IEpUserFilled /></el-icon><span>角色列表</span>
               </el-menu-item>
               <el-menu-item index="user">
                  <el-icon><IEpAvatar /></el-icon><span>用户列表</span>
               </el-menu-item>
            </el-sub-menu>

            <el-menu-item index="setting">
               <el-icon><IEpSetting /></el-icon><span>系统设置</span>
            </el-menu-item>
         </el-menu>
      </el-scrollbar>
   </el-aside>
</template>

<style lang="scss" scoped>
.el-aside {
   background-color: #a0cfff;
   height: 100vh;
   width: auto;

   .el-menu {
      background-color: #a0cfff;
      border-right: none;
      width: 200px;
      // 当el-menu的class有el-menu--collapse样式时, 修改侧边栏宽度
      // &表示父元素
      &.el-menu--collapse {
         width: 60px;
         // 不显示h1标签
         & h1 {
            display: none;
         }
      }

      .logo {
         display: flex; // 弹性
         justify-content: center; // 水平居中显示
         align-items: center;
         height: 60px;
         text-decoration: none; // 取消a标签的下划线
         img {
            width: 32px;
            height: 32px;
         }
      }
   }
}
</style>
