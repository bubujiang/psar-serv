package main

import (
	"os"
	"os/signal"
	"psar/serv/conf"
	"syscall"
)

func main() {
	pserver := &server{}
	pserver.run(getConf())
	watch(pserver)
}

/**
 * 文件形式配置文件
 */
func getConf() *conf.Config {
	return (&conf.FileConf{}).Get()
}

/**
 * 监控信号
 */
func watch(s *server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		syscall.SIGINT,//退出
		syscall.SIGUSR1,//重新加载配置文件
	)

	select {
	case sig := <-sigs:
		//s.load()
		if sig == syscall.SIGINT {
			s.stop()
		} else if sig == syscall.SIGUSR1 {
			s.load()
		}
	}
}
