<template>
  <el-container>
    <el-aside
        class="side-bar"
        width="collapse"
    >
      <SideBar
          class="side-menu"
          :menu-list="sideMenu"
          @selectMenu="handleSelect"
      >
      </SideBar>
    </el-aside>

    <el-main
        class="main-container"
    >
      <router-view/>
    </el-main>

  </el-container>
</template>

<script setup>

import SideBar from "@/components/SideBar.vue";
import {router} from "@/router/index.js";

//侧边菜单内容
const sideMenu = [
  {name: 'Profile', label: "个人资料", path: "/home/userCenter/profile", icon: "UserFilled"},
  {name: 'AccountSetting', label: "账号设置", path: "/home/userCenter/accountSetting", icon: "Edit"},
];
//点击侧边栏菜单
function handleSelect(menu){
  router.push(menu.path)
}
</script>

<style scoped>
.side-menu {
  height: 100%;
  border: 0 !important;
}
.main-container{
  padding: 0;
}
</style>