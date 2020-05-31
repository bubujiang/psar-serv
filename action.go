package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index (con *gin.Context) {
	showData := gin.H{
		//"moduleTags": mconfig.Cnf.Log.ModulesTags,
	}

	con.HTML(http.StatusOK, "index.html", showData)
}


