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
  <router-view v-slot="{ Component }">
    <keep-alive :include="cachedViews">
      <component :is="Component" @addTab="addTab" @removeTab="removeTab"/>
    </keep-alive>
  </router-view>
</template>

<script setup>

import {router} from "@/router/index.js";
import {ref, watch} from "vue";
import {objectToUrl} from "@/utils/objectToUrl"

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
const cachedViews = ref([prop.defaultTab.name]); // 缓存的视图

//暴露函数，可供父组件调用
defineExpose({
  addTab,
  removeTab
});

router.push(prop.defaultTab.path);
//监听当前tab界面，发生改变时自动切换路由
watch(() => activeTab.value, (newValue) => {
  router.push(newValue);
});

//添加tab界面
function addTab(name, label, path, params){
  //如果有参数则将其加入到路径中
  if(params !== undefined){
    path = objectToUrl({
      path: path,
      query: params
    })
  }
  // 检查这个路由是否已经在 tabList 中
  if (!tabList.value.some(tab => tab.path === path)) {
    // 如果不在，就添加到 tabList 中
    tabList.value.push({
      name: name,
      label: label, // 设置标签的名称
      path: path,
      closable: true, // tab可以被关掉
    });
    // 将新的视图添加到缓存中
    cachedViews.value.push(name);
  }
  // 设置这个路由为当前活动的 tab
  activeTab.value = path;
}

//删除tab
function removeTab(currentTab){
  let tabs = tabList.value;
  let a = activeTab.value;
  let currentTabName;
  //如果当前路径和你要删除的路径一致
  if (currentTab === a) {
    tabs.forEach((tab, index) => {
      //要删除的路径和tabList的路径匹配
      if (tab.path === currentTab) {
        currentTabName = tab.name
        const nextTab = tabs[index + 1] || tabs[index - 1];
        //如果存在，获取nextTab路径
        if (nextTab) {
          a = nextTab.path;
        }
      }
    });
  }
  // 重新赋值activeTab
  activeTab.value = a;
  // 过滤删除后的tabList
  tabList.value = tabList.value.filter((tab) => tab.path !== currentTab);
  // 从缓存中移除视图
  cachedViews.value = cachedViews.value.filter((view) => view !== currentTabName);
}
</script>

<style scoped>

</style>