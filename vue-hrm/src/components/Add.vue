<template>
   <div class="container">
      <!-- 当点击对话框右上角关闭时, 向外部发送canel事件 -->
      <el-dialog
         v-model="dialogVisible"
         :title="employee?.name ? '修改' : '新增'"
         width="45%"
         draggable
         :close-on-press-escape="isEscapeClose"
         :show-close="isShowClose"
         :close-on-click-modal="isClickModalToClose">
         <el-form>
            <!-- 第1行 -->
            <el-row class="ml-6">
               <el-col :span="6">
                  <el-form-item label="姓名" prop="name">
                     <el-input v-model="form.name"> </el-input>
                  </el-form-item>
               </el-col>
               <el-col :span="6" class="mx-5">
                  <el-form-item label="性别" prop="gender">
                     <el-select v-model="form.gender" placeholder="">
                        <el-option
                           v-for="item in genderOptions"
                           :key="item.value"
                           :label="item.label"
                           :value="item.value" />
                     </el-select>
                  </el-form-item>
               </el-col>
               <el-col :span="8">
                  <el-form-item label="生日" prop="birthday">
                     <el-date-picker
                        v-model="form.birthday"
                        type="date"
                        placeholder="" />
                  </el-form-item>
               </el-col>
            </el-row>
            <!-- 第2行 -->
            <el-row class="ml-6">
               <el-col :span="6">
                  <el-form-item label="身高" prop="height">
                     <el-input v-model="form.height" placeholder="">
                        <template #append>cm</template>
                     </el-input></el-form-item
                  >
               </el-col>
               <el-col :span="6" class="mx-5">
                  <el-form-item label="体重" prop="weight">
                     <el-input v-model="form.weight" placeholder="">
                        <template #append>kg</template>
                     </el-input></el-form-item
                  >
               </el-col>
            </el-row>
            <!-- 第3行 -->
            <el-row class="ml-3">
               <el-col :span="20">
                  <el-form-item label="现住址" prop="current_address">
                     <el-input v-model="form.current_address"> </el-input>
                  </el-form-item>
               </el-col>
            </el-row>
            <!-- 第4行 -->
            <el-row>
               <el-col :span="8">
                  <el-form-item label="政治面貌" prop="political_status">
                     <el-select v-model="form.political_status" placeholder="">
                        <el-option
                           v-for="item in politicalStatusOptions"
                           :key="item.value"
                           :label="item.label"
                           :value="item.value" />
                     </el-select>
                  </el-form-item>
               </el-col>
               <el-col :span="6" class="mx-5">
                  <el-form-item label="学历" prop="degree">
                     <el-select v-model="form.degree" placeholder="">
                        <el-option
                           v-for="item in degreeOptions"
                           :key="item.value"
                           :label="item.label"
                           :value="item.value" />
                     </el-select>
                  </el-form-item>
               </el-col>
               <el-col :span="6">
                  <el-form-item label="社保" prop="social_security">
                     <el-select v-model="form.social_security" placeholder="">
                        <el-option
                           v-for="item in socialSecurityOptions"
                           :key="item.value"
                           :label="item.label"
                           :value="item.value" />
                     </el-select>
                  </el-form-item>
               </el-col>
            </el-row>
            <!-- 第5行 -->
            <el-row class="ml-3">
               <el-col :span="10">
                  <el-form-item label="身份证" prop="identifier">
                     <el-input
                        v-model="form.identifier"
                        :formatter="identifierFormatter"
                        placeholder="">
                     </el-input>
                  </el-form-item>
               </el-col>
               <el-col :span="8" class="ml-5">
                  <el-form-item label="电话" prop="phone">
                     <el-input
                        v-model="form.phone"
                        :formatter="phoneFormatter"
                        placeholder="">
                     </el-input>
                  </el-form-item>
               </el-col>
            </el-row>
            <!-- 第6行 -->
            <el-row class="ml-3">
               <el-col :span="10">
                  <el-form-item label="原单位" prop="former_employer">
                     <el-input v-model="form.former_employer"> </el-input>
                  </el-form-item>
               </el-col>
               <el-col :span="10" class="ml-5">
                  <el-form-item label="参加工作时间" prop="first_work_time">
                     <el-date-picker
                        v-model="form.first_work_time"
                        type="date"
                        placeholder="" />
                  </el-form-item>
               </el-col>
            </el-row>
            <!-- 第7行 -->
            <el-row class="ml-6">
               <el-col :span="6">
                  <el-form-item label="岗位" prop="post">
                     <el-select v-model="form.post" placeholder="">
                        <el-option
                           v-for="item in postOptions"
                           :key="item.value"
                           :label="item.label"
                           :value="item.value" />
                     </el-select>
                  </el-form-item>
               </el-col>
               <el-col :span="8" class="mx-5">
                  <el-form-item label="工资" prop="salary">
                     <el-input
                        v-model="form.salary"
                        :formatter="salaryFormatter"
                        placeholder="">
                        <template #append>¥</template>
                     </el-input>
                  </el-form-item>
               </el-col>
               <el-col :span="8">
                  <el-form-item label="保安证" prop="security_card">
                     <el-input
                        v-model="form.security_card"
                        :formatter="securityCardFormatter"
                        placeholder="">
                     </el-input>
                  </el-form-item>
               </el-col>
            </el-row>
            <el-row>
               <el-col :span="20">
                  <el-form-item label="备注" prop="comments">
                     <el-input
                        v-model="form.comments"
                        maxlength="800"
                        placeholder="填写了解的情况"
                        show-word-limit
                        type="textarea" />
                  </el-form-item>
               </el-col>
            </el-row>
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
import { ref, computed, watch } from "vue"
import Employee from "../class/Employee"

