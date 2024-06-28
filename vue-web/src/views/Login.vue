<template>
  <div class="login-body">
    <div class="login-container">
      <div class="head">
        <img class="logo" src="../res/WMS-Logo.png"  alt="原神，启动！"/>
        <div class="name">
          <div class="title">Go-SimpleWMS</div>
          <div class="tips">简行云仓库</div>
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
          <el-button type="primary" text style="margin-left: 25px" @click="state.settingFormVisible = true">API设置</el-button>
          <el-button type="primary" text style="margin-left: 20px" @click="
            state.forgetForm.account = state.ruleForm.account;
            state.forgetFormVisible = true"
          >
            找回密码
          </el-button>
          <el-button style="width: 100%; margin-top: 20px" type="primary" @click="submitForm(loginForm)" :loading="state.loading" round>
            <el-text style="color: white">立即登录</el-text>
          </el-button>
        </el-form-item>

      </el-form>
    </div>
  </div>
  <el-dialog
      v-model="state.settingFormVisible"
      title="API设置"
      width="500"
      center
  >

    <el-form :model="state.settingForm" :rules="state.rules" ref="mySettingForm" label-position="top" status-icon>

      <el-form-item
          label="API"
          prop="api"
      >
        <el-input
            v-model.trim="state.settingForm.api"
            autocomplete="off"
        />

      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="state.settingFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitSettingForm">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>

  <el-dialog
      v-model="state.forgetFormVisible"
      title="找回密码"
      width="500"
      center
  >

    <el-form :model="state.forgetForm" :rules="state.rules" ref="myForgetForm" label-position="top" status-icon>

      <el-form-item
          label="账号"
          prop="account"
      >
        <el-input
            v-model.trim="state.forgetForm.account"
            autocomplete="off"
        />
      </el-form-item>

      <el-form-item
          label="邮箱"
          prop="email"
      >
        <el-input
            v-model.trim="state.forgetForm.email"
            autocomplete="off"
        />
      </el-form-item>

      <el-form-item
          label="验证码"
          prop="code"
      >
      <div
          class="login-verification-input"
      >
        <el-input
            v-model.trim="state.forgetForm.code"
            autocomplete="off"
        />
        <el-button
            type="primary"
            style="margin-left: 30px"
            :loading="state.isVerificationBtnLoading"
            :disabled="state.countDown > 0"
            @click="sendVerificationCode(myForgetForm)"
            plain
        >
          发送验证码
          {{state.countDown > 0 ? state.countDown : ''}}
        </el-button>
      </div>
      </el-form-item>

      <el-form-item
          label="新密码"
          prop="new_password"
      >
        <el-input
            type="password"
            v-model.trim="state.forgetForm.new_password"
            autocomplete="off"
        />
      </el-form-item>

      <el-form-item
          label="确认密码"
          prop="confirm"
      >
        <el-input
            type="password"
            v-model.trim="state.forgetForm.confirm"
            autocomplete="off"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="state.forgetFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForgetForm">
          确定
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref, reactive, inject, onMounted} from "vue";
import {router} from "@/router/index.js";
import { ElMessage } from 'element-plus'
import {Lock, User} from "@element-plus/icons-vue";
import {axiosPost} from "@/utils/axiosUtil.js";
import {isPasswordValid, isSame} from "@/utils/validator.js";

const axios = inject("axios")
const loginForm = ref(null)
const myForgetForm = ref(null)

onMounted(initialize)
const state = reactive({
  ruleForm: {
    account: '',
    password: ''
  },
  settingForm:{
    api: axios.defaults.baseURL
  },
  forgetForm:{
    account: '',
    email: '',
    code: '',
    new_password: '',
    confirm: ''
  },
  countDown: 0,
  remember: false,
  loading: false,
  isVerificationBtnLoading: false,
  settingFormVisible: false,
  forgetFormVisible: false,
  rules: {
    account: [
      { required: 'true', message: '账户不能为空', trigger: 'blur' }
    ],
    password: [
      { required: 'true', message: '密码不能为空', trigger: 'blur' },
      { validator: isPasswordValid, trigger: 'blur' }
    ],
    api: [
      { required: 'true', message: 'API不能为空', trigger: 'blur' }
    ],
    email: [
      { required: 'true', message: '邮箱不能为空', trigger: 'blur' }
    ],
    new_password: [
      { required: 'true', message: '密码不能为空', trigger: 'blur' },
      { validator: isPasswordValid, trigger: 'blur' }
    ],
    confirm: [
      { validator: (rule, value, callback) => isSame(rule, value, callback, state.forgetForm.new_password), trigger: 'blur' }
    ],
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

function submitSettingForm(){
  state.settingFormVisible = false
  axios.defaults.baseURL = state.settingForm.api
  localStorage.setItem("url", state.settingForm.api)
  ElMessage.success("API设置成功!")
}

async function sendVerificationCode(form){
  if (!form) return
  await form.validateField(['account', 'email'], async (valid) => {
    if (valid) {
      state.isVerificationBtnLoading = true
      const result = await axiosPost({
        url: '/user/psw/reset',
        data: state.forgetForm,
        name: 'login-sendVerification'
      })
      state.isVerificationBtnLoading = false
      if(result){
        ElMessage.success("验证码已发送")
        state.countDown = 60
        let intervalId = null
        intervalId = setInterval(() => {
          if (state.countDown > 0) {
            state.countDown--
          } else {
            clearInterval(intervalId)
          }
        }, 1000);
      }
    }
  })
}

async function submitForgetForm() {
  state.forgetFormVisible = false
  const result = await axiosPost({
    url: '/user/psw/verify',
    data: state.forgetForm,
    name: 'login-resetPassword'
  })
  if(result){
    ElMessage.success("密码重置成功")
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
        const result = await axiosPost({url: '/user/login', data: data, name: 'login'})
        if(result && result.data.code === 201){
          // 需要将返回的数据存入Store中
          localStorage.setItem("token", result.data.data.token)
          localStorage.setItem("user", JSON.stringify(result.data.data.user))
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
          await router.push("/home/homePage")
        }
        else if(result && result.data.code === 202){
          ElMessage.error("账号或密码错误！")
        }
      }
      else {
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
  width: 100vw;
  height: 100vh;
  background: url("@/res/login-background.png") no-repeat center;
  background-size: cover;
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
  font-size: 18px;
  color: #999;
}
.login-form {
  width: 70%;
  margin: 0 auto;
}
.login-verification-input{
  width: 100%;
  display: flex;
  justify-content: flex-start;
}
</style>