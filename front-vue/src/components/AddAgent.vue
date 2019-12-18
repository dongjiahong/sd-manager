<template>
  <div class="add-agent">
    <Modal
      v-model="show"
      title="添加代理"
      :loading="loading"
      @on-ok="asyncOK('agentFormValidate')"
    >
      <Form
        ref="agentFormValidate"
        :model="agentFormValidate"
        :label-width="80"
      >
        <FormItem
          label="姓名"
          prop="name"
          :rules="{ required: true, message: '姓名必填', trigger: 'blur' }"
        >
          <i-input
            v-model="agentFormValidate.name"
            placeholder="输入代理姓名"
          ></i-input>
        </FormItem>
        <FormItem
          label="账户"
          prop="account"
          :rules="{
            required: true,
            message: '账户必填，推荐姓名拼音如：小王->xiaowang',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="agentFormValidate.account"
            placeholder="输入代理账户"
          ></i-input>
        </FormItem>
        <FormItem
          label="密码"
          prop="password"
          :rules="{
            required: true,
            min: 6,
            message: '密码必填不少于6位',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="agentFormValidate.password"
            placeholder="输入代理密码"
          ></i-input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "AddAgent",
  props: {
    showModal: Boolean
  },
  data() {
    return {
      show: false, // show 控制模态框是否展示
      loading: true,
      agentFormValidate: {
        name: "", // 代理姓名
        account: "", // 代理账户
        password: "" // 账户密码
      }
    };
  },
  methods: {
    ...mapActions(["addAgent"]),
    asyncOK(name) {
      this.$refs[name].validate(valid => {
        if (!valid) {
          this.$Message.warning("请检查你填写的数据");
        } else {
          this.addAgent({
            me: this,
            data: {
              agent_name: this.agentFormValidate.name,
              agent_account: this.agentFormValidate.account,
              agent_password: this.agentFormValidate.password
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
