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
