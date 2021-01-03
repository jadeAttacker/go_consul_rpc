#  protoc的使用
## 假设现在所在的目录是$GOPATH/src/helloworld/helloworld，我们将通过如下命令生成gRPC对应的GoLang代码

```
<!-- # 1. 编写proto文件 -->
<!-- # 2. 生成 go文件 -->
protoc --go_out=plugins=grpc:. helloworld.proto

<!-- # 3. 实现接口 -->
```

## 此时，将在目录下生成helloworld.pb.go文件


## 不指定目录：protoc -I=目录1 --java_out=目录2 文件3  
```
<!-- # 参考https://www.jianshu.com/p/9362c28eb539?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation -->

protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

# RPC调用




# docker命令文档
```
<!-- 先打开docker -->
https://blog.csdn.net/wangmx1993328/article/details/81735070
```

# consul命令文档
```
<!-- 参考 https://blog.csdn.net/y435242634/article/details/78769147#t15 -->

consul members

consul catalog nodes -http-addr=192.168.1.107:8500
```

## 启动docker作为consul节点
```
docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul:latest agent -server -bootstrap-expect 2 -ui -bind=0.0.0.0 -client=0.0.0.0

docker run --name consul2 -d -p 8501:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join 172.17.0.2
docker run --name consul3 -d -p 8502:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join 172.17.0.2

<!-- 
-d：表示后台运行 image-name:指定运行的镜像名称
-p： 表示各种端口。8500:web ui端口 
image-name:tag： 指定运行的镜像名称
-bootstrap-expect： 期待集群有2个节点
docker inspect --format '{{ .NetworkSettings.IPAddress }}' consul2： 获取 consul server1 的 ip 地址

-->
```
## 查看节点信息,打开:
`http://192.168.1.107:8500/ui/dc1/nodes/0b8ff3e89bac`



## 参考：
`https://www.cnblogs.com/chaselogs/p/11462954.html` 

`https://www.cnblogs.com/lonelyxmas/p/10880717.html`