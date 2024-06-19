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
import {ElInput, ElMessage, ElMessageBox, ElSwitch} from "element-plus";
import {axiosPut} from "@/utils/axiosUtil.js";
import {h, ref} from "vue";

//对外事件列表
const emit = defineEmits(["addTab"]);

//自定义点击函数
const audit = (tid, isAudit) =>{
  if(!isAudit){
    const passed = ref(false)
    const comment = ref('')
    ElMessageBox({
      title: '调拨审核',
      message: () =>
          h('p', null, [
            h('div', null, [
              `调拨单id：${tid}`,
            ]),
            h('div', null, [
              `是否通过审核(${passed.value ? '是' : '否'})`,
              h(ElSwitch, {
                modelValue: passed.value,
                style: 'margin-left: 10px',
                'onUpdate:modelValue': (val) => {passed.value = val},
              }),
            ]),
            h('div', null, [
              h(ElInput, {
                modelValue: comment.value,
                placeholder: "审核备注",
                'onUpdate:modelValue': (val) => {comment.value = val},
              }),
            ]),
         ])
    })
    .then(async () => {
      const result = await axiosPut({
        url: '/trans/audit',
        data: {
          tid: tid,
          passed: passed.value,
          comment: comment.value
        },
        name: 'audit'
      })
      if(result){
        ElMessage.success("审核成功")
      }
    })
  }
  else{
    ElMessageBox.confirm(
        `你确定要撤销${tid}的审核吗？`,
        '注意',
        {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning',
        }
    )
    .then(async () => {
      const result = await axiosPut({
        url: '/trans/audit/revoke',
        data: {
          tid: tid
        },
        name: 'audit-revoke'
      })
      if(result){
        ElMessage.success("审核已撤销")
      }
    })
  }
}

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
  {property: "goods_list", label: "货品信息", sortable: false, width:100, isFixed: true, isExpand: true, children:[
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
  {property: "audited", label: "审核状态", sortable: false, operable: true, operationEvent: audit, width: 100, isMapping: true, mappingList:[
      {label: '已审核', value: true},
      {label: '待审核', value: false},
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
  {property: "comment", label: "操作员备注", width: 240, sortable: false},
  {property: "auditor", label: "审核人", width: 100, sortable: false, isFK: true, FKData:{
      url: "/staff/list",
      property: "sid",
      label: "name"
    }},
  {property: "audited_time", label: "审核日期", width: 180, isDateFormat: true, sortable: false},
  {property: "audit_comment", label: "审核人备注", width: 240, sortable: false},
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