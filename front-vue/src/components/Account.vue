<template>
  <div class="account">
    <Table
      border
      :stripe="true"
      size="small"
      :highlight-row="true"
      :columns="account_label"
      :data="getAccountsWithState($route.params.state)"
    >
      <!-- <template slot-scope="{ row }" slot="account_no">
        <strong>{{ row.account_no }}</strong>
      </template> -->
      <template slot-scope="{ row }" slot="action">
        <Button type="primary" size="small" style="margin-right: 5px" @click="edit(row)">编辑</Button>
        <Button type="error" size="small" @click="remove(row)">删除</Button>
      </template>
    </Table>
    <EditAccount :showModal.sync="showEditAccount" :accountDetail="activeRow"></EditAccount>
  </div>
</template>

<script>
import EditAccount from "@/components/EditAccount";
import ExpandRow from "@/components/AccountTableExpand";
import util from "@/util/util";
import { mapState, mapGetters, mapActions } from "vuex";

export default {
  name: "Account",
  components: {
    // ExpandRow,
    EditAccount
  },
  data() {
    return {
      showEditAccount: false,
      activeRow: {},
      account_label: [
        {
          type: "expand",
          width: 50,
          render: (h, params) => {
            return h(ExpandRow, {
              props: {
                row: params.row
              }
            });
          }
        },
        {
          // title: "账号代号",
          width: 110,
          key: "account_no",
          // sortable: true,
          // slot: "account_no",
          renderHeader: (h) => {
            return h("div", [
              h('strong', '账户代号'),
              h('i-input')
            ]);
          }
        },
        {
          title: "机器代号",
          sortable: true,
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
          title: "代理人",
          key: "agent_name"
        },
        {
          title: "账户到期时间",
          key: "end_date",
          width: 130,
          sortable: true
        },
        {
          title: "剩余(天)",
          key: "time_left",
          width: 100,
          sortable: true
        },
        {
          title: "备注",
          key: "tip"
        },
        {
          title: "Action",
          slot: "action",
          width: 130,
          fixed: "right",
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
