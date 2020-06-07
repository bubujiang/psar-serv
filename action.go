package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Index (con *gin.Context, hosts map[string]string) {
	k:=""
	for k = range hosts{ break }
	host := con.DefaultQuery("h",k)
	hostUrl := hosts[host]

	showData := gin.H{
		"hosts": hosts,
		"host": host,
		"hostUrl": template.URL(hostUrl),
	}

	con.HTML(http.StatusOK, "index.html", showData)
}


