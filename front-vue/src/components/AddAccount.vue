<template>
  <div class="add-account">
    <Modal
      v-model="show"
      title="添加账户"
      :loading="loading"
      @on-ok="asyncOK('accountFormValidate')"
    >
      <Form
        ref="accountFormValidate"
        :model="accountFormValidate"
        :label-width="120"
      >
        <FormItem
          label="账户代号"
          prop="account_no"
          :rules="{
            required: true,
            message: '账户代号, 如：a101',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="accountFormValidate.account_no"
            placeholder="输入账户代号"
          ></i-input>
        </FormItem>
        <FormItem
          label="账户邮箱"
          prop="account_mail"
          :rules="{
            required: true,
            message: '账户邮箱, 如：1234555@gmail.com',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="accountFormValidate.account_mail"
            placeholder="输入账户邮箱"
          ></i-input>
        </FormItem>
        <FormItem
          label="账户密码"
          prop="account_password"
          :rules="{ required: true, trigger: 'blur' }"
        >
          <i-input
            v-model="accountFormValidate.account_password"
            placeholder="输入账户密码"
          ></i-input>
        </FormItem>
        <FormItem
          label="验证邮箱"
          prop="verify_mail"
          :rules="{ required: true, trigger: 'blur' }"
        >
          <i-input
            v-model="accountFormValidate.verify_mail"
            placeholder="输入验证邮箱"
          ></i-input>
        </FormItem>
        <FormItem label="代理人">
          <i-input
            v-model="accountFormValidate.agent_name"
            placeholder="输入代理名字"
          ></i-input>
        </FormItem>
        <FormItem label="代理时间" prop="agent_date">
          <DatePicker
            type="date"
            placeholder="代理的授权日期"
            v-model="accountFormValidate.agent_date"
          ></DatePicker>
        </FormItem>
        <FormItem label="代理截止日期" prop="end_date">
          <DatePicker
            type="date"
            placeholder="代理的截止日期"
            v-model="accountFormValidate.end_date"
          ></DatePicker>
        </FormItem>
        <FormItem label="选择机器" prop="machine_no">
          <Select
            v-model="accountFormValidate.machine_no"
            placeholder="选择机器"
            clearable
          >
            <Option
              v-for="m in getMachinesWithState('useful')"
              :value="m.machine_no"
              :key="m.machine_no"
              >{{ m.machine_no }}</Option
            >
          </Select>
        </FormItem>
        <FormItem label="输入备注">
          <i-input
            v-model="accountFormValidate.tip"
            placeholder="输入备注"
          ></i-input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
import { mapState, mapGetters, mapActions } from "vuex";
export default {
  name: "AddAccount",
  props: {
    showModal: Boolean
  },
  data() {
    return {
      show: false, // show 控制模态框是否展示
      loading: true,
      accountFormValidate: {
        account_no: "", // 账户代号
        account_mail: "", // 账户邮箱
        account_password: "", // 账户密码
        verify_mail: "", // 验证邮箱
        agent_name: "", // 账户代理
        agent_date: "", // 代理时间
        end_date: "", // 代理的截止日期
        machine_no: "", // 机器
        tip: "" // 备注
      }
    };
  },
  computed: {
    ...mapState(["agents"]),
    ...mapGetters(["getMachinesWithState"])
  },
  methods: {
    ...mapActions(["addAccount"]),
    asyncOK(name) {
      this.$refs[name].validate(valid => {
        if (!valid) {
          this.$Message.warning("请检查你的输入内容");
        } else {
          this.addAccount({
            me: this,
            data: {
              account_no: this.accountFormValidate.account_no,
              account_mail: this.accountFormValidate.account_mail,
              account_password: this.accountFormValidate.account_password,
              verify_mail: this.accountFormValidate.verify_mail,
              agent_name: this.accountFormValidate.agent_name,
              agent_date: this.accountFormValidate.agent_date,
              end_date: this.accountFormValidate.end_date,
              machine_no: this.accountFormValidate.machine_no,
              tip: this.accountFormValidate.tip
            }
          });
        }
      });
      this.show = false;

      this.$refs[name].resetFields();
    }
  },
  watch: {
    showModal: function() {
      this.show = this.showModal;
    },
    show: function() {
      this.$emit("update:showModal", this.show);
    }
  }
};
</script>

<style lang="scss" scoped></style>
