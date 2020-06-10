package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"psar/serv/conf"
	"strconv"
	"time"
)


type server struct {
	Ip string
	Port uint64
	wserv *http.Server
	Hosts map[string]string
}

func (s *server)run(c *conf.Config)  {
	s.Ip = c.Ip
	s.Port = c.Port
	s.Hosts = c.Hosts
	//s.wserv = &http.Server{}
	s._start()
}

func (s *server) _start() {
	r := gin.Default()
	r.Static("/dist", "html/dist")
	r.Static("/plugins", "html/plugins")
	//r.LoadHTMLGlob("html/*")
	r.LoadHTMLFiles("html/index.html")
	r.GET("/", func(c *gin.Context) {
		Index(c,s.Hosts)
	})

	//s.wserv.Addr = s.Ip+":"+strconv.FormatUint(s.Port,10)
	//s.wserv.Handler = r
	s.wserv = &http.Server{
		Addr:    s.Ip+":"+strconv.FormatUint(s.Port,10),
		Handler: r,
	}

	//if err := s.wserv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//	log.Fatalf("listen: %s\n", err)
	//}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.wserv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (s *server) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.wserv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func (s *server) load() {
	s.stop()
	s._start()
}

