# 系统概述

- 前端 Node.js18 + Vite4 + Vue3 + WindiCSS + Element Plus + TypeScript + Axios
- 后端 Golang + Sqlite + Restful API

# 后端

## go mod 命令说明

| 命令                 | 说明                                                   |
| :------------------- | :----------------------------------------------------- |
| go mod download      | 下载依赖的 module 到本地 cache                         |
| go mod edit          | 编辑 go.mod                                            |
| go mod graph         | 打印模块依赖图                                         |
| go mod init          | 在当前目录下初始化 go.mod(就是会新建一个 go.mod 文件)  |
| go mod tidy          | 整理依赖关系，会添加丢失的 module，删除不需要的 module |
| go mod vender        | 将依赖复制到 vendor 下                                 |
| go mod verify        | 校验依赖                                               |
| go mod why           | 解释为什么需要依赖                                     |
| go list -m -json all | JSON 格式显示所有 Import 库信息                        |

## 项目中使用 Go Modules, 并构建项目

```
linux
mkdir -p /opt/go/gopath
echo "export PATH=${PATH}:/opt/go" >> ~/.profile
echo "export GOROOT=/opt/go" >> ~/.profile
echo "export GOPATH=/opt/go/gopath" >> ~/.profile
echo "export GOPROXY=https://goproxy.io" >> ~/.profile
echo "export GO111MODULE=on" >> ~/.profile
source ~/.profile
cd hrm
go mod init hrm
go mod tidy
go get -u golang.org/x/net
gofmt -l -w .
go build .
```

## Windows 下 golang 使用 sqlite

- , 开发环境需要安装 GCC, 因为 go 在编译 github.com/mattn/go-sqlite3(sqlite 数据库驱动)需要 CGO_ENABLED=1
- [MingGW](https://github.com/niXman/mingw-builds-binaries) 有 MinGWX64 的安装器和整个离线包在 Releases 中
  https://www.mingw-w64.org/ 下载安装对应操作系统位的 mingw 或者其它 Windows 上的 GCC
  配置 gcc 的 bin 到 PATH
  配置 CGO_ENABLED=1 到系统环境

## 使用 sqlite

- golang 交叉编译 go-sqlite3 后就可以单独使用了, 不需要 sqlite-dll-win64-x64(本体)和 sqlite-tools-win32-x86(命令行工具)
- https://sqlite.org/download.html 下载 sqlite-dll-win64-x64(本体)和 sqlite-tools-win32-x86(命令行工具), 放在任意目录下
  打开 cmd,切换到安装的目录,执行`sqlite test.db`可以在当前创建数据库文件,然后就可以把 test.db 放到工程中单独使用了
- [SQLiteStudio](https://sqlitestudio.pl/)

# 前端

### Node.js 最好使用 LTS 版本, 可以使用 nvm 对 node.js 版本进行管理

- [Node.js](https://nodejs.org/zh-cn/) 后端使用 Libuv, C++编写, npm 是前端包管理工具, 安装 node 会自动安装 npm
- [pnpm](https://nodejs.org/zh-cn/), pnpm 是 npm 改进版本
- [Vue3](https://cn.vuejs.org/) 比 Vue2 更快
- [Vite](https://cn.vitejs.dev/) 是比 Webpack 更快更先进的前端项目构建工具
- [Vue-Router](https://router.vuejs.org/zh/installation.html)
- [WindiCSS](https://cn.windicss.org/integrations/vite.html) SCSS
- [axios](http://axios-js.com/)
- [Element Plus](https://element-plus.gitee.io/zh-CN/guide/installation.html) 是基于 Vue3 的 UI 库
- TypeScript 是 JavaScript 的超集, 支持面向对象的特性

### VScode 开发需要安装插件

- Live Server
- Prettier - Code formatter
- Vue Language Features (Volar)
- TypeScript Vue Plugin (Volar)
- WindiCSS IntelliSense

### 项目初始化

```
# 设置镜像源
npm config set registry https://registry.npm.taobao.org
pnpm config set registry https://registry.npm.taobao.org

# 查看镜像源
npm config get registry
pnpm config get registry

# 创建项目
npm init vue@latest
pnpm create vue@latest

# 初始化项目, install命令会根据 package.json 中 dependencies 和 devDependenciesr 的描述安装相应的依赖
npm install
pnpm install
```

### 安装组件

```
cd path/to/project/root

# 安装Vue的路由组件
npm install vue-router -S

# 安装windicss
npm i -D vite-plugin-windicss windicss

# 安装axios
npm install axios

# 安装element-plus
npm install element-plus -S
pnpm install element-plus


# 自动导入模块插件
npm install -D unplugin-vue-components unplugin-auto-import
```

### 将项目运行起来

```
# dev 是 package.json 中 scripts 的一个命令
nmp run dev
pnpm run dev
```

# 打包

[打包工具 Inno Setup](http://www.dayanzai.me/inno-setup.html)

# 安装

# 系统初始化

hrm.exe -init -u admin -p 123456
