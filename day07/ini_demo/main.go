package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// INI配置文件解析器

// MysqlConfig MySql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedislConfig Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `int:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") // 新建一个错误
		return
	}
	// 0.2 传进来的data参数必须是结构体指针类型(因为配置文件中各种键值对对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}
	// 1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b) // 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	// 2. 一行一行得到数据
	for _, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头就表示是节(section)
		if strings.HasPrefix(line, "[") {

		}
		// 2.3 如果不是[开头就是=分割的键值对

	}

	return
}
func main() {
	var mc MysqlConfig
	err := loadIni("./config.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed, err%v\n", err)
		return
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
