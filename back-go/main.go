package main

import (
	//"github.com/dongjiahong/gotools"
	"github.com/gin-gonic/gin"

	"back-go/api"
	"back-go/fake"
)

type Conf struct {
}

func main() {
	r := gin.Default()

	r.GET("/api/load", api.Load)

	r.GET("/api/fake/load", fake.Load)
	r.GET("/api/fake/getallaccounts", fake.GetAllAccounts)
	r.GET("/api/fake/getallagents", fake.GetAllAgents)
	r.GET("/api/fake/getallmachines", fake.GetAllMachines)
	r.GET("/api/fake/getallinfo", fake.GetAllInfo)

	r.GET("api/getallinfo", api.GetAllInfo)
	r.GET("api/backup", api.Backup)
	r.GET("api/getbackupfiles", api.GetBackupFile)
	r.POST("/api/add/agent", api.AddAgent)
	r.POST("/api/add/machine", api.AddMachine)
	r.POST("/api/add/account", api.AddAccount)
	r.POST("/api/del/account", api.DelAccount)
	r.POST("/api/del/machine", api.DelMachine)
	r.POST("/api/edit/account", api.EditAccount)
	r.POST("/api/edit/machine", api.EditMachine)

	r.Run(":8081")
}
