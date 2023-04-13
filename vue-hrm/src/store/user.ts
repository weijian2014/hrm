import { defineStore } from "pinia"

export const useUserStore = defineStore("user", () => {
   const tokenJson = ref("")
   const data = reactive({})
   const employeeCloumnSettings = ref<EmployeeColumnSetting[]>([
      {
         label: "姓名",
         visible: true,
         disable: true,
      },
      {
         label: "性别",
         visible: true,
         disable: true,
      },
      {
         label: "生日",
         visible: true,
         disable: true,
      },
      {
         label: "参加工作时间",
         visible: true,
         disable: true,
      },
      {
         label: "工资",
         visible: true,
         disable: true,
      },
      {
         label: "岗位",
         visible: true,
         disable: true,
      },
      {
         label: "社保",
         visible: true,
         disable: true,
      },
      {
         label: "电话",
         visible: true,
         disable: true,
      },
      {
         label: "原单位",
         visible: true,
         disable: true,
      },
      {
         label: "身高",
         visible: true,
         disable: true,
      },
      {
         label: "体重",
         visible: true,
         disable: true,
      },
      {
         label: "学历",
         visible: true,
         disable: true,
      },
      {
         label: "政治面貌",
         visible: true,
         disable: false,
      },
      {
         label: "身份证",
         visible: true,
         disable: false,
      },
      {
         label: "保安证",
         visible: true,
         disable: false,
      },
      {
         label: "现住址",
         visible: true,
         disable: false,
      },
      {
         label: "备注",
         visible: true,
         disable: false,
      },
   ])

   const tokenInfo = computed<TokenInfo>(() => {
      try {
         // 先读Store的tokenJson, 空再从Local Storage读取
         return JSON.parse(tokenJson.value || window.localStorage.getItem("TokenInfo") || "{}")
      } catch (err) {
         ElMessage.error("保存token时json格式错误")
         window.localStorage.setItem("TokenInfo", "")
         throw err
      }
   })

   function saveToken(data: string) {
      tokenJson.value = data
      // 浏览器刷新后Store就清空了, 所以需要将数据保存到Local Storage
      window.localStorage.setItem("TokenInfo", data)
   }

   return { tokenInfo, saveToken, data, employeeCloumnSettings }
})
