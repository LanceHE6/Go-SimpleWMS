<template>
  <DataShowView
      key-data="unid"
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
  {property: "name", label: "计量单位名", sortable: false, width: 140},
  {property: "unid", label: "计量单位ID", sortable: true, width: 140},
]

/**
 * 编辑数据时所用到的表单对象
 * */
const editForm = {
  data :{
    unid:'',
    name:'',
  },
  dataType:{
    name:'String',
  },
  dataNum: 2,
  rules: {

  },
  item:[
    {label: '计量单位名', prop: 'name', dataName: 'name', isInput: true},
  ],
}

/**
 * 添加数据时所用到的对象
 * */
const addForm = {
  data :{
    name:'',
  },
  dataType:{
    name:'String',
  },
  dataNum: 1,
  rules: {
    name:[
      { required: 'true', message: '请输入计量单位名', trigger: 'blur' }
    ]
  },
  item:[
    {label: '计量单位名', prop: 'name', dataName: 'name', isInput: true},
  ],
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/unit/list",
  deleteData: "/unit/delete",
  addData: "/unit/add",
  updateData: "/unit/update",
  uploadData: "/unit/upload",
}

</script>

<style scoped>

</style>