package etcd

import (
	"context"
	"fmt"
	"github.com/xgpc/dsg/v2/pkg/util"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type Handler struct {
	Client *clientv3.Client
	Conf   Config
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

	return &Handler{Client: client, Conf: conf}
}

func (p *Handler) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return p.Client.Get(context.Background(), key, opts...)
}

func (p *Handler) Put(key, value string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return p.Client.Put(context.Background(), key, value, opts...)
}

// Service 服务结构体
type Service struct {
	Name    string
	Address string
}

func (s *Service) GetUrl() string {
	url := "http://" + s.Address
	return url
}

// DiscoverServices 发现服务
func (p *Handler) DiscoverServices(name string) ([]Service, error) {
	key := "/service/" + name

	resp, err := p.Client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	var list []Service
	for _, kv := range resp.Kvs {
		serverName, address, err := parseServiceKey(string(kv.Key))
		if err != nil {
			fmt.Println(err)
			continue
		}
		node := Service{
			Name:    serverName,
			Address: address,
		}

		list = append(list, node)
	}

	return list, nil
}

// 解析key
func parseServiceKey(key string) (string, string, error) {
	var name, address string

	list := strings.Split(key, "/")

	if len(list) != 4 {
		return "", "", fmt.Errorf("%s 服务器解析失败", key)
	}

	name = list[2]
	address = list[3]

	return name, address, nil
}

// RegisterServiceDefault  注册服务默认配置
func (p *Handler) RegisterServiceDefault() error {
	return p.RegisterService(p.Conf.Name, p.Conf.Address, p.Conf.Port)
}

// RegisterService 注册服务
func (p *Handler) RegisterService(name string, address string, port int) error {
	key := fmt.Sprintf("/service/%s/%s:%d", name, address, port)

	value := util.RandomNumber(10)

	// 创建租约
	lease, err := p.Client.Grant(context.Background(), int64(p.Conf.DefLeaseSecond))
	if err != nil {
		panic(err)
	}

	// 注册服务
	_, err = p.Client.Put(context.Background(), key, value, clientv3.WithLease(lease.ID))
	if err != nil {
		panic(err)
	}

	// 自动续租
	ch, err := p.Client.KeepAlive(context.Background(), lease.ID)
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
