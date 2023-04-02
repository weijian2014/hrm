import { createRouter, createWebHistory } from "vue-router"
import HomeView from "@/views/Home.vue"
import LayoutVue from "@/components/layout/Layout.vue"

const router = createRouter({
   history: createWebHistory(import.meta.env.BASE_URL),
   routes: [
      {
         path: "/",
         name: "",
         // 当前脚本执行立即加载
         component: LayoutVue,
         children: [
            {
               path: "",
               name: "home",
               // 赖加载
               component: () => import("@/components/summary/Summary.vue"),
            },
            {
               path: "/employee",
               name: "employee",
               // 赖加载
               component: () => import("@/components/employee/Employee.vue"),
            },
            {
               path: "/:xxx(.*)*",
               name: "error",
               component: () => import("@/components/Error.vue"),
            },
         ],
      },
      {
         path: "/login",
         name: "login",
         component: () => import("@/views/Login.vue"),
      },
   ],
})

export default router
