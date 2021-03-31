# Day01课上回顾

## GO安装

`$GOPATH`: 你写Go代码的工作区，保存你的Go代码的。

`go env` : 

GOROOT=c:\go	安装路径

GOPATH=D:\work\go	项目路径

`GOPATH/bin` 添加到环境变量：`go install`命令会把生成的二进制文件拷贝到`GOPATH/bin`

## Go命令

`go build`：编译Go程序

`go build -o "xxx.exe"`:指定编译后的文件名

`go run main.go`: 像执行脚本一样执行main.go

`go install`:先编译后执行(拷贝到GOPATH/bin下)

## Go语言文件基础语法

存放Go源代码的文件名后缀名是`.go`

文件第一行：`package`关键字声明包名

如果要编译可执行文件，必须要有mian包和main函数(入口函数)

```go
// 单行注释

/*
多行注释
*/
```

Go语言函数外的语句必须以关键字开头

函数内部定义的变量必须使用

## 变量

三种声明方式：

1、`var name string`

2、`var name = "hello"`

3、函数内部专属：`name := 'h'`

匿名变量(哑元变量)：

当有些数据必须用变量接受但是又不使用它时，就可以用`_`来接受这个值。

## 常量

`const PI = 3.1415926`

`const  UserNotExisErr = 10000`

iota：实现枚举

三个要点：

1、`iota`在`const`关键字出现时将被重置为0

2、`const`中每新增一行常量声明，`iota`累加1

## 流程控制

### if判断

``` go
var age = 19
    if age > 18 {
        fmt.Println("成年了")
    }else if age > 7 {
        fmt.Println("上小学")
    }else {
        fmt.Println("最快乐的时光")
    }
```

### for循环

for循环各种写法

``` go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

``` go
var i = 0
for ; i < 10; i++{
    fmt.Println(i)
}
```

```go
var j = 0
for j < 10 {
    fmt.Println(j)
    j++
}
```

```go
for {
    fmt.Println("无限循环")
}
```

```go
s := "hello"
fmt.Println(len(s)) // 5
for i, v := range s {
    fmt.Println(i,v)
}
```

