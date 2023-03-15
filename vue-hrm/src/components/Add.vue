<template>
   <div class="container">
      <!-- 当点击对话框右上角关闭时, 向外部发送canel事件 -->
      <el-dialog
         v-model="dialogVisible"
         @close="$emit('canel')"
         :title="employee?.name ? '修改' : '新增'"
         width="30%"
         draggable
         :close-on-press-escape="isEscapeClose"
         :show-close="isShowClose"
         :close-on-click-modal="isClickModalToClose">
         <el-form>
            <el-form-item label="姓名" prop="name">
               <el-input v-model="form.name"> </el-input>
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

<script setup lang="ts">
import { ref, computed, defineProps, Ref, watch } from "vue"
import Employee from "./../class/Employee"

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps({
   isShow: Boolean,
   id: Number,
   employee: Employee,
})

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "canel"])

const form: Ref<Employee> = ref<Employee>({
   id: 0,
   name: "",
   gender: "",
   age: 0,
   work_time: "",
   salary: 0,
   post: "",
   social_security: "",
   phone: "",
   former_employer: "",
   height: 0,
   weight: 0,
   diploma: "",
   political_status: "",
   identifier: "",
   security_card: "",
   current_address: "",
   comments: "",
})

// 是否按ESC关闭
const isEscapeClose = ref<boolean>(false)
// 是否显示右上角的关闭
const isShowClose = ref<boolean>(false)
// 是否可以通过点击 modal 关闭 Dialog (对话框以外的任意位置)
const isClickModalToClose = ref<boolean>(false)

// 当props中的employee变化到newEmployee时, 给form赋值, 而form被绑定给了el-form, 从而达到外部传送值给el-form的目的
watch(
   () => props.employee,
   (newEmployee) => {
      if (newEmployee) {
         form.value = {
            id: newEmployee.id,
            name: newEmployee.name,
            gender: newEmployee.gender,
            age: newEmployee.age,
            work_time: newEmployee.work_time,
            salary: newEmployee.salary,
            post: newEmployee.post,
            social_security: newEmployee.social_security,
            phone: newEmployee.phone,
            former_employer: newEmployee.former_employer,
            height: newEmployee.height,
            weight: newEmployee.weight,
            diploma: newEmployee.diploma,
            political_status: newEmployee.political_status,
            identifier: newEmployee.identifier,
            security_card: newEmployee.security_card,
            current_address: newEmployee.current_address,
            comments: newEmployee.comments,
         }
      }
   }
)

// 当props.isShow变化时会传递给dialogVisible, 而dialogVisible被绑定给了el-dialog, 从而达到外部控制显示隐藏el-dialog的目的
const dialogVisible = computed(() => props.isShow)

const handleSave = () => {
   // 向外发送save(保存)事件
   emits("save", "保存成功")
}

const handleCancel = () => {
   // 向外发送canel(取消)事件
   emits("canel")
}
</script>
<style lang="scss" scoped></style>
