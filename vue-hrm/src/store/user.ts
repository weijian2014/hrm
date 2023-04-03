import { defineStore } from "pinia"

export const useUserStore = defineStore("user", () => {
   const tokenJson = ref("")
   const data = reactive({})

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

   function saveToken(info: TokenInfo) {
      tokenJson.value = JSON.stringify(info)
      // 浏览器刷新后Store就清空了, 所以需要将数据保存到Local Storage
      window.localStorage.setItem("TokenInfo", JSON.stringify(info))
   }

   return { tokenInfo, saveToken, data }
})
