# GoFrame 框架学习

主要功能：使用go语言的GoFramework框架做的注册登陆程序

## 使用技术点

- mysql 数据模型操作
- redis 缓存操作
- lru内存缓存做三级缓存
- 资源打包到可执行文件做内存文件系统
- 从apollo配置中心获取配置信息并免重启更新
- 代码编译打包成docker镜像
- 用grpc做rpc调用
- 用etcd做rpc的服务发现
- jaeger 做链路跟踪

## 操作命令

### generated
generated文件夹下的内容为自动生成的代码，包含：

- 根据数据库表生成entity实体类
- do 数据库表描述类
- 对数据库表操作的dao类
- rpc的protobuf文件结合go语言和grpc插件生成的执行类

生成命令：
- 生成dao相关
  `gf gen dao`
- 从logic逻辑层生成service接口层代码
  `gf gen service -s internal/logic/user -d internal/service/user -f Snake`
- 把*.proc rpc描述文件生成对应语言的可执行文件
  `protoc --go_out=./generated/user/protobuf --go-grpc_out=./generated/user/protobuf ./api/rpc/v1/user/smc/smc.proto`

- 把资源文件打包成go文件，进而成为内存文件系统
  `gf pack manifest/config,resource packed/data.go -n packed`

- 从代码编译出预设置的平台和处理器架构的可执行文件并打包成docker镜像
  `gf docker main.go -tn fanghailiang:v1`

- jaeger链路跟踪
````shell
docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.24.0
````

- apollo配置中心
````shell
docker run -d --name apollo-quick-start \
-e APOLLO_CONFIG_DB_USERNAME=apollo \
-e APOLLO_CONFIG_DB_PASSWORD=apollo \
-e APOLLO_PORTAL_DB_USERNAME=apollo \
-e APOLLO_PORTAL_DB_PASSWORD=apollo \
-e APOLLO_PORTAL_DB_USERNAME=apollo \
-e JAVA_OPTS='-Xms100m -Xmx1000m -Xmn100m -Xss256k -XX:MetaspaceSize=10m -XX:MaxMetaspaceSize=250m' \
-p 8080:8080/tcp \
-p 8070:8070/tcp \
-p 8090:8090/tcp \
nobodyiam/apollo-quick-start
```` 
  
## 联系我
- 邮箱 **hailiang_fang@163.com**
- 微信 **fanghailiang2023**