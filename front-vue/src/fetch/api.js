import async from "./fetch";

export default {
  // 拉取所有信息
  getAllInfo() {
    return async("api/getallinfo");
    // return async("api/fake/getallinfo");
  },
  // 添加账户
  addAccount(item) {
    return async("api/add/account", item, "post");
  },
  // 修改账户
  editAccount(item) {
    return async("api/edit/account", item, "post");
  },
  // 修改机器
  editMachine(item) {
    return async("api/edit/machine", item, "post");
  },
  // 添加机器
  addMachine(item) {
    return async("api/add/machine", item, "post");
  },
  // 添加代理
  addAgent(item) {
    return async("api/add/agent", item, "post");
  },
  // 删除账户
  delAccount(item) {
    return async("api/del/account", item, "post");
  },
  // 删除机器
  delMachine(item) {
    return async("api/del/machine", item, "post");
  },
  // 备份
  backup() {
    return async("api/backup");
  },
  // 拉取备份文件
  getBackupFiles() {
    return async("api/getbackupfiles");
  },
  // 登录
  load(user, pwd) {
    let token = window.btoa(user + "::" + pwd);
    return async("/api/fake/load?token=" + token);
  }
};
