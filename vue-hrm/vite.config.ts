import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import { ElementPlusResolver } from "unplugin-vue-components/resolvers"
import WindiCSS from "vite-plugin-windicss"

export default defineConfig({
   plugins: [
      vue(),
      WindiCSS({
         scan: {
            dirs: ["./src"], // ./src目录下所有文件
            fileExtensions: ["vue", "js", "ts", "tsx"], // 同时启用扫描vue/js/ts/tsx
         },
      }),
      AutoImport({
         resolvers: [ElementPlusResolver()],
      }),
      Components({
         resolvers: [ElementPlusResolver()],
      }),
      WindiCSS(),
   ],
})
