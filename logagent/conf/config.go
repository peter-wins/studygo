package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	PathConf  `int:"path"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type PathConf struct {
	Logjam string `ini:"logjam"`
}
