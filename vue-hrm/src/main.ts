import { createApp } from "vue"
import App from "./App.vue"
import router from "./router/index"
import ElementPlus from "element-plus"
import locale from "element-plus/lib/locale/lang/zh-cn"
import axios from "axios"
import "virtual:windi.css"

const app = createApp(App)
app.config.globalProperties.$http = axios
// 全局注册图标
import * as ElementPlusIconsVue from "@element-plus/icons-vue"
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
   app.component(key, component)
}

app.use(ElementPlus, { locale }).use(router).mount("#app")
