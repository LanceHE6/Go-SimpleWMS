<template>
<div
    v-loading="state.isLoading"
    class="profile-body"
>
  <el-card class="profile-title-card">
    <el-container>
      <el-aside class="title-aside">
        <el-avatar :size="90" src="https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png" />
      </el-aside>
      <el-main class="title-main">
        <div class="title-nickname">
          <b>{{user.nickname}}</b>
          <el-tag
            size="small"
            :type="getTagType(user.permission)"
            style="margin-left: 20px"
          >
            {{getTagContent(user.permission)}}
          </el-tag>
          <el-button
              type="primary"
              size="small"
              style="margin-left: auto"
              icon="Edit"
              @click="openEditForm"
          >
            编辑
          </el-button>
        </div>
        <div class="small-text" style="margin-bottom: 10px">{{`用户ID：${user.uid}`}}</div>
        <div class="small-text">{{`创建时间：${new Date(user.created_at).toLocaleString()}`}}</div>
      </el-main>
    </el-container>
  </el-card>

  <el-card class="profile-body-card">
    <div class="title-nickname">
      <b>个人资料</b>
    </div>
    <el-divider></el-divider>
    <div
        class="body-data"
        v-for="item in dataCol"
    >
      <span class="data-label">{{item.label}}</span>
      <span v-if="user[item.property]" class="data-prop">{{mapping(item, user[item.property])}}</span>
      <span v-else class="data-prop-null">未填写</span>
      <span v-if="item.isBind">
        <el-button
          type="primary"
          size="small"
          style="margin-left: 15px"
          @click="openEmailForm"
          text
        >
          {{user[item.property] ? `换绑${item.label}` : `绑定${item.label}`}}
        </el-button>
      </span>
    </div>
  </el-card>
</div>

<el-dialog
    v-model="editFormVisible"
    title="个人信息编辑"
    width="500"
    center
>

  <el-form :model="editForm.data" :rules="editForm.rules" ref="myEditForm" label-position="top" status-icon>

    <el-form-item
        v-for="item in editForm.item"
        :label="item.label"
        :prop="item.prop"
    >
      <el-input
          v-if="item.isInput"
          v-model.trim="editForm.data[item.dataName]"
          :type="item.type"
          :show-password="item.type === 'password'"
          autocomplete="off"
      />
    </el-form-item>
  </el-form>

  <template #footer>
    <div class="dialog-footer">
      <el-button @click="editFormVisible = false">取消</el-button>
      <el-button type="primary" @click="submitEditForm(myEditForm)">
        确定
      </el-button>
    </div>
  </template>

</el-dialog>
<el-dialog
    v-model="emailFormVisible"
    title="邮箱绑定"
    width="500"
    center
>

  <el-form :model="emailForm.data" :rules="emailForm.rules" ref="myEmailForm" label-position="top" status-icon>

    <el-form-item
        v-for="item in emailForm.item"
        :label="item.label"
        :prop="item.prop"
    >
      <el-input
          v-if="item.isInput"
          v-model.trim="emailForm.data[item.dataName]"
          :type="item.type"
          :show-password="item.type === 'password'"
          autocomplete="off"
      />
      <div
        class="verification-input"
        v-if="item.isVerification"
      >
        <el-input
            v-model.trim="emailForm.data[item.dataName]"
            :type="item.type"
            :show-password="item.type === 'password'"
            autocomplete="off"
        />
        <el-button
          type="primary"
          style="margin-left: 30px"
          :loading="state.isVerificationBtnLoading"
          :disabled="countDown > 0"
          @click="sendVerificationCode(myEmailForm)"
          plain
        >
          发送验证码
          {{countDown > 0 ? countDown : ''}}
        </el-button>
      </div>
    </el-form-item>
  </el-form>

  <template #footer>
    <div class="dialog-footer">
      <el-button @click="emailFormVisible = false">取消</el-button>
      <el-button type="primary" @click="submitEmailForm(myEmailForm)">
        确定
      </el-button>
    </div>
  </template>

