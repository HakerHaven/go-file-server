# go-file-server 文件上传服务

## 快速构建

```shell
go mod init
go mod tidy
# 启动
go run cmd/main.go -c config/config.toml
# 查看项目启动使用的配置文件
go run cmd/main.go -c config/config.toml -p
```
