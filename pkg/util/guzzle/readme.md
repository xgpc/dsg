封装golang http请求

# 设置host发送请求 
```golang
...

client := NewClient(WithHost("http://localhost"))

response, err := client.Get("/api/test")
if err != nil {
    return err
}

fmt.Println(response.String())

data := map[string]interface{}
err = response.JSON(data)
if err != nil {
    return err
}

...
```

# 不设置host发送请求
```golang
...
client := NewClient()

response, err := client.Get("http://localhost/api/test")
if err != nil {
    return err
}

fmt.Println(response.String())

data := map[string]interface{}
err = response.JSON(data)
if err != nil {
    return err
}

...
```

# 设置query请求
```golang
...

client := NewClient()

response, err := client.RequestQuery(map[string][]string{
    "id": {"1", "2", "3"},
    "name": {"zhangsan"},
    }).Get("/api/test")
if err != nil {
    return err
}

fmt.Println(response.String())

data := map[string]interface{}
err = response.JSON(data)
if err != nil {
    return err
}

...
```

# 设置Json请求体
```golang
...

client := NewClient()

response, err := client.RequestJSON(map[string]interface{}{
        "id":   "1",
        "name": "zhangsan",
    }).Post("/api/test")    
if err != nil {
    return err
}

fmt.Println(response.String())

data := map[string]interface{}
err = response.JSON(data)
if err != nil {
    return err
}

...
```