</el-dialog>
</template>

<script setup>
import {CURRENT_USER, setUser} from "@/utils/appManager.js";
import {reactive, ref} from "vue";
import {isSame, isPasswordValid, isNotSame, isNotEmptyCondition, isEmail} from "@/utils/validator.js";
import {axiosGet, axiosPost, axiosPut} from "@/utils/axiosUtil.js";
import {ElMessage} from "element-plus";

const user = ref(CURRENT_USER.value)

const myEditForm = ref(null)
const editFormVisible = ref(false)

const myEmailForm = ref(null)
const emailFormVisible = ref(false)

//发送验证码计时器
const countDown = ref(0);

//用户信息显示格式列表
const dataCol = [
  {property: "nickname", label: "昵称"},
  {property: "account", label: "账号"},
  {property: "uid", label: "用户ID"},
  {property: "phone", label: "手机号"},
  {property: "email", label: "邮箱", isBind: true},
  {property: "updated_at", label: "最近登录", isDate: true},
]

//状态
const state = reactive({
  isLoading: false,  //页面是否正在加载
  isVerificationBtnLoading: false,  //发送验证码按钮是否正在加载
})

//用户信息编辑窗口数据
const editForm = reactive({
  data :{
    uid:'',
    permission: 3,
    old_password:'',
    new_password:'',
    confirm:'',
    nickname:'',
    phone:''
  },
  dataType:{
    uid:'String',
    permission: 'Int',
    old_password:'String',
    new_password:'String',
    confirm:'String',
    nickname:'String',
    phone:'String'
  },
  dataNum: 7,
  rules: {
    old_password: [
      { min: 6, max: 16, message: '密码长度需要在6-16之间', trigger: 'blur' },
      { validator: isPasswordValid, trigger: 'blur' },
      { validator: (rule, value, callback) => isNotEmptyCondition(rule, value, callback, editForm.data.new_password, '旧密码不能为空'), trigger: 'blur' },
    ],
    new_password: [
      { min: 6, max: 16, message: '密码长度需要在6-16之间', trigger: 'blur' },
      { validator: isPasswordValid, trigger: 'blur' },
      { validator: (rule, value, callback) => isNotSame(rule, value, callback, editForm.data.old_password, '新密码与旧密码不能相同'), trigger: 'blur' },
    ],
    confirm: [
      { validator: (rule, value, callback) => isSame(rule, value, callback, editForm.data.new_password, '两次输入的密码不一致'), trigger: 'blur' },
    ],
    nickname:[
      { required: 'true', message: '请输入用户昵称', trigger: 'blur' },
      { min: 1, max: 8, message: '昵称长度需要在1-8个字符之间', trigger: 'blur' },
    ]
  },
  item:[
    {label: '昵称', prop: 'nickname', dataName: 'nickname', isInput: true,},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
    {label: '旧密码', prop: 'old_password', dataName: 'old_password', isInput: true, type: 'password'},
    {label: '新密码', prop: 'new_password', dataName: 'new_password', isInput: true, type: 'password'},
    {label: '确认密码', prop: 'confirm', dataName: 'confirm', isInput: true, type: 'password'},
  ],
})

//邮箱绑定窗口数据
const emailForm = reactive({
  data :{
    uid:'',
    email:'',
    code:''
  },
  dataType:{
    uid:'String',
    email:'String',
    code:'String'
  },
  dataNum: 3,
  rules: {
    email: [
      { required: 'true', message: '邮箱不能为空', trigger: 'blur' },
      { validator: (rule, value, callback) => isEmail(rule, value, callback), trigger: 'blur' },
    ],
    code: [
      { required: 'true', message: '验证码不能为空', trigger: 'blur' },
    ],
  },
  item:[
    {label: '邮箱', prop: 'email', dataName: 'email', isInput: true,},
    {label: '验证码', prop: 'code', dataName: 'code', isVerification: true},
  ],
})

const mapping = (item, data) => {
  if(item.isDate){
    return new Date(data).toLocaleString()
  }
  else{
    return data
  }
}

