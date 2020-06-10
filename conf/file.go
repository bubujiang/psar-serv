package conf

import (
	"flag"
	"gopkg.in/ini.v1"
)

type FileConf struct {

}

func (f *FileConf)Get() *Config {
	path := flag.String("c","conf.ini","配置文件")
	flag.Parse()

	cfg, err := ini.Load(*path)
	if err != nil {
		panic(err)
	}

	pc := &Config{}
	pc.Ip = cfg.Section("server").Key("ip").String()
	pc.Hosts = cfg.Section("host").KeysHash()
	pc.Port,err = cfg.Section("server").Key("port").Uint64()
	if err != nil {
		panic(err)
	}

	return pc
}