# dsg

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

# use AES and JWT

```go
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "time"
    "github.com/xgpc/dsg/v2"
)

func main() {
    // init
    //Read configuration file
    dsg.Load("config.yaml")
    dsg.Default(
        dsg.OptionAes(dsg.Conf.AesKey), // use aes
        dsg.OptionJwt(dsg.Conf.JwtKey), // use jwt
    )
    // Service code
    // aes
    aesStr := dsg.AESEnCode("test")

    deStr := dsg.AESDeCode(aesStr)
    fmt.Println(deStr)

    // jwt
    var userID uint32 = 1
    token := dsg.CreateToken(userID, time.Hour*24*7)
    fmt.Println(token)

    MapClaims := dsg.ParseToken(token)
    fmt.Println(MapClaims)

    if MapClaims.UserID == userID {
        fmt.Println("success")
    }
    
    // Use with iris
    api := iris.New()
    
    
    // api Route loading
    // dsg.Login is a middleware that verifies the token
    // dsg.Login After successful verification, the user information is stored in the context
    // dsg.NewBase(ctx) is used to get the user ID in the token
    api.Get("/test", dsg.Login, func(ctx iris.Context) {
        
        p := dsg.NewBase(ctx)
        if p.MyId() != UserID {
            panic("error")
        }
        fmt.Println(userID)
    })
    
    // api.Run(iris.Addr(":8080")) ....
}
```

## mysql cond
### Database operation
### cond use
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
    // conditionQuery
    // Paging query
    var list []User
    err = dsg.DB().Scopes(cond.Page(1, 10)).Find(&list).Error
    if err != nil {
        panic(err)
    }

    var ctx iris.Context
    // PageByQuery 
    // 从URLParamIntDefault中获取page和page_size
    dsg.DB().Scopes(cond.PageByQuery(ctx)).Find(&list)
    
    // PageByParams
    // 从GetIntDefault中获取page和page_size
    dsg.DB().Scopes(cond.PageByParams(ctx)).Find(&list)
    
    // cond.Eq id = 1
    dsg.DB().Scopes(cond.Eq("id", 1)).Find(&list)
    
    // cond.NotEq id != 1
    dsg.DB().Scopes(cond.NotEq("id", 1)).Find(&list)
    
    // cond.Gt id > 1
    dsg.DB().Scopes(cond.Gt("id", 1)).Find(&list)
    
    // cond.Gte id >= 1
    dsg.DB().Scopes(cond.Gte("id", 1)).Find(&list)
    
    // cond.Lt id < 1
    dsg.DB().Scopes(cond.Lt("id", 1)).Find(&list)
    
    // cond.Lte id <= 1
    dsg.DB().Scopes(cond.Lte("id", 1)).Find(&list)
    
    // cond.Like name like '%test%'
    dsg.DB().Scopes(cond.Like("name", "test")).Find(&list)
    
    // cond.Starting name like 'test%'
    dsg.DB().Scopes(cond.Starting("name", "test")).Find(&list)
    
    // cond.Ending name like '%test'
    dsg.DB().Scopes(cond.Ending("name", "test")).Find(&list)
    
    // cond.In id in (1,2,3)
    dsg.DB().Scopes(cond.In("id", []int{1, 2, 3})).Find(&list)
    
    // cond.NotIn  id not in (1,2,3)
    dsg.DB().Scopes(cond.NotIn("id", []int{1, 2, 3})).Find(&list)
    
    // cond.Between id between 1 and 10
    dsg.DB().Scopes(cond.Between("id", 1, 10)).Find(&list)
    
    
}

type User struct {
    ID   int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
    Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

func (User) TableName() string {
    return "user"
}      


```

## Error handling
```go
// 报错
// {
//	"code" : exce.CodeSysBusy
//	"msg" : "错误信息"
// }
exce.ThrowSys(exce.CodeSysBusy, "错误信息")

// 中间件 使用
// ExceptionLog 会捕获异常并返回
api := iris.Default()
api.Use(middleware.ExceptionLog)
```



>gRPC
 ```shell
go install google.golang.org/protobuf/cmd/remote-gen-go@v1.26
go install google.golang.org/grpc/cmd/remote-gen-go-grpc@v1.1

```

You can find my updates on https://github.com/xgpc/dsg/blob/main/README.md