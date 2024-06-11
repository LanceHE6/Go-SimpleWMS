<template>
  <submit-view>

  </submit-view>
</template>

<script setup>
import SubmitView from "@/components/SubmitView.vue";
import {reactive} from "vue";

/**
 * 添加数据时所用到的对象
 * */
const form = reactive({
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
</script>

<style scoped>

</style>