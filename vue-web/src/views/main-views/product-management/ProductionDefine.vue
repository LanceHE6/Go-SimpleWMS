<template>
  <DataShowView
      key-data="gid"
      search-data="name"
      :table-col-list="tableColList"
      :add-form="addForm"
      :edit-form="editForm"
      :urls="urls"
      has-refresh-event-bus
      large
      delete
      download
      print
      upload-img
      upload-file
  />
</template>

<script setup>
import {reactive} from "vue";
import DataShowView from "@/components/DataShowView.vue";
import { pNumValidatorNRequire } from "@/utils/validator.js"

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "name", label: "货品名称", width: 180, sortable: false},
  {property: "images", label: "图片", sortable: false, isImage: true, width: 120},
  {property: "goods_code", label: "货品编码", width: 120, sortable: false},
  {property: "model", label: "规格型号", width: 120, sortable: false},
  {property: "goods_type", label: "货品类型", width: 120, sortable: false, isFK: true, FKData:{
      url: "/gt/list",
      property: "gtid",
      label: "name"
    }},
  {property: "unit", label: "计量单位", width: 100, sortable: false, isFK: true, FKData:{
      url: "/unit/list",
      property: "unid",
      label: "name"
    }},
  {property: "unit_price", label: "单价", width: 80, sortable: true},
  {property: "manufacturer", label: "生产厂商", width: 150, sortable: false},
  {property: "files", label: "附件", sortable: false, isFile: true, width: 100},
]

/**
 * 编辑数据时所用到的表单对象
 * */
const editForm = reactive({
  data : {
    gid:'',
    name:'',
    model:'',
    goods_code:'',
    goods_type:'',
    manufacturer:'',
    unit:'',
    unit_price:''
  },
  dataType:{
    gid:'String',
    name:'String',
    model:'String',
    goods_code:'String',
    goods_type:'String',
    manufacturer:'String',
    unit:'String',
    unit_price:'Float'
  },
  dataNum: 7,
  rules: {
    unit_price:[
      { validator: pNumValidatorNRequire, trigger: 'blur' }
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
    {label: '计量单位', prop: 'unit', dataName: 'unit', isFK: true,
      FKData:{
        url: "/unit/list",
        property: "unid",
        label: "name"
      }},
    {label: '单价', prop: 'unit_price', dataName: 'unit_price', isInput: true},
    {label: '制造商', prop: 'manufacturer', dataName: 'manufacturer', isInput: true},
  ],
})

/**
 * 添加数据时所用到的对象
 * */
const addForm = reactive({
  data : {
    name:'',
    model:'',
    goods_code:'',
    goods_type:'',
    manufacturer:'',
    unit:'',
    unit_price:''
  },
  dataType:{
    name:'String',
    model:'String',
    goods_code:'String',
    goods_type:'String',
    manufacturer:'String',
    unit:'String',
    unit_price:'Float'
  },
  dataNum: 7,
  rules: {
    name:[
      { required: 'true', message: '请输入货品名称', trigger: 'blur' }
    ],
    goods_code:[
      { required: 'true', message: '请输入货品编码', trigger: 'blur' }
    ],
    goods_type:[
      { required: 'true', message: '请选择货品类型', trigger: 'blur' }
    ],
    unit:[
      { required: 'true', message: '请选择计量单位', trigger: 'blur' }
    ],
    unit_price:[
      { validator: pNumValidatorNRequire, trigger: 'blur' }
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
    {label: '计量单位', prop: 'unit', dataName: 'unit', isFK: true,
      FKData:{
        url: "/unit/list",
        property: "unid",
        label: "name"
      }},
    {label: '单价', prop: 'unit_price', dataName: 'unit_price', isInput: true},
    {label: '制造商', prop: 'manufacturer', dataName: 'manufacturer', isInput: true,},
  ],
})

/**
 * 网络请求url
 * */
const urls = {
  getData: "/goods/search",
  deleteData: "/goods/delete",
  addData: "/goods/add",
  updateData: "/goods/update",
  uploadImage: "/upload/goods_img",
  uploadFile: "/upload/goods_file"
}
</script>

<style scoped>

</style>