<template>
  <DataShowView
      :table-col-list="tableColList"
      :key-data="'iid'"
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
  {property: "number", label: "单号", sortable: false},
  {property: "inventory_type", label: "类型", sortable: false, isFK: true, FKData:{
      url: "/invt/list",
      property: "itid",
      label: "name"
    }},
  {property: "date", label: "单据日期", sortable: false},
  {property: "warehouse", label: "仓库名称", sortable: false, isFK: true, FKData:{
      url: "/warehouse/list",
      property: "wid",
      label: "name"
    }},
  {property: "operator", label: "操作员", sortable: false, isFK: true, FKData:{
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