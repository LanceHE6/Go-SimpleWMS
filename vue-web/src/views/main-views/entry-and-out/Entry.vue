<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'iid'"
      :add-form="addForm"
      :urls="urls"
      large
      download
      print
  />
</template>

<script setup>
import {reactive} from "vue";
import DataShowView from "@/components/DataShowView.vue";

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "goods", label: "货品名称", sortable: false, isParent: true, child:{
      property: "name",
      isParent: false
    }},
  {property: "amount", label: "数量", sortable: true},
  {property: "manufacturer", label: "生产厂商", sortable: false, width: 150},
  {property: "warehouse", label: "所属仓库", sortable: false, isFK: true, FKData:{
      url: "/warehouse/list",
      property: "wid",
      label: "name"
    }},
  {property: "inventory_type", label: "出入库类型", sortable: false, isFK: true, FKData:{
      url: "/invt/list",
      property: "itid",
      label: "name"
    }},
  {property: "operator", label: "操作员", sortable: false, isFK: true, FKData:{
      url: "/staff/list",
      property: "sid",
      label: "name"
    }},
  {property: "comment", label: "备注", sortable: false},
]

/**
 * 添加数据时所用到的对象
 * */
const addForm = reactive({
  data : {
    gid: '',
    name: '',
    amount: '',
    unit_price: '',
    inventory_type: '',
    warehouse: '',
    manufacturer: '',
    operator: '',
    comment: '',
  },
  dataType:{
    gid: 'String',
    name:'String',
    amount: 'Int',
    unit_price: 'Float',
    inventory_type: 'String',
    warehouse: 'String',
    manufacturer: 'String',
    operator: 'String',
    comment: 'String',
  },
  dataNum: 9,
  rules: {
    amount:[
      { required: 'true', message: '请输入货品数量', trigger: 'blur' }
    ],
    inventory_type:[
      { required: 'true', message: '请选择出入库类型', trigger: 'blur' }
    ],
    warehouse:[
      { required: 'true', message: '请选择所属仓库', trigger: 'blur' }
    ],
    operator:[
      { required: 'true', message: '请选择操作员', trigger: 'blur' }
    ],
  },
  item:[
    {label: '货品id', prop: 'gid', dataName: 'gid', isInput: true,},
    {label: '货品名称', prop: 'name', dataName: 'name', isInput: true,},
    {label: '数量', prop: 'amount', dataName: 'amount', isInput: true, type: 'number'},
    {label: '单价', prop: 'unit_price', dataName: 'unit_price', isInput: true},
    {label: '出入库类型', prop: 'inventory_type', dataName: 'inventory_type', isFK: true,
      FKData:{
        url: "/invt/list",
        property: "itid",
        label: "name"
      }},
    {label: '所属仓库', prop: 'warehouse', dataName: 'warehouse', isFK: true,
      FKData:{
        url: "/warehouse/list",
        property: "wid",
        label: "name"
      }},
    {label: '操作员', prop: 'operator', dataName: 'operator', isFK: true,
      FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
    {label: '制造商', prop: 'manufacturer', dataName: 'manufacturer', isInput: true,},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true,},
  ],
})

/**
 * 网络请求url
 * */
const urls = {
  getData: "/inv/search",
  addData: "/inv/add",
}
</script>

<style scoped>

</style>