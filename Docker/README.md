## 本地编写好 Dockerfile 文件执行下面命令构建docker镜像
```bash
# docker build <镜像名称>:<版本号> .
docker image build -t whatismyip:latest .
```
## 构建好的镜像通过下面命令手动启动镜像
```bash
docker run -d --name whatismyip_docker --restart=always -p 80:8080 gongyanhui66/whatismyip:2.5.4
# gongyanhui66/whatismyip:2.5.4 是我已经推送到镜像仓库的镜像 
# 本地刚构建的镜像应该为 whatismyip:latest
```

## 通过下述命令登录docker将构建好的镜像推送到自己的 dockerHub
```bash
# 登录到自己的docker账号 需要服务器能够跨境访问海外DockerHub
docker login
# 给镜像打上tag
docker image tag whatismyip:latest whatismyip:2.5.4
# 推送镜像 需要指定推送的仓库名 gongyanhui66
docker image push gongyanhui66/whatismyip:2.5.4
```

## 安装 docker compose
```bash
mkdir -p ~/.docker/cli-plugins/
curl -SL https://github.com/docker/compose/releases/download/v2.26.1/docker-compose-linux-x86_64 -o ~/.docker/cli-plugins/docker-compose
chmod +x ~/.docker/cli-plugins/docker-compose
# 验证安装
docker compose version
# Docker Compose version v2.26.1
# 此版本支持 docker compose 文件 version: '3.9'
# 此版本命令 docker compose 旧版本为 docker-compose
```

## 编写好 docker-compose 文件后通过下述命令启动
```bash
docker compose up &
# '&' 表示后台运行 停止容器执行下述命令 把up换成down
docker compose down
```