<script setup lang="ts">
import zhCn from "element-plus/lib/locale/lang/zh-cn"
</script>

<template>
   <div>
      <el-config-provider :locale="zhCn">
         <router-view />
      </el-config-provider>
   </div>
</template>

<style lang="scss"></style>
