### Day06课上笔记

------



#### 内容回顾

##### 包(package)

包的定义-->`package`关键字,包名通常是和目录名一致，不能包含`-`

- 一个文件夹就是一个包
- 文件夹里面放的是`.go`文件

包的导入-->`import`

- 包导入路径是从`$GOPATH/src`
- 单行导入
- 多行导入
- 给导入的包起别名
- 匿名导入-->`sql`包导入时会讲这个
- Go不支持循环导入

包中`标识符`(变量名、函数名、结构体名、接口名、常量)可见性 -->标识符首字母大写表示对外可见

`init()`

- 包导入的时候会自动执行
- 一个包里面只有一个`init()`
- `init()`没有参数没有返回值也不能调用它
- 多个包的`init()`执行顺序
- 一般用于做一些初始化的操作

##### 接口(interface)

------

接口是一种类型，一种抽象的类型.

接口就是你要实现的方法的清单.

接口的定义

```go
type mover interface {
    方法的签名(参数)(返回值)
}
```

###### 接口的实现

实现了接口的所有方法就是实现了这个接口.

就实现了接口就可以当成这个接口类型的变量

###### 接口的变量

实现了一个变量，可以保存所有实现了我这个接口的类型的值.

###### 空接口

`interface{}`

接口中没有定义任何方法，也就是说任意类型都实现了空接口--> 任何类型都可以存到空接口变量中

- 作为函数参数-->`fmt.Println()`

- `map[string]interface{}`

###### 接口底层

- 动态类型
- 动态值

###### 类型断言

做类型断言的前提是一定要是一个接口类型的变量

`x.(T)`

![image-20210126161427794](C:\Users\97152\AppData\Roaming\Typora\typora-user-images\image-20210126161427794.png)

使用`switch`做类型断言

![image-20210126161441691](C:\Users\97152\AppData\Roaming\Typora\typora-user-images\image-20210126161441691.png)





##### 文件操作

------

###### 打开文件和关闭文件

![image-20210126163424536](C:\Users\97152\AppData\Roaming\Typora\typora-user-images\image-20210126163424536.png)

###### 读文件

`file.Read()`

Read方法定义如下：

```go
func (f *File) Read(b []byte) (n int, err error)
```

`bufio`

bufio是在file的基础上封装了一层API，支持更多的功能。

`ioutil`

`io/ioutil`包的`ReadFile`方法能够读取完整的文件，只需要将文件名作为参数传入

###### 写文件

`os.OpenFile()`函数能够以指定模式打开文件，从而实现文件写入相关功能。

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}
```

其中：

`name`：要打开的文件名 `flag`：打开文件的模式。 模式有以下几种：

| 模式          | 含义     |
| :------------ | :------- |
| `os.O_WRONLY` | 只写     |
| `os.O_CREATE` | 创建文件 |
| `os.O_RDONLY` | 只读     |
| `os.O_RDWR`   | 读写     |
| `os.O_TRUNC`  | 清空     |
| `os.O_APPEND` | 追加     |

`perm`：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。



#### 今日内容

------



#### 日志库

##### 需求分析

1、支持终端输出和写入文件的两种方式

2、区分日志级别

- Debug
- Trace
- Info
- Warning
- Error
- Fatal

3、日志需要支持开关控制，比如说开发阶段所有级别的日志都能输出，线上环境只有INFO级别往下的日志才能输出

4、日志格式：时间+行号+文件名+日志级别+日志内容

5、日志文件切割

- 按文件大小切割
- 按日期切割