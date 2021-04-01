# Day04课上笔记

## 结构体

基本的数据类型：表示现实中的物件有局限性.(编程是用代码解决现实生活中的问题)

```go
var name = "保德路"
```

结构体定义：结构体是一种数据类型，一种我们自己造的可以保存多个维度的数据类型.

```go
type person struct{
    name string
    id int64
}
```

匿名结构体 ：多用于临时场景

```go
var a = struct{
    x int
    y int
}{10,20}
```

结构体的初始化

```go
// 定义一个person的结构体类型
type person struct{
    name string
    age int
}
// 结构体实例化1
var p1 person
p1.name = "保德路"
p1.age = 19
// 结构体实例化2
p2 := person{"保德路"，18}
```

## 构造函数

```go
// 构造结构体变量的函数，返回值是对应的结构体类型
func newPerson(n string, i int) person{
    return person{
        name: n,
        age: i,
    }
}
// 调用构造函数生成person类型变量
p3 := newPerson("nazha", 18)

```



## 方法和接收者

## 结构体的嵌套

## 结构体的匿名字段

## 序列化和发序列化

## JSON序列化与反序列化

