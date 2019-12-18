package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var dataPath = "./data/"
var dataFile = "sd.db"

type AllInfo struct {
	Accounts []Account `json:"accounts"`
	Agents   []Agent   `json:"agents"`
	Machines []Machine `json:"machines"`
}

// 账户信息
type Account struct {
	AccountId  string `json:"id"`          // 账户id--自增长
	AccountNo  string `json:"account_no"`  // 账户代号
	CreateDate string `json:"create_date"` // 创建日期: 2019-02-05
	AgentDate  string `json:"agent_date"`  // 授权日期: 2019-03-06
	EndDate    string `json:"end_date"`    // 截止日期: 2019-04-05

	AgentId   string `json:"agent_id"`   // 代理的id，闲置账户不绑定
	MachineId string `json:"machine_id"` // 机器id,闲置账户可以不绑定
}

// 修改账户结构体
type ModifyAccount struct {
	Account
	DstMachineId string `json:"dst_machine_id"`
	DstAgentId   string `json:"dst_agent_id"`
	Ext          string `json:"ext"` // 保留字段用来添加额外信息
}

// 代理信息
type Agent struct {
	AgentId       string `json:"id"`             // 代理的id -- 自增长
	AgentName     string `json:"agent_name"`     // 代理的昵称： 董帅
	AgentAccount  string `json:"agent_account"`  // 代理的账户： dongshuai
	AgentPassword string `json:"agent_password"` // 代理的密码: 1234567
}

// 机器信息
type Machine struct {
	MachineId         string `json:"id"`                  // 自增id
	AccountId         string `json:"account_id"`          // **绑定的账户，0表示闲置机器**
	MachineNo         string `json:"machine_no"`          // 机器代号: a109
	MachineIP         string `json:"machine_ip"`          // 机器的ip
	MachinePassword   string `json:"machine_password"`    // 机器的密码
	MachineCreateDate string `json:"machine_create_date"` // 机器创建时间
	MachineEndDate    string `json:"machine_end_date"`    // 机器的到期时间
}

// 查询所有代理
func QueryUsers(dbFile string) ([]Agent, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("[QueryUsers] Open db file: ", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from agent")
	if err != nil {
		log.Println("[QueryUsers] query agent table: ", err)
		return nil, err
	}
	defer rows.Close()

	as := make([]Agent, 0)
	for rows.Next() {
		var a Agent
		rows.Scan(&a.AgentId, &a.AgentName, &a.AgentAccount, &a.AgentPassword)
		as = append(as, a)
	}

	return as, nil
}

// 查询某个代理
func QueryUserWithName(name, password, dbFile string) string {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("[QueryUserWithName] open db file: ", err)
		return err.Error()
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("select * from agent where account='%s';", name))
	if err != nil {
		log.Println("[QueryUserWithName] open db file: ", err)
		return err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		var a Agent
		rows.Scan(&a.AgentId, &a.AgentName, &a.AgentAccount, &a.AgentPassword)

		if a.AgentName == name && a.AgentPassword == password {
			return "ok"
		}
	}

	return "没有查询到该用户"
}

