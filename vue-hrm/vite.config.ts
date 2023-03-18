import { fileURLToPath, URL } from "node:url"

import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import { ElementPlusResolver } from "unplugin-vue-components/resolvers"
import WindiCSS from "vite-plugin-windicss"

// https://vitejs.dev/config/
export default defineConfig({
   plugins: [
      vue(),
      AutoImport({
         resolvers: [ElementPlusResolver()],
      }),
      Components({
         resolvers: [ElementPlusResolver()],
      }),
      WindiCSS(),
   ],
   resolve: {
      alias: {
         "@": fileURLToPath(new URL("./src", import.meta.url)),
      },
   },
   server: {
      open: true, // 自动打开浏览器
   },
})
