<template>
  <DataShowView
    :table-col-list="tableColList"
    :key-data="'wid'"
    :search-data="'name'"
    :add-form="addForm"
    :edit-form="editForm"
    :delete-data-body="deleteDataBody"
    :urls="urls"
  />
</template>

<script setup>
import {reactive} from "vue";
import DataShowView from "@/components/DataShowView.vue";

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "name", label: "仓库名", sortable: false},
  {property: "comment", label: "备注", sortable: false},
  {property: "manager", label: "管理人", sortable: false},
  {property: "status", label: "状态", sortable: false, isMapping: true, mappingList:[
      {label: "禁用", value: 0},
      {label: "启用", value: 1},
    ]},
  {property: "created_at", label: "创建时间", sortable: true},
  {property: "wid", label: "仓库ID", sortable: true, width: 120},
]

/**
 * 编辑仓库时所用到的表单对象
 * */
const editForm = reactive({
  data : {
    wid:'',
    name:'',
    comment:'',
    manager:'',
    status:'',
  },
  dataType:{
    wid:'String',
    name:'String',
    comment:'String',
    manager:'String',
    status:'Int',
  },
  dataNum: 5,
  rules: {
    status:[
      { required: 'true', message: '请选择仓库状态', trigger: 'blur' }
    ],
    name:[
      { required: 'true', message: '请输入仓库名', trigger: 'blur' },
    ],
    manager:[
      { required: 'true', message: '请选择管理人', trigger: 'blur' },
    ]
  },
  item:[
    {label: '仓库名', prop: 'name', dataName: 'name', isInput: true,},
    {label: '是否启用', prop: 'status', dataName: 'status', isSelect: true,
      selectOptions: [
        {label: '禁用', value: 0},
        {label: '启用', value: 1},
      ]},
    {label: '管理人', prop: 'manager', dataName: 'manager', isFK: true,
      FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true},
  ],
})

/**
 * 添加仓库时所用到的对象
 * */
const addForm = reactive({
  data :{
    name:'',
    comment:'',
    manager:'',
    status:'',
  },
  dataType:{
    name:'String',
    comment:'String',
    manager:'String',
    status:'Int',
  },
  dataNum: 4,
  rules: {
    name: [
      { required: 'true', message: '仓库名不能为空', trigger: 'blur' },
    ],
    status:[
      { required: 'true', message: '请选择仓库状态', trigger: 'blur' }
    ],
    manager:[
      { required: 'true', message: '请选择管理人', trigger: 'blur' }
    ]
  }
  ,
  item:[
    {label: '仓库名', prop: 'name', dataName: 'name', isInput: true},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true},
    {label: '仓库状态', prop: 'status', dataName: 'status', isSelect: true,
      selectOptions: [
        {label: '禁用', value: 0},
        {label: '启用', value: 1},
      ]},
    {label: '管理人', prop: 'manager', dataName: 'manager', isFK: true,
      FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
  ],
})

/**
 * 删除请求体
 * */
const deleteDataBody = {
  wid: ""
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/warehouse/list",
  deleteData: "/warehouse/delete",
  addData: "/warehouse/add",
  updateData: "/warehouse/update",
  uploadData: "/warehouse/upload",
}
</script>

<style scoped>

</style>