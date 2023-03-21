import { defineStore } from "pinia"

export const useMainStore = defineStore("main", {
   state: () => ({
      data: "",
      roles: [],
      menus: [],
   }),
   actions: {
      reset() {
         this.data = ""
         this.roles = []
         this.menus = []
      },
   },
})