const getTagType = (permission) => {
  switch (permission){
    case 1:{
      return 'primary'
    }
    case 2:{
      return 'warning'
    }
    case 3:{
      return 'danger'
    }
    default:{
      return 'info'
    }
  }
}
const getTagContent = (permission) => {
  switch (permission){
    case 1:{
      return '普通用户'
    }
    case 2:{
      return '管理员'
    }
    case 3:{
      return '超级管理员'
    }
    default:{
      return '未知'
    }
  }
}

const openEditForm = () => {
  for(const i in editForm.data){
    if(user.value[i]){
      editForm.data[i] = user.value[i]
    }
  }
  editFormVisible.value = true
}

const submitEditForm = async (form) => {
  if (!form) return
  await form.validate(async (valid) => {
    if (valid) {
      editFormVisible.value = false
      const result = await axiosPut({
        url:'/user/update',
        data: editForm.data,
        name: 'userCenter-edit'
      })
      if(result){
        ElMessage.success("用户信息修改成功")
        await refresh()
      }
    }
  })
}

const openEmailForm = () => {
  for(const i in emailForm.data){
    if(user.value[i]){
      emailForm.data[i] = user.value[i]
    }
  }
  emailFormVisible.value = true
}

const submitEmailForm = async (form) => {
  if (!form) return
  await form.validate(async (valid) => {
    if (valid) {
      emailFormVisible.value = false
      const result = await axiosPost({
        url:'/user/email/verify',
        data: emailForm.data,
        name: 'userCenter-bind'
      })
      if(result){
        ElMessage.success("邮箱绑定成功")
        emailForm.data.code = ''
        await refresh()
      }
    }
  })
}

const refresh = async () => {
  state.isLoading = true
  const result = await axiosGet({
    url: '/user/info',
    name: 'userCenter-refresh'
  })
  if(result && result.data && result.data.user){
    user.value = result.data.user
    setUser(result.data.user)
  }
  state.isLoading = false
}

const sendVerificationCode = async (form) => {
  if (!form) return
  await form.validateField(['email'], async (valid) => {
    if (valid) {
      state.isVerificationBtnLoading = true
      const result = await axiosPost({
        url: '/user/email/bind',
        data: emailForm.data,
        name: 'userCenter-sendVerification'
      })
      state.isVerificationBtnLoading = false
      if(result){
        ElMessage.success("验证码已发送")
        countDown.value = 60
        let intervalId = null
        intervalId = setInterval(() => {
          if (countDown.value > 0) {
            countDown.value--
          } else {
            clearInterval(intervalId)
          }
        }, 1000);
      }
    }
  })
}
</script>

<style scoped>
.profile-body {
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 子元素在父容器中垂直分布 */
  align-items: center;
  background-image: url("@/res/profile-background.png");
  width: 100%;
  height: 100%;
  background-size: cover;
}
.profile-title-card{
  margin: 30px 0 0 0;
  width: 940px;
}
.profile-body-card{
  margin: 10px 0 0 0;
  width: 940px;
}
.title-aside{
  margin-left: 15px;
  width: auto;
}
.title-main{
  justify-content: flex-start;
  padding: 0 0 0 15px;
}
.title-nickname{
  display: flex;
  justify-content: space-between;
  align-items: center; /* 垂直居中 */
  font-size: 16px;
  margin-bottom: 10px;
}
.small-text{
  font-size: 12px;
  color: #8c939d;
}
.body-data{
  display: flex;
  margin: 40px 0 40px 10px;
  justify-content: flex-start;
}
.data-label{
  display: flex;
  align-items: center; /* 垂直居中 */
  width: 100px;
  font-size: 15px;
}
.data-prop{
  display: flex;
  align-items: center; /* 垂直居中 */
  font-size: 15px;
}
.data-prop-null{
  display: flex;
  align-items: center; /* 垂直居中 */
  color: #8c939d;
}
.verification-input{
  width: 100%;
  display: flex;
  justify-content: flex-start;
}
</style>