<script setup lang="ts">
import { postListApi } from "@/utils/post"
import PostAddOrEdit from "./PostAddOrEdit.vue"
import moment from "moment"
import { TableColumnCtx } from "element-plus"

const state = reactive<{
   // 弹窗
   isShow: boolean
   title: string
   rowData: Post

   posts: Post[] // 从数据库读取出来的所有菜单
}>({
   isShow: false,
   title: "新增菜单",
   rowData: {} as Post,
   posts: [],
})

const { isShow, title, rowData, posts } = toRefs(state)

postListApi()
   .then((res) => {
      if (res.code === 200) {
         console.log(res)
         posts.value = res.data.posts
      }
   })
   .catch((res) => {
      console.log(res)
   })

function dateFormatter(row: any, column: TableColumnCtx<any>, cellValue: any, index: number) {
   var date = row[column.property]
   if (date == undefined) {
      return "未知"
   }

   return moment(date).format("YYYY-MM-DD")
}

const handleAdd = () => {
   console.log("新增岗位")
   title.value = "新增岗位"
   rowData.value = {} as Post
   isShow.value = true
}

const handleEdit = (index: number, row: Post | undefined) => {
   console.log("handleEdit", index, row)
   if (row) {
      rowData.value = row
      console.log("handleEdit", rowData.value)
      title.value = "修改岗位"
      isShow.value = true
   }
}

const handleDelete = (index: number, row: Post) => {
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
                  <span style="vertical-align: middle">新增岗位</span>
               </el-button>
            </div>
         </template>
         <el-table :data="posts" border style="width: 100%">
            <el-table-column prop="id" label="序号" align="center" />
            <el-table-column prop="name" label="岗位名称" align="center" />
            <el-table-column prop="updated_at" label="最后更新日期" :formatter="dateFormatter" align="center" />
            <el-table-column width="180" label="操作" align="center">
               <template #default="scope">
                  <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">修改</el-button>
                  <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
               </template>
            </el-table-column>
         </el-table>
      </el-card>
      <PostAddOrEdit :isShow="isShow" :title="title" :formData="rowData" @save="handleSave" @cancel="handleCancel">
      </PostAddOrEdit>
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
