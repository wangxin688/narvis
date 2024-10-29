服务端打包

1. bootstrap
当模板相关的内容更新时
```
cd server/cmd/bootstrap
go-bindata -o asset/asset.go -pkg=asset appdata/... templates/...
```

2. server
