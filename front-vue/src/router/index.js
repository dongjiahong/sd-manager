import Vue from "vue";
import Router from "vue-router";

import Home from "@/components/Home";
import Load from "@/components/Load";
import Backup from "@/components/Backup";
import Account from "@/components/Account";
import Agent from "@/components/Agent";
import Machine from "@/components/Machine";

Vue.use(Router);

export default new Router({
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
          path: "agents/:name",
          component: Agent
        },
        {
          path:"machines/:state",
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
