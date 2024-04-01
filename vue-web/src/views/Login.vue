<template>
  <div class="login-body">
    <div class="login-container">
      <div class="head">
        <img class="logo" src="../res/WMS-Logo.png"  alt="原神，启动！"/>
        <div class="name">
          <div class="title">Go-SimpleWMS</div>
          <div class="tips">仓库后台管理系统</div>
        </div>
      </div>
      <el-form label-position="top"
               :rules="state.rules"
               :model="state.ruleForm"
               status-icon
               label-width="auto"
               ref="loginForm"
               class="login-form">

        <el-form-item label="账号" prop="account">
          <el-input
                    type="text"
                    v-model.trim="state.ruleForm.account"
                    autocomplete="off"
                    :prefix-icon="User">

          </el-input>
        </el-form-item>

        <el-form-item label="密码" prop="password">
          <el-input
                    type="password"
                    v-model.trim="state.ruleForm.password"
                    autocomplete="off"
                    :prefix-icon="Lock">

          </el-input>
        </el-form-item>

        <el-form-item>
          <br>
          <el-checkbox v-model="state.remember" @change="!state.remember">记住密码</el-checkbox>
          <el-button style="width: 100%; margin-top: 20px" type="primary" @click="submitForm(loginForm)" :loading="state.loading" round>
            <el-text style="color: white">立即登录</el-text>
          </el-button>
        </el-form-item>

      </el-form>
    </div>
  </div>
</template>

<script setup>
import {ref, reactive, inject, onMounted} from "vue";
import {UserStore} from "@/store/UserStore";
import {router} from "@/router/index.js";
import { ElMessage } from 'element-plus'

const axios = inject("axios")
const loginForm = ref(null)
import {Lock, User} from "@element-plus/icons-vue";

onMounted(initialize)
const state = reactive({
  ruleForm: {
    account: '',
    password: ''
  },
  remember: false,
  loading: false,
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
  account: localStorage.getItem("account") || '',
  password: localStorage.getItem("password") || '',
  remember: localStorage.getItem("remember") || '0',
})

function initialize(){
  if(user.remember === '1') {
    state.ruleForm.account = user.account
    state.ruleForm.password = user.password
    state.remember = true
  }
}


const submitForm = async (form) => {
  form.validate(async (valid) => {
    if (valid) {
      state.loading = true
      const data = {
        account: state.ruleForm.account,
        password: state.ruleForm.password
      };

      if (state.ruleForm.account && state.ruleForm.password) {
        await axios.post('/user/login', data)
            .then(async result => {
              if (result.status === 200) {
                // 需要将返回的数据存入Store中
                UserStore.token = result.data.token
                localStorage.setItem("token", UserStore.token)
                console.log(localStorage.getItem("token"))

                // 记住账号密码
                if (state.remember) {
                  localStorage.setItem("account", state.ruleForm.account)
                  localStorage.setItem("password", state.ruleForm.password)
                  localStorage.setItem("remember", state.remember ? '1' : '0')
                } else {
                  localStorage.setItem("account", '')
                  localStorage.setItem("password", '')
                  localStorage.setItem("remember", '0')
                }

                ElMessage.success("登录成功")
                await router.push("/home/productManagement")
              } else {
                localStorage.setItem("token", "")
                ElMessage.error("账号或密码错误")
              }
              console.log(result)
            })
            .catch(error => {
                  ElMessage.error("网络请求出错了！")
                  console.error(error)
                  state.loading = false
                }
            )
      } else {
        ElMessage.error("账号和密码不能为空！")
      }
      state.loading = false
    }
  })
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
  height: 450px;
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
  color: #16cbcb;
  font-weight: bold;
}
.head .tips {
  font-size: 12px;
  color: #999;
}
.login-form {
  width: 70%;
  margin: 0 auto;
}
</style>