<template>
  <div class="edit-account">
    <Modal
      v-model="show"
      title="编辑账户"
      :loading="loading"
      @on-ok="asyncOK('editAccountFormValidate')"
    >
      <Form
        ref="editAccountFormValidate"
        :model="editAccountFormValidate"
        :label-width="80"
      >
        <FormItem
          label="账户代号"
          prop="account_no"
          :rules="{ required: true, trigger: 'blur' }"
        >
          <i-input
            v-model="editAccountFormValidate.account_no"
            placeholder="输入账户代号"
            readonly
          ></i-input>
        </FormItem>
        <FormItem
          label="账户邮箱"
          prop="account_mail"
          :rules="{ required: true, trigger: 'blur' }"
        >
          <i-input
            v-model="editAccountFormValidate.account_mail"
            placeholder="输入账户邮箱"
          ></i-input>
        </FormItem>
        <FormItem
          label="账户密码"
          prop="account_password"
          :rules="{ required: true, trigger: 'blur' }"
        >
          <i-input
            v-model="editAccountFormValidate.account_password"
            placeholder="输入账户密码"
          ></i-input>
        </FormItem>
        <FormItem
          label="验证邮箱"
          prop="verify_mail"
          :rules="{ required: true, trigger: 'blur' }"
        >
          <i-input
            v-model="editAccountFormValidate.verify_mail"
            placeholder="输入验证邮箱"
          ></i-input>
        </FormItem>
        <FormItem label="代理人" prop="agent_name">
          <i-input
            v-model="editAccountFormValidate.agent_name"
            placeholder="输入代理人名称"
          ></i-input>
        </FormItem>
        <FormItem label="服务器">
          <Select
            v-model="editAccountFormValidate.machine_no"
            placeholder="选择服务器"
            clearable
            style="width: 250px"
          >
            <Option
              v-for="m in getEditMachines(accountDetail.machine_no)"
              :value="m.machine_no"
              :key="m.machine_no"
              >{{ m.machine_no }}</Option
            >
          </Select>
        </FormItem>
        <FormItem label="创建日期">
          <!-- 创建日期不可更改 -->
          <i-input
            v-model="editAccountFormValidate.create_date"
            readonly
          ></i-input>
        </FormItem>

        <FormItem label="授权时间">
          <DatePicker
            type="date"
            placeholder="代理的授权日期"
            v-model="editAccountFormValidate.agent_date"
          ></DatePicker>
        </FormItem>

        <FormItem label="代理时间">
          <DatePicker
            type="date"
            placeholder="代理的到期日期"
            v-model="editAccountFormValidate.end_date"
          ></DatePicker>
        </FormItem>

        <FormItem label="剩余时间">
          <i-input
            v-model="editAccountFormValidate.time_left"
            placeholder="输入该账号的使用时间"
            readonly
          ></i-input>
        </FormItem>
        <FormItem label="备注信息">
          <i-input
            v-model="editAccountFormValidate.tip"
            placeholder="输入备注"
          ></i-input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
import util from "@/util/util";
import { mapGetters, mapActions } from "vuex";
export default {
  name: "EditAccount",
  props: {
    showModal: Boolean,
    // ********* 注意 ******** //
    // 在js中对象和数组是引用类型，指向同一个内存，如果prop是对象或者数组
    // 在子组件内部改变它会影响到父组件, 为了不影响父组件最好使用深度copy,
    // 如： JSON.parse(JSON.stringif(obj)), 用法见下面的accountDetail
    accountDetail: Object
  },
  data() {
    return {
      show: false, // show 控制模态框是否展示
      loading: true,
      editAccountFormValidate: {}
    };
  },
  computed: {
    ...mapGetters(["getEditMachines"])
  },
  methods: {
    ...mapActions(["editAccount"]),
    asyncOK(name) {
      console.log(
        "===> editAccountFormValidate data: ",
        this.editAccountFormValidate
      );
      this.$refs[name].validate(valid => {
        if (!valid) {
          this.$Message.warning("请检查你的输入内容");
        } else {
          this.editAccountFormValidate.agent_date = util.formatDate(
            this.editAccountFormValidate.agent_date
          );
          this.editAccountFormValidate.end_date = util.formatDate(
            this.editAccountFormValidate.end_date
          );
          this.editAccount({ me: this, data: this.editAccountFormValidate });
        }
      });
      this.show = false;
    }
  },
  watch: {
    showModal: function() {
      this.show = this.showModal;
    },
    show: function() {
      this.$emit("update:showModal", this.show);
    },
    accountDetail: function(newVal) {
      // 由于是动态传参，初始化时accountdetail是没有值得，这里监控变化
      this.editAccountFormValidate = util.deepCopy(newVal);
    }
  }
};
</script>

<style lang="scss" scoped></style>
