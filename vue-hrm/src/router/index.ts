import { createRouter, createWebHistory } from "vue-router"
import HomeView from "@/views/Home.vue"
import AppLayout from "@/components/Layout.vue"

const router = createRouter({
   history: createWebHistory(import.meta.env.BASE_URL),
   routes: [
      {
         path: "/",
         name: "home",
         component: AppLayout,
      },
      {
         path: "/login",
         name: "login",
         component: () => import("@/views/Login.vue"),
      },
      {
         path: "/employee",
         name: "employee",
         component: () => import("@/components/employee/Employee.vue"),
      },
   ],
})

export default router
