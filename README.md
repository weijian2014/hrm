# go mod命令说明
| 命令 | 说明 |
| :- | :- |
| go mod download | 下载依赖的module到本地cache
| go mod edit | 编辑go.mod
| go mod graph | 打印模块依赖图
| go mod init | 在当前目录下初始化go.mod(就是会新建一个 go.mod 文件)
| go mod tidy | 整理依赖关系，会添加丢失的module，删除不需要的 module
| go mod vender | 将依赖复制到vendor下
| go mod verify | 校验依赖
| go mod why | 解释为什么需要依赖
| go list -m -json all | JSON格式显示所有Import库信息

# 项目中使用Go Modules, 并构建项目
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

# 安装Postgresql(https://www.postgresql.org/download/)
    #https://blog.csdn.net/rudy5348/article/details/79299162
    yum -y install https://download.postgresql.org/pub/repos/yum/10/redhat/rhel-7-x86_64/pgdg-centos10-10-2.noarch.rpm
    yum -y install postgresql10 postgresql10-server
    /usr/pgsql-10/bin/postgresql-10-setup initdb && systemctl enable postgresql-10 && systemctl start postgresql-10
    su - postgres
    psql
    ALTER USER postgres WITH PASSWORD 'wj303101610';
    \q
    exit

    vim /var/lib/pgsql/10/data/pg_hba.conf(perr--->md5, ident--->md5)
    local   all             postgres                                peer
    local   all             postgres                                md5<---

    vim /var/lib/pgsql/10/data/postgresql.conf
    listen_addresses = 'localhost'
    listen_addresses = '*'

    systemctl restart postgresql-10

    psql -U postgres #打开postgresql命令行

# 创建数据库和表
    psql -h localhost -d postgres -U postgres -f /opt/create_table.sql

# 批量导入数据(序列不用在文件中给出，插入需要插入的字段即可)
    psql -U postgres
    COPY t_parts(vehicle_name, parts_name, parts_no, parts_count, parts_price, parts_selling_price) FROM '/tmp/data.txt' DELIMITER ',';

# 导出数据库
    pg_dump --host=192.168.199.123 --port=5432 --username=postgres --password=123456 --dbname=db_hi -C -F -v -f ~/hi.sql
    pg_dump -h 192.168.199.123 -p 5432 -U postgres -W -C -F -f ~/hi.sql db_hi

# 打包
    cd hiweb/src
    tar -czf hiweb.tar.gzip ../config/ web/ main

# 安装
    mkdir -p /opt/hiweb
    tar -zxvf hiweb.tar.gzip -C /opt/hiweb
    /opt/hiweb/main -f /opt/hiweb/config/config.json

# 从SQL文件中创建并恢复数据
    psql -U postgres --set ON_ERROR_STOP=ON -f /root/weijian/hiweb/doc/hi_20180328232626.sql
