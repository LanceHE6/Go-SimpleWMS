<template>
  <submit-view
    :key-data="'gid'"
    :submit-form="submitForm"
    :main-table-col-list="mainTableColList"
    :add-table-col-list="addTableColList"
    :urls="urls"
    @remove-tab="removeTab"
  >
  </submit-view>
</template>

<script setup>
import SubmitView from "@/components/SubmitView.vue";

//对外事件列表
const emit = defineEmits(["removeTab"]);

const removeTab = () => {
  const CURRENT_PATH = window.location.hash  // 获取当前路径，例如 "#/page/subpage"
  const pageList = CURRENT_PATH.split('#')
  const currentPage = pageList[pageList.length - 1]
  emit("removeTab", currentPage)
}

/**
 * 主页面中的列表表头
 * */
const mainTableColList = [
  {property: "image", label: "图片", sortable: false, isImage: true, width: 60},
  {property: "goods_code", label: "货品编码", sortable: false, width: 150},
  {property: "name", label: "货品名称", sortable: false, width: 150},
  {property: "model", label: "规格型号", sortable: false, width: 120},
  {property: "goods_type", label: "货品类型", sortable: false, width: 120, isFK: true, FKData:{
      url: "/gt/list",
      property: "gtid",
      label: "name"
    }},
  {property: "unit", label: "计量单位", sortable: false, width: 120, isFK: true, FKData:{
      url: "/unit/list",
      property: "unid",
      label: "name"
    }},
  {property: "unit_price", label: "单价", sortable: false, width: 120},
  {property: "amount", label: "数量", sortable: false, width: 140, isInput: true, type: 'number', required: true},
  {property: "comment", label: "备注", sortable: false, isInput: true},
]

/**
 * 添加窗口中的列表表头
 * */
const addTableColList = [
  {property: "image", label: "图片", sortable: false, isImage: true, width: 80},
  {property: "name", label: "货品名称", sortable: false, width: 150},
  {property: "model", label: "规格型号", sortable: false, width: 120},
  {property: "manufacturer", label: "生产厂商", sortable: false, width: 120},
  {property: "goods_code", label: "货品编码", sortable: false, width: 150},
  {property: "goods_type", label: "货品类型", sortable: false, width: 120, isFK: true, FKData:{
      url: "/gt/list",
      property: "gtid",
      label: "name"
    }},
  {property: "unit", label: "计量单位", sortable: false, width: 120, isFK: true, FKData:{
      url: "/unit/list",
      property: "unid",
      label: "name"
    }},
  {property: "unit_price", label: "单价", sortable: false, width: 120},
  {property: "amount", label: "数量", sortable: false},
]

/**
 * 提交表单时所用到的对象
 * */
const submitForm = {
  data : {
    date: '',
    number: '',
    department: '',
    inventory_type: '',
    warehouse: '',
    manufacturer: '',
    operator: '',
    comment: '',
  },
  listDataTemplate:[
    {prop: 'goods_list', dataName: 'goods_list', children:[
        {prop: 'goods', dataName: 'gid'},
        {prop: 'amount', dataName: 'amount'},
        {prop: 'comment', dataName: 'comment'},
    ]}
  ],
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
      { required: 'true', message: '请选择仓库', trigger: 'blur' }
    ],
    operator:[
      { required: 'true', message: '请选择操作员', trigger: 'blur' }
    ],
  },
  item:[
    {label: '仓库', prop: 'warehouse', dataName: 'warehouse', isFK: true, FKData:{
        url: "/warehouse/list",
        property: "wid",
        label: "name"
      }},
    {label: '单据日期', prop: 'date', dataName: 'date', isDate: true},
    {label: '单号', prop: 'number', dataName: 'number', isInput: true},
    {label: '类型', prop: 'inventory_type', dataName: 'inventory_type', isFK: true, FKData:{
        url: "/invt/list",
        property: "itid",
        label: "name"
      }},
    {label: '所属部门', prop: 'department', dataName: 'department', isFK: true, FKData:{
        url: "/dept/list",
        property: "did",
        label: "name"
      }},
    {label: '制造商', prop: 'manufacturer', dataName: 'manufacturer', isInput: true},
    {label: '操作员', prop: 'operator', dataName: 'operator', isFK: true, FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true, isLong: true},
  ],
}

const urls = {
  getData: "/goods/search",
  addData: "/inv/add"
}

</script>

<style scoped>

</style>