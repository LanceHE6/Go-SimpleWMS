<template>
  <el-menu
      :collapse="state.isCollapse"
      active-text-color="#ffd04b"
      background-color="#545c64"
      class="side-menu"
      text-color="#fff"
      @select="handleSideMenu"
  >
    <!--第一个为折叠按钮-->
    <el-menu-item
        @click="collapse"
    >
      <el-icon v-if="state.isCollapse"><DArrowRight /></el-icon>
      <el-icon v-if="!state.isCollapse"><DArrowLeft /></el-icon>
      <span>折叠</span>
    </el-menu-item>

    <!--其他选项动态设置-->
    <el-menu-item
        v-for="item in menuList"
        :index="item.path"
    >
      <component :is="item.icon" style="margin-right: 5px; height: 20px; width: 20px"></component>
      <span>{{item.label}}</span>
    </el-menu-item>

  </el-menu>
</template>

<script setup>
import {reactive} from "vue";
  import {DArrowLeft, DArrowRight} from "@element-plus/icons-vue";

  const state = reactive({
    isCollapse: false,  //侧边栏是否折叠
  })

  const prop = defineProps({
    menuList: {
      type: Array,
      default: () => [],
      description: '侧边栏的菜单列表，该组件通过v-for遍历这个列表来绘制侧边栏。' +
          '\n列表子元素格式为：{name: "菜单对应vue文件名称", label: "子菜单名称", path: "路由地址", icon: "图标名称(参考el-plus文档)"}'
    }
  });


  function collapse() {
    state.isCollapse = !state.isCollapse
  }

  const emit = defineEmits(["selectMenu"]);

  //点击侧边栏, 返回menu, 包含菜单的显示名以及路由
  const handleSideMenu = (path) => {
    //如果点到折叠按钮则不触发
    if(path === ""){
      return;
    }

    const item = prop.menuList.find(menu => menu.path === path)
    const menu = {
      name: item.name,
      label: item.label,  //菜单名字
      path: path,  //路由
    }
    emit("selectMenu", menu);
  }

</script>

<style scoped>

</style>