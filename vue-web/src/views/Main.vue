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
        <el-text size="large" style="color: white; margin-right: 10px;">
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
            用户
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
import axios from "axios";
import {ElMessage, ElMessageBox} from "element-plus";

onMounted(initialize)

const state = reactive({
  nowMenuActive: '',  //当前首部栏界面
  defaultPage: '',  //当前路由界面
  headMenuOffset: 0,  //首部菜单锚点偏移量
})

function help(){
  window.open('https://sr.mihoyo.com/', '_blank');
}

async function about(){
  let go_simpleWMS_version = 'unknown'
  let go_server_version = 'unknown'
  let vue_web_version = 'v0.0.2.20240610_Alpha'
  await axios.get('/ping')
      .then(result => {
        console.log("ping:", result)
        if(result.status === 200){
          go_server_version = result.data.data.version
        }
      })
      .catch(error => {
        ElMessage.error("获取后端版本号失败！")
        console.error("ping:", error.message)
      })
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

  const token="bearer "+localStorage.getItem("token");
  await axios.get('/auth', {
    headers: {
      'Authorization': token
    }
  })
      .then(result => {
        console.log("auth:", result)
        if(result.status === 200){
          //ignore
        }
        else{
          ElMessage.error("操作失败，请先登录！")
          router.push('/')
        }
      })
      .catch(error => {
        ElMessage.error("操作失败，请先登录！")
        console.error("auth:", error.message)
        router.push('/')
      })
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