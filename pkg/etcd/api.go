package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	client *clientv3.Client
	conf   Config
}

func New(conf Config) *Handler {
	config := clientv3.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: conf.DialTimeout * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	return &Handler{client: client, conf: conf}
}

func (p *Handler) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return p.client.Get(context.Background(), key, opts...)
}

func (p *Handler) Put(key, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return p.client.Put(context.Background(), key, value, opts...)
}

// Service 服务结构体
type Service struct {
	Name    string
	Address string
	Port    int
}

// DiscoverServices 发现服务
func (p *Handler) DiscoverServices(name string) ([]Service, error) {
	key := name

	resp, err := p.client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	var list []Service
	for _, kv := range resp.Kvs {
		serverName, address, port, err := parseServiceKey(string(kv.Key))
		if err != nil {
			fmt.Println(err)
			continue
		}
		node := Service{
			Name:    serverName,
			Address: address,
			Port:    port,
		}

		list = append(list, node)
	}

	return list, nil
}

// 解析key
func parseServiceKey(key string) (string, string, int, error) {
	var name, address string
	var port int

	index := strings.LastIndex(key, "/")

	name = key[0:index]

	split := strings.Split(key[index+1:], ":")
	if len(split) != 2 {
		fmt.Println(split, "大于2个")
		return "", "", 0, fmt.Errorf("key 解析失败:%s", key[index:])
	}

	address = split[0]

	atoi, err := strconv.Atoi(split[1])
	if err != nil {
		return "", "", 0, fmt.Errorf("strconv.Atoi(%s)", split[1])
	}
	port = atoi

	return name, address, port, nil
}

// RegisterServiceDefault  注册服务默认配置
func (p *Handler) RegisterServiceDefault() error {
	return p.RegisterService(p.conf.Name, p.conf.Address, p.conf.Port)
}

// RegisterService 注册服务
func (p *Handler) RegisterService(name string, address string, port int) error {
	key := fmt.Sprintf("%s/%s:%d", name, address, port)

	value := ""

	// 创建租约
	lease, err := p.client.Grant(context.Background(), 10)
	if err != nil {
		panic(err)
	}

	// 注册服务
	_, err = p.client.Put(context.Background(), key, value, clientv3.WithLease(lease.ID))
	if err != nil {
		panic(err)
	}

	// 自动续租
	ch, err := p.client.KeepAlive(context.Background(), lease.ID)
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
