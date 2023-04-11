<script setup lang="ts">
import { menuListApi } from "@/utils/menu"
import type { MenuListResponse } from "@/utils/menu"

const allMenus = ref<MenuListResponse[]>()

menuListApi()
   .then((res) => {
      if (res.code === 200) {
         console.log(res)
         allMenus.value = res.data
      }
   })
   .catch((res) => {
      console.log(res)
   })
</script>

<template>
   <el-card class="box-card">
      <template #header>
         <div class="card-header">
            <el-button type="primary">
               <IEpPlus />
               <span style="vertical-align: middle">增加菜单</span>
            </el-button>
         </div>
      </template>
      <el-table :data="allMenus" border style="width: 100%">
         <el-table-column prop="date" label="Date" width="180" />
         <el-table-column prop="name" label="Name" width="180" />
         <el-table-column prop="address" label="Address" />
      </el-table>
   </el-card>
</template>

<style lang="scss" scoped>
.card-header {
   display: flex;
   justify-content: space-between;
   align-items: center;
}

.text {
   font-size: 14px;
}

.item {
   margin-bottom: 18px;
}

.box-card {
   width: auto;
}
</style>
