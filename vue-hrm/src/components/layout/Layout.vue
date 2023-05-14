<script setup lang="ts">
import Aside from "./Aside.vue"
</script>

<template>
   <div class="common-layout">
      <el-container>
         <Aside></Aside>
         <el-container class="header-and-main">
            <el-main>
               <el-scrollbar>
                  <RouterView></RouterView>
               </el-scrollbar>
            </el-main>
         </el-container>
      </el-container>
   </div>
</template>

<style lang="scss" scoped>
.header-and-main {
   flex-direction: column; // 垂直显示
   height: 100vh;
}

.el-main {
   background-color: #ffffff;
}
</style>
