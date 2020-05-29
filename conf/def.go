package conf

type Config struct {
	Ip string
	Port uint64
}

type Op interface {
	Get() *Config //获得具体配置
}


