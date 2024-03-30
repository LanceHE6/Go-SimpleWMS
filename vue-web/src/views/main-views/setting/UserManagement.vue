<template>
  <my-table
    ref="myTable"
    v-loading="data.isLoading"
    :table-col-list="tableColList"
    :default-data="data.userArray"
    @add="add"
    @upload="upload"
    @download="download"
    @del="del"
    @edit="edit"
  >
  </my-table>

  <el-dialog
      v-model="editFormVisible"
      title="用户编辑"
      width="700"
      center
  >

    <el-form :model="editForm.user" :rules="editForm.rules" ref="myEditForm" label-position="top" status-icon>

      <el-container>

        <el-aside style="width: 50%; padding: 20px">
          <el-form-item label="昵称" :label-width="formLabelWidth" prop="nickname">
            <el-input
                v-model.trim="editForm.user.nickname"
                autocomplete="off"
                maxlength="10"
                show-word-limit/>
          </el-form-item>
          <el-form-item label="用户权限" :label-width="formLabelWidth" type="number" prop="permission">
            <el-select v-model.trim="editForm.user.permission" placeholder="请选择">
              <el-option label="普通用户" value="1" />
              <el-option label="管理员" value="2" />
              <el-option label="超级管理员" value="3" />
            </el-select>
          </el-form-item>
          <el-form-item label="手机" :label-width="formLabelWidth">
            <el-input
                v-model.trim="editForm.user.phone"
                autocomplete="off"
                type="number"
                maxlength="20"
                show-word-limit/>
          </el-form-item>
        </el-aside>

        <el-main style="width: 50%">
          <el-form-item label="新密码" :label-width="formLabelWidth" prop="password">
            <el-input
                v-model.trim="editForm.user.password"
                autocomplete="off"
                type="password"
                maxlength="16"
                show-password/>
          </el-form-item>
        </el-main>

      </el-container>
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
      v-model="addFormVisible"
      title="新增用户"
      width="700"
      center
  >

    <el-form :model="addForm.user" :rules="addForm.rules" ref="myAddForm" label-position="top" status-icon>

      <el-container>

        <el-aside style="width: 50%; padding: 20px">
          <el-form-item label="昵称" :label-width="formLabelWidth" prop="nickname">
            <el-input
                v-model.trim="addForm.user.nickname"
                autocomplete="off"
                maxlength="10"
                show-word-limit/>
          </el-form-item>
          <el-form-item label="用户权限" :label-width="formLabelWidth" type="number" prop="permission">
            <el-select v-model.trim="addForm.user.permission" placeholder="请选择">
              <el-option label="普通用户" value="1" />
              <el-option label="管理员" value="2" />
              <el-option label="超级管理员" value="3" />
            </el-select>
          </el-form-item>
          <el-form-item label="手机" :label-width="formLabelWidth">
            <el-input
                v-model.trim="addForm.user.phone"
                autocomplete="off"
                type="number"
                maxlength="20"
                show-word-limit/>
          </el-form-item>
        </el-aside>

        <el-main style="width: 50%">
          <el-form-item label="账号" :label-width="formLabelWidth" prop="account">
            <el-input
                v-model.trim="addForm.user.account"
                autocomplete="off"
                type="text"
                maxlength="16"/>
          </el-form-item>
          <el-form-item label="密码" :label-width="formLabelWidth" prop="password">
            <el-input
                v-model.trim="addForm.user.password"
                autocomplete="off"
                type="password"
                maxlength="16"
                show-password/>
          </el-form-item>
        </el-main>

      </el-container>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="addFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitAddForm(myAddForm)">
          确定
        </el-button>
      </div>
    </template>

  </el-dialog>
</template>

<script setup>
import axios from "axios";
import {ElMessage} from "element-plus";
import MyTable from "@/components/MyTable.vue";
import {onMounted, reactive, ref} from "vue";

//table组件, 可以调用里面的暴露函数来更新里面的内容
const myTable = ref(null)

//编辑窗口的表单组件
const myEditForm = ref(null)

//新增窗口的表单组件
const myAddForm = ref(null)

//编辑用户窗口是否可见
const editFormVisible = ref(false)

//新增用户窗口是否可见
const addFormVisible = ref(false)

const formLabelWidth = '80px'

const tableColList = [
  {property: "uid", label: "用户ID", sortable: true, width: 100},
  {property: "nickname", label: "昵称", sortable: false},
  {property: "account", label: "账号", sortable: false},
  {property: "phone", label: "手机号码", sortable: false},
  {property: "permission", label: "权限", sortable: true},
  {property: "register_time", label: "注册时间", sortable: true}
]

onMounted(() => {
  userList()
})

function add(){
  addFormVisible.value = true
}
function upload(){}
function download(){}

//点击子组件的删除按钮
function del(row){
  userDelete(row.uid)
}

