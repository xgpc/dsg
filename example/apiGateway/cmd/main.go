package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	etcdEndpoints = "localhost:2379" // etcd的地址
	serviceKey    = "/services/my-service"
)

var (
	localIP string // 本地服务器的IP地址
)

func main() {
	// 获取本地服务器的IP地址
	localIP = getLocalIP()
	if localIP == "" {
		log.Fatal("Failed to get local IP address")
	}

	// 创建etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdEndpoints},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 注册服务
	if err := registerService(cli); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Service registered.")

	// 创建HTTP服务器
	router := gin.Default()
	router.GET("/hello", handleHello)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// 监听系统信号，用于优雅退出
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	// 等待系统信号
	<-stopCh

	// 取消服务的租约
	if err := revokeLease(cli); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Service unregistered.")

	// 关闭HTTP服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

// 注册服务
func registerService(cli *clientv3.Client) error {
	// 创建租约
	leaseResp, err := cli.Grant(context.Background(), 5) // 租约时间设为5秒
	if err != nil {
		return err
	}

	// 注册服务，将服务的IP和端口存储到etcd中
	_, err = cli.Put(context.Background(), serviceKey, localIP+":8080", clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	return nil
}

// 刷新服务的租约
func keepAlive(cli *clientv3.Client) {
	// 创建一个周期性的定时器，每隔一段时间刷新一次租约
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		_, err := cli.KeepAliveOnce(context.Background(), clientv3.LeaseID(cli.Lease()))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 取消服务的租约
func revokeLease(cli *clientv3.Client) error {
	_, err := cli.Revoke(context.Background(), cli.Lease())
	if err != nil {
		return err
	}
	return nil
}

// 处理Hello请求
func handleHello(c *gin.Context) {
	// 查询服务的IP和端口
	resp, err := cli.Get(context.Background(), serviceKey)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历所有服务实例，检查是否本地服务器或远程服务器
	for _, kv := range resp.Kvs {
		if string(kv.Value) == localIP+":8080" {
			// 本地服务器，直接处理请求
			c.String(http.StatusOK, "Hello from local server")
			return
		}
	}

	// 远程服务器，通过外网调用
	c.String(http.StatusOK, "Hello from remote server")
}

// 获取本地服务器的IP地址
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return ""
}
