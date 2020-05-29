package main

import (
	"fmt"
	//"context"
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"sync"

	//mconfig "log-server/config"
	//"strconv"

	//pools "github.com/jolestar/go-commons-pool/v2"
	//mdb "log-server/db/mongodb"
	"net/http"
	//"strings"
	//"time"
)

type worker struct {
	sync.Mutex
	host   string
	source chan interface{}
	//quit   chan struct{}
}

func (w *worker)clean()  {
	w.Lock()
	defer w.Unlock()
	for {
		select {
		case <-w.source:
			break
		default:
			return
		}
	}
}

func (w *worker) Start() {
	w.source = make(chan interface{})
	go func() {
		for {
			msg := <-w.source
			fmt.Println("==========>> ", w.name, msg)
		}
	}()
}

type threadSafeSlice struct {
	sync.Mutex
	workers []*worker
}

func (slice *threadSafeSlice) Push(w *worker) {
	slice.Lock()
	defer slice.Unlock()

	slice.workers = append(slice.workers, w)
}

func (slice *threadSafeSlice) Remove(w *worker) {
	slice.Lock()
	defer slice.Unlock()

	for i, worker := range slice.workers {
		if worker == w{
			slice.workers = append(slice.workers[:i], slice.workers[i+1:]...)
		}
	}
}

func (slice *threadSafeSlice) Iter(routine func(*worker)) {
	slice.Lock()
	defer slice.Unlock()

	for _, worker := range slice.workers {
		routine(worker)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Index (con *gin.Context) {
	showData := gin.H{
		//"moduleTags": mconfig.Cnf.Log.ModulesTags,
	}

	con.HTML(http.StatusOK, "index.html", showData)
}

//只获得数据
func GData(con *gin.Context) {
	c, err := upgrader.Upgrade(con.Writer, con.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	tss := &threadSafeSlice{}
	for {
		//读取网络数据

		mt, message, err := c.ReadMessage()
		if err != nil {
			//todo 错误处理
			continue
			//log.Println("read:", err)
			//break
		}
		//写入worker
		tss.Iter(func(w *worker) { w.source <- message })
		//log.Printf("recv: %s", message)
		//err = c.WriteMessage(mt, message)
		//if err != nil {
		//	log.Println("write:", err)
		//	break
		//}
	}
}

//只发送数据
//先接收具体要显示的服务器ip
func PData(c *gin.Context) {
	wc, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer wc.Close()
	//
	w := &worker{
		//host: string(message),
		//quit: globalQuit,
	}
	for{
		mt, message, err := wc.ReadMessage()
		if err != nil {
			//todo 错误处理
			return
			//continue
			//log.Println("read:", err)
			//break
		}
		w.host = string(message)
		w.clean()
	}

	//创建worker

	//读取worker里的数据输出
}

