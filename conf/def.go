package conf

type Config struct {
	Ip string
	Port uint64
	Hosts map[string]string
}

type Op interface {
	Get() *Config //获得具体配置
}


