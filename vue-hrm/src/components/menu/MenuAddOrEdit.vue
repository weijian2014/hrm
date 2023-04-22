<script setup lang="ts">
import type { Menu } from "@/utils/menu"

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps<{
   isShow: boolean
   title: string
   formData: Menu
   allMenu: Menu[]
}>()

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "cancel"])

const state = reactive<{
   isEscapeClose: boolean
   isShowClose: boolean
   isClickModalToClose: boolean
   dialogVisible: boolean
   formLabelWidth: string
   rawFormData: Menu
   menus: Menu[]
}>({
   isEscapeClose: false, // 是否按ESC关闭
   isShowClose: false, // 是否显示右上角的关闭
   isClickModalToClose: false, // 是否可以通过点击 modal 关闭 Dialog (对话框以外的任意位置)
   dialogVisible: false, // 是否显示对话框
   formLabelWidth: "10px",
   rawFormData: {} as Menu,
   menus: [],
})

// 解构
const { isEscapeClose, isShowClose, isClickModalToClose, dialogVisible, formLabelWidth, rawFormData, menus } =
   toRefs(state)

// watch写法上支持一个或者多个监听源, 这些监听源必须只能是getter/effect函数, ref数据, reactive对象或者数组类型
watch(
   () => props.formData,
   (newValue) => {
      // 当props.formData变化为newValue时, 拷贝一份原始数据,
      // 不要直接使用props.formData, 否则会影响父组件的数据
      rawFormData.value = { ...props.formData }
   }
)

watch(
   () => props.allMenu,
   (newValue) => {
      menus.value.push({ id: 0, name: "顶级菜单", parent_id: 0, icon: "", url: "", updated_at: "" })
      props.allMenu.forEach((item) => {
         if (item.parent_id === 0) {
            menus.value.push(item)
         }
      })
   }
)

// 当props.isShow变化时会传递给dialogVisible, 而dialogVisible被绑定给了el-dialog, 从而达到外部控制显示隐藏el-dialog的目的
watch(
   () => props.isShow,
   (newValue) => {
      dialogVisible.value = newValue as boolean
   }
)

const handleSave = () => {
   // todo 保存到数据库
   // 向外发送save(保存)事件
   emits("save", "保存成功")
}

const handleCancel = () => {
   // 向外发送cancel(取消)事件
   emits("cancel", "已取消")
}
</script>

<template>
   <div class="container">
      <!-- 当点击对话框右上角关闭时, 向外部发送canel事件 -->
      <el-dialog
         width="45%"
         draggable
         v-model="dialogVisible"
         :title="title"
         :close-on-press-escape="isEscapeClose"
         :show-close="isShowClose"
         :close-on-click-modal="isClickModalToClose">
         <el-form :model="rawFormData">
            <el-form-item label="菜单名" prop="name">
               <el-input v-model="rawFormData.name"> </el-input>
            </el-form-item>
            <el-form-item label="父菜单" prop="parent_id">
               <el-select v-model="rawFormData.parent_id" placeholder="">
                  <el-option v-for="item in menus" :key="item.id" :label="item.name" :value="item.name" />
               </el-select>
            </el-form-item>
            <el-form-item label="图标" prop="icon">
               <el-input v-model="rawFormData.icon"> </el-input>
            </el-form-item>
            <el-form-item label="链接" prop="url">
               <el-input v-model="rawFormData.url"> </el-input>
            </el-form-item>
         </el-form>
         <template #footer>
            <span class="dialog-footer">
               <el-button type="primary" @click="handleSave">保存</el-button>
               <el-button type="danger" @click="handleCancel">取消</el-button>
            </span>
         </template>
      </el-dialog>
   </div>
</template>

<style lang="scss" scoped></style>