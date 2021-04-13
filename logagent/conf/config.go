package conf

type AppConf struct {
	KafkaConf `mapstructure:"kafka"`
	PathConf  `mapstructure:"path"`
}
type KafkaConf struct {
	Address string `mapstructure:"address"`
	Topic   string `mapstructure:"topic"`
}
type PathConf struct {
	Logjam string `mapstructure:"logjam"`
}
