package fake

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"back-go/db"
)

type AllInfo struct {
	Accounts []db.Account `json:"accounts"`
	Agents   []db.Agent   `json:"agents"`
	Machines []db.Machine `json:"machines"`
}

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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
	log.Println("=====> data: ", string(b))
	return b
}

func Load(c *gin.Context) {
	token := c.Query("token")
	var user, pwd string

	log.Println("======> token: ", token)

	if info, err := base64.StdEncoding.DecodeString(token); err != nil {
		c.JSON(200, gin.H{
			"status": "decode err: " + err.Error(),
			"msg":    token,
		})
		return
	} else {
		if ss := strings.Split(string(info), "::"); len(ss) != 2 {
			c.JSON(200, gin.H{
				"status": "decode err: " + err.Error(),
				"msg":    token,
			})
			return
		} else {
			user, pwd = ss[0], ss[1]
		}
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"msg":    fmt.Sprintf("user: %s, pwd: %s", user, pwd),
	})
	return
}

func GetAllAccounts(c *gin.Context) {
	accounts := make([]db.Account, 0)
	accountsStr := `[{
		"id": "1",
		"account_no": "a101",
		"create_date": "2019-02-05",
		"agent_date": "2019-03-06",
		"end_date": "2019-04-05",
		"account_tip": "备注",
		"agent_id": "1",
		"machine_id": "2"
		},
		{
		"id": "2",
		"account_no": "a102",
		"create_date": "2019-03-05",
		"agent_date": "2019-04-06",
		"end_date": "2019-05-05",
		"account_tip": "备注",
		"agent_id": "2",
		"machine_id": "3"
		},
		{
		"id": "3",
		"account_no": "a103",
		"create_date": "2019-03-05",
		"agent_date": "2019-04-03",
		"end_date": "2019-05-02",
		"account_tip": "备注",
		"agent_id": "2",
		"machine_id": "4"
		}]`
	if err := json.Unmarshal([]byte(accountsStr), &accounts); err != nil {
		log.Println("======> getallaccounts: err: ", err)
	}
	log.Println("======> getallaccounts: ", accounts[1])

	c.Data(200, "application/json", newResp("ok", accounts).toJson())
}
func GetAllAgents(c *gin.Context) {
	agents := make([]db.Agent, 0)
	agentsStr := `[{
		"id": "1",
		"agent_name": "董帅",
		"agent_account": "dongshuai",
		"agent_password": "123455"
		},
		{
		"id": "2",
		"agent_name": "董明",
		"agent_account": "dongming",
		"agent_password": "123451"
		},
		{
		"id": "3",
		"agent_name": "董凤娟",
		"agent_account": "dongfengjuan",
		"agent_password": "1234577"
		}]`
	if err := json.Unmarshal([]byte(agentsStr), &agents); err != nil {
		log.Println("======> getallagents: err: ", err)
	}

	c.Data(200, "application/json", newResp("ok", agents).toJson())
}
func GetAllMachines(c *gin.Context) {
	machines := make([]db.Machine, 0)
	machinesStr := `[{
		"id": "1",
		"machine_no": "a101",
		"machine_ip": "192.168.12.21:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2019-06-12",
		"machine_tip": "备注",
		"account_id": "1"
		},
		{
		"id": "2",
		"machine_no": "a103",
		"machine_ip": "192.168.12.22:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2019-06-12",
		"machine_tip": "备注",
		"account_id": "2"
		},
		{
		"id": "3",
		"machine_no": "a102",
		"machine_ip": "192.168.12.23:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2019-06-12",
		"machine_tip": "备注",
		"account_id": "3"
		},
		{
		"id": "4",
		"machine_no": "a104",
		"machine_ip": "192.168.12.24:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2019-06-12",
		"machine_tip": "备注",
		"account_id": ""
		}]`
	if err := json.Unmarshal([]byte(machinesStr), &machines); err != nil {
		log.Println("======> getallmachines: err: ", err)
	}

	c.Data(200, "application/json", newResp("ok", machines).toJson())
}

func GetAllInfo(c *gin.Context) {
	allStr := `{
	"accounts": [{
		"id": "1",
		"account_no": "a101",
		"create_date": "2019-02-05",
		"agent_date": "2019-03-06",
		"end_date": "2019-04-05",
		"account_tip": "备注",
		"agent_id": "1",
		"machine_id": "2"
		},
		{
		"id": "2",
		"account_no": "a102",
		"create_date": "2019-03-05",
		"agent_date": "2019-04-06",
		"end_date": "2020-05-05",
		"account_tip": "备注",
		"agent_id": "2",
		"machine_id": "3"
		},
		{
		"id": "3",
		"account_no": "a103",
		"create_date": "2019-03-05",
		"agent_date": "2019-04-03",
		"end_date": "2019-05-02",
		"account_tip": "备注",
		"agent_id": "2",
		"machine_id": "4"
		},
		{
		"id": "4",
		"account_no": "a104",
		"create_date": "2019-03-05",
		"agent_date": "",
		"end_date": "",
		"account_tip": "备注",
		"agent_id": "1",
		"machine_id": ""
		}],
	"agents":[{
		"id": "1",
		"agent_name": "董帅",
		"agent_account": "dongshuai",
		"agent_password": "123455"
		},
		{
		"id": "2",
		"agent_name": "董明",
		"agent_account": "dongming",
		"agent_password": "123451"
		},
		{
		"id": "3",
		"agent_name": "董凤娟",
		"agent_account": "dongfengjuan",
		"agent_password": "1234577"
		}],
	"machines":[{
		"id": "1",
		"machine_no": "a101",
		"machine_ip": "192.168.12.21:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2020-06-12",
		"machine_tip": "备注",
		"account_id": "1"
		},
		{
		"id": "2",
		"machine_no": "a103",
		"machine_ip": "192.168.12.22:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2020-06-12",
		"machine_tip": "备注",
		"account_id": "2"
		},
		{
		"id": "3",
		"machine_no": "a102",
		"machine_ip": "192.168.12.23:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2019-06-12",
		"machine_tip": "备注",
		"account_id": "3"
		},
		{
		"id": "4",
		"machine_no": "a104",
		"machine_ip": "192.168.12.24:6666",
		"machine_password": "31231342",
		"machine_create_date": "2019-03-05",
		"machine_end_date": "2020-06-12",
		"machine_tip": "备注",
		"account_id": ""
		}]}`
	var allInfo AllInfo
	if err := json.Unmarshal([]byte(allStr), &allInfo); err != nil {
		log.Println("======> getallaccounts: err: ", err)
	}

	c.Data(200, "application/json", newResp("ok", allInfo).toJson())
}
