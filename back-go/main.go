package main

import (
	"github.com/gin-gonic/gin"

	"back-go/api"
)

func main() {
	r := gin.Default()
	//r.Use(api.ManagerCheckMiddleware)

	r.GET("/api/load", api.Load)

	r.GET("/api/getallinfo", api.GetAllInfo)
	r.GET("/api/backup", api.Backup)
	r.GET("/api/getbackupfiles", api.GetBackupFile)
	r.POST("/api/add/machine", api.AddMachine)
	r.POST("/api/add/account", api.AddAccount)
	r.POST("/api/del/account", api.DelAccount)
	r.POST("/api/del/machine", api.DelMachine)
	r.POST("/api/edit/account", api.EditAccount)
	r.POST("/api/edit/machine", api.EditMachine)

	r.Run(":8081")
}
