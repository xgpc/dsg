package dsgEtcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Endpoints    []string `yaml:"endpoints"`
	DialTimeout  int      `yaml:"dial_timeout"`
	ServerRouter string   `yaml:"server_router"`
	Host         string   `yaml:"host"`
	Port         int      `yaml:"port"`
}

type Hand struct {
	conf Config
}

func main() {
	// 链接etcd
	config := clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 10 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	// 查询key
	get, err := client.Get(context.Background(), "1", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	for _, v := range get.Kvs {
		fmt.Printf("key:%s, value:%s\n", v.Key, v.Value)
	}

	// 注册服务
	err = registerService(client, "my-service", "192.168.31.245", 8080)
	if err != nil {
		panic(err)
	}

	// 发现服务
	for {
		services, err := discoverServices(client, "my-service")
		if err != nil {
			panic(err)
		}

		fmt.Println(services)
		time.Sleep(time.Second * 10)
	}

	select {}

	fmt.Println("exit")
}

// 服务结构体
type Service struct {
	Name    string
	Address string
	Port    int
}

// 发现服务
func discoverServices(client *clientv3.Client, name string) ([]Service, error) {
	key := fmt.Sprintf("/services/%s", name)

	resp, err := client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	var list []Service
	for _, kv := range resp.Kvs {
		address, port, err := parseServiceKey(string(kv.Key))
		if err != nil {
			fmt.Println(err)
			continue
		}
		node := Service{
			Name:    name,
			Address: address,
			Port:    port,
		}

		list = append(list, node)
	}

	return list, nil
}

// 解析key
func parseServiceKey(key string) (string, int, error) {
	var address string
	var port int

	index := strings.LastIndex(key, "/")

	split := strings.Split(key[index+1:], ":")
	if len(split) != 2 {
		fmt.Println(split, "大于2个")
		return "", 0, fmt.Errorf("key 解析失败:%s", key[index:])
	}

	address = split[0]

	atoi, err := strconv.Atoi(split[1])
	if err != nil {
		return "", 0, fmt.Errorf("strconv.Atoi(%s)", split[1])
	}
	port = atoi

	return address, port, nil
}

// 注册服务
func registerService(client *clientv3.Client, name string, address string, port int) error {
	key := fmt.Sprintf("/services/%s/%s:%d", name, address, port)

	value := ""

	// 创建租约
	lease, err := client.Grant(context.Background(), 10)
	if err != nil {
		panic(err)
	}

	// 注册服务
	_, err = client.Put(context.Background(), key, value, clientv3.WithLease(lease.ID))
	if err != nil {
		panic(err)
	}

	// 自动续租
	ch, err := client.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case e, ok := <-ch:
				if !ok {
					fmt.Println(e)
					return
				}
			}
		}
	}()

	return nil
}
