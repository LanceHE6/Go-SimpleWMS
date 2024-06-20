<template>
  <DataShowView
      key-data="iid"
      :table-col-list="tableColList"
      :urls="urls"
      :extra-params="extraParams"
      @addTab="addTab"
      has-refresh-event-bus
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
  type: 2
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
  {property: "number", label: "单号", width: 200, sortable: false},
  {property: "inventory_type", label: "类型", width: 120, sortable: false, isFK: true, FKData:{
      url: "/invt/list",
      property: "itid",
      label: "name"
    }},
  {property: "date", label: "单据日期", width: 180, isDateFormat: true, sortable: false},
  {property: "warehouse", label: "仓库名称", width: 150, sortable: false, isFK: true, FKData:{
      url: "/warehouse/list",
      property: "wid",
      label: "name"
    }},
  {property: "operator", label: "操作员", width: 100, sortable: false, isFK: true, FKData:{
      url: "/staff/list",
      property: "sid",
      label: "name"
    }},
  {property: "manufacturer", label: "制造商", sortable: false, width: 150},
  {property: "comment", label: "备注", sortable: false},
]

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