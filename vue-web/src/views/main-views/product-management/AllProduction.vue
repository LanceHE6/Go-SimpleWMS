<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'gid'"
      :add-form="addForm"
      :edit-form="editForm"
      :delete-data-body="deleteDataBody"
      :urls="urls"
      large
  />
</template>

<script setup>
import {reactive} from "vue";
import DataShowView from "@/components/DataShowView.vue";

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "name", label: "货品名称", sortable: false},
  {property: "image", label: "图片", sortable: false, isImage: true},
  {property: "goods_code", label: "货品编码", sortable: false},
  {property: "model", label: "规格型号", sortable: false},
  {property: "goods_type", label: "货品类型", sortable: false, isFK: true, FKData:{
      url: "/gt/list",
      property: "gtid",
      label: "name"
    }},
  {property: "warehouse", label: "所属仓库", sortable: false, isFK: true, FKData:{
      url: "/warehouse/list",
      property: "wid",
      label: "name"
    }},
  {property: "unit", label: "计量单位", sortable: false, isFK: true, FKData:{
      url: "/unit/list",
      property: "unid",
      label: "name"
    }},
  {property: "quantity", label: "单位数", sortable: true},
  {property: "manufacturer", label: "生产厂商", sortable: false, width: 150},
]

/**
 * 编辑仓库时所用到的表单对象
 * */
const editForm = reactive({
  data : {
    gid:'',
    name:'',
    model:'',
    goods_code:'',
    goods_type:'',
    warehouse:'',
    manufacturer:'',
    unit:'',
    quantity:''
  },
  dataType:{
    gid:'String',
    name:'String',
    model:'String',
    goods_code:'String',
    goods_type:'String',
    warehouse:'String',
    manufacturer:'String',
    unit:'String',
    quantity:'Int'
  },
  dataNum: 9,
  rules: {

  },
  item:[
    {label: '货品名称', prop: 'name', dataName: 'name', isInput: true,},
    {label: '规格类型', prop: 'model', dataName: 'model', isInput: true,},
    {label: '货品编码', prop: 'goods_code', dataName: 'goods_code', isInput: true,},
    {label: '货品类型', prop: 'goods_type', dataName: 'goods_type', isFK: true,
      FKData:{
        url: "/gt/list",
        property: "gtid",
        label: "name"
      }},
    {label: '所属仓库', prop: 'warehouse', dataName: 'warehouse', isFK: true,
      FKData:{
        url: "/warehouse/list",
        property: "wid",
        label: "name"
      }},
    {label: '计量单位', prop: 'unit', dataName: 'unit', isFK: true,
      FKData:{
        url: "/unit/list",
        property: "unid",
        label: "name"
      }},
    {label: '单位数', prop: 'quantity', dataName: 'quantity', isInput: true, type: 'number'},
    {label: '制造商', prop: 'manufacturer', dataName: 'manufacturer', isInput: true,},
  ],
})

/**
 * 添加仓库时所用到的对象
 * */
const addForm = reactive({
  data : {
    name:'',
    model:'',
    goods_code:'',
    goods_type:'',
    warehouse:'',
    manufacturer:'',
    unit:'',
    quantity:''
  },
  dataType:{
    name:'String',
    model:'String',
    goods_code:'String',
    goods_type:'String',
    warehouse:'String',
    manufacturer:'String',
    unit:'String',
    quantity:'Int'
  },
  dataNum: 8,
  rules: {
    name:[
      { required: 'true', message: '请输入货品名称', trigger: 'blur' }
    ],
    goods_type:[
      { required: 'true', message: '请选择货品类型', trigger: 'blur' }
    ],
    warehouse:[
      { required: 'true', message: '请选择所属仓库', trigger: 'blur' }
    ],
    unit:[
      { required: 'true', message: '请选择计量单位', trigger: 'blur' }
    ],
  },
  item:[
    {label: '货品名称', prop: 'name', dataName: 'name', isInput: true,},
    {label: '规格类型', prop: 'model', dataName: 'model', isInput: true,},
    {label: '货品编码', prop: 'goods_code', dataName: 'goods_code', isInput: true,},
    {label: '货品类型', prop: 'goods_type', dataName: 'goods_type', isFK: true,
      FKData:{
        url: "/gt/list",
        property: "gtid",
        label: "name"
      }},
    {label: '所属仓库', prop: 'warehouse', dataName: 'warehouse', isFK: true,
      FKData:{
        url: "/warehouse/list",
        property: "wid",
        label: "name"
      }},
    {label: '计量单位', prop: 'unit', dataName: 'unit', isFK: true,
      FKData:{
        url: "/unit/list",
        property: "unid",
        label: "name"
      }},
    {label: '单位数', prop: 'quantity', dataName: 'quantity', isInput: true, type: 'number'},
    {label: '制造商', prop: 'manufacturer', dataName: 'manufacturer', isInput: true,},
  ],
})

/**
 * 删除请求体
 * */
const deleteDataBody = {
  gid: ""
}

/**
 * 网络请求url
 * */
const urls = {
  getData: "/goods/search",
  deleteData: "/goods/delete",
  addData: "/goods/add",
  updateData: "/goods/update",
  uploadData: "/goods/upload",
}
</script>

<style scoped>

</style>