package main
// logAgent 入口程序
import (
	"fmt"
	"github.com/peter-wins/studygo/logagent/conf"
	"github.com/peter-wins/studygo/logagent/kafka"
	"github.com/peter-wins/studygo/logagent/taillog"
	"gopkg.in/ini.v1"
	"time"
)
var (
	cfg = new(conf.AppConf)
)
func run (){
	//1.读取日志
	for {
		select {
		case line := <- taillog.ReadChan():
			//2.发送到Kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
func main(){
	//0、加载配置文件
	err := ini.MapTo(cfg,"conf/config.ini")
	if err !=nil {
		fmt.Printf("load ini failed, err%v\n", err)
		return
	}
	fmt.Println("加载配置文件成功！")
	//1、初始化Kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil{
		fmt.Printf("init Kafka failed, err%v\n", err)
		return
	}
	fmt.Println("Init kafka success.")
	//2、打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil{
		fmt.Printf("Init taillog failed, err%v\n", err)
		return
	}
	fmt.Println("Init taillog success.")
	run()
}