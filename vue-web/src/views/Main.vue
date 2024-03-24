<template>

  <div class="main-body">

  <el-container
    class="side-container"
  >
    <el-aside
        class="side-bar"
        width="collapse"
    >
    <!--侧边栏-->
      <el-menu
          active-text-color="#ffd04b"
          background-color="#545c64"
          class="el-menu-vertical-demo"
          default-active="2"
          text-color="#fff"
          :collapse="state.isCollapse"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-menu-item
            ref="sideMenu"
            index="1"
            @click="collapse"
        >
          <el-icon v-if="state.isCollapse"><DArrowRight /></el-icon>
          <el-icon v-if="!state.isCollapse"><DArrowLeft /></el-icon>
          <span>折叠</span>
        </el-menu-item>

        <el-menu-item
            index="2"
        >
          <el-icon><lock /></el-icon>
          <span>One</span>
        </el-menu-item>

        <el-menu-item
            index="3"
        >
          <el-icon><document /></el-icon>
          <span>Two</span>
        </el-menu-item>

        <el-menu-item
            index="4"
        >
          <el-icon><setting /></el-icon>
          <span>Three</span>
        </el-menu-item>

      </el-menu>
    </el-aside>

    <el-container
      class="main-container"
    >
      <!--头部-->
      <el-header
        class="head-container"
      >
        <el-menu
            :router="true"
            :default-active="'/home/productManagement'"
            :ellipsis="false"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
        >
          <!--左侧头部-->

            <el-menu-item
                index="/home"
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
                index="/home/entryAndOut"
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
                index="/"
            >
              分析统计
            </el-menu-item>

            <el-menu-item
                index="/"
            >
              系统设置
            </el-menu-item>

          <!--右侧头部-->
          <div class="head-container-right"/>
            <a href="https://ys.mihoyo.com/" title="进入官网" target="_blank">
              <img
                  class="main-logo"
                  src="../res/WMS-Logo.png"
                  alt="Element logo"
              />
            </a>

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
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>

  </el-container>
</div>

</template>





<script setup>
import {onMounted, reactive, ref} from 'vue'
import {router} from "@/router/index.js";
import {DArrowLeft, DArrowRight, Document, HomeFilled, Lock, Setting, User} from "@element-plus/icons-vue";

onMounted(initialize)

const sideMenu = ref(null)
const state = reactive({
  isCollapse: false,
  nowSideIndex: 2,
})

const handleOpen = (key, keyPath) => {
  //console.log(key, keyPath)
}

const handleClose = (key, keyPath) => {
  //console.log(key, keyPath)
}

const handleSelect = () => {

}

function collapse() {
  state.isCollapse = !state.isCollapse
  console.log(sideMenu.value)
  sideMenu.value
}
async function initialize(){
  if((localStorage.getItem("account") || '' ) === '' || (localStorage.getItem("token") || '') === ''){
    await router.push("/")
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
.side-container{
  display: flex;
  max-width: 100%;
}
.side-container .side-bar {
  height: 100vh;
  max-width: 10%;
}
.main-container{
}
.head-container{
  padding: 0;
  width: 100%;
}
.main-logo {
  width: 40px;
  height: 40px;
  margin-top: 8px;
}
.head-container-right {
  flex-grow: 1;
}
.el-menu-vertical-demo {
  height: 100%;
}
</style>