//点击子组件的编辑按钮
function edit(row){
  editForm.user.uid = row.uid
  editForm.user.permission = (row.permission === 1) ? "普通用户" : (row.permission === 2) ? "管理员" : (row.permission === 3) ? "超级管理员" : "0"
  editForm.user.nickname = row.nickname
  editForm.user.phone = row.phone

}

async function submitEditForm(form){
  if (!form) return
  await form.validate((valid) => {
    if (valid) {
      editForm.user.permission = Number(editForm.user.permission)
      userUpdate(editForm.user)
      editFormVisible.value = false
    }
  })
}

async function submitAddForm(form){
  if (!form) return
  await form.validate((valid) => {
    if (valid) {
      addForm.user.permission = Number(addForm.user.permission)
      userRegister(addForm.user)
      addFormVisible.value = false
    }
  })
}

const data =  reactive({
  userArray: [],  //接收用户对象的列表
  isLoading: true  //是否正在加载
})

/**
 * 编辑用户时所用到的表单对象
 * */
const editForm = reactive({
  user :{
    uid:'',
    password:'',
    permission:0,
    nickname:'',
    phone:''
  },
  rules: {
    password: [
      { min: 6, max: 16, message: '密码长度需要在6-16之间', trigger: 'blur' },
    ],
    permission:[
      { required: 'true', message: '请选择用户权限', trigger: 'blur' }
    ],
    nickname:[
      { required: 'true', message: '请输入用户昵称', trigger: 'blur' }
    ]
  }
})

/**
 * 注册用户时所用到的对象
 * */
const addForm = reactive({
  user :{
    account:'',
    password:'',
    permission:'',
    nick_name:'',
    phone:''
  },
  rules: {
    account: [
      { required: 'true', message: '用户名不能为空', trigger: 'blur' },
      { min: 6, max: 16, message: '用户名长度需要在6-16之间', trigger: 'blur' },
    ],
    password: [
      { required: 'true', message: '密码不能为空', trigger: 'blur' },
      { min: 6, max: 16, message: '密码长度需要在6-16之间', trigger: 'blur' },
    ],
    permission:[
      { required: 'true', message: '请选择用户权限', trigger: 'blur' }
    ],
    nickname:[
      { required: 'true', message: '请输入用户昵称', trigger: 'blur' }
    ]
  }
})

const token="bearer "+localStorage.getItem("token");

/**
 * userList()
 * 获取用户的信息的请求
 * 打开网页时自动调用一次
 * 结果：userArray中包含一到多个user对象
 * */

const userList = async () => {
  data.isLoading = true
  await axios.get('/user/list', {
    headers: {
      'Authorization': token
    }
  })
  .then(result => {
    console.log("userList:", result)
    if (result && result.data) {
      for (let i = 0; i < result.data.rows.length; i++){
        result.data.rows[i].register_time = new Date(Number(result.data.rows[i].register_time) * 1000).toLocaleString()
      }
      data.userArray = result.data.rows;
    }
  })
  .catch(error => {
    ElMessage.error("网络请求出错了！")
    console.error("userList:", error.message)
  })
  data.isLoading = false
}

/**
 * userDelete()
 * 删除用户的请求，param: uid 类型:String
 * */
const userDelete=async (uid) => {
  await axios.delete('/user/delete', {
        headers: {
          'Authorization': token
        },
        data: {
          "uid": uid
        }
      }
  )
  .then( message => {
    ElMessage.success("用户已被删除！")
    console.log("userDelete:", message)
  })
  .catch( error => {
    ElMessage.error("网络请求出错了！")
    console.error("userDelete:", error.message)
  })
  await userList()

}

/**
 * userRegister()
 * 新增用户 param: userNew对象
 * userNew={
 *     account:'',
 *     password:'',
 *     permission:'',
 *     nickname:'',
 *     phone:''
 * }
 * */
const userRegister=async (user) => {
  await axios.post('/user/register', user, {
    headers: {
      'Authorization': token
    }
  })
  .then( message => {
    ElMessage.success("用户注册成功！")
    console.log("userRegister:", message)
  })
  .catch( error => {
    ElMessage.error("网络请求出错了！")
    console.error("userRegister:", error.message)
  })
  await userList()
}

/**
 * userUpdate()
 * 修改用户  param: userNew对象
 * userNew={
 *     uid:'',
 *     password:'',
 *     permission:'',
 *     nickname:'',
 *     phone:''
 * */
const userUpdate=async (userNew) => {
  await axios.put('/user/update', userNew, {
        headers: {
          'Authorization': token
        },
      }
  ).then( message =>{
    ElMessage.success("用户信息修改成功！")
    console.log("userUpdate:", message)
  })
  .catch( error => {
    ElMessage.error("网络请求出错了！")
    console.error("userUpdate:", error.message)
  })
  await userList()
}
</script>

<style scoped>

</style>