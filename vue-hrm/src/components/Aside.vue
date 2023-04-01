<script setup lang="ts">
import { reactive, toRefs } from "vue"

const state = reactive({
   index: "",
})

const { index } = toRefs(state)
</script>

<template>
   <el-aside>
      <el-menu>
         <el-menu-item>
            <el-icon><IEpSetting /></el-icon><span>员工管理</span>
         </el-menu-item>

         <el-sub-menu :index="index">
            <template #title
               ><el-icon><IEpSetting /></el-icon><span>权限管理</span>
            </template>
            <el-menu-item>
               <el-icon><IEpSetting /></el-icon><span>菜单列表</span>
            </el-menu-item>
            <el-menu-item>
               <el-icon><IEpSetting /></el-icon><span>角色列表</span>
            </el-menu-item>
            <el-menu-item>
               <el-icon><IEpSetting /></el-icon><span>用户列表</span>
            </el-menu-item>
         </el-sub-menu>

         <el-menu-item>
            <el-icon><IEpSetting /></el-icon><span>系统设置</span>
         </el-menu-item>
      </el-menu>
   </el-aside>
</template>

<style lang="scss" scoped></style>
