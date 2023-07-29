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

#运行hertz-http-server
$ cd hertz-http-server
$ go run .

#运行kitex
$ cd microservices/student-service/
$ go run .


```

##### 接口协议
**学生注册接口**
请求地址
> POST http://127.0.0.1:8080/student/register

请求参数
> name:Xiao Hong
> age:18
> email:xiaohong@tsinghua.edu.cn
> address:Tsinghua University, 30 Shuangqing Road, Haidian District, Beijing 100084, China.

响应结果
> {
>     "status": "success",
>     "id": "10001"
> }


**查询学生信息接口**
请求地址
> POST http://127.0.0.1:8080/student/get

请求参数
> id:10001

响应结果
> {
>     "name": "Xiao Hong",
>     "age": "18",
>     "email": "xiaohong@tsinghua.edu.cn",
>     "address": "Tsinghua University, 30 Shuangqing Road, Haidian District, Beijing 100084, China."
> }

