
# Day03上课总结

## 函数

#### 函数的定义

```go
func 函数名(参数1，参数2...)返回值{
    函数体
}
```

#### 函数进阶

高阶函数：函数可以作为参数，也可以作为返回值

```go
func f1(name string){
    fmt.Println("Hello",name)
}
func f2(f func(string)){		// 函数作为参数
    f()
}
func f3() func(int, int) int{	// 函数作为返回值
    return func(x, y int) int{
        return x + y
    }
}
func main(){
    f2(f1) // 调用函数f2,f1作为参数传给f2
    ret := f3()
    fmt.Printf("%T\n", ret)
    sum := ret(10, 20)
    fmt.Println(sum)
}
```

闭包：函数和其外部变量的引用

```go
func low(f func){
    f()
}
// 闭包
func bi(f func(string), name string) func(){
    return func{
        f(name)
    }
}
func main(){
    // 闭包示例
    fc := bi(f3, "xxx")
    low(fc)
}
```

defer：延迟调用，多用于处理资源释放

内置函数：`panic`和`recover`

```go
func f1(){
    defer func(){
        err := recover() // 收集当前的错误
        fmt.Println(err)     
    }()
    panic("犯了不可原谅的错误")
    fmt.Println("f1")
}
func f2(){
    fmt.Println("f2")
}
func mian(){
    f1()
    f2()
}
```



