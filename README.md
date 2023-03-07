# go mod 命令说明

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

# 项目中使用 Go Modules, 并构建项目

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

# Windows 下 golang 使用 sqlite, 开发环境需要安装 GCC, 因为 go 在编译 github.com/mattn/go-sqlite3(sqlite 数据库驱动)需要 CGO_ENABLED=1

https://github.com/niXman/mingw-builds-binaries 有 MinGWX64 的安装器和整个离线包在 Releases 中
https://www.mingw-w64.org/ 下载安装对应操作系统位的 mingw 或者其它 Windows 上的 GCC
配置 gcc 的 bin 到 PATH
配置 CGO_ENABLED=1 到系统环境

# 使用 sqlite

golang 交叉编译 go-sqlite3 后就可以单独使用了, 不需要 sqlite-dll-win64-x64(本体)和 sqlite-tools-win32-x86(命令行工具)
https://sqlite.org/download.html 下载 sqlite-dll-win64-x64(本体)和 sqlite-tools-win32-x86(命令行工具), 放在任意目录下
打开 cmd,切换到安装的目录,执行`sqlite test.db`可以在当前创建数据库文件,然后就可以把 test.db 放到工程中单独使用了

# 打包 (Inno Setup)

打包工具: http://www.dayanzai.me/inno-setup.html

# 安装

# 系统初始化

hrm.exe -init -u admin -p 123456
