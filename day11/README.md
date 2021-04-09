# Day11课上笔记

## 依赖管理 go module

如果使用是Go1.11和Go1.12的版本，需要手动开启go module支持：

```go
Windows：set GO111MODULE=on
Mac：export GO111MODULE=on
```

### goproxy

设置代理

```shell
export GOPROXY=https://goproxy.cn	// mac
SET GOPROXY=https://goproxy.cn	// windows
```

### go.mod文件

记录了当前项目依赖的第三方包信息和版本信息

第三方的依赖包都下载到了`GOPATH/pkg/mod`目录下

### go.sum文件

详细包名和版本信息

### 常用的命令

```bash
go mod init [包名]	// 初始化项目
go mod tidy			// 检查代码里的依赖去更新go.mod文件中的依赖
go get
go mod download
```



## context

如何优雅的控制子Goroutine退出

### 两个默认值

```go
context.Background()
context.TODO()
```

### 四个方法

```go
context.WithCancel(context.Background())
context.WithDeadline(context.Background(), time.Time)
context.WithTimeout(context.Background(), time.Duration)
context.WithValue(context.Background(), key, value)
```

