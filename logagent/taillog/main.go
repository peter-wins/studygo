package taillog
// 专门从日志文件收集日志的模块
import (
	"fmt"
	"github.com/hpcloud/tail"
)
var (
	tailObj *tail.Tail
)
func Init(fileName string)(err error){
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 //　是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v\n", err)
		return
	}
	return
}
func ReadChan()<-chan *tail.Line{
	return tailObj.Lines
}