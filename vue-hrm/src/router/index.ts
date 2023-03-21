import { createRouter, createWebHistory } from "vue-router"
import HomeView from "@/views/Home.vue"

const router = createRouter({
   history: createWebHistory(import.meta.env.BASE_URL),
   routes: [
      {
         path: "/",
         name: "index",
         component: HomeView,
      },
      {
         path: "/home",
         name: "home",
         component: HomeView,
      },
      {
         path: "/login",
         name: "login",
         component: () => import("@/views/Login.vue"),
      },
   ],
})

export default router
