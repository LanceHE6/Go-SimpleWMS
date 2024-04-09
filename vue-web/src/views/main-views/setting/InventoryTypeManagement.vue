<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'gtid'"
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
  {property: "name", label: "出入库类型名", sortable: false},
  {property: "type_code", label: "出入库类型编号", sortable: true},
  {property: "created_at", label: "创建时间", sortable: true},
  {property: "itid", label: "出入库类型ID", sortable: true, width: 130},
]

/**
 * 编辑用户时所用到的表单对象
 * */
const editForm = {
  data :{
    itid:'',
    name:'',
    type_code:''
  },
  dataType:{
    gtid:'String',
    name:'String',
    type_code:'String'
  },
  dataNum: 3,
  rules: {
    name:[
      { required: 'true', message: '请输入货品类型名', trigger: 'blur' }
    ]
  },
  item:[
    {label: '货品类型名', prop: 'name', dataName: 'name', isInput: true},
    {label: '货品类型编码', prop: 'type_code', dataName: 'type_code', isInput: true},
  ],
  key: 'itid'
}

/**
 * 注册用户时所用到的对象
 * */
const addForm = {
  data :{
    name:'',
    type_code:''
  },
  dataType:{
    name:'String',
    type_code:'String'
  },
  dataNum: 2,
  rules: {
    name:[
      { required: 'true', message: '请输入出入库类型名', trigger: 'blur' }
    ]
  },
  item:[
    {label: '出入库类型名', prop: 'name', dataName: 'name', isInput: true},
    {label: '出入库类型编码', prop: 'type_code', dataName: 'type_code', isInput: true},
  ],
}

/**
 * 删除请求体
 * */
const deleteDataBody = {
  itid: ""
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/invt/list",
  deleteData: "/invt/delete",
  addData: "/invt/add",
  updateData: "/invt/update",
  uploadData: "/invt/upload",
}

</script>

<style scoped>

</style>