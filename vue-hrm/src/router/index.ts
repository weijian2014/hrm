import { createRouter, createWebHistory } from "vue-router"
import HomeView from "@/views/Home.vue"
import LayoutVue from "@/components/layout/Layout.vue"
import { useUserStore } from "@/store/user"

const router = createRouter({
   history: createWebHistory(import.meta.env.BASE_URL),
   routes: [
      {
         path: "/login",
         name: "login",
         component: () => import("@/views/Login.vue"),
      },
      {
         path: "/",
         name: "",
         // 当前脚本执行立即加载
         component: LayoutVue,
         // 首页及其子页需要登录才能显示
         meta: { requiresAuth: true },
         children: [
            {
               path: "",
               name: "home",
               // 赖加载
               component: () => import("@/components/summary/Summary.vue"),
            },
            {
               path: "/recruitment",
               name: "recruitment",
               component: () => import("@/components/recruitment/Recruitment.vue"),
            },
            {
               path: "/post",
               name: "post",
               component: () => import("@/components/post/Post.vue"),
            },
            {
               path: "/menu",
               name: "menu",
               component: () => import("@/components/menu/Menu.vue"),
            },
            {
               path: "/role",
               name: "role",
               component: () => import("@/components/role/Role.vue"),
            },
            {
               path: "/user",
               name: "user",
               component: () => import("@/components/user/User.vue"),
            },
            {
               path: "/employee",
               name: "employee",
               component: () => import("@/components/employee/Employee.vue"),
            },
            {
               path: "/menu",
               name: "menu",
               component: () => import("@/components/menu/Menu.vue"),
            },
            {
               // 路由不匹配时显示此页面
               path: "/:xxx(.*)*",
               name: "error",
               component: () => import("@/components/Error.vue"),
            },
         ],
      },
   ],
})

// 路由拦截(路由变化时调用)
router.beforeEach((to, from, next) => {
   // some遍历每一个路由是否需要登录后才能显示
   if (to.matched.length > 0 && to.matched.some((r) => r.meta?.requiresAuth)) {
      const store = useUserStore()
      if (!store.tokenInfo.access_token) {
         // 跳转到login页, 登录成功后跳转回to.fullPath
         next({ name: "login", query: { redirect: to.fullPath } })
         return
      }
   }

   // 继续跳转
   next()
})

export default router
