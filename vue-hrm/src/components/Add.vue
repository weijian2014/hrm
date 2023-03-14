<template>
   <div class="container">
      <el-dialog v-model="dialogVisible" :title="employee?.name ? '修改' : '新增'" width="30%">
         <el-form>
            <el-form-item label="姓名" prop="name">
               <el-input v-model="form.name"> </el-input>
            </el-form-item>
         </el-form>
         <template #footer>
            <span class="dialog-footer">
               <el-button @click="handleSave">保存</el-button>
               <el-button @click="handleCancel">取消</el-button>
            </span>
         </template>
      </el-dialog>
   </div>
</template>

<script setup lang="ts">
import { ref, computed, defineProps, Ref, watch } from "vue"
import Employee from "./../class/Employee"

// defineProps定义了使用当前组件需要传入的对象
const props = defineProps({
   isShow: Boolean,
   id: Number,
   employee: Employee,
})

// 当props中的employee变化到newEmployee时, 给form赋值
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

const dialogVisible = computed(() => props.isShow)
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
const handleSave = () => {}

const handleCancel = () => {}
</script>
<style lang="scss" scoped></style>
