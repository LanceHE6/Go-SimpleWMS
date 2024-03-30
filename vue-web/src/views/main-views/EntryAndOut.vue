<!--出入库管理-->
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
  {name: "入库单", index: "/home/entryAndOut/entry", icon: "FolderAdd"},
  {name: "出库单", index: "/home/entryAndOut/out", icon: "FolderRemove"},
  {name: "盘点单", index: "/home/entryAndOut/check", icon: "FolderChecked"},
  {name: "调拨单", index: "/home/entryAndOut/allocate", icon: "Folder"},
];
//初始tab内容
const defaultTab = {
    label: '入库单',
    path: '/home/entryAndOut/entry',
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