

### 数据库
1. 数据库：COLLATE=utf8mb4_0900_ai_ci
2. 建表：ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 
3. 项目启动时，会检测当前数据库无user表情况下，创建user表，并在models文件夹生成表对应结构体文件model.go
4. 生成model.go使用models下cmd/gen.go，可根据genModel.go进行配置


### 功能

#### 1. gRPC
- 通过id查询用户
- 通过id列表查询用户
- 通过Token查询用户id和openID
- 通过手机号查询用户id
- 发送短信
- 校验验证码
- 登录（手机号未注册则自动注册
  - 参数：手机号，验证码，设备信息，登录方式（电脑web和app、公众号和小程序），其中token时效暂定web为2小时，其他为90天
- 刷新token
- 微信
  - 登录
    - 参数：系统标识编码、code、appID
    - 返回值：token，可以根据token获取用户id和openID，如果userID=0，说明该用户未绑定手机号
  - 绑定手机号
    - 小程序直接绑定，给公众号需要验证码校验

#### 缓存
- 默认缓存（非Redis
  - 系统版本
  - RSA秘钥

---

>gRPC
 ```shell
go install google.golang.org/protobuf/cmd/remote-gen-go@v1.26
go install google.golang.org/grpc/cmd/remote-gen-go-grpc@v1.1

```