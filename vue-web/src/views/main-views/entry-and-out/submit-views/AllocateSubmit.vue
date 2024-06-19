<template>
  <submit-view
    :key-data="['gid', 'goods']"
    :submit-form="submitForm"
    :main-table-col-list="mainTableColList"
    :add-table-col-list="addTableColList"
    :urls="urls"
    :extra-params="extraParams"
    @remove-tab="removeTab"
  >
  </submit-view>
</template>

<script setup>
import SubmitView from "@/components/SubmitView.vue";
import {strValidatorNEqual} from "@/utils/validator.js";

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
  {property: "goods", label: "图片", sortable: false, width: 60, isImage: true, isParent: true,
    child:{
      property: "images", isParent: false
    }},
  {property: "goods", label: "货品编码", sortable: false, width: 120, isParent: true,
    child:{
      property: "goods_code", isParent: false
    }},
  {property: "goods", label: "货品名称", sortable: false, width: 120, isParent: true,
    child:{
      property: "name", isParent: false
    }},
  {property: "goods", label: "规格型号", sortable: false, width: 120, isParent: true,
    child:{
      property: "model", isParent: false
    }},
  {property: "goods", label: "计量单位", sortable: false, width: 120, isParent: true,
    child:{
      property: "unit", isParent: false, isFK: true, FKData:{
        url: "/unit/list",
        property: "unid",
        label: "name"
      },
    }},
  {property: "goods", label: "单价", sortable: false, width: 120, isParent: true,
    child:{
      property: "unit_price", isParent: false
    }},
  {property: "quantity", label: "库存数量", sortable: false, width: 140},
  {property: "amount", label: "调拨数量", sortable: false, width: 140, isInput: true, type: 'number', required: true},
  {property: "comment", label: "备注", sortable: false, isInput: true},
]

/**
 * 添加窗口中的列表表头
 * */
const addTableColList = [
  {property: "goods", label: "图片", sortable: false, width: 120, isImage: true, isParent: true,
    child:{
      property: "images", isParent: false
    }},
  {property: "goods", label: "货品编码", sortable: false, width: 120, isParent: true,
    child:{
      property: "goods_code", isParent: false
    }},
  {property: "goods", label: "货品名称", sortable: false, width: 120, isParent: true,
    child:{
      property: "name", isParent: false
    }},
  {property: "goods", label: "规格型号", sortable: false, width: 120, isParent: true,
    child:{
      property: "model", isParent: false
    }},
  {property: "goods", label: "计量单位", sortable: false, width: 120, isParent: true,
    child:{
      property: "unit", isParent: false, isFK: true, FKData:{
        url: "/unit/list",
        property: "unid",
        label: "name"
      },
    }},
  {property: "goods", label: "单价", sortable: false, width: 120, isParent: true,
    child:{
      property: "unit_price", isParent: false
    }},
  {property: "quantity", label: "数量", sortable: false, width: 80},
  {property: "comment", label: "备注", sortable: false, width: 240},
]

/**
 * 提交表单时所用到的对象
 * */
const submitForm = {
  data : {
    date: '',
    number: '',
    source_warehouse: '',
    destination_warehouse: '',
    operator: '',
    comment: '',
  },
  listDataTemplate:[
    {prop: 'goods_list', dataName: 'goods_list', children:[
        {prop: 'goods', dataName: ['goods', 'gid']},
        {prop: 'amount', dataName: 'amount'},
        {prop: 'comment', dataName: 'comment'},
      ]}
  ],
  dataType:{
    date:'String',
    number: 'String',
    goods_list: 'Array',
    source_warehouse: 'String',
    destination_warehouse: 'String',
    operator: 'String',
    comment: 'String',
  },
  dataNum: 7,
  rules: {
    source_warehouse:[
      { required: 'true', message: '请选择源仓库', trigger: 'blur' },
      { validator: (rule, value, callback) =>
        strValidatorNEqual(rule, value, callback, submitForm.data.destination_warehouse, '源仓库与目标仓库不能相同'),
        trigger: 'blur'
      }
    ],
    destination_warehouse:[
      { required: 'true', message: '请选择目标仓库', trigger: 'blur' },
      { validator: (rule, value, callback) =>
        strValidatorNEqual(rule, value, callback, submitForm.data.source_warehouse, '源仓库与目标仓库不能相同'),
        trigger: 'blur'
      }
    ],
    operator:[
      { required: 'true', message: '请选择操作员', trigger: 'blur' }
    ],
  },
  item:[
    {label: '源仓库', prop: 'source_warehouse', dataName: 'source_warehouse', isKey: true, isFK: true, FKData:{
        url: "/warehouse/list",
        property: "wid",
        label: "name"
      }},
    {label: '目标仓库', prop: 'destination_warehouse', dataName: 'destination_warehouse', isFK: true, FKData:{
        url: "/warehouse/list",
        property: "wid",
        label: "name"
      }},
    {label: '单据日期', prop: 'date', dataName: 'date', isDate: true},
    {label: '单号', prop: 'number', dataName: 'number', isInput: true},
    {label: '操作员', prop: 'operator', dataName: 'operator', isFK: true, FKData:{
        url: "/staff/list",
        property: "sid",
        label: "name"
      }},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true, isLong: true},
  ],
}

const urls = {
  getData: "/stock/get",
  addData: "/trans/add"
}

const extraParams = [
  {label: '源仓库', prop: 'warehouse', dataName: 'source_warehouse'}
]

</script>

<style scoped>

</style>