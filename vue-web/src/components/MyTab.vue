<template>
  <el-tabs
      v-model="activeTab"
      type="card"
      @tab-remove="removeTab"
      :closable="tabList.length > 1"
  >
    <el-tab-pane
        v-for="item in tabList"
        :label="item.label"
        :name="item.path"
    />
  </el-tabs>
  <keep-alive>
    <router-view/>
  </keep-alive>
</template>

<script setup>

import {router} from "@/router/index.js";
import {ref, watch} from "vue";

const prop = defineProps({
  defaultTab: {
    type: Object,
    default: () => '',
    description: '默认tab界面对象, 传参时可当做是初始显示界面的路由'+
        '\n对象格式为：  {label: "标签名", path: "标签路由"}'
  }
});

const tabList = ref([prop.defaultTab]);
const activeTab = ref(prop.defaultTab.path)

//暴露函数，可供父组件调用
defineExpose({
  addTab
});

router.push(prop.defaultTab.path);
//监听当前tab界面，发生改变时自动切换路由
watch(() => activeTab.value, (newValue) => {
  router.push(newValue);
});

//添加tab界面
function addTab(name, path){
  // 检查这个路由是否已经在 tabList 中
  if (!tabList.value.some(tab => tab.path === path)) {
    // 如果不在，就添加到 tabList 中
    tabList.value.push({
      label: name, // 设置标签的名称
      path: path,
      closable: true, // tab可以被关掉
    });
  }
  // 设置这个路由为当前活动的 tab
  activeTab.value = path;
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

</style>