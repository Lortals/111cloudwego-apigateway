### 系统部署

##### 系统环境

> 系统环境 ubuntu-20.04-desktop-amd64
> 系统配置 2核CPU 4G内存 40G硬盘
> Go版本 1.20.6

##### 安装环境

安装Go

```shell
#下载go
$ wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz

#解压缩
$ tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz

#设置环境变量
$ vim /etc/profile

#最后一行增加
export PATH=$PATH:/usr/local/go/bin

#更新环境变量
$ source /etc/profile

#检查go版本
$ go version

```

安装Git

```shell
$ sudo apt install git
```

安装etcd

```shell
# 下载
$ wget https://github.com/etcd-io/etcd/releases/download/v3.4.27/etcd-v3.4.27-linux-amd64.tar.gz

# 解压缩
$ tar zxf etcd-v3.4.27-linux-amd64.tar.gz

# 复制到系统目录
$ cd etcd-v3.4.27-linux-amd64/
$ cp -r etcd* /usr/local/bin/
```

安装Hertz

```shell
$ go install github.com/cloudwego/hertz/cmd/hz@latest
```

安装Kitex

```shell
$ go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
$ go install github.com/cloudwego/thriftgo@latest
$ kitex --version
$ thriftgo --version
```

##### 下载代码

从github获取代码

```shell
$ git clone https://github.com/Lortals/111cloudwego-apigateway.git
```

##### 运行项目

```shell
#启动ectd
$ etcd --log-level debug

#运行hzsvr-http
$ cd hzsvr-http
$ go run .

#运行kitex
$ cd rpc-services/student-service/
$ go run .


```

##### 接口协议

**学生注册接口**
请求地址

> POST http://127.0.0.1:8080/add-student-info

请求参数

> {"id": 99, "name":"zhong",
> "sex":"male", "age":20, "college":
> {"name": "ee college", "address":
> "NJU"}}

响应结果

> {"success":true,"message":"Information added successfully."}

**查询学生信息接口**
请求地址

> POST http://127.0.0.1:8080/query?id=99

请求参数

> id=99

响应结果

> {"id": 99, "name":"zhong",
> "sex":"male", "age":20, "college":
> {"name": "ee college", "address":
> "NJU"}}

#### IDL更新

在IDL-manage可以进行idl的更新，将更新的内容加入gateway_api.thrift文件中
