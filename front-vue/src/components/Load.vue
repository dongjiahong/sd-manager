<template>
  <div class="load">
    <h1 class="title">远程科技管理系统</h1>
    <div class="from">
      <Form ref="loadForm" :model="loadData" :label-width="80">
        <FormItem
          label="用户名"
          prop="user"
          :rules="{
            required: true,
            min: 1,
            message: '用户名不能为空',
            trigger: 'blur'
          }"
        >
          <i-input
            type="text"
            v-model="loadData.user"
            placeholder="请输入用户名"
          ></i-input>
        </FormItem>
        <FormItem
          label="密码"
          prop="password"
          :rules="{
            required: true,
            min: 6,
            message: '密码最少6位',
            trigger: 'blur'
          }"
        >
          <i-input
            type="password"
            v-model="loadData.password"
            placeholder="请输入密码"
          ></i-input>
        </FormItem>
      </Form>
    </div>
    <div class="btn">
      <Button class="lod" type="primary" @click="loadOK">登录</Button>
      <Button>取消</Button>
    </div>
  </div>
</template>

<script>
import api from "@/fetch/api";

export default {
  name: "Load",
  data() {
    return {
      loadData: {
        user: "",
        password: ""
      }
    };
  },
  methods: {
    loadOK() {
      api.load(this.loadData.user, this.loadData.password).then(data => {
        if (data.status === "ok") {
          this.$router.push("/home");
        } else {
          console.log("==> load info: ", data.msg, " status: ", data.status )
          this.$Message.warning({
            content: "登录失败请检查你的用户名和密码",
            duration: 2
          });
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.load {
  height: 420px;
  width: 420px;
  margin: 100px auto;
  border: 2px solid #555;
  border-radius: 10px;
  text-align: center;
  .title {
    margin: 50px;
  }
  .btn {
    margin-top: 60px;
    .lod {
      margin-right: 80px;
    }
  }
  .from {
    width: 350px;
  }
}
</style>
