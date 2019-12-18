# sd manager 的后端系统

1. 使用go modules进行包管理
2. 使用gin网络库
3. sqlite3 `https://github.com/mattn/go-sqlite3`

## sql 建表

```sql
CREATE TABLE "accounts" (
	"accountid"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"accountno"	TEXT NOT NULL UNIQUE,
	"createdate"	TEXT NOT NULL,
	"agentdate"	TEXT,
	"enddate"	TEXT,
	"agentid"	TEXT,
	"machineid"	TEXT
)
CREATE TABLE "agents" (
	"agentid"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"agentname"	TEXT NOT NULL UNIQUE,
	"agentaccount"	TEXT NOT NULL UNIQUE,
	"agentpassword"	TEXT NOT NULL
);
CREATE TABLE "machines" (
	"machineid"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"machineno"	TEXT NOT NULL UNIQUE,
	"machineip"	TEXT NOT NULL,
	"machinepassword"	TEXT NOT NULL,
	"machinecreatedate"	TEXT NOT NULL,
	"machineenddate"	TEXT NOT NULL,
	"accountid"	INTEGER
);
```
