### 1、预装Gcc docker环境  还有docker-compose


### 2、启动项目所依赖的环境  根据实际情况配置是否需要国内代理

```
docker engine 配置 

{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "experimental": false,
  "features": {
    "buildkit": true
  },
  "registry-mirrors": [
    "https://docker.registry.cyou",
    "https://docker-cf.registry.cyou",
    "https://dockercf.jsdelivr.fyi",
    "https://docker.jsdelivr.fyi",
    "https://dockertest.jsdelivr.fyi",
    "https://mirror.aliyuncs.com",
    "https://dockerproxy.com",
    "https://mirror.baidubce.com",
    "https://docker.m.daocloud.io",
    "https://docker.nju.edu.cn",
    "https://docker.mirrors.sjtug.sjtu.edu.cn",
    "https://docker.mirrors.ustc.edu.cn",
    "https://mirror.iscas.ac.cn",
    "https://docker.rainbond.cc"
  ]
}



#### 步骤1: 部署redis集群
$ docker-compose -f docker-compose-redis.yml up -d

#### 步骤2：启动redis集群
$ docker exec -it redis-1 redis-cli --cluster create 172.20.99.11:6381 172.20.99.12:6382 172.20.99.13:6383 172.20.99.14:6384 172.20.99.15:6385 172.20.99.16:6386 --cluster-replicas 1

#### 步骤2：部署环境中间件
$ docker-compose -f docker-compose-env.yml up -d

#### 步骤3：部署etcd
$ docker-compose -f docker-compose-etcd.yml up -d

#### 步骤4: 部署Td
$ docker-compose -f docker-compose-tdengine.yml up -d

#### 步骤5: 部署asynqmon
$ docker-compose -f docker-compose-asynqmon.yml up -d

#### 步骤6: 部署asynqmon
$ docker-compose -f docker-compose-asynqmon.yml up -d





$ docker exec -it mysql mysql -uroot -p
##输入密码：PXDNA999999
$ use mysql;
$ update user set host='%' where user='root';
$ FLUSH PRIVILEGES;