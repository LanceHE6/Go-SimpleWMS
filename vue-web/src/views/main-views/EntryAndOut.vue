<!--出入库管理-->
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


    <el-main>
      <el-tabs
          v-model="activeTab"

          type="card"
          class="demo-tabs"
          @tab-remove="removeTab"
          :closable="tabList.length > 1"
      >
        <el-tab-pane
            v-for="item in tabList"
            :label="item.label"
            :name="item.path"
        >
          <router-view></router-view>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<script setup>

import {onMounted, ref, watch} from "vue";
import {router} from "@/router/index.js";
import SideBar from "@/components/SideBar.vue";

onMounted(init)

//侧边菜单内容
const sideMenu = [
  {name: "入库单", index: "/home/entryAndOut/entry", icon: "FolderAdd"},
  {name: "出库单", index: "/home/entryAndOut/out", icon: "FolderRemove"},
  {name: "盘点单", index: "/home/entryAndOut/check", icon: "FolderChecked"},
  {name: "调拨单", index: "/home/entryAndOut/allocate", icon: "Folder"},
];

const activeTab = ref('/home/entryAndOut/entry')
const tabList = ref([
  {
    label: '入库单',
    path: '/home/entryAndOut/entry',
  }
])

function init(){
  router.push(activeTab.value);
}

//监听当前tab界面，发生改变时自动切换路由
watch(activeTab, (newValue) => {
  router.push(newValue);
});

//点击侧边栏菜单
function handleSelect(menu){

  // 检查这个路由是否已经在 tabList 中
  if (!tabList.value.some(tab => tab.path === menu.key)) {
    // 如果不在，就添加到 tabList 中
    tabList.value.push({
      label: menu.name, // 这只是一个示例，你应该根据你的实际情况来设置标签的名称
      path: menu.key,
      closable: true, // 这些 tab 可以被关掉
    });
  }
  // 设置这个路由为当前活动的 tab
  activeTab.value = menu.key;
}

//删除tab
function removeTab(currentTab){
  let tabs = tabList.value;
  let a = activeTab.value;
  //如果当前路径和你要删除的路径一致
  if (currentTab === a) {
    tabs.forEach((tab, index) => {
      //要删除的路径和tabList的路径匹配
      if (tab.path === currentTab) {
        const nextTab = tabs[index + 1] || tabs[index - 1];
        //如果存在，获取nextTab路径
        if (nextTab) {
          a = nextTab.path;
        }
      }
    });
  }
  //重新赋值activeTab
  activeTab.value = a;
  //过滤删除后的tabList
  tabList.value = tabList.value.filter((tab) => tab.path !== currentTab);
}

</script>

<style scoped>
.side-menu {
  height: 100%;
  border: 0 !important;
}
</style>