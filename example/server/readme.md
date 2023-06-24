使用swagger时， 需要添加下方扩展
```shell
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/iris-contrib/swagger/v12
```

安装`swag`之后， windows环境下需要设置环境变量

在命令行中执行 `swag -v`能看到swag的版本， 说明swag安装成功

生产swagger文档时， 在`main.go`同级目录命令行中执行 `swag init`。 看到在`main.go`同级目录中出现docs文件夹， 说明swag文档生成成功
