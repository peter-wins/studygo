package main

// logAgent 入口程序
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/peter-wins/studygo/logagent/conf"
	"github.com/peter-wins/studygo/logagent/kafka"
	"github.com/peter-wins/studygo/logagent/taillog"
	"github.com/spf13/viper"
	"time"
)

var (
	//cfg = new(conf.AppConf)
	cfg = new(conf.AppConf)
)

func run() {
	//1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.发送到Kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
func main() {
	//0、加载配置文件
	viper.SetConfigFile("conf/config.ini") // 指定配置文件路径
	err := viper.ReadInConfig()            // 读取配置信息
	if err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s\n", err))
	}
	// 将读取的配置信息保存到全局变量Conf
	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s\n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(cfg); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	//1、初始化Kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err%v\n", err)
		return
	}
	fmt.Println("Init kafka success.")
	//2、打开日志文件准备收集日志
	err = taillog.Init(cfg.PathConf.Logjam)
	if err != nil {
		fmt.Printf("Init taillog failed, err%v\n", err)
		return
	}
	fmt.Println("Init taillog success.")
	run()
}
