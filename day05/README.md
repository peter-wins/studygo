# Day05上课回顾

### 包(package)

包的定义-->`package`关键字,包名通常是和目录名一致，不能包含`-`

- 一个文件夹就是一个包
- 文件夹里面放的都是`.go`文件

包的导入-->`import`

- 单行导入
- 多行导入
- 给导入的包起别名
- 匿名导入
- 包导入路径是从`$GOPATH/SRC`后面的路径开始写起
- Go不支持循环导入

包中`标识符`(变量名 函数名 结构体名 接口名 常量..)可见性 -->标识符首字母大写表示对外可见

`init()`

- 包导入的是时候会自动执行
- 一个包里面只有一个`init()`
- `init()`没有参数也没有返回值也不能调用它
- 多个包的`init`执行顺序
- 一般用于做一些初始化的操作

### 接口

接口是一种类型，一种抽象的类型。接口就是你要实现的方法的清单

#### 接口的定义

```go
type mover interface {
    方法的签名(参数)(返回值)
}
```

#### 接口的实现

实现了接口的所有方法就实现了接口。

实现了接口就可以当成这个接口类型的变量

#### 接口的变量

实现了一个万能的变量，可以保存所有实现了我这个接口的类型的值。通常是作为函数的参数出现

#### 空接口

`interface{}`:空接口

接口中没有定义任何方法，也就是说任意类型都实现了空接口，任何类型都可以存到空接口变量中

- 作为函数参数-->`fmt.Println()`

- ```go
  map[string]interface{}
  ```

#### 接口底层

- 动态类型
- 动态值

#### 类型断言

做类型断言的前提是一定要是一个接口类型的变量。

`X.(T)`

```go
func main() {
    var a interface{}	// 定义一个空接口变量
    a = 100
    // 如何判断a保存的值的具体类型是什么
    // 类型断言
    v, ok := a.(int8)
    if ok {
        fmt.Println("猜对了，a是int8", v)
    }else {
        fmt.Println("猜错了")
    }
}
```

使用`switch`来做类型断言

```go
switch v2 := a.(type) {
case int8:
    fmt.Println("int8", v2)
case int16:
    fmt.Println("int16", v2)
case string:
    fmt.Println("string", v2)
case int:
    fmt.Println("int", v2)
default:
    fmt.Println("滚~")
}
```

### 文件操作

#### [打开文件和关闭文件](https://github.com/peter-wins/studygo/blob/master/day05/12file_open/main.go)

```go
os.Open("./main.go")
```

```go
defer fileObj.Close()
```



#### 读文件

#### 写文件