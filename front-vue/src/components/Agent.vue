<template>
  <div class="agent">
    <Table border :columns="account_label" :data="getAccountsWithName($route.params.name)" >
      <template slot-scope="{ row }" slot="account_no">
        <strong>{{ row.account_no }}</strong>
      </template>
      <template slot-scope="{ row, index }" slot="action">
        <Button type="error" size="small" @click="remove(index)">删除</Button>
        <EditAccount
          :showModal.sync="showEditAccount"
          :accountDetail="activeAccount"
        ></EditAccount>
      </template>
    </Table>
  </div>
</template>

<script>
import EditAccount from "@/components/EditAccount";
import { mapState, mapGetters } from "vuex";

export default {
  name: "Account",
  components: {
    EditAccount
  },
  data() {
    return {
      showEditAccount: false,
      activeAccount: {},
      account_label: [
        {
          title: "账号代号",
          slot: "account_no"
        },
        {
          title: "代理人",
          key: "agent_name"
        },
        {
          title: "机器代号",
          key: "machine_no"
        },
        {
          title: "机器IP",
          key: "machine_ip"
        },
        {
          title: "机器密码",
          key: "machine_password"
        },
        {
          title: "账户创建时间",
          key: "create_date",
          sortable: true
        },
        {
          title: "账户授权日期",
          key: "agent_date",
          sortable: true
        },
        {
          title: "账户到期时间",
          key: "end_date",
          sortable: true
        },
        {
          title: "授权剩余时间(天)",
          key: "time_left",
          sortable: true
        },
        {
          title: "Action",
          slot: "action",
          width: 150,
          align: "center"
        }
      ]
    };
  },
  computed: {
    ...mapState(["accounts"]),
    ...mapGetters(["getAccountsWithName"])
  },
  methods: {
    edit(index) {
      this.showEditAccount = true;
      this.activeAccount = this.getAccounts[index];
    },
    remove(index) {
      this.account_detail.splice(index, 1);
    }
  }
};
</script>

<style lang="scss" scoped>
.agent{
  height: 100%;
}
</style>
