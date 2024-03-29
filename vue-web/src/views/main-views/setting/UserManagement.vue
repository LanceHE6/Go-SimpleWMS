<template>
用户管理
</template>

<script setup>
import axios from "axios";
import {ElMessage} from "element-plus";

/**
 *接收用户对象的列表
 * */
let userArray;
/**
 * 用于用户注册的对象
 * */
const user={
  account:'',
  password:'',
  permission:'',
  nick_name:'',
  phone:''
}

/**
 * 修改用户时所用到的对象
 * */
const userNew={
  account:'',
  password:'',
  permission:'',
  nick_name:'',
  phone:''
}


const token="bearer "+localStorage.getItem("token");

/**
 * userList()
 * 获取用户的信息的请求
 * 打开网页时自动调用一次
 * 结果：userArray中包含一到多个user对象
 * */

const userList = async () => {
  const array = await axios.get('/user/list', {
    headers: {
      'Authorization': token
    }
  }).catch(error => {
    ElMessage.error("网络请求出错了！")
    console.log(error.message)
  })

  if (array && array.data) {
    userArray = array.data.rows;
  }
  console.log(userArray)
}
userList()

/**
 * userDelete()
 * 删除用户的请求，param: uid 类型:String
 * */
const userDelete=async (uid) => {
  let result = await axios({
    method: 'delete',
    url: '/user/delete',
    headers: {
      'Content-Type': 'multipart/form-data',
      'Authorization': token
    },
    data: {
      "uid":uid
    }
  }).catch( error => {
    ElMessage.error("网络请求出错了！")
    console.log(error.message)
  })
  console.log(result)
}

/**
 * userRegister()
 * 新增用户 param: userNew对象
 * userNew={
 *     account:'',
 *     password:'',
 *     permission:'',
 *     nick_name:'',
 *     phone:''
 * }
 * */
const userRegister=async (user) => {
  let result = await axios({
    method: 'post',
    url: '/user/register',
    headers: {
      'Content-Type': 'multipart/form-data',
      'Authorization': token
    },
    data: user
  }).then(()=>{
        userNew.account=''
        userNew.nick_name=''
        userNew.password=''
        userNew.phone=''
        userNew.permission=''
      }
  ).catch( error => {
    ElMessage.error("网络请求出错了！")
    console.log(error.message)
  })
  console.log(result)
}

/**
 * userUpdate()
 * 修改用户  param: userNew对象
 * userNew={
 *     account:'',
 *     password:'',
 *     permission:'',
 *     nick_name:'',
 *     phone:''
 * */
const userUpdate=async (userNew) => {
  let result = await axios.put('/user/update', userNew, {
        headers: {
          'Authorization': token
        },
      }
  ).catch( error => {
    ElMessage.error("网络请求出错了！")
    console.log(error.message)
  })
  console.log(result)
}
</script>

<style scoped>

</style>