# sd manager 的后端系统

1. 使用go modules进行包管理
2. 使用gin网络库
3. sqlite3 `https://github.com/mattn/go-sqlite3`
4. golang版本**1.11**以上，才能支持modules

## sql 建表

```sql
#建表accounts
CREATE TABLE "accounts" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"accountno"	TEXT NOT NULL UNIQUE,
	"accountmail"	TEXT NOT NULL,
	"accountpassword"	TEXT NOT NULL,
	"verifymail"	TEXT DEFAULT "",
	"createdate"	TEXT NOT NULL,
	"agentdate"	TEXT DEFAULT "",
	"enddate"	TEXT DEFAULT "",
	"agentname"	TEXT DEFAULT "",
	"machineno"	TEXT DEFAULT "",
	"tip"	TEXT DEFAULT ""
);
#设置触发器，如果删除account信息，会触发对应的machine信息的修改
CREATE TRIGGER after_account_del AFTER DELETE ON accounts
BEGIN
	UPDATE machines set accountno='' where machineno=OLD.machineno;
END
#设置触发器，如果更新account信息，会触发对应的machine的更改
CREATE TRIGGER after_account_up AFTER UPDATE ON accounts
BEGIN
	UPDATE machines set accountno='' where machineno=OLD.machineno;
	UPDATE machines set accountno=OLD.accountno where machineno=NEW.machineno;
END
#建表machines
CREATE TABLE "machines" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"machineno"	TEXT NOT NULL UNIQUE,
	"machineip"	TEXT NOT NULL,
	"machinepassword"	TEXT NOT NULL,
	"machinecreatedate"	TEXT NOT NULL,
	"machineenddate"	TEXT NOT NULL,
	"accountno"	TEXT
);
#建表managers
CREATE TABLE "managers" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"managername"	TEXT NOT NULL UNIQUE,
	"manageraccount"	TEXT NOT NULL UNIQUE,
	"managerpassword"	TEXT NOT NULL
);
```
