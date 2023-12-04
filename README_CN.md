



# 介绍
dsg是基于iris的一个开源工具包，主要用于快速开发，提供了一些常用的功能，如：日志、配置、数据库、缓存、gRPC等

# 安装
```shell
go get -u github.com/xgpc/dsg/v2
```

# 使用

```go

// 初始化
package main

import (
    "github.com/kataras/iris/v12"
    "github.com/xgpc/dsg/v2"
)

func main() {

    // 初始化
    // 读取配置文件
    dsg.Load("config.yaml")
    dsg.Default(
        dsg.OptionMysql(dsg.Conf.DBInfo), // 启动mysql
        dsg.OptionRedis(dsg.Conf.Redis),  // 启动redis
        dsg.OptionAes(dsg.Conf.AesKey),   // 启动aes
        dsg.OptionJwt(dsg.Conf.JwtKey),   // 启动jwt
        dsg.OptionEtcd(dsg.Conf.Etcd),    // 启动etcd
    )
    // 业务代码

    api := iris.New()
    
    // api 路由加载
    
    
    if dsg.Conf.TLS != "" {
        api.Run(iris.TLS(":8080", "server.crt", "server.key"))
    }else {
        api.Run(iris.Addr(":8080"))    
    }
    
}

```


## 使用mysql

```go
package main

import "github.com/xgpc/dsg/v2"

func main() {
    // 初始化
    // 读取配置文件
    dsg.Load("config.yaml")
    dsg.Default(
        dsg.OptionMysql(dsg.Conf.DBInfo), // 启动mysql
    )
    // 业务代码
    // 获取mysql连接
    // 查
    var user User
    err := dsg.DB().Model(user).First(&user, userID).Error
    if err != nil {
        panic(err)
    }

    // 增
    user := User{
        Name: "test",
    }
    err := dsg.DB().Model(user).Create(&user).Error
    if err != nil {
        panic(err)
    }

    // 删
    err := dsg.DB().Model(user).Delete(&user, userID).Error
    if err != nil {
        panic(err)
    }

    // 改
    err := dsg.DB().Model(user).Where("id", userID).Update(&user).Error
    if err != nil {
        panic(err)
    }
}

type User struct {
    ID   int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
    Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (User) TableName() string {
    return "user"
}

```


## 使用redis

```go
package main

import (
    "context"
    "github.com/xgpc/dsg/v2"
)

func main() {
    redisConn := dsg.Redis()
    background := context.Background()

    redisConn.Set(background, "test", "test", 0)
    redisConn.Get(background, "test")

    redisConn.HSet(background, "test", "test", "test")
    redisConn.HGet(background, "test", "test")

    redisConn.HMSet(background, "test", map[string]interface{}{"test": "test"})
    redisConn.HMGet(background, "test", "test")

    redisConn.HGetAll(background, "test")

    redisConn.LPush(background, "test", "test")
    redisConn.LRange(background, "test", 0, 1)

    redisConn.SAdd(background, "test", "test")
    redisConn.SMembers(background, "test")

    redisConn.ZAdd(background, "test", 1, "test")
    redisConn.ZRange(background, "test", 0, 1)
    
}
```

>gRPC
 ```shell
go install google.golang.org/protobuf/cmd/remote-gen-go@v1.26
go install google.golang.org/grpc/cmd/remote-gen-go-grpc@v1.1

```