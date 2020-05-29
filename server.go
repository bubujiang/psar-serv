package main

import (
	"fmt"
	"os"
	"os/signal"
	"psar/serv/conf"
	"syscall"
)

type server struct {
	Ip string
	Port uint32
}

func (s *server)run(c *conf.Config)  {
	s.Ip = c.Ip
	s.Port = c.Port
	s._start()
}

func (s *server) _start() {
	//启动服务
	//监控信号
}

func (s *server) stop() {

}

func (s *server) load() {

}

