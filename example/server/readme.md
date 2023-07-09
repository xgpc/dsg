docker run -d -v /usr/share/ca-certificates/:/etc/ssl/certs -p 4001:4001 -p 2380:2380 -p 2379:2379 \
--name etcd etcd /usr/local/bin/etcd \
-name etcd0 \
-advertise-client-urls http://192.168.3.3:2379,http://192.168.3.3:4001 \
-listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
-initial-advertise-peer-urls http://192.168.3.3:2380 \
-listen-peer-urls http://0.0.0.0:2380 \
-initial-cluster-token etcd-cluster-1 \
-initial-cluster etcd0=http://192.168.3.3:2380 \
-initial-cluster-state new



# 启动


```shell
# 启动本地etcd
etcd.exe --name bgdsgetcd
```

``` shell

  etcd --name bgpc --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380

  etcdctl put 111 222

  etcdctl get 111
  -> 111
  -> 222

## 获取版本信息
etcdctl version
## 获取所有键值对
etcdctl get --prefix ""
## 添加键值对
etcdctl put zhangsan hello
## 删除键值对
etcdctl del zhangsan
## 添加一个过期时间为20s的租约
etcdctl lease grant 20
## 获取所有租约
etcdctl lease list
## 添加键值对，并为该键指定租约
etcdctl put lisi world --lease="3f3574057fe0e61c"
## 查看某个租约的keepalived时间
etcdctl lease keep-alive 3f3574057fe0e61c
## 续租
etcdctl lease timetolive 3f3574057fe0e61c --keys
## 回收租约
etcdctl lease revoke 3f3574057fe0e61c





```

```shell
--data-dir #指定节点的数据存储目录，这些数据包括节点ID，集群ID，集群初始化配置，Snapshot文件，若未指定—wal-dir，还会存储WAL文件；
--wal-dir  #指定节点的was文件的存储目录，若指定了该参数，wal文件会和其他数据文件分开存储。
--name # 节点名称
--initial-advertise-peer-urls # 告知集群其他节点url.(对于集群内提供服务的url)
--listen-peer-urls # 监听URL，用于与其他节点通讯
--advertise-client-urls # 告知客户端url, 也就是服务的url(对外提供服务的utl)
--initial-cluster-token # 集群的ID
--initial-cluster # 集群中所有节点

--addr       #公布的 IP 地址和端口；默认为 127.0.0.1:2379
--bind-addr   #用于客户端连接的监听地址；默认为–addr 配置
--peers       #集群成员逗号分隔的列表；例如 127.0.0.1:2380,127.0.0.1:2381
--peer-addr   #集群服务通讯的公布的 IP 地址；默认为 127.0.0.1:2380
-peer-bind-addr  #集群服务通讯的监听地址；默认为-peer-addr 配置
--wal-dir         #指定节点的 wal 文件的存储目录，若指定了该参数 wal 文件会和其他数据文件分开存储
--listen-client-urls #监听 URL；用于与客户端通讯
--listen-peer-urls   #监听 URL；用于与其他节点通讯
--initial-advertise-peer-urls  #告知集群其他节点 URL
--advertise-client-urls  #告知客户端 URL
--initial-cluster-token  #集群的 ID
--initial-cluster        #集群中所有节点
--initial-cluster-state new  #表示从无到有搭建 etcd 集群
--discovery-srv  #用于 DNS 动态服务发现，指定 DNS SRV 域名
--discovery      #用于 etcd 动态发现，指定 etcd 发现服务的 URL
```




# 安装
```linux
ETCD_VER=v3.4.26

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
```

```docker
rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
  docker rmi gcr.io/etcd-development/etcd:v3.4.26 || true && \
  docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
  --name etcd-gcr-v3.4.26 \
  gcr.io/etcd-development/etcd:v3.4.26 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr

docker exec etcd-gcr-v3.4.26  /usr/local/bin/etcd --version
docker exec etcd-gcr-v3.4.26  /usr/local/bin/etcdctl version
docker exec etcd-gcr-v3.4.26  /usr/local/bin/etcdctl endpoint health
docker exec etcd-gcr-v3.4.26  /usr/local/bin/etcdctl put foo bar
docker exec etcd-gcr-v3.4.26  /usr/local/bin/etcdctl get foo
```