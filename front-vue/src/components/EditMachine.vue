<template>
  <div class="edit-machine">
    <Modal
      v-model="show"
      title="编辑机器"
      :loading="loading"
      @on-ok="asyncOK('editMachineFormValidate')"
    >
      <Form
        ref="editMachineFormValidate"
        :model="editMachineFormValidate"
        :label-width="80"
      >
        <FormItem
          label="机器代号"
          prop="machine_no"
          :rules="{
            required: true,
            message: '机器代号, 如：a101',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="editMachineFormValidate.machine_no"
            placeholder="输入机器代号"
            readonly
          ></i-input>
        </FormItem>
        <FormItem
          label="机器ip"
          prop="machine_ip"
          :rules="{
            required: true,
            message: '机器ip, 如：192.168.1.1:6666',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="editMachineFormValidate.machine_ip"
            placeholder="输入机器ip"
          ></i-input>
        </FormItem>
        <FormItem
          label="机器密码"
          prop="machine_password"
          :rules="{
            required: true,
            message: '机器密码, 如：123456',
            trigger: 'blur'
          }"
        >
          <i-input
            v-model="editMachineFormValidate.machine_password"
            placeholder="输入机器密码"
          ></i-input>
        </FormItem>
        <FormItem label="创建日期">
          <!-- 创建日期不可更改 -->
          <i-input
            v-model="editMachineFormValidate.machine_create_date"
            readonly
          ></i-input>
        </FormItem>
        <FormItem label="到期日期">
          <DatePicker
            type="date"
            placeholder="机器的到期日期"
            v-model="editMachineFormValidate.machine_end_date"
          ></DatePicker>
        </FormItem>

        <FormItem label="剩余时间">
          <i-input
            v-model="editMachineFormValidate.machine_time_left"
            placeholder="输入该账号的使用时间"
            readonly
          ></i-input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
import util from "@/util/util";
import { mapActions } from "vuex";
export default {
  name: "EditMachine",
  props: {
    showModal: Boolean,
    // ********* 注意 ******** //
    // 在js中对象和数组是引用类型，指向同一个内存，如果prop是对象或者数组
    // 在子组件内部改变它会影响到父组件, 为了不影响父组件最好使用深度copy,
    machineDetail: Object
  },
  data() {
    return {
      show: false, // show 控制模态框是否展示
      loading: true,
      editMachineFormValidate: {}
    };
  },
  methods: {
    ...mapActions(["editMachine"]),
    asyncOK(name) {
      this.$refs[name].validate(valid => {
        if (!valid) {
          this.$Message.warning("请检查你的输入内容");
        } else {
          this.editMachineFormValidate.machine_end_date = util.formatDate(
            this.editMachineFormValidate.machine_end_date
          );
          console.log("=====> end date: ", this.editMachineFormValidate.machine_end_date)
          this.editMachine({ me: this, data: this.editMachineFormValidate });
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
    machineDetail: function(newVal) {
      // 由于是动态传参，初始化时accountdetail是没有值得，这里监控变化
      this.editMachineFormValidate = newVal;
    }
  }
};
</script>

<style lang="scss" scoped></style>
