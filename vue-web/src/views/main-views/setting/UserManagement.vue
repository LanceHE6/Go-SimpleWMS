<template>
  <my-table
    ref="myTable"
    :table-col-list="tableColList"
    :default-data="data.userArray"
    :search-data="'nickname'"
    :add-data-template="addForm"
    :edit-data-template="editForm"
    :loading="data.isLoading"
    @add="add"
    @upload="upload"
    @del="del"
    @edit="edit"
  >
  </my-table>
</template>

<script setup>
import axios from "axios";
import {ElMessage, ElNotification} from "element-plus";
import MyTable from "@/components/MyTable.vue";
import {onMounted, reactive} from "vue";

//表头属性列表
const tableColList = [
  {property: "account", label: "账号", sortable: false},
  {property: "nickname", label: "昵称", sortable: false},
  {property: "permission", label: "权限", sortable: true},
  {property: "phone", label: "手机号码", sortable: false},
  {property: "register_time", label: "注册时间", sortable: true},
  {property: "uid", label: "用户ID", sortable: true, width: 100},
]

//初始化函数
onMounted(() => {
  userList()
})

//点击子组件的添加按钮, 子组件处理完返回的可提交表单
function add(form){
  userRegister(form)
}

//点击子组件的上传按钮, 子组件处理完返回的可提交表单
function upload(form){
  userUpload(form)
}

//点击子组件的删除按钮
function del(row){
  userDelete(row.uid)
}

//点击子组件的编辑按钮, 子组件处理完返回的可提交表单
function edit(form){
  userUpdate(form)
}

const data =  reactive({
  userArray: [],  //接收用户对象的列表
  isLoading: true  //是否正在加载
})

/**
 * 编辑用户时所用到的表单对象
 * */
const editForm = reactive({
  data :{
    uid:'',
    password:'',
    permission:0,
    nickname:'',
    phone:''
  },
  dataNum: 5,
  rules: {
    password: [
      { min: 6, max: 16, message: '密码长度需要在6-16之间', trigger: 'blur' },
    ],
    permission:[
      { required: 'true', message: '请选择用户权限', trigger: 'blur' }
    ],
    nickname:[
      { required: 'true', message: '请输入用户昵称', trigger: 'blur' },
      { min: 1, max: 10, message: '昵称长度需要在1-10个字符之间', trigger: 'blur' },
    ]
  },
  item:[
    {label: '昵称', prop: 'nickname', dataName: 'nickname', isInput: true,},
    {label: '权限', prop: 'permission', dataName: 'permission', isSelect: true,
      selectOptions: [
        {label: '普通用户', value: 1},
        {label: '管理员', value: 2},
        {label: '超级管理员', value: 3},
      ]},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
    {label: '新密码', prop: 'password', dataName: 'password', isInput: true, type: 'password'},
  ],
  key: 'uid'
})

/**
 * 注册用户时所用到的对象
 * */
const addForm = reactive({
  data :{
    account:'',
    password:'',
    permission:'',
    nickname:'',
    phone:''
  },
  dataNum: 5,
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
  ,
  item:[
    {label: '账号', prop: 'account', dataName: 'account', isInput: true, type: 'text'},
    {label: '密码', prop: 'password', dataName: 'password', isInput: true, type: 'password'},
    {label: '昵称', prop: 'nickname', dataName: 'nickname', isInput: true},
    {label: '权限', prop: 'permission', dataName: 'permission', isSelect: true,
      selectOptions: [
        {label: '普通用户', value: 1},
        {label: '管理员', value: 2},
        {label: '超级管理员', value: 3},
      ]},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
  ],
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
      data.warehouseArray = result.data.rows;
    }
  })
  .catch(error => {
    ElMessage.error("网络请求出错了！")
    console.error("userList:", error)
  })
  data.isLoading = false
}

/**
 * userDelete()
 * 删除用户的请求，param: uid 类型:String
 * */
const userDelete=async (uid) => {
  data.isLoading = true
  await axios.delete('/user/delete', {
        headers: {
          'Authorization': token
        },
        data: {
          "uid": uid
        }
      }
  )
  .then( async message => {
    ElMessage.success("用户已被删除！")
    console.log("userDelete:", message)
    await userList()
  })
  .catch( error => {
    ElMessage.error("网络请求出错了！")
    console.error("userDelete:", error)
  })
  data.isLoading = false
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
  data.isLoading = true
  await axios.post('/user/register', user, {
    headers: {
      'Authorization': token
    }
  })
  .then( async message => {
    ElMessage.success("用户注册成功！")
    console.log("userRegister:", message)
    await userList()
  })
  .catch( error => {
    ElMessage.error("网络请求出错了！用户名是否已被注册？")
    console.error("userRegister:", error)
  })
  data.isLoading = false
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
  data.isLoading = true
  await axios.put('/user/update', userNew, {
        headers: {
          'Authorization': token
        },
      }
  ).then( async message => {
    ElMessage.success("用户信息修改成功！")
    console.log("userUpdate:", message)
    await userList()
  })
  .catch( error => {
    ElMessage.error("网络请求出错了！")
    console.error("userUpdate:", error)
  })
  data.isLoading = false
}


const userUpload=async (list) => {
  console.log(list)
  data.isLoading = true
  await axios.post('/user/register-batch', {user_list: list}, {
    headers: {
      'Authorization': token
    }
  })
      .then( async message => {
        if(message.status === 200) {
          ElMessage.success("批量导入成功！")
        }
        else if(message.status === 217){
          ElMessage.warning("部分导入成功")
          let errMessage = ""
          for(const item in message.data.detail){
            if(message.data.detail[item].code === 402){
              errMessage += (message.data.detail[item].message + "\n")
              ElNotification({
                title: '失败信息',
                message: message.data.detail[item].message,
                type: 'error',
                duration: 30000,
              })
            }
          }
        }
        console.log("userUpload:", message)
        await userList()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！用户名是否已被注册？")
        console.error("userUpload:", error)
      })
  data.isLoading = false
}
</script>

<style scoped>

</style>