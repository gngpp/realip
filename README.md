### 简介
> 获取请求主机真实IP，适用于DDNS获取外网IP
### 使用方式
- 运行测试
```shell
go run app.go
# 默认服务端口8080，需要更换端口加上-p参数
// example: go run app.go -p=9999
// open web browser: http://127.0.0.1:9999/ip
// response: {"proto":"HTTP/1.1","real_ip":"127.0.0.1","uri":"/ip","request_url":"http://127.0.0.1:8080/ip","method":"GET","time":"1.24µs"}

```
- 手动编译
```shell
go build -o app app.go
```
- 打包docker镜像
```shell
docker build -t realip .
```
