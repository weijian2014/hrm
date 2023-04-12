<script setup lang="ts">
import { menuListApi } from "@/utils/menu"
import type { Menu } from "@/utils/menu"
import AddOrEdit from "./AddOrEdit.vue"

const state = reactive<{
   // 弹窗
   isShow: boolean
   title: string
   rowData: Menu

   menus: Menu[] // 从数据库读取出来的所有菜单
}>({
   isShow: false,
   title: "新增菜单",
   rowData: {} as Menu,
   menus: [],
})

const { isShow, title, rowData, menus } = toRefs(state)

menuListApi()
   .then((res) => {
      if (res.code === 200) {
         console.log(res)
         menus.value = res.data.menus
      }
   })
   .catch((res) => {
      console.log(res)
   })

const handleAdd = () => {
   console.log("新增菜单")
   title.value = "新增菜单"
   rowData.value = {} as Menu
   isShow.value = true
}

const handleEdit = (index: number, row: Menu | undefined) => {
   console.log("handleEdit", index, row)
   if (row) {
      rowData.value = row
      console.log("handleEdit", rowData.value)
      title.value = "修改菜单"
      isShow.value = true
   }
}

const handleDelete = (index: number, row: Menu) => {
   console.log("handleDelete", index, row)
}

// 组件发送的保存事件
const handleSave = (message: string) => {
   // todo 更新表格数据
   isShow.value = false
   console.log("handleSave", message, rowData.value)
   ElMessage.success(message)
}

// 组件发送的消息事件
const handleCancel = (message: string) => {
   isShow.value = false
   console.log("handleCancel", message, rowData.value)
   ElMessage.info(message)
}
</script>

<template>
   <div>
      <el-card class="box-card">
         <template #header>
            <div class="card-header">
               <el-button type="primary" @click="handleAdd">
                  <IEpPlus />
                  <span style="vertical-align: middle">新增菜单</span>
               </el-button>
            </div>
         </template>
         <el-table :data="menus" border style="width: 100%">
            <el-table-column prop="id" label="序号" align="center" />
            <el-table-column prop="name" label="菜单名" align="center" />
            <el-table-column prop="parent_id" label="父菜单序号" align="center" />
            <el-table-column prop="icon" label="菜单图标" align="center" />
            <el-table-column prop="url" label="链接" align="center" />
            <el-table-column width="180" label="操作" align="center">
               <template #default="scope">
                  <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
                  <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
               </template>
            </el-table-column>
         </el-table>
      </el-card>
      <AddOrEdit :isShow="isShow" :title="title" :formData="rowData" @save="handleSave" @cancel="handleCancel">
      </AddOrEdit>
   </div>
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
