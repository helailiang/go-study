# go-zero 学习


## user center
1. 先要在usercenter.api中定义好service中的方法
2. goctl生成api代码
   a. 切换到api所在的目录，运行` goctl api go -api *.api -dir ../  -style=goZero`
    解释上面命令：
      * 根据 api 文件生成 Go HTTP 代码, 
      * api：api 文件路径 ；
      * dir：代码输出目录
      * style： 输出文件和目录的命名风格格式化符号，详情见 文件风格
        https://go-zero.dev/docs/tutorials/cli/api#go
        goctl rpc protoc greet.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --client=true 
   
# 单个 rpc 服务生成示例指令
```shell
cd usercenter/cmd/rpc/pb
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ -style=goZero
sed -i "" 's/,omitempty//g' *.pb.go
```


## 生成model
cd usercenter/cmd/rpc
goctl model mysql datasource -url="root:123456@tcp(10.0.51.35:3306)/zero" -table="user,user_auth"  -dir="./model" -cache=true --style=goZero


```shell
cd usercenter/cmd/api/desc
goctl api plugin -plugin goctl-swagger="swagger -filename usercenter.json -host 127.0.0.0 -basepath /api -schemes http" -api usercenter.api -dir .
```