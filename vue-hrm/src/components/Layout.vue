<script setup lang="ts">
import Aside from "./Aside.vue"
</script>

<template>
   <div class="common-layout">
      <el-container>
         <Aside></Aside>
         <el-container>
            <el-header>Header</el-header>
            <el-main>Main</el-main>
         </el-container>
      </el-container>
   </div>
</template>

<style lang="scss" scoped></style>
