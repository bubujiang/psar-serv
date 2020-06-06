package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index (con *gin.Context, hosts map[string]string) {
	k:=""
	for k = range hosts{ break }

	showData := gin.H{
		"hosts": hosts,
		"host": con.DefaultQuery("h",k),
	}

	con.HTML(http.StatusOK, "index.html", showData)
}


