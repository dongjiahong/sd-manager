<template>
  <div class="account">
    <Table
      border
      size="small"
      :highlight-row="true"
      :columns="account_label"
      :data="getAccountsWithState($route.params.state)"
    >
      <template slot-scope="{ row }" slot="account_no">
        <strong>{{ row.account_no }}</strong>
      </template>
      <template slot-scope="{ row }" slot="action">
        <Button
          type="primary"
          size="small"
          style="margin-right: 5px"
          @click="edit(row)"
          >编辑</Button
        >
        <Button type="error" size="small" @click="remove(row)">删除</Button>
      </template>
    </Table>
    <EditAccount
      :showModal.sync="showEditAccount"
      :accountDetail="activeRow"
    ></EditAccount>
  </div>
</template>

<script>
import EditAccount from "@/components/EditAccount";
import util from "@/util/util";
import { mapState, mapGetters, mapActions } from "vuex";

export default {
  name: "Account",
  components: {
    EditAccount
  },
  data() {
    return {
      showEditAccount: false,
      activeRow: {},
      account_label: [
        {
          title: "账号代号",
          width: 90,
          fixed: 'left',
          slot: "account_no"
        },
        {
          title: "账户邮箱",
          key: "account_mail"
        },
        {
          title: "账户密码",
          key: "account_password"
        },
        {
          title: "验证邮箱",
          key: "verify_mail"
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
          title: "剩余(天)",
          key: "time_left",
          width: 80,
          sortable: true
        },
        {
          title: "备注",
          key: "tips"
        },
        {
          title: "Action",
          slot: "action",
          width: 130,
          fixed: 'right',
          align: "center"
        }
      ]
    };
  },
  computed: {
    ...mapState(["accounts"]),
    ...mapGetters(["getAccountsWithState"])
  },
  methods: {
    ...mapActions(["delAccount"]),
    edit(row) {
      this.showEditAccount = true;
      // 这里使用深拷贝传值，以保障每次编辑都能触发更新
      this.activeRow = util.deepCopy(row);
    },
    remove(item) {
      console.log("=====> remove item:", item);
      this.delAccount({ me: this, data: item });
    }
  }
};
</script>

<style lang="scss" scoped>
.account {
  height: 100%;
}
</style>
