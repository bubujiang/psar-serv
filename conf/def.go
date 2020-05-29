package conf

type Config struct {
	Ip string
	Port uint32
}

type Op interface {
	Get() *Config //获得具体配置
}


