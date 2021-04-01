

# Day02上课总结

## 运算符

### 算术运算符

`+`  `-`  `*`  `/`

### 逻辑运算符

`&&` `||` `!`

### 位运算符

`>>` `<<` `|` `^` `&`

### 赋值运算符

`=` `+=`  `-=`

`++` 和`--`独立的语句，不属于赋值运算符

### 比较运算符

`<` `<=` `!=`

## 数组(Array)

`var ages [30]int`

`var names [30]string`

`var nums [30]int`

数组包含元素的类型和个数。元素的个数(数组的长度)属于数据类型的一部分。

数组是值类型

```go
func f1(a [3]int){
    // Go语言中函数传递的都是值(Ctrl+C Ctrl+V)
    a[1] =100	//此处修改的都是副本的值 
}

func main(){
    x := [3]int{1,2,3}
    fmt.Println(x)	// {1 2 3}
    f1(x)
    fmt.Println(x)  // {1 2 3}
}
```

数组的初始化：

```go
func main(){
    var ages [30]int	//声明了变量ages,它是[30]int类型
    ages = [30]int{1,2,3,4,5,6}	// 给变量ages赋值
    var ages2 = [...]int{1,2,3,4,90,100} // 自动计算数组长度
    var ages3 = [...]int{1: 100,99: 200} // 索引是1的位置值是100，索引是99的位置值是200
    fmt.Println(ages，ages2,ages3)
}
```

多维数组：

```go
var a1 [...][2]int	// [[1,2] [3,4] [5,6]]
a1 = [...][2]int{
    [2]int{1, 2},
    [2]int{3, 4},
    [2]int{5, 6},
}
// 多维数组只有最外层可以用...
```

## 切片(Slice)

切片的定义和初始化

切片不存值，它就像一个框，去底层数组框值。

```go
var s1 []int		// 没有分配内存 nil
s1 = []int{1,2,3}	// 初始化
s2 := make([]bool, 2, 4)	// make初始化 分配内存
s3 := make([]int,0, 4)
```

切片：指针、长度、容量

切片的扩容策略：

1、如果小于1024，那么就直接两倍

2、如果大于1024，就按照1.25倍去扩容

3、如果申请的容量大于两倍，那就直接扩容至新申请的容量

4、具体存储的值类型不同，扩容策略也有一定的不同。

append

```go
var s1 []int
s1 = append(s1,1)	//自动初始化切片 添加元素
```

## 指针

只需要记住两个符号：`&` 和 `*`  Go语言里面的指针只能读不能修改，不能修改指针变量指向的地址

```go
addr := "沙河"
addrP := &addr	// 查看内存地址
addrV := *addrP // 根据内存地址找值
```

## map

map存储的是键值对

```go
var m1 map[string]int			// 定义一个map
m1 = make(map[string]int, 10)	// make初始化map
m1["shasha"] = 100
fmt.Println("ji")				// 如果key不存在返回的是value对应的零值
```

```go
delete(m1, "lidad") // 删除的key 不存在什么都不干
```

```go
// 如果返回的是布尔值，我们通常用ok去接收
score, ok := m1["ji"]
if !ok {
    fmt.Println("没有ji这个人")
}else {
    fmt.Println("ji的分数是"，score)
}
```

