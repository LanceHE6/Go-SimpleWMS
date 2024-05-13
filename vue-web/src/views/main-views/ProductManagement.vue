<!--货品管理-->
<template>
  <el-container>
    <el-aside
        class="side-bar"
        width="collapse"
    >
      <side-bar
          class="side-menu"
          :menu-list="sideMenu"
          @selectMenu="handleSelect"
      >
      </side-bar>

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

import {onMounted, ref} from "vue";
import SideBar from "@/components/SideBar.vue";
import MyTab from "@/components/MyTab.vue";

onMounted(init)

//侧边菜单内容
const sideMenu = [
  {label: "全部货品", path: "/home/productManagement/allProduction", icon: "Box"},
];
//初始tab内容
const defaultTab = {
  label: "全部货品",
  path: "/home/productManagement/allProduction"
}

let myTab = ref(null);

//初始化函数
function init(){}

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