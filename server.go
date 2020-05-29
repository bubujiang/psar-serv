package main

import (
	"github.com/gin-gonic/gin"
	"psar/serv/conf"
	"strconv"
)


type server struct {
	Ip string
	Port uint64
}

func (s *server)run(c *conf.Config)  {
	s.Ip = c.Ip
	s.Port = c.Port
	s._start()
}

func (s *server) _start() {
	r := gin.Default()
	r.Static("/dist", "html/dist")
	r.Static("/plugins", "html/plugins")
	r.GET("/", Index)
	r.GET("/gdata",GData)
	r.GET("/pdata",PData)
	r.Run(s.Ip+":"+strconv.FormatUint(s.Port,10))
}

func (s *server) stop() {

}

func (s *server) load() {

}

