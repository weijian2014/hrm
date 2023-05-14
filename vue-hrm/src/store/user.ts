import { defineStore } from "pinia"

const tokenInfoKey = "TokenInfo"

export const useUserStore = defineStore("user", () => {
   const tokenJson = ref("")
   const logo = ref("/src/assets/img/user_logo.png")
   const data = reactive({})

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

   return { tokenInfo, saveToken, logo, data }
})
