<template>
  <DataShowView
      key-data="sid"
      search-data="name"
      :table-col-list="tableColList"
      :add-form="addForm"
      :edit-form="editForm"
      :urls="urls"
      delete
      download
  />
</template>

<script setup>
import DataShowView from "@/components/DataShowView.vue";

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "name", label: "员工名", sortable: false, width: 120},
  {property: "phone", label: "手机号码", sortable: false, width: 120},
  {property: "department", label: "所属部门", sortable: false, width: 180, isFK: true, FKData: {
      url: "/dept/list",
      property: "did",
      label: "name"
    }},
  {property: "created_at", label: "注册时间", isDateFormat: true, sortable: true, width: 240},
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
    department:[
      { required: 'true', message: '请选择员工所属部门', trigger: 'blur' }
    ]
  },
  item:[
    {label: '员工名', prop: 'name', dataName: 'name', isInput: true},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
    {label: '所属部门', prop: 'department', dataName: 'department', isFK: true,
      FKData: {
        url: "/dept/list",
        property: "did",
        label: "name"
      }},
  ],
}

/**
 * 添加员工时所用到的对象
 * */
const addForm = {
  data :{
    name:'',
    phone:'',
    department:'',
  },
  dataType:{
    name:'String',
    phone:'String',
    department:'String',
  },
  dataNum: 3,
  rules: {
    name:[
      { required: 'true', message: '请输入员工名', trigger: 'blur' }
    ],
    department:[
      { required: 'true', message: '请选择员工所属部门', trigger: 'blur' }
    ]
  },
  item:[
    {label: '员工名', prop: 'name', dataName: 'name', isInput: true},
    {label: '电话', prop: 'phone', dataName: 'phone', isInput: true, type: 'number'},
    {label: '所属部门', prop: 'department', dataName: 'department', isFK: true,
      FKData: {
        url: "/dept/list",
        property: "did",
        label: "name"
      }},
  ],
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