服务端打包

1. bootstrap
当模板相关的内容更新时
```
cd server/cmd/bootstrap
go-bindata -o asset/asset.go -pkg=asset appdata/... templates/...
```

2. server

linux:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o narvis -ldflags="-w -s"
macos:
go build -o narvis -ldflags="-w -s"


Docker镜像制作

Server, 替换version
```
docker buildx build --platform linux/amd64 \
-t jeffry688/narvis-server:latest -t jeffry688/narvis-server:all-0.0.12 \
-f server/Dockerfile . \
--build-arg GOFLAGS=-ldflags="-w -s" \
--push

```

Client
```
docker buildx build --platform linux/amd64 \
-t jeffry688/narvis-proxy:latest -t jeffry688/narvis-proxy:all-0.0.6 \
-f client/Dockerfile . \
--build-arg GOFLAGS=-ldflags="-w -s" \
--push
```

Bootstrap
```
docker buildx build --platform linux/amd64 \
-t jeffry688/narvis-bootstrap:latest -t jeffry688/narvis-bootstrap:all-0.0.8 \
-f server/cmd/bootstrap/Dockerfile . \
--build-arg GOFLAGS=-ldflags="-w -s" \
--push
```