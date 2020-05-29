package conf

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type FileConf struct {

}

func (f *FileConf)Get() *Config {
	path := flag.String("c","conf.ini","配置文件")
	flag.Parse()

	cfg, err := ini.Load(path)
	if err != nil {
		//todo 错误处理
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	pc := &Config{}
	pc.Ip = cfg.Section("server").Key("ip").String()
	port,_ := cfg.Section("server").Key("port").Uint64()
	pc.Port = uint32(port)

	return pc
}