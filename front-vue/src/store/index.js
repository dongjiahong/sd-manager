import Vue from "vue";
import Vuex from "vuex";
import util from "@/util/util";
import api from "@/fetch/api";

Vue.use(Vuex);

// vuex实例
export default new Vuex.Store({
  state: {
    keepAlive: false, // 是否登录
    accounts: [],
    machines: [],
    backups: []
  },
  getters: {
    // 获取登录状态
    isLoad: state => {
      return state.keepAlive;
    },
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
          m => m.machine_time_left > 0 && m.account_id === ""
        );
      } else {
        console.log("===> getMachinesWithState unknown state: ", info);
      }
    },
    getEditMachines: state => machine_no => {
      return state.machines.filter(
        m =>
          (m.machine_time_left > 0 && m.account_id === "") ||
          m.machine_no == machine_no
      );
    }
  },
  mutations: {
    // 修改是否登录状态
    changeLoad(state, load) {
      state.keepAlive = load;
    },
    // 修改账户信息
    changeAccount(state, payload) {
      for (let index in state.accounts) {
        if (state.accounts[index].account_id === payload.item.account_id) {
          switch (payload.operate) {
            case "del": // 删除数据
              state.accounts.splice(index, 1);
              return;
            case "ch": // 修改数据
              if (state.accounts[index].end_date != payload.item.end_date) {
                // 修改的内容是时间
                // TODO 时间的加减
              }
              state.accounts[index] = payload.item;
              return;
          }
        }
      }
    },
    initMachines(state, items) {
      for (let index in items) {
        let newValue = {
          id: items[index].id,
          machine_no: items[index].machine_no,
          machine_ip: items[index].machine_ip,
          machine_password: items[index].machine_password,
          machine_create_date: items[index].machine_create_date,
          machine_end_date: items[index].machine_end_date,
          account_id: items[index].account_id,
          machine_time_left: util.dateDifference(items[index].machine_end_date)
        };
        Vue.set(state.machines, index, newValue);
      }
    },
    initAccounts(state, items) {
      for (let index in items) {
        let agent_name = "";
        let machine_no = "";
        let machine_ip = "";
        let machine_password = "";
        for (let i in state.agents) {
          if (state.agents[i].id == items[index].agent_id) {
            agent_name = state.agents[i].agent_name;
          }
        }
        for (let i in state.machines) {
          if (state.machines[i].id == items[index].machine_id) {
            machine_no = state.machines[i].machine_no;
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
          create_date: items[index].create_date,
          agent_date: items[index].agent_date,
          end_date: items[index].end_date,
          agent_id: items[index].agent_id,
          agent_name: agent_name,
          machine_id: items[index].machine_id,
          machine_no: machine_no,
          machine_ip: machine_ip,
          machine_password: machine_password,
          time_left: util.dateDifference(items[index].end_date)
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
      let agent_name = "";
      let machine_no = "";
      let machine_ip = "";
      let machine_password = "";
      for (let i in state.agents) {
        // 获取代理名称
        if (state.agents[i].id == item.agent_id) {
          agent_name = state.agents[i].agent_name;
          break;
        }
      }
      for (let i in state.machines) {
        if (state.machines[i].id == item.machine_id) {
          // 获取机器信息
          machine_no = state.machines[i].machine_no;
          machine_ip = state.machines[i].machine_ip;
          machine_password = state.machines[i].machine_password;
          state.machines[i].account_id = item.id;
          break;
        }
      }
      let newValue = {
        id: item.id,
        account_no: item.account_no,
        create_date: item.create_date,
        agent_date: item.agent_date,
        end_date: item.end_date,
        agent_id: item.agent_id,
        agent_name: agent_name,
        machine_id: item.machine_id,
        machine_no: machine_no,
        machine_ip: machine_ip,
        machine_password: machine_password,
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
      if (item.machine_id != "") {
        for (let index in state.machines) {
          if (item.machine_id == state.machines[index].id) {
            let copyMachine = util.deepCopy(state.machines[index]);
            copyMachine.account_id = "";
            // state.machines[index] = copyMachine;
            Vue.set(state.machines, index, copyMachine);
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
    changeLoad(context, load) {
      console.log("=======>   context: ", context)
      context.commit("changeLoad", load)
    },
    initAllInfo({commit}, me) {
      // if (state.keepAlive == false) {
      //   me.$router.push({name: 'load'});
      //   return
      // }
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
    editAccount({ dispatch, state }, payload) {
      let item = payload.data;
      // 首先判断是否修改了机器
      if (item.machine_id == "") {
        if (item.machine_no != "") {
          // 添加
          for (let index in state.machines) {
            if (item.machine_no == state.machines[index].machine_no) {
              item.dst_machine_id = state.machines[index].id;
              item.ext = "add";
              break;
            }
          }
        }
      } else {
        if (item.machine_no == "" || item.machine_no == undefined) {
          // 释放
          item.ext = "release";
        } else {
          for (let index in state.machines) {
            if (item.machine_no == state.machines[index].machine_no) {
              if (item.machine_id != state.machines[index].id) {
                // 变更
                item.ext = "modify";
                // 这里不规范，使用no来保存需要变更的目标id
                item.dst_machine_id = state.machines[index].id;
                break;
              }
            }
          }
        }
      }
      // 处理代理
      for (let index in state.agents) {
        if (item.agent_name == state.agents[index].agent_name) {
          if (item.agent_id != state.agents[index].id) {
            item.dst_agent_id = state.agents[index].id;
            break;
          }
        }
      }
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
