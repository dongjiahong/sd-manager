<template>
  <div class="add-machine">
    <Modal
      v-model="show"
      title="添加机器"
      :loading="loading"
      @on-ok="asyncOK('machineFormValidate')"
    >
      <Form
        ref="machineFormValidate"
        :model="machineFormValidate"
        :label-width="120"
      >
        <FormItem
          label="机器代号"
          prop="machine_no"
          :rules="{
            required: true,
            message: '账户代号, 如：a101',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="machineFormValidate.machine_no"
            placeholder="输入机器代号"
          ></i-input>
        </FormItem>
        <FormItem
          label="机器IP"
          prop="machine_ip"
          :rules="{
            required: true,
            message: '机器IP必填，如：192.168.1.1:6666',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="machineFormValidate.machine_ip"
            placeholder="输入代理机器IP"
          ></i-input>
        </FormItem>
        <FormItem
          label="机器密码"
          prop="machine_password"
          :rules="{
            required: true,
            min: 6,
            message: '密码必填不少于6位',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="machineFormValidate.machine_password"
            placeholder="输入机器密码"
          ></i-input>
        </FormItem>
        <FormItem
          label="机器创建日期"
          prop="machine_create_date"
          :rules="{ required: true }"
        >
          <DatePicker
            type="date"
            placeholder="机器创建日期"
            v-model="machineFormValidate.machine_create_date"
          ></DatePicker>
        </FormItem>
        <FormItem
          label="机器截止日期"
          prop="machine_end_date"
          :rules="{ required: true }"
        >
          <DatePicker
            type="date"
            placeholder="代理的截止日期"
            v-model="machineFormValidate.machine_end_date"
          ></DatePicker>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "AddMachine",
  props: {
    showModal: Boolean
  },
  data() {
    return {
      show: false, // show 控制模态框是否展示
      loading: true,
      machineFormValidate: {
        machine_no: "", // 账户代号
        machine_ip: "", // 机器ip
        machine_password: "", // 机器密码
        machine_create_date: "", // 机器创建时间
        machine_end_date: "" // 机器截止时间
      }
    };
  },
  methods: {
    ...mapActions(["addMachine"]),
    asyncOK(name) {
      this.$refs[name].validate(valid => {
        if (!valid) {
          this.$Message.warning("请检查你填写的数据");
        } else {
          this.addMachine({
            me: this,
            data: {
              machine_no: this.machineFormValidate.machine_no,
              machine_ip: this.machineFormValidate.machine_ip,
              machine_password: this.machineFormValidate.machine_password,
              machine_create_date: this.machineFormValidate.machine_create_date,
              machine_end_date: this.machineFormValidate.machine_end_date
            }
          });
        }
      });

      this.$refs[name].resetFields();
      this.show = false;
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
