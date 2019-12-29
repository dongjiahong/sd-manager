import Vue from "vue";
import Router from "vue-router";

import Home from "@/components/Home";
import Load from "@/components/Load";
import Backup from "@/components/Backup";
import Account from "@/components/Account";
import Machine from "@/components/Machine";

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: "/",
      name: "load",
      component: Load
    },
    {
      path: "/home",
      name: "home",
      component: Home,
      children: [
        {
          path: "accounts/:state",
          component: Account
        },
        {
          path: "machines/:state",
          component: Machine
        }
      ]
    },
    {
      path: "/backup",
      name: "backup",
      components: Backup
    }
  ]
});
// 使用router.beforeEach 注册全局前置守卫，判断是否登录
router.beforeEach((to, from, next) => {
  if (to.path === "/") {
    next();
  } else {
    let token = localStorage.getItem("Authorization");
    let authDate = localStorage.getItem("AuthDate");
    if (
      token === "null" ||
      token === "" ||
      token === undefined ||
      authDate === "" ||
      authDate === undefined
    ) {
      next("/");
    } else {
      // 获取上次的认证时间
      let begin = new Date(authDate);
      let end = new Date();
      let diff = end.getTime() - begin.getTime();
      if (Math.floor(diff / (3600 * 1000)) > 24) {
        // 相差24小时以上
        next("/");
      } else {
        next();
      }
    }
  }
});
export default router;
