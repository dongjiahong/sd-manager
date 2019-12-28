<template>
  <div class="home">
    <div class="layout">
      <Layout>
        <Header>
          <Menu mode="horizontal" theme="dark" active-name="1">
            <div class="layout-logo">
              <MenuItem name="4">
                <Button type="success" @click="backup">备份数据</Button>
              </MenuItem>
            </div>
            <div class="layout-nav">
              <MenuItem name="1">
                <Button type="warning" @click="addAccount">添加账户</Button>
              </MenuItem>
              <MenuItem name="2">
                <Button type="primary" @click="addMachine">添加机器</Button>
              </MenuItem>
              <!-- 添加信息的模态框 -->
              <AddAccount :showModal.sync="showAddAccount"></AddAccount>
              <AddMachine :showModal.sync="showAddMachine"></AddMachine>
            </div>
          </Menu>
        </Header>
        <Layout>
          <Sider hide-trigger :style="{ background: '#fff', height: '100%' }">
            <Menu theme="light" width="auto" :open-names="['1']">
              <Submenu name="1">
                <template slot="title">
                  <Icon type="ios-navigate"></Icon>账户信息
                </template>
                <MenuItem name="1-1" to="/home/accounts/all">所有账户</MenuItem>
                <MenuItem name="1-2" to="/home/accounts/used">使用中的账户</MenuItem>
                <MenuItem name="1-3" to="/home/accounts/unuse">闲置账户</MenuItem>
              </Submenu>
              <Submenu name="2">
                <template slot="title">
                  <Icon type="ios-analytics"></Icon>机器管理
                </template>
                <MenuItem name="2-1" to="/home/machines/formal">有效的机器</MenuItem>
                <MenuItem name="2-2" to="/home/machines/useful">有效未用的机器</MenuItem>
                <MenuItem name="2-3" to="/home/machines/expose">过期的机器</MenuItem>
              </Submenu>
              <Submenu name="3">
                <template slot="title">
                  <Icon type="md-medkit"></Icon>备份数据
                </template>
                <MenuItem v-for="b in backups" v-bind:key="b" :name="b">{{ b }}</MenuItem>
              </Submenu>
            </Menu>
          </Sider>
          <!-- 属性里有冒号表示后面是表达式，没有表示为字面量 -->
          <Layout :style="{ padding: '0 24px 24px' }">
            <Content :style="{ padding: '24px', background: '#fff', height: '100%' }">
              <router-view />
            </Content>
          </Layout>
        </Layout>
      </Layout>
    </div>
  </div>
</template>

<script>
import AddAccount from "@/components/AddAccount";
import AddMachine from "@/components/AddMachine";
import { mapState, mapActions } from "vuex";
import api from "@/fetch/api";

export default {
  name: "Home",
  components: {
    AddAccount,
    AddMachine
  },
  data() {
    return {
      showAddAccount: false,
      showAddMachine: false
    };
  },
  created: function() {
    // 创建完成后拉取数据
    this.fetchAllInfo();
  },
  computed: {
    ...mapState(["backups"])
  },
  methods: {
    ...mapActions(["backupFile", "initAllInfo"]),
    addAccount() {
      this.showAddAccount = true;
    },
    addMachine() {
      this.showAddMachine = true;
    },
    backup() {
      this.backupFile(this);
    },
    fetchAllInfo: function() {
      this.initAllInfo(this);
      api.getBackupFiles().then(data => {
        console.log("======> 获取备份文件目录: ", data);
        if (data.message === "ok") {
          this.$store.commit("initBackupFiles", data.data);
        } else {
          this.$Message.warning(data.message);
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
.home {
  height: 100%;
  .layout {
    .ivu-layout {
      height: 100%; // ivu-layout 是iview的类，这里设置是为了充满整个网页
    }
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    height: 100%;
    border-radius: 4px;
    overflow: scroll;
  }
  .layout-logo {
    float: left;
  }
  .layout-nav {
    width: 420px;
    margin: 0 auto;
    margin-right: 20px;
  }
}
</style>
