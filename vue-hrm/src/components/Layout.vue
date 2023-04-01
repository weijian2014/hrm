<script setup lang="ts">
import Aside from "./Aside.vue"
import Header from "./Header.vue"
</script>

<template>
   <div class="common-layout">
      <el-container>
         <Aside></Aside>
         <el-container class="header-and-main">
            <Header>Header</Header>
            <el-main>Main</el-main>
         </el-container>
      </el-container>
   </div>
</template>

<style lang="scss" scoped>
.header-and-main {
   flex-direction: column; // 垂直显示
}
</style>
