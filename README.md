



# recommend
dsg is an open source toolkit based on iris, which is mainly used for rapid development and provides some common functions such as: logging, configuration, database, cache, g RPC, etc

# install
```shell
go get -u github.com/xgpc/dsg/v2
```

# use
```go

//initialize
package main

import (
    "github.com/kataras/iris/v12"
    "github.com/xgpc/dsg/v2"
)

func main() {

    
    // Read the configuration file
    dsg.Load("config.yaml")
    dsg.Default(
        dsg.OptionMysql(dsg.Conf.DBInfo), // start mysql
        dsg.OptionRedis(dsg.Conf.Redis),  // start redis
        dsg.OptionAes(dsg.Conf.AesKey),   // start aes
        dsg.OptionJwt(dsg.Conf.JwtKey),   // start jwt
        dsg.OptionEtcd(dsg.Conf.Etcd),    // start etcd
    )
    // Service code

    api := iris.New()
    
    // api Route loading
    
    
    if dsg.Conf.TLS != "" {
        api.Run(iris.TLS(":8080", "server.crt", "server.key"))
    }else {
        api.Run(iris.Addr(":8080"))    
    }
    
}

```


## use mysql

```go
package main

import "github.com/xgpc/dsg/v2"

func main() {
    // init
    //Read configuration file
    dsg.Load("config.yaml")
    dsg.Default(
        dsg.OptionMysql(dsg.Conf.DBInfo), // use mysql
    )
    // Service code
    // select
    var user User
    err := dsg.DB().Model(user).First(&user, userID).Error
    if err != nil {
        panic(err)
    }

    // install
    user := User{
        Name: "test",
    }
    err := dsg.DB().Model(user).Create(&user).Error
    if err != nil {
        panic(err)
    }

    // delete
    err := dsg.DB().Model(user).Delete(&user, userID).Error
    if err != nil {
        panic(err)
    }

    // update
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


## use redis

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