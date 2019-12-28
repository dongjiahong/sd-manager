package api

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"back-go/db"
)

var dataPath = "./data/"
var dataFile = "sd.db"
var dbFileReg *regexp.Regexp

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func init() {
	// sd.2006-01-02 15:03:04.db
	dbFileReg = regexp.MustCompile(`sd.\S+ \S+.db`)
}

func newResp(msg string, data interface{}) *Resp {
	return &Resp{
		Message: msg,
		Data:    data,
	}
}

func (r *Resp) toJson() []byte {
	if r == nil {
		return nil
	}
	b, _ := json.Marshal(r)
	return b
}

func ManagerCheckMiddleware(c *gin.Context) {
	// TODO
	au := c.GetHeader("Authorization")
	if len(au) == 0 {
	}
}

func Load(c *gin.Context) {

	token := c.Query("token")
	var user, pwd string

	if info, err := base64.StdEncoding.DecodeString(token); err != nil {
		log.Println("[Load] decode token err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	} else {
		if ss := strings.Split(string(info), "::"); len(ss) != 2 {
			log.Println("[Load] split wrong token:  ", token)
			c.Data(200, "application/json", newResp("split token err", nil).toJson())
			return
		} else {
			user, pwd = ss[0], ss[1]
		}
	}

	m, err := db.QueryUserWithName(user, pwd, dataPath+dataFile)

	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	c.Data(200, "application/json", newResp(m, nil).toJson())
	return
}

// 添加机器
func AddMachine(c *gin.Context) {
	var machine db.Machine
	if err := c.ShouldBindJSON(&machine); err != nil {
		log.Println("[AddMachine] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	if res, err := db.AddMachine(machine, dataPath+dataFile); err != nil {
		log.Println("[AddMachine] inert to db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return
}

// 添加账户
func AddAccount(c *gin.Context) {
	var acc db.Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Println("[AddAccount] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	if res, err := db.AddAccount(acc, dataPath+dataFile); err != nil {
		log.Println("[AddAccount] inert to db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return
}

// 删除账户
func DelAccount(c *gin.Context) {
	var acc db.Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Println("[DelAccount] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	if res, err := db.DelAccount(acc, dataPath+dataFile); err != nil {
		log.Println("[DelAccount] inert to db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return
}

// 更新账户
func EditAccount(c *gin.Context) {
	var ma db.ModifyAccount
	if err := c.ShouldBindJSON(&ma); err != nil {
		log.Println("[EditAccount] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	if err := db.EditAccount(ma, dataPath+dataFile); err != nil {
		log.Println("[EditAccount] update to db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", nil).toJson())
	}
	return
}

// 更新机器
func EditMachine(c *gin.Context) {
	var m db.Machine
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Println("[EditMachine] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	if res, err := db.UpdateMachine(m, dataPath+dataFile); err != nil {
		log.Println("[EditMachine] update to db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return
}

// 拉取所有的机器
func GetAllMachines(c *gin.Context) {}

// 删除机器
func DelMachine(c *gin.Context) {
	var m db.Machine
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Println("[DelMachine] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	if res, err := db.DelMachine(m, dataPath+dataFile); err != nil {
		log.Println("[DelAccount] inert to db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return

}

// 拉取所有信息
func GetAllInfo(c *gin.Context) {
	if res, err := db.QueryAllInfo(dataPath + dataFile); err != nil {
		log.Println("[GetAllInfo] get all info err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return
}

// 备份数据库
func Backup(c *gin.Context) {
	if res, err := db.Backup(dataPath + dataFile); err != nil {
		log.Println("[Backup] backup db err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", res).toJson())
	}
	return
}

func GetBackupFile(c *gin.Context) {
	dbfs, err := getBackupFileNames(dataPath)
	if err != nil {
		log.Println("[GetBackupFile] get names err: ", err)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
	} else {
		c.Data(200, "application/json", newResp("ok", dbfs).toJson())
	}
	return
}

// 获取备份文件
func getBackupFileNames(dbPath string) ([]string, error) {
	fileNames, err := filepath.Glob(filepath.Join(dbPath, "*"))
	if err != nil {
		return nil, err
	}

	dbfiles := make([]string, 0)
	for _, dbName := range fileNames {
		log.Println("======> dbName: ", dbName)
		if f := dbFileReg.FindString(dbName); len(f) > 0 {
			log.Println("======> f: ", f)
			dbfiles = append(dbfiles, f)
		}
	}
	return dbfiles, nil
}
