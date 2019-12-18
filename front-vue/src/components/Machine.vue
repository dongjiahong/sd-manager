<template>
  <div class="machine">
    <Table
      border
      :columns="machine_label"
      :data="getMachinesWithState($route.params.state)"
    >
      <template slot-scope="{ row }" slot="machine_no">
        <strong>{{ row.machine_no }}</strong>
      </template>
      <template slot-scope="{ row }" slot="action">
        <Button
        type="primary"
        size="small"
        style="margin-right: 5px"
        @click="edit(row)"
        >编辑</Button>
        <Button type="error" size="small" @click="remove(row)">删除</Button>
      </template>
    </Table>
    <EditMachine
      :showModal.sync="showEditMachine"
      :machineDetail="activeMachine"
    ></EditMachine>
  </div>
</template>

<script>
import EditMachine from "@/components/EditMachine";
import util from "@/util/util"
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Machine",
  components: {
    EditMachine
  },
  data() {
    return {
      showEditMachine: false,
      activeMachine: {},
      machine_label: [
        {
          title: "机器代号",
          slot: "machine_no"
        },
        {
          title: "绑定账户id",
          key: "account_id"
        },
        {
          title: "机器ip",
          key: "machine_ip"
        },
        {
          title: "机器密码",
          key: "machine_password"
        },
        {
          title: "机器创建时间",
          key: "machine_create_date",
          sortable: true
        },
        {
          title: "机器到期时间",
          key: "machine_end_date",
          sortable: true
        },
        {
          title: "机器可用时间",
          key: "machine_time_left",
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
    ...mapGetters(["getMachinesWithState"])
  },
  methods: {
    ...mapActions(['eidtMachine', 'delMachine']),
    edit(row) {
      this.showEditMachine = true;
      // js的对象时引用类型，当我们修改内容时，表格里的内容也会改动，所以这里深度拷贝
      this.activeMachine =  util.deepCopy(row)
    },
    remove(item){
      if (item.account_id != '' && item.account_id != null) {
        this.$Notice.info({
          title: "删除警告",
          desc: "请先解除该机器上的账户, 绑定的账户id: "+item.account_id
        })
        return
      }
      this.delMachine({me: this, data: item})
    }
  }
};
</script>
<style lang="scss" scoped>
.machine {
  height: 100%;
}
</style>
