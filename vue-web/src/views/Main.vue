<template>

<div class="main-body">
  <el-backtop :right="10" :bottom="50" />
  <!--总容器-->
  <el-container>
    <!--头部-->
    <el-header class="head-container">
      <el-affix
          :offset="state.headMenuOffset"
      >
        <el-menu
          :router="true"
          :ellipsis="false"
          :default-active="state.defaultPage"
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
        <el-text size="large" style="color: white; margin-right: 10px; white-space: nowrap">
          简行云仓库
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
            index="/home/productManagement/productionDefine"

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
            {{state.user?.nickname || '用户'}}
          </template>
          <el-menu-item @click="help">帮助中心</el-menu-item>
          <el-menu-item @click="about">关于</el-menu-item>
          <el-menu-item @click="logout">退出登录</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-affix>

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
import {h, onMounted, reactive} from 'vue'
import {router} from "@/router/index.js";
import {HomeFilled, User} from "@element-plus/icons-vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {axiosGet} from "@/utils/axiosUtil.js";

onMounted(initialize)

const state = reactive({
  nowMenuActive: '',  //当前首部栏界面
  defaultPage: '',  //当前路由界面
  headMenuOffset: 0,  //首部菜单锚点偏移量
  user: null,  //用户
})

function help(){
  window.open('https://sr.mihoyo.com/', '_blank');
}

async function about(){
  let go_simpleWMS_version = 'v0.0.0_Alpha'
  let go_server_version = 'unknown'
  let vue_web_version = 'v0.5.0.20240619_Alpha'
  const result = await axiosGet({url: '/ping'})
  if(result){
    go_server_version = result.data.version
  }
  else{
    ElMessage.error("获取后端版本号失败！")
  }

  await ElMessageBox({
    title: '关于',
    message: h('p', null, [
        h('div', null,[
          h('span', null, 'Go-SimpleWMS '),
          h('b', {style: 'color: teal'}, go_simpleWMS_version)
        ]),
        h('div', null,[
          h('span', null, 'Vue-Web '),
          h('b', {style: 'color: teal'}, vue_web_version)
        ]),
        h('div', null,[
          h('span', null, 'Go-Server '),
          h('b', {style: 'color: teal'}, go_server_version)
        ])
    ]),
    confirmButtonText: '确定',
  })
}

function logout(){
  localStorage.setItem("token", "");
  router.push("/")
}

//点击首部菜单
const handleSelect = (key) => {
  state.nowMenuActive = key;
  pageChange(key)
}

const pageChange = (key = '') => {
  if(key === ''){
    const CURRENT_PATH = window.location.hash  // 获取当前路径，例如 "#/page/subpage"
    const pageList = CURRENT_PATH.split('#')
    state.defaultPage = pageList[pageList.length - 1]
    state.nowMenuActive = state.defaultPage
  }
  else{
    state.nowMenuActive = key
  }
  //homePage中的首部菜单固定在首部, 其他界面不固定
  state.headMenuOffset = state.nowMenuActive === '/home/homePage' ? 0 : -65535
  console.log('pageChange', state.nowMenuActive)
}

//初始化
async function initialize(){
  pageChange()
  const userJson = localStorage.getItem("user") || '';
  if(!userJson){
    ElMessage.error("用户信息获取失败，请重新登录！")
    await router.push('/')
  }
  else{
    state.user = JSON.parse(userJson)
  }
  const result = await axiosGet({url: '/auth', name: 'auth'})
  if(!result){
    await router.push('/')
  }
}
</script>


<style scoped>
.main-body {
  display: flex;
  justify-content: center;
  min-width: 100vh;
  min-height: 100vh;
  background: rgb(255, 255, 255);
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
  width: 100vw;
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