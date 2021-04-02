

# Day04课上笔记

### 自定义类型和类型别名

```go
type MyInt int		// 自定义类型
type newInt = int	// 类型别名
```

类型别名只在代码编写过程中有效，编译完之后就不存在，内置的`byte`和`rune`都属于类型别名。

### 结构体

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

### 构造函数

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

### 方法和接收者

方法是有接收者的函数，接收者指的是哪个类型的变量可以调用函数，结构体是值类型。

```go
// 指定了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream(){
    fmt.Println("%的梦想是学好Go语言"，p.name)
}
// 调用
p1.dream()
```

指针接收者
1、需要修改结构体变量的值时要使用指针接收者
2、结构体本身比较大，拷贝的内存开销比较大时也要使用指针接收者
3、保持一致性，如果有一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者

```go
func (p *person) guonian(){
    p.age++
}
```

### 结构体的嵌套

```go
type addr struct{
    province string
    city string
}
type student struct{
    name string
    address addr	// 嵌套别的结构体
}
```

### 结构体的匿名字段

```go
type student struct{
    name string
    addr	// 匿名嵌套别的结构体
}
```



### 序列化和反序列化

要注意的问题：

1、结构体内部的字段首字母要大写，不大写别人是访问不到的

```go
// 序列化
type point struct{
    X int `json: "n"`
    Y int `json: "m"`
}
p1 := point{100, 200}
b, err := json.Marshal(p1)
if err != nil {
    fmt.Printf("marshal failed, err:%v\n", err)
}
fmt.Println(string(b))
```

2、反序列化时要传递指针

```go
// 反序列化 由字符串-->Go中的结构体变量
str1 := `{"n":10, "m":20}`
var p2 point	//造一个结构体变量，准备接受反序列化的值
err = json.Unmarshal([]byte(str1), &p2)
if err != nil {
    fmt.Printf("unmarshal failed,err:%v\n", err)
}
fmt.Println(p2)
```

