import Vue from "vue";
import Vuex from "vuex";
import util from "@/util/util";
import api from "@/fetch/api";

Vue.use(Vuex);

// vuex实例
export default new Vuex.Store({
  state: {
    accounts: [],
    machines: [],
    backups: []
  },
  getters: {
    getAccountsWithState: state => info => {
      if (info === "all") {
        return state.accounts;
      } else if (info === "used") {
        return state.accounts.filter(a => a.time_left > 0);
      } else if (info === "unuse") {
        return state.accounts.filter(a => a.time_left <= 0);
      } else {
        console.log("===> getAccountsWithState unknown state: ", info);
      }
    },
    getAccountsWithName: state => name => {
      return state.accounts.filter(a => a.agent_name === name);
    },
    // 获取机器信息
    getMachinesWithState: state => info => {
      if (info === "formal") {
        // 有效机器
        return state.machines.filter(m => m.machine_time_left > 0);
      } else if (info === "expose") {
        // 过期机器
        return state.machines.filter(m => m.machine_time_left <= 0);
      } else if (info === "useful") {
        return state.machines.filter(
          m => (m.machine_time_left > 0 && m.account_no === "")
        );
      } else if (info === "all") {
        return state.machines;
      } else {
        console.log("===> getMachinesWithState unknown state: ", info);
      }
    },
    getEditMachines: state => machine_no => {
      return state.machines.filter(
        m =>
          (m.machine_time_left > 0 && m.account_no === "") ||
          m.machine_no == machine_no
      );
    }
  },
  mutations: {
    initMachines(state, items) {
      for (let index in items) {
        let newValue = {
          id: items[index].id,
          machine_no: items[index].machine_no,
          machine_ip: items[index].machine_ip,
          machine_password: items[index].machine_password,
          machine_create_date: items[index].machine_create_date,
          machine_end_date: items[index].machine_end_date,
          account_no: items[index].account_no,
          machine_time_left: util.dateDifference(items[index].machine_end_date)
        };
        Vue.set(state.machines, index, newValue);
      }
    },
    initAccounts(state, items) {
      for (let index in items) {
        let machine_ip = "";
        let machine_password = "";
        for (let i in state.machines) {
          if (state.machines[i].machine_no == items[index].machine_no) {
            machine_ip = state.machines[i].machine_ip;
            machine_password = state.machines[i].machine_password;
          }
        }

        // ****** 注意注意 ****
        // 由于vue不能监控数组的变化，所以我们如果使用a[i] = b这样的形式是不会触发视图更新的，
        // 使用Vue.set(a, i, b)来给数组添加成员就可以触发视图更新了
        let newValue = {
          id: items[index].id,
          account_no: items[index].account_no,
          account_mail: items[index].account_mail,
          account_password: items[index].account_password,
          verify_mail: items[index].verify_mail,
          create_date: items[index].create_date,
          agent_name: items[index].agent_name,
          agent_date: items[index].agent_date,
          end_date: items[index].end_date,
          machine_no: items[index].machine_no,
          machine_ip: machine_ip,
          machine_password: machine_password,
          time_left: util.dateDifference(items[index].end_date),
          tip: items[index].tip
        };
        Vue.set(state.accounts, index, newValue);
      }
    },
    initBackupFiles(state, items) {
      for (let index in items) {
        Vue.set(state.backups, index, items[index]);
      }
    },
    addMachine(state, item) {
      console.log("=======>  addMachine item: ", item);
      let newValue = {
        id: item.id,
        machine_no: item.machine_no,
        machine_ip: item.machine_ip,
        machine_password: item.machine_password,
        machine_create_date: item.machine_create_date,
        machine_end_date: item.machine_end_date,
        machine_time_left: util.dateDifference(item.machine_end_date)
      };
      Vue.set(state.machines, state.machines.length, newValue);
    },
    // 增加账户
    addAccount(state, item) {
      console.log("=======>  addAccount item: ", item);
      let machine_ip = "";
      let machine_password = "";
      for (let i in state.machines) {
        if (state.machines[i].machine_no == item.machine_no) {
          // 获取机器信息
          machine_ip = state.machines[i].machine_ip;
          machine_password = state.machines[i].machine_password;
          state.machines[i].account_no = item.account_no;
          break;
        }
      }
      let newValue = {
        id: item.id,
        account_no: item.account_no,
        account_mail: item.account_mail,
        account_password: item.account_password,
        verify_mail: item.verify_mail,
        create_date: item.create_date,
        agent_date: item.agent_date,
        agent_name: item.agent_name,
        end_date: item.end_date,
        machine_no: item.machine_no,
        machine_ip: machine_ip,
        machine_password: machine_password,
        tip: item.tip,
        time_left: util.dateDifference(item.end_date)
      };
      Vue.set(state.accounts, state.accounts.length, newValue);
    },
    // 删除账户
    delAccount(state, item) {
      // 更新账户和机器列表
      for (let index in state.accounts) {
        if (item.id === state.accounts[index].id) {
          state.accounts.splice(index, 1);
          break;
        }
      }
      console.log("====> delAccount item:", item);
      if (item.machine_no != "") {
        for (let index in state.machines) {
          if (item.machine_no == state.machines[index].machine_no) {
            let copyMachine = util.deepCopy(state.machines[index]);
            copyMachine.account_no = "";
            Vue.set(state.machines, index, copyMachine);
            console.log("====> delAccount copy ok");
            break;
          }
        }
      }
    },
    // 删除机器
    delMachine(state, item) {
      for (let index in state.machines) {
        if (item.id == state.machines[index].id) {
          state.machines.splice(index, 1);
          break;
        }
      }
    },
    backupFile(state, item) {
      Vue.set(state.backups, state.backups.length, item);
    },
    // 修改账户
    // editAccount(state, item) {
    //
    // },
    // 修改机器
    editMachine(state, item) {
      console.log("=====> editMachine item: ", item);
      for (let index in state.machines) {
        if (item.id == state.machines[index].id) {
          let newValue = {
            id: item.id,
            machine_ip: item.machine_ip,
            machine_no: item.machine_no,
            machine_password: item.machine_password,
            machine_create_date: item.machine_create_date,
            machine_end_date: item.machine_end_date,
            account_id: item.account_id,
            machine_time_left: util.dateDifference(item.machine_end_date)
          };
          Vue.set(state.machines, index, newValue);
          break;
        }
      }
    }
  },
  actions: {
    // addAgent({commit}, item) {
    //   // TODO 请求后台
    //   commit('addAgent', item);
    // }
    initAllInfo({ commit }, me) {
      api.getAllInfo().then(data => {
        console.log("=====> 初始化的信息数据: ", data);
        if (data.message == "ok") {
          commit("initMachines", data.data.machines);
          commit("initAccounts", data.data.accounts);
        } else {
          me.$Notice.error({
            title: "初始化数据失败",
            desc: data.message
          });
        }
      });
    },
    addAccount(context, payload) {
      let item = payload.data;
      api.addAccount(item).then(data => {
        console.log("====> add account: ", data);
        if (data.message == "ok") {
          context.commit("addAccount", data.data);
        } else {
          payload.me.$Notice.error({
            title: "添加账户失败",
            desc: data.message
          });
        }
      });
    },
    addMachine(context, payload) {
      let item = payload.data;
      api.addMachine(item).then(data => {
        console.log("====> add machine: ", data);
        if (data.message == "ok") {
          context.commit("addMachine", data.data);
        } else {
          payload.me.$Notice.error({
            title: "添加机器失败",
            desc: data.message
          });
        }
      });
    },
    delAccount(context, payload) {
      let item = payload.data;
      api.delAccount(item).then(data => {
        console.log("=====> del account: ", data);
        if (data.message == "ok") {
          context.commit("delAccount", data.data);
        } else {
          payload.me.$Notice.error({
            title: "删除账户错误",
            desc: data.message
          });
        }
      });
    },
    delMachine(context, payload) {
      let item = payload.data;
      api.delMachine(item).then(data => {
        if (data.message == "ok") {
          context.commit("delMachine", data.data);
        } else {
          payload.me.$Notice.error({
            title: "删除机器错误",
            desc: data.message
          });
        }
      });
    },
    editAccount({ dispatch }, payload) {
      let item = payload.data;
      api.editAccount(item).then(data => {
        console.log("======> edit account: ", data);
        if (data.message === "ok") {
          // commit("", data.data);
          dispatch("initAllInfo", payload.me);
        } else {
          payload.me.$Notice.error({
            title: "更新账户错误",
            desc: data.message
          });
        }
      });
    },
    editMachine({ commit }, payload) {
      let item = payload.data;
      console.log("=======> edit machine: ", item);
      api.editMachine(item).then(data => {
        if (data.message == "ok") {
          commit("editMachine", data.data);
        } else {
          payload.me.$Notice.error({
            title: "更新机器错误",
            desc: data.message
          });
        }
      });
    },
    backupFile(context, me) {
      api.backup().then(data => {
        console.log("======> backup file: ", data);
        if (data.message === "ok") {
          context.commit("backupFile", data.data);
        } else {
          me.$Message.info("备份: ", data.message);
        }
      });
    }
  }

  // ### 示例 ###
  // state: {
  //   name: 'tom'
  // },
  // getters: {
  // // 调用 ...mapGetters([ "getname" ])
  //   getname: state => {
  //     // return state.name;
  //   }
  // }
  // mutations: {
  // // 调用 this.$store.commit('setname')
  //   setname(state) {
  //     state.name = "july";
  //   },
  //   setname2(state, name) {
  //     state.name = name
  //   },
  // },
  // actions: {
  //  // 调用this.$store.dispatch('setname', 'xiaogang')
  //   setname(context, name) {
  //     context.commit('setname2', name)
  //   }
  // },
});
