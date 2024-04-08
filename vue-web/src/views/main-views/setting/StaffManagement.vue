<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'sid'"
      :search-data="'name'"
      :add-form="addForm"
      :edit-form="editForm"
      :delete-data-body="deleteDataBody"
      :urls="urls"
  />
</template>

<script setup>
import DataShowView from "@/components/DataShowView.vue";

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "name", label: "员工名", sortable: false},
  {property: "phone", label: "手机号码", sortable: false},
  {property: "department", label: "所属部门", sortable: false},
  {property: "created_at", label: "注册时间", sortable: true},
  {property: "sid", label: "员工ID", sortable: true, width: 120},
]

/**
 * 编辑员工时所用到的表单对象
 * */
const editForm = {
  data :{
    sid:'',
    name:'',
    phone:'',
    department:'',
  },
  dataType:{
    sid:'String',
    name:'String',
    phone:'String',
    department:'String',
  },
  dataNum: 4,
  rules: {
    name:[
      { required: 'true', message: '请输入员工名', trigger: 'blur' }
    ],
    dept_id:[
      { required: 'true', message: '请选择员工所属部门', trigger: 'blur' }
    ]
  },
  item:[
    {label: '员工名', prop: 'name', dataName: 'name', isInput: true},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
    {label: '所属部门', prop: 'dept_id', dataName: 'dept_id', isFK: true,
      FKData: {
        url: "/dept/list",
        property: "did",
        label: "name"
      }},
  ],
  key: 'sid'
}

/**
 * 注册员工时所用到的对象
 * */
const addForm = {
  data :{
    name:'',
    phone:'',
    dept_id:'',
  },
  dataType:{
    name:'String',
    phone:'String',
    dept_id:'String',
  },
  dataNum: 3,
  rules: {
    name:[
      { required: 'true', message: '请输入员工名', trigger: 'blur' }
    ],
    dept_id:[
      { required: 'true', message: '请选择员工所属部门', trigger: 'blur' }
    ]
  },
  item:[
    {label: '员工名', prop: 'name', dataName: 'name', isInput: true},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
    {label: '所属部门', prop: 'dept_id', dataName: 'dept_id', isFK: true,
      FKData: {
        url: "/dept/list",
        property: "did",
        label: "name"
      }},
  ],
}

/**
 * 删除请求体
 * */
const deleteDataBody = {
  sid: ""
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/staff/list",
  deleteData: "/staff/delete",
  addData: "/staff/add",
  updateData: "/staff/update",
  uploadData: "/staff/upload",
}

</script>

<style scoped>

</style>