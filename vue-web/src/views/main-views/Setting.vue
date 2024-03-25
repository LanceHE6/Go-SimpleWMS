<template>



</template>



<script setup>

import {ElMessage} from "element-plus";
import axios from "axios";


/**
 *接收用户对象的列表
 * */
let userArray;
let array;
/**
 * 获取用户的信息的请求
 * 打开网页时自动调用一次
 * */

const userList=async () => {
    console.log("bearer "+localStorage.getItem("token"))
    array = await axios.get('/user/list', {
        headers: {
            'Authorization': "bearer " + localStorage.getItem("token")
        }
    }).catch( error => {
        ElMessage.error("网络请求出错了！")
        console.log(error.message)
    })

    if (array && array.data) {
        userArray = array.data.row;
    }
}
userList()
</script>

<style scoped>

</style>