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

#### 微服务添加

在IDL-manage可以进行idl的更新，将更新的内容加入gateway_api.thrift文件中

hzsvr-http文件夹包含用于接受HTTP请求的代码和API网关实现的主要业务代码。

utils文件夹包含文件utils.go，提供了许多有用的功能来帮助使用API网关。

IDL-manage文件夹包含thrift接口定义语言文件。这些文件用于为Hertz HTTP服务器以及所开发的任何微服务生成基础架构代码。

添加服务文件夹包含用于RPC服务器的代码，即添加微服务。
这些不是网关本身的一部分，而更像是如何使用它的示例。现在将实现第三个服务—division来如何展示使用API网关构建一个新的微服务并将其添加到现有的微服务基础上。

在执行以下步骤之前，将库克隆到本地计算机上，并确保位于GO根目录中

根据thrift的规范完成idl文件的编写之后，进入./rpc-service目录运行

```
kitex -module github.com/Lortals/111cloudwego-apigateway -service Division ../IDL-manage/xxx_management.thrift
```

其中xxx_management.thrift是新增加的微服务功能

使用Kitex生成脚手架后，在handler.go和main.go中加入kitex服务器的逻辑

完成后进入./hzsvr-http并打开终端运行

```
hz update -idl ../IDL-manage/gateway_api.thrift
```

接下来在hzsvr-http\biz\handler\api\gateway.go中完成业务逻辑

微服务的添加完成
