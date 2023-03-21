import { defineStore } from "pinia"

export const useMainStore = defineStore("main", {
   state: () => ({
      counter: 0,
      name: "Eduardo",
   }),
   // optional getters
   getters: {
      doubleCounter: (state) => state.counter * 2,
      doubleCounterPlusOne(): number {
         return this.doubleCounter + 1
      },
   },
   // optional actions
   actions: {
      reset() {
         this.counter = 0
      },
   },
})
