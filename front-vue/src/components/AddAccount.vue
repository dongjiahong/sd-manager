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
        <FormItem label="选择代理" prop="agent_id" :rules="{ required: true }">
          <Select
            v-model="accountFormValidate.agent_id"
            placeholder="选择账户的代理"
          >
            <Option v-for="a in agents" :value="a.id" :key="a.agent_name">{{
              a.agent_name
            }}</Option>
          </Select>
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
        <FormItem label="选择机器" prop="machine_id">
          <Select
            v-model="accountFormValidate.machine_id"
            placeholder="选择机器"
            clearable
          >
            <Option
              v-for="m in getMachinesWithState('useful')"
              :value="m.id"
              :key="m.machine_no"
              >{{ m.machine_no }}</Option
            >
          </Select>
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
        agent_id: "", // 账户代理
        agent_date: "", // 代理时间
        end_date: "", // 代理的截止日期
        machine_id: "" // 机器
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
              agent_id: this.accountFormValidate.agent_id,
              agent_date: this.accountFormValidate.agent_date,
              end_date: this.accountFormValidate.end_date,
              machine_id: this.accountFormValidate.machine_id
            }
          });
        }
      });
      this.show = false;

      this.$refs[name].resetFields();
    },
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