const genderOptions = [
   {
      value: "男",
      label: "男",
   },
   {
      value: "女",
      label: "女",
   },
]

const socialSecurityOptions = [
   {
      value: "有",
      label: "有",
   },
   {
      value: "无",
      label: "无",
   },
]

const politicalStatusOptions = [
   {
      value: "群众",
      label: "群众",
   },
   {
      value: "无党派人士",
      label: "无党派人士",
   },
   {
      value: "民主党派成员",
      label: "民主党派成员",
   },
   {
      value: "共青团员",
      label: "共青团员",
   },
   {
      value: "党员",
      label: "党员",
   },
]

const degreeOptions = [
   {
      value: "小学",
      label: "小学",
   },
   {
      value: "初中",
      label: "初中",
   },
   {
      value: "高中",
      label: "高中",
   },
   {
      value: "中专",
      label: "中专",
   },
   {
      value: "大专",
      label: "大专",
   },
   {
      value: "专升本",
      label: "专升本",
   },
   {
      value: "本科",
      label: "本科",
   },
   {
      value: "研究生",
      label: "研究生",
   },
   {
      value: "博士",
      label: "博士",
   },
   {
      value: "博士后",
      label: "博士后",
   },
]

const postOptions = [
   {
      value: "前台",
      label: "前台",
   },
   {
      value: "巡逻",
      label: "巡逻",
   },
   {
      value: "消防",
      label: "消防",
   },
   {
      value: "教练",
      label: "教练",
   },
   {
      value: "保安",
      label: "保安",
   },
   {
      value: "银保",
      label: "银保",
   },
]

const phoneFormatter = (value: string | number) => {
   console.log("phoneFormatter", value, form)
}

const salaryFormatter = (value: string | number) => {
   console.log("salaryFormatter", value, form)
}

const identifierFormatter = (value: string | number) => {
   console.log("identifierFormatter", value, form)
}

const securityCardFormatter = (value: string | number) => {
   console.log("securityCardFormatter", value, form)
}

// defineProps定义了当前组件的属性, 外部组件使用当前组件可以绑定传递进来
const props = defineProps({
   isShow: Boolean,
   id: Number,
   employee: Employee,
})

// defineEmits定义了当前组件的事件, 可以向外部发送(通知外部组件)
const emits = defineEmits(["save", "cancel"])

const form = ref<Employee>(new Employee())

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
            birthday: newEmployee.birthday,
            first_work_time: newEmployee.first_work_time,
            salary: newEmployee.salary,
            post: newEmployee.post,
            social_security: newEmployee.social_security,
            phone: newEmployee.phone,
            former_employer: newEmployee.former_employer,
            height: newEmployee.height,
            weight: newEmployee.weight,
            degree: newEmployee.degree,
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
   // 向外发送cancel(取消)事件
   emits("cancel", "已取消")
}
</script>
<style lang="scss" scoped></style>
