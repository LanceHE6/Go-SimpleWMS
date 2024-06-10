<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'iid'"
      :add-form="addForm"
      :urls="urls"
      :extra-params="extraParams"
      @addTab="addTab"
      has-submit-page
      large
      download
      print
  />
</template>

<script setup>
import {reactive} from "vue";
import DataShowView from "@/components/DataShowView.vue";

//对外事件列表
const emit = defineEmits(["addTab"]);

/**
 * 新增tab的方法
 * */
const addTab = () => {
  const data = {
    name: 'InventorySubmit',
    label: '新增出入库记录',
    path: '/home/entryAndOut/inventorySubmit',
  }
  emit("addTab", data.name, data.label, data.path)
}

/**
 * getData的额外请求体
 * */
const extraParams = {
  type: 1
}

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
    date: '',
    number: '',
    department: '',
    inventory_type: '',
    goods_list: [],
    warehouse: '',
    manufacturer: '',
    operator: '',
    comment: '',
  },
  dataType:{
    date:'String',
    number: 'String',
    department: 'String',
    inventory_type: 'String',
    goods_list: 'Array',
    warehouse: 'String',
    manufacturer: 'String',
    operator: 'String',
    comment: 'String',
  },
  dataNum: 9,
  rules: {
    inventory_type:[
      { required: 'true', message: '请选择出入库类型', trigger: 'blur' }
    ],
    warehouse:[
      { required: 'true', message: '请选择所属仓库', trigger: 'blur' }
    ],
    operator:[
      { required: 'true', message: '请选择操作员', trigger: 'blur' }
    ],
  }
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