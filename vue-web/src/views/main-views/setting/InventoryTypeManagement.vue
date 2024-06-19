<template>
  <DataShowView
      key-data="itid"
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
  {property: "name", label: "出入库类型名", sortable: false, width: 180},
  {property: "type_code", label: "出入库类型编号", width: 180, sortable: true},
  {property: "type", label: "类型", sortable: false, width: 80, isMapping: true,
    mappingList:[
      {label: "入库", value: 1},
      {label: "出库", value: 2},
    ]},
  {property: "created_at", label: "创建时间", isDateFormat: true, sortable: true, width: 240},
  {property: "itid", label: "出入库类型ID", sortable: true, width: 140},
]

/**
 * 编辑数据时所用到的表单对象
 * */
const editForm = {
  data :{
    itid:'',
    name:'',
    type_code:'',
    type:'',
  },
  dataType:{
    gtid:'String',
    name:'String',
    type_code:'String',
    type:'Int',
  },
  dataNum: 4,
  rules: {
    name:[
      { required: 'true', message: '请输入货品类型名', trigger: 'blur' }
    ]
  },
  item:[
    {label: '货品类型名', prop: 'name', dataName: 'name', isInput: true},
    {label: '货品类型编码', prop: 'type_code', dataName: 'type_code', isInput: true},
    {label: '出入库类型', prop: 'type', dataName: 'type', isSelect: true,
      selectOptions: [
        {label: '入库', value: 1},
        {label: '出库', value: 2},
      ]},
  ],
}

/**
 * 添加数据时所用到的对象
 * */
const addForm = {
  data :{
    name:'',
    type_code:'',
    type:'',
  },
  dataType:{
    name:'String',
    type_code:'String',
    type:'Int',
  },
  dataNum: 3,
  rules: {
    name:[
      { required: 'true', message: '请输入出入库类型名', trigger: 'blur' }
    ],
    type:[
      { required: 'true', message: '请选择出入库类型', trigger: 'blur' }
    ]
  },
  item:[
    {label: '出入库类型名', prop: 'name', dataName: 'name', isInput: true},
    {label: '出入库类型编码', prop: 'type_code', dataName: 'type_code', isInput: true},
    {label: '出入库类型', prop: 'type', dataName: 'type', isSelect: true,
      selectOptions: [
        {label: '入库', value: 1},
        {label: '出库', value: 2},
      ]},
  ],
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