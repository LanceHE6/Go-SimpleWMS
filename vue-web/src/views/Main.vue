<template>

<div class="main-body">
  <!--总容器-->
  <el-container>
    <!--头部-->
    <el-header class="head-container">
      <el-menu
          :router="true"
          :ellipsis="false"
          class="head-menu"
          mode="horizontal"
          @select="handleSelect"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b"
      >
        <!--左侧头部-->
        <a href="https://ys.mihoyo.com/" title="进入官网" target="_blank">
          <img
              class="main-logo"
              src="../res/WMS-Logo.png"
              alt="Logo"
          />
        </a>
        <el-text size="large" style="color: white; margin-right: 10px;">
          SimpleWMS
        </el-text>

        <el-menu-item
            index="/home/homePage"
        >
          <el-icon>
            <home-filled />
          </el-icon>
          首页
        </el-menu-item>

        <el-menu-item
            index="/home/productManagement"

        >
          货品管理
        </el-menu-item>

        <el-menu-item
            index="/home/entryAndOut/entry"
        >
          出入库管理
        </el-menu-item>

        <el-menu-item
            index="/home"
        >
          订单管理
        </el-menu-item>

        <el-menu-item
            index="/home"
        >
          往来单位
        </el-menu-item>

        <el-menu-item
            index="/home"
        >
          分析统计
        </el-menu-item>

        <el-menu-item
            index="/home/setting/userManagement"
        >
          系统设置
        </el-menu-item>

        <!--右侧头部-->
        <div class="head-container-right"/>

        <el-button
            style="margin-top: 18px; margin-left: 15px"
            size="small"
            :disabled="true"
            round
        >
          白金典藏版
        </el-button>
        <el-sub-menu
            index="/"
        >
          <template #title>
            <el-icon>
              <user />
            </el-icon>
            用户
          </template>
          <el-menu-item index="/">one</el-menu-item>
          <el-menu-item index="/">two</el-menu-item>
          <el-menu-item index="/">three</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-header>
      <!--主体-->
      <el-main
        class="main-container"
      >
        <router-view></router-view>
      </el-main>

  </el-container>

</div>

</template>



<script setup>
import {onMounted, reactive} from 'vue'
import {router} from "@/router/index.js";
import {HomeFilled, User} from "@element-plus/icons-vue";

onMounted(initialize)

const state = reactive({
  nowMenuActive: '/home/homePage',  //当前首部栏界面
})


//点击首部菜单
const handleSelect = (key) => {
  state.nowMenuActive = key;
}

//初始化
async function initialize(){
  if((localStorage.getItem("token") || '') === ''){
    await router.push("/")
  }
  else{

  }
}
</script>


<style scoped>
.main-body {
  display: flex;
  justify-content: center;
  min-width: 100vh;
  min-height: 100vh;
  background: linear-gradient(rgba(255, 255, 255, 0.5), rgba(255, 255, 255, 0.5)) no-repeat center;
  background-size: contain;
}
.head-container{
  padding: 0;
  width: 100%;
}
.main-container{
  display: flex;
  padding: 0;
  max-width: 100%;
}
.head-menu {
  border: 0 !important;
}
.main-logo {
  width: 40px;
  height: 40px;
  margin-top: 8px;
  margin-left: 10px;
  margin-right: 10px;
}
.head-container-right {
  flex-grow: 1;
}
</style>