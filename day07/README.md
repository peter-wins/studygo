## Day07课上笔记

## 内容回顾

### 并发之goroutine

#### 并发和并行的区别

#### goroutine的启动

将要并发执行的任务包装成一个函数，调用函数的时候前面加上`go`关键字，就能够开启一个goroutine去执行该函数的任务。

goroutine对应的函数执行完，该goroutine就结束了。

程序启动的时候就会自动创建一个goroutine去执行main函数

main函数结束了，那么程序就结束了，由该程序启动的所有其他goroutine也都结束了。

#### goroutine的本质

goroutine的调度模型：`GMP`

`m:n` 把m个goroutine分配给n个操作系统线程

#### goroutine与操作系统线程(OS线程)的区别

goroutine是用户态的线程，比内核态的线程更轻量级，初始时只占用2KB的栈空间。

可以轻松开启数十万的goroutine也不会崩内存。

#### runtime.GOMAXPROCS

Go1.5之后默认就是操作系统的逻辑核心数，默认跑满cpu。

例如：`runtime.GOMAXPROCS(1)`  表示只占用1核

#### work pool模式

开启一定数量的goroutine去干活

#### sync.WaitGroup

`var wg sync.WaitGroup`

- wg.Add(1)：计数器+1
- wg.Done()：计数器-1
- wg.Wait()：等

### channel

#### 为什么需要channel？

通过channel实现多个goroutine之间的通信。

`CSP`:通过通信来共享内存

channel是一种类型，一种引用类型， make函数初始化之后才能用(slice map channel)

#### channel的声明

`var ch chan 元素类型`

#### channel的初始化

`ch = make(chan 元素类型，[缓冲区大小])`

#### channel的操作

- 发送：`ch <- 100`
- 接受：`x := <- ch`
- 关闭：`close(ch)`

#### 带缓冲区的通道和不带缓冲区的通道：

快递员送快递的示例：有缓冲区就是有快递柜

从通道中取值：

#### 单项通道：

通常是用作函数的参数，只读通道`<-chan int`和只写通道`chan<- int`

#### 通道的各种考虑情况

![channel异常总结](https://www.liwenzhou.com/images/Go/concurrence/channel01.png)



### select

同一时刻有多个通道要操作的场景，使用`select`.

使用`select`语句能提高代码的可读性。

- 可以处理一个或者多个channel的发送/接受操作
- 如果多个`case`同时满足，`select`会随机选择一个
- 对于没有`case`的`select{}`会一直等待，可用于阻塞main函数