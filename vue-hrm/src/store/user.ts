import { defineStore } from "pinia"

const tokenInfoKey = "TokenInfo"
const employeeTableSettingsKey = "EmployeeTableSettings"

export const useUserStore = defineStore("user", () => {
   const tokenJson = ref("")
   const data = reactive({})
   const employeeTableSettings = ref<Map<string, boolean>>()

   const tokenInfo = computed<TokenInfo>(() => {
      try {
         // 先读Store的tokenJson, 空再从Local Storage读取
         return JSON.parse(tokenJson.value || window.localStorage.getItem(tokenInfoKey) || "{}")
      } catch (err) {
         ElMessage.error("保存token时json格式错误")
         window.localStorage.setItem(tokenInfoKey, "")
         throw err
      }
   })

   function saveToken(data: string) {
      tokenJson.value = data
      // 浏览器刷新后Store就清空了, 所以需要将数据保存到Local Storage
      window.localStorage.setItem(tokenInfoKey, data)
   }

   // const employeeTableSettings = computed<Map<string, boolean>>(() => {
   //    try {
   //       return JSON.parse(
   //          employeeTableSettingsJson.value || window.localStorage.getItem(employeeTableSettingsKey) || "{}"
   //       )
   //    } catch (err) {
   //       ElMessage.error("保存EmployeeTableSettings时json格式错误")
   //       window.localStorage.setItem(employeeTableSettingsKey, "")
   //       throw err
   //    }
   // })
   function saveEmployeeTableSettings(settings: Map<string, boolean>) {
      employeeTableSettings.value = settings
      const settingsString = JSON.stringify(settings)
      console.log("settingsString", settingsString)
      window.localStorage.setItem(employeeTableSettingsKey, settingsString)
   }

   return { tokenInfo, saveToken, data, employeeTableSettings, saveEmployeeTableSettings }
})
