<template>
  <DataShowView
      key-data="tid"
      :table-col-list="tableColList"
      :urls="urls"
      @addTab="addTab"
      has-submit-page
      large
      download
      print
  />
</template>

<script setup>
import DataShowView from "@/components/DataShowView.vue";

//对外事件列表
const emit = defineEmits(["addTab"]);

/**
 * 新增tab的方法
 * */
const addTab = () => {
  const data = {
    name: 'AllocateSubmit',
    label: '新增调拨记录',
    path: '/home/entryAndOut/allocateSubmit',
  }
  emit("addTab", data.name, data.label, data.path)
}

/**
 * 表头属性列表
 * */
const tableColList = [
  {property: "goods_list", label: "货品信息", sortable: false, width:100, isExpand: true, children:[
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
      {property: "amount", label: "数量", sortable: false, width: 120},
      {property: "goods", label: "单价", sortable: false, width: 120, isParent: true,
        child:{
          property: "unit_price", isParent: false
        }},
      {property: "comment", label: "备注", sortable: false, width: 240},
    ]},
  {property: "number", label: "单号", width: 180, sortable: false},
  {property: "date", label: "单据日期", width: 180, isDateFormat: true, sortable: false},
  {property: "source_warehouse", label: "源仓库", width: 150, sortable: false, isFK: true, FKData:{
      url: "/warehouse/list",
      property: "wid",
      label: "name"
    }},
  {property: "destination_warehouse", label: "目标仓库", width: 150, sortable: false, isFK: true, FKData:{
      url: "/warehouse/list",
      property: "wid",
      label: "name"
    }},
  {property: "operator", label: "操作员", width: 100, sortable: false, isFK: true, FKData:{
      url: "/staff/list",
      property: "sid",
      label: "name"
    }},
  {property: "passed", label: "通过状态", sortable: false, width: 100 , isMapping: true, mappingList:[
      {label: '通过', value: true},
      {label: '未通过', value: false},
    ]},
  {property: "comment", label: "备注", sortable: false},
]

/**
 * 网络请求url
 * */
const urls = {
  getData: "/trans/search",
  addData: "/trans/add",
}
</script>

<style scoped>

</style>