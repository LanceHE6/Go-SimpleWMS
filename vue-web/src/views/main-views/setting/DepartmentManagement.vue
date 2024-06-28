<template>
  <DataShowView
      key-data="did"
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
  {property: "name", label: "部门名", sortable: false, width: 120},
  {property: "created_at", label: "添加时间", isDateFormat: true, sortable: true, width: 240},
]

/**
 * 编辑部门时所用到的表单对象
 * */
const editForm = {
  data :{
    did:'',
    name:'',
  },
  dataType:{
    did:'String',
    name:'String',
  },
  dataNum: 2,
  rules: {
    name: [
      { required: 'true', message: '部门名不能为空', trigger: 'blur' },
    ],
  },
  item:[
    {label: '部门名称', prop: 'name', dataName: 'name', isInput: true},
  ],
}

/**
 * 添加部门时所用到的对象
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
    name: [
      { required: 'true', message: '部门名不能为空', trigger: 'blur' },
    ],
  },
  item:[
    {label: '部门名称', prop: 'name', dataName: 'name', isInput: true},
  ],
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/dept/list",
  deleteData: "/dept/delete",
  addData: "/dept/add",
  updateData: "/dept/update",
  uploadData: "/dept/upload",
}

</script>

<style scoped>

</style>