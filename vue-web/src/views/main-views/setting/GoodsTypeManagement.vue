<template>
  <DataShowView
      key-data="gtid"
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
  {property: "name", label: "货品类型名", sortable: false, width: 180},
  {property: "type_code", label: "货品类型编号", sortable: true, width: 180},
  {property: "created_at", label: "创建时间", isDateFormat: true, sortable: true, width: 240},
]

/**
 * 编辑数据时所用到的表单对象
 * */
const editForm = {
  data :{
    gtid:'',
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
}

/**
 * 添加数据时所用到的对象
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
      { required: 'true', message: '请输入货品类型名', trigger: 'blur' }
    ]
  },
  item:[
    {label: '货品类型名', prop: 'name', dataName: 'name', isInput: true},
    {label: '货品类型编码', prop: 'type_code', dataName: 'type_code', isInput: true},
  ],
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/gt/list",
  deleteData: "/gt/delete",
  addData: "/gt/add",
  updateData: "/gt/update",
  uploadData: "/gt/upload",
}

</script>

<style scoped>

</style>