func QueryAllAccount(dbFile string) ([]Account, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("[QueryAllAccount] open db file: ", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from accounts;")
	if err != nil {
		log.Println("[QueryAllAccount] open db file: ", err)
		return nil, err
	}
	defer rows.Close()

	as := make([]Account, 0)
	for rows.Next() {
		var a Account
		rows.Scan(&a.AccountId, &a.AccountNo, &a.CreateDate, &a.AgentDate, &a.EndDate, &a.AgentId, &a.MachineId)

		as = append(as, a)
	}
	return as, nil
}

func QueryAllAgent(dbFile string) ([]Agent, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("[QueryAllAgent] open db file: ", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from agents;")
	if err != nil {
		log.Println("[QueryAllAgent] query sql err: ", err)
		return nil, err
	}
	defer rows.Close()

	as := make([]Agent, 0)
	for rows.Next() {
		var a Agent
		rows.Scan(&a.AgentId, &a.AgentName, &a.AgentAccount, &a.AgentPassword)

		as = append(as, a)
	}
	return as, nil
}

func AddAgent(a Agent, dbFile string) (*Agent, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("insert into agents(agentname, agentaccount, agentpassword) values('%s','%s','%s');",
		a.AgentName, a.AgentAccount, a.AgentPassword)

	if res, err := db.Exec(sqlStr); err != nil {
		return nil, err
	} else {
		id, _ := res.LastInsertId()
		a.AgentId = strconv.Itoa(int(id))
	}

	return &a, err
}

func AddMachine(m Machine, dbFile string) (*Machine, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	m.MachineCreateDate = formatDate(m.MachineCreateDate)
	m.MachineEndDate = formatDate(m.MachineEndDate)

	sqlStr := fmt.Sprintf("insert into machines(machineno, machineip, machinepassword, machinecreatedate, machineenddate) values('%s','%s','%s','%s','%s');",
		m.MachineNo, m.MachineIP, m.MachinePassword, m.MachineCreateDate, m.MachineEndDate)

	if res, err := db.Exec(sqlStr); err != nil {
		return nil, err
	} else {
		id, _ := res.LastInsertId()
		m.MachineId = strconv.Itoa(int(id))
	}

	return &m, err
}

func AddAccount(a Account, dbFile string) (*Account, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	a.CreateDate = time.Now().Format("2006-01-02")
	a.AgentDate = formatDate(a.AgentDate)
	a.EndDate = formatDate(a.EndDate)

	sqlStr := fmt.Sprintf("insert into accounts(accountno, createdate, agentid, agentdate, enddate, machineid) values('%s','%s','%s','%s','%s','%s');",
		a.AccountNo, a.CreateDate, a.AgentId, a.AgentDate, a.EndDate, a.MachineId)

	if res, err := db.Exec(sqlStr); err != nil {
		return nil, err
	} else {
		id, _ := res.LastInsertId()
		a.AccountId = strconv.Itoa(int(id))
	}
	// 在创建时就赋予了机器
	if len(a.MachineId) > 0 {
		if err := UpdateMachineAccountId(a.MachineId, a.AccountId, dbFile); err != nil {
			return nil, errors.New("[AddAccount] update machine's account id err: " + err.Error())
		}
	}

	return &a, err

}

func QueryAllMachine(dbFile string) ([]Machine, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("[QueryAllMachine] open db file: ", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from machines;")
	if err != nil {
		log.Println("[QueryAllMachine] open db file: ", err)
		return nil, err
	}
	defer rows.Close()

	ms := make([]Machine, 0)
	for rows.Next() {
		var m Machine
		rows.Scan(&m.MachineId, &m.MachineNo, &m.MachineIP, &m.MachinePassword, &m.MachineCreateDate, &m.MachineEndDate, &m.AccountId)

		log.Printf("=====> m: %+v", m)
		ms = append(ms, m)
	}
	return ms, nil
}

func QueryAllInfo(dbFile string) (*AllInfo, error) {
	var ai AllInfo
	if as, err := QueryAllAccount(dbFile); err != nil {
		return nil, err
	} else {
		ai.Accounts = as
	}
	if as, err := QueryAllAgent(dbFile); err != nil {
		return nil, err
	} else {
		ai.Agents = as
	}
	if ms, err := QueryAllMachine(dbFile); err != nil {
		return nil, err
	} else {
		ai.Machines = ms
	}

	return &ai, nil

}

func DelAccount(a Account, dbFile string) (*Account, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if len(a.MachineId) > 0 { // 这个账户绑定了机器
		sqlStr := fmt.Sprintf("update machines set accountid='' where machineid=%s;", a.MachineId)
		if _, err := db.Exec(sqlStr); err != nil {
			return nil, err
		}
	}

	sqlStr := fmt.Sprintf("delete from accounts where accountid=%s;", a.AccountId)

	if res, err := db.Exec(sqlStr); err != nil {
		return nil, err
	} else {
		num, _ := res.RowsAffected()
		if num != 1 {
			log.Println("[DelAccount] 删除账户的语句：", sqlStr, " 受影响的数量: ", num)
		}
	}

	return &a, nil
}

func UpdateAccount(a Account, dbFile string) (*Account, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("update accounts set agentdate='%s', enddate='%s', agentid='%s', machineid='%s' where accountid=%s;",
		a.AgentDate, a.EndDate, a.AgentId, a.MachineId, a.AccountId)

	if _, err := db.Exec(sqlStr); err != nil {
		return nil, errors.New(" update account err: " + err.Error() + " sql: " + sqlStr)
	}
	return &a, nil
}

func UpdateMachineAccountId(mId, aId, dbFile string) error {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("update  machines set accountid='%s' where machineid=%s;", aId, mId)

	if _, err := db.Exec(sqlStr); err != nil {
		return errors.New(" update machineAccountid err: " + err.Error() + " sql: " + sqlStr)
	}
	return nil
}

func UpdateMachine(m Machine, dbFile string) (*Machine, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("update  machines set machineip='%s', machinepassword='%s', machineenddate='%s' where machineid=%s;",
		m.MachineIP, m.MachinePassword, m.MachineEndDate, m.MachineId)

	if _, err := db.Exec(sqlStr); err != nil {
		return nil, errors.New(" update machine err: " + err.Error() + " sql: " + sqlStr)
	}
	return &m, nil
}

func DelAgent(a Agent, dbFile string) (*Agent, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("delete from agents where agentid=%s;", a.AgentId)

	if res, err := db.Exec(sqlStr); err != nil {
		return nil, err
	} else {
		num, _ := res.RowsAffected()
		if num != 1 {
			return nil, errors.New("删除代理数据库受影响的原因")
		}
	}

	sqlStr = fmt.Sprintf("update accounts set agentid='' where agentid=%s;", a.AgentId)
	if _, err := db.Exec(sqlStr); err != nil {
		return nil, err
	}

	return &a, err
}

func DelMachine(m Machine, dbFile string) (*Machine, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if len(m.AccountId) > 0 { // 这台机器绑定了账户
		sqlStr := fmt.Sprintf("update accounts set machineid='' where accountid=%s;", m.AccountId)
		if _, err := db.Exec(sqlStr); err != nil {
			return nil, err
		}
	}

	sqlStr := fmt.Sprintf("delete from machines where machineid=%s;", m.MachineId)

	if res, err := db.Exec(sqlStr); err != nil {
		return nil, err
	} else {
		num, _ := res.RowsAffected()
		if num != 1 {
			return nil, errors.New("删除账户数据库受影响的原因")
		}
	}

	return &m, err
}

func EditAccount(ma ModifyAccount, dbFile string) error {
	log.Printf("==========> ma: %+v", ma)
	// 账户信息
	switch ma.Ext {
	case "add": // 添加了机器
		if _, err := UpdateAccount(ma.Account, dbFile); err != nil {
			return err
		}
		if err := UpdateMachineAccountId(ma.DstMachineId, ma.AccountId, dbFile); err != nil {
			return err
		}
	case "release": // 释放了机器
		if err := UpdateMachineAccountId(ma.MachineId, "", dbFile); err != nil {
			return err
		}
		ma.MachineId = ""
		if _, err := UpdateAccount(ma.Account, dbFile); err != nil {
			return err
		}
	case "modify": // 修改了机器
		// 原机器释放绑定
		if err := UpdateMachineAccountId(ma.MachineId, "", dbFile); err != nil {
			return err
		}
		// 新机器绑定
		if err := UpdateMachineAccountId(ma.DstMachineId, ma.AccountId, dbFile); err != nil {
			return err
		}
		// 更改账户绑定的机器
		ma.MachineId = ma.DstMachineId
		if _, err := UpdateAccount(ma.Account, dbFile); err != nil {
			return err
		}
	}
	// 处理代理
	if ma.AgentId != ma.DstAgentId {
		ma.AgentId = ma.DstAgentId
		if _, err := UpdateAccount(ma.Account, dbFile); err != nil {
			return err
		}
	}
	return nil
}

func Backup(dbFile string) (string, error) {
	backupFile := fmt.Sprintf("sd.%s.db", time.Now().Format("2006-01-02 15:04:05"))

	srcFile, err := os.Open(dbFile)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dataPath+backupFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", err
	}
	return backupFile, nil
}

func formatDate(d string) string {
	if len(d) == 0 {
		return ""
	}
	ts := strings.Split(d, "T")
	return ts[0]
}
