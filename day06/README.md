# Day06课上笔记

### Time

`2006-01-02 15:04:05.000`

### 时间类型

- `time.Time`：`time.Now()`

- 时间戳：

  `time.Now().Unix()`：1970.1.1到现在的秒数

  `time.Now().UnixNano()`：1970.1.1到现在的纳秒数

### 时间间隔类型

- `time.Duration`：时间间隔类型

  `time.Second`

```go
const (
	Nanosecond Duration = 1
    Microsecond			= 1000 * Nanosecond
    Millisecond			= 1000 * Microsecond
    Second				= 1000 * Millisecond
    Minute				= 60 * Second
    Hour				= 60 * Minute
)
```

### 时间操作

时间对象+/-一个时间间隔对象

```go
//now + 24小时
fmt.Println(now.Add(24 * time.Hour))
fmt.Println("------------------------------------")
//Sub两个时间相减
nextDay, err := time.Parse("2006-01-02", "2021-01-27")
if err != nil {
	fmt.Printf("parse time failed, err:%v\n", err)
	return
}
d := nextDay.Sub(now)
```

### 时间格式化

`fmt.Println(now.Format("2006-01-02"))`

### 定时器

```go
timer:=.tick
timer := time.Tick(time.Second) //定义一个1秒间隔的定时器
for i := range timer {
fmt.Println(i) //每秒都会执行的任务
```

