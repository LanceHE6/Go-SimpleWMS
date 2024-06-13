<template>
  <DataShowView
    :table-col-list="tableColList"
    :key-data="'uid'"
    :search-data="'nickname'"
    :add-form="addForm"
    :edit-form="editForm"
    :delete-data-body="deleteDataBody"
    :urls="urls"
    upload
    download
  />
</template>

<script setup>
import DataShowView from "@/components/DataShowView.vue";

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "account", label: "账号", sortable: false, width: 180},
  {property: "nickname", label: "昵称", sortable: false, width: 120},
  {property: "permission", label: "权限", sortable: true, width: 120, isMapping: true, mappingList:[
      {label: '普通用户', value: 1},
      {label: '管理员', value: 2},
      {label: '超级管理员', value: 3},
    ]},
  {property: "phone", label: "手机号码", sortable: false, width: 120},
  {property: "created_at", label: "注册时间", sortable: true, width: 240},
  {property: "uid", label: "用户ID", sortable: true, width: 120},
]

/**
 * 编辑用户时所用到的表单对象
 * */
const editForm = {
  data :{
    uid:'',
    password:'',
    permission:0,
    nickname:'',
    phone:''
  },
  dataType:{
    uid:'String',
    password:'String',
    permission:'Int',
    nickname:'String',
    phone:'String'
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
      { min: 1, max: 8, message: '昵称长度需要在1-8个字符之间', trigger: 'blur' },
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
}

/**
 * 注册用户时所用到的对象
 * */
const addForm = {
  data :{
    account:'',
    password:'',
    permission:'',
    nickname:'',
    phone:''
  },
  dataType:{
    account:'String',
    password:'String',
    permission:'Int',
    nickname:'String',
    phone:'String'
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
      { required: 'true', message: '请输入用户昵称', trigger: 'blur' },
      { min: 1, max: 8, message: '昵称长度需要在1-8个字符之间', trigger: 'blur' },
    ]
  },
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
}

/**
 * 删除请求体
 * */
const deleteDataBody = {
  uid: ""
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/user/list",
  deleteData: "/user/delete",
  addData: "/user/register",
  updateData: "/user/update",
  uploadData: "/user/upload",
}

</script>

<style scoped>

</style>