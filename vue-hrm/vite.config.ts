import { defineConfig, loadEnv } from "vite"
import { fileURLToPath, URL } from "node:url"
import vue from "@vitejs/plugin-vue"
import Icons from "unplugin-icons/vite"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import IconsResolver from "unplugin-icons/resolver"
import { ElementPlusResolver } from "unplugin-vue-components/resolvers"
import WindiCSS from "vite-plugin-windicss"
import path from "node:path"

const pathSrc = path.resolve(__dirname, "src")

// https://vitejs.dev/config/
export default defineConfig({
   plugins: [
      vue(),
      WindiCSS(),
      AutoImport({
         imports: [
            // 自动导入 Vue 相关函数，如：ref, reactive, toRef 等
            "vue",
         ],
         resolvers: [
            ElementPlusResolver(),
            // 自动导入图标组件
            IconsResolver({
               prefix: "Icon",
            }),
         ],
         dts: path.resolve(pathSrc, "auto-imports.d.ts"),
      }),
      Components({
         resolvers: [
            // 自动导入 Element Plus 组件
            ElementPlusResolver(),
            // 自动注册图标组件
            IconsResolver({
               enabledCollections: ["ep"],
            }),
         ],
         dts: path.resolve(pathSrc, "components.d.ts"),
      }),
      Icons({
         autoInstall: true,
      }),
   ],
   resolve: {
      alias: {
         "@": fileURLToPath(new URL("./src", import.meta.url)),
      },
   },
   server: {
      open: true, // 自动打开浏览器
      cors: true,
      proxy: {
         "^/api": {
            target: "http://0.0.0.0:8080",
            changeOrigin: true, // 允许跨域
            rewrite: (path) => path.replace(/^\/api/, ""), // 将api替换为空
         },
      },
   },
})
