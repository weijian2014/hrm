import { createApp, useAttrs } from "vue"
import ElementPlus from "element-plus"
import "virtual:windi.css"
import router from "./router/index"
import App from "./App.vue"
import { createPinia } from "pinia"
import axios from "./api/request"
import locale from "element-plus/lib/locale/lang/zh-cn"

let app = createApp(App)
app.config.globalProperties.$http = axios

// 全局注册图标
import * as ElementPlusIconsVue from "@element-plus/icons-vue"
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
   app.component(key, component)
}
const pinia = createPinia()
app.use(ElementPlus, { locale }).use(pinia).use(router).mount("#app")
