<script setup lang="ts">
const state = reactive({ isCollapse: false })

const { isCollapse } = toRefs(state)
</script>

<template>
   <el-header>
      <el-icon
         ><IEpExpand v-show="isCollapse" /> <IEpFold v-show="!isCollapse" />
      </el-icon>

      <!-- 面包屑 -->
      <el-breadcrumb separator="/">
         <el-breadcrumb-item :to="{ path: '/' }">homepage</el-breadcrumb-item>
         <el-breadcrumb-item
            ><a href="/">promotion management</a></el-breadcrumb-item
         >
         <el-breadcrumb-item>promotion list</el-breadcrumb-item>
         <el-breadcrumb-item>promotion detail</el-breadcrumb-item>
      </el-breadcrumb>

      <!-- 下拉菜单 -->
      <el-dropdown>
         <span class="el-dropdown-link">
            <el-avatar
               :size="32"
               :src="'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'" />
            <el-icon class="el-icon--right">
               <i-ep-arrow-down />
            </el-icon>
         </span>
         <template #dropdown>
            <el-dropdown-menu>
               <el-dropdown-item>用户名</el-dropdown-item>
               <el-dropdown-item divided>退出</el-dropdown-item>
            </el-dropdown-menu>
         </template>
      </el-dropdown>
   </el-header>
</template>

<style lang="scss" scoped>
.el-header {
   display: flex; // 让子元素水平显示
   align-items: center;
   background-color: #ccc;

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
