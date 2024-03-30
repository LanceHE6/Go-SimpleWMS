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
      <my-tab
        ref="myTab"
        :default-tab="defaultTab"
      >
      </my-tab>
    </el-main>

  </el-container>
</template>



<script setup>
import SideBar from "@/components/SideBar.vue";
import {onMounted, ref} from "vue";
import MyTab from "@/components/MyTab.vue";

onMounted(init)

//侧边菜单内容
const sideMenu = [
  {name: "用户管理", index: "/home/setting/userManagement", icon: "User"},
];
//初始tab内容
const defaultTab = {
    label: '用户管理',
    path: '/home/setting/userManagement',
  }

let myTab = ref(null);

//初始化函数
function init() {}

//点击侧边栏菜单
function handleSelect(menu){
  myTab.value.addTab(menu.name, menu.key)
}

</script>

<style scoped>
.side-menu {
  height: 100%;
  border: 0 !important;
}
.main-container{
  padding: 5px;
}
</style>