<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'did'"
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
  {property: "name", label: "部门名", sortable: false},
  {property: "created_at", label: "添加时间", sortable: true},
  {property: "did", label: "部门ID", sortable: true, width: 120},
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
  key: 'did'
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
 * 删除请求体
 * */
const deleteDataBody = {
  did: ""
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