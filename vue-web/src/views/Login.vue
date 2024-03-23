<template>
  <div class="login-body">
    <div class="login-container">
      <div class="head">
        <img class="logo" src=""  alt="原神，启动！"/>
        <div class="name">
          <div class="title">Go-SimpleWMS</div>
          <div class="tips">仓库后台管理系统</div>
        </div>
      </div>
      <el-form label-position="top" :rules="state.rules" :model="state.ruleForm" ref="loginForm" class="login-form">
        <el-form-item label="账号" prop="account">
          <el-input class="my-input" type="text" v-model.trim="state.ruleForm.account" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input class="my-input" type="password" v-model.trim="state.ruleForm.password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <br>
          <el-button style="width: 100%" type="primary" @click="submitForm">立即登录</el-button>
          <el-checkbox v-model="state.remember" @change="!state.remember">记住密码</el-checkbox>
          <el-checkbox style="margin-left: 120px" v-model="state.autoLogin" @change="!state.autoLogin">自动登录</el-checkbox>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {ref, reactive, inject} from "vue";
import {UserStore} from "@/store/UserStore";
import {router} from "@/router/index.js";
import { ElMessage } from 'element-plus'

const axios = inject("axios")
const message = inject("message")
const loginForm = ref(null)
import md5 from 'js-md5'
const state = reactive({
  ruleForm: {
    account: '',
    password: ''
  },
  autoLogin: false,
  remember: false,
  rules: {
    account: [
      { required: 'true', message: '账户不能为空', trigger: 'blur' }
    ],
    password: [
      { required: 'true', message: '密码不能为空', trigger: 'blur' }
    ]
  }
})

const user = reactive({
  account: localStorage.getItem("account") || "",
  password: localStorage.getItem("password") || "",
  remember: localStorage.getItem("remember") === '0',
  autoLogin: localStorage.getItem("autoLogin") === '0',
})

const submitForm = async () => {
  const data = {
    account: state.ruleForm.account,
    password: md5.md5(state.ruleForm.password)
  };

  const formData = new URLSearchParams();

  for(let key in data) {
    formData.append(key, data[key]);
  }
  let result
  if (state.ruleForm.account && state.ruleForm.password) {
    result = axios.post('/user/login', formData, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    })
  } else {
    ElMessage.error("账号或密码不能为空！")
    console.error('账号或密码为空');
  }

  if (result.data.code === 200) {
    console.log("登录成功")

    // 需要将返回的数据存入Store中
    UserStore.token = result.data.token
    localStorage.setItem("token", UserStore.token)

    // 记住账号密码
    if (user.remember) {
      localStorage.setItem("account", state.ruleForm.account)
      localStorage.setItem("password", state.ruleForm.password)
      localStorage.setItem("remember", user.remember ? '1' : '0')
      localStorage.setItem("autoLogin", user.autoLogin ? '1' : '0')
    } else {
      localStorage.setItem("account", "")
      localStorage.setItem("password", "")
      localStorage.setItem("remember", "")
      localStorage.setItem("autoLogin", "")
    }

    ElMessage.success("登录成功, token:" + UserStore.token)
    await router.push("/home")
  } else {
    console.log("登录失败")
    ElMessage.error("账号或密码错误")
  }
}
</script>

<style scoped>
.login-body {
  display: flex;
  justify-content: center;
  align-items: center;
  min-width: 100vh;
  min-height: 100vh;
  background: linear-gradient(rgba(255, 255, 255, 0.5), rgba(255, 255, 255, 0.5)), url("../res/login-background.jpeg") no-repeat center;
  background-size: contain;
}
.login-container {
  width: 420px;
  height: 400px;
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 21px 41px 0 rgba(0, 0, 0, 0.2);
}
.head {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 0 20px 0;
}
.head img {
  width: 100px;
  height: 100px;
  margin-right: 20px;
}
.head .title {
  font-size: 28px;
  color: #1BAEAE;
  font-weight: bold;
}
.head .tips {
  font-size: 12px;
  color: #999;
}
.login-form {
  width: 70%;
  margin: -20px auto;
}
.my-input {
  border: 100px aqua;
}
</style>