<template>
<el-card
  class="card"
>
  <template #header>
    <div class="card-header">
      <el-text><b>类型分布</b></el-text>
      <el-select
          v-model="selectOption"
          placeholder="Select"
          size="small"
          style="width: 150px; margin-left: 20px"
      >
        <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
      </el-select>
    </div>
  </template>
  <v-chart
      v-if="optionAmountData.length > 0"
      :option="option_column"
      class="my-chart"
      autoresize
  />
  <el-empty
      v-else
      description="暂无数据"
      style="width: 100%"
  />

</el-card>
</template>

<script setup>

import {ref, watch} from "vue";

const prop = defineProps({
  optionAmountData:{
    type: Array,
    default: () => [],
    description: '图表中的数据(数量占比)'
  },
  optionPriceData:{
    type: Array,
    default: () => [],
    description: '图表中的数据(金额占比)'
  }
})

const selectOption = ref(0)
const amountData = ref(prop.optionAmountData)
const priceData = ref(prop.optionPriceData)

//监听数据变化并实时更新
watch(() => prop.optionAmountData, (newValue) => {
  amountData.value = newValue
  if(selectOption.value === 0){
    option_column.value.series[0].data = amountData.value
  }
});
watch(() => prop.optionPriceData, (newValue) => {
  priceData.value = newValue
  if(selectOption.value === 1) {
    option_column.value.series[0].data = priceData.value
  }
});

const option_column = ref({
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)'
  },
  legend: {
    orient: 'vertical',
    left: 'left'
  },
  series: [
    {
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 5,
        borderColor: '#fff',
        borderWidth: 2
      },
      label: {
        formatter: '{d}%',
      },
      data: amountData.value //根据情况切换data中的内容
    }
  ]
})

// 监听 selectOption 的变化
watch(selectOption, (newValue) => {
  if (newValue === 0) {
    option_column.value.series[0].data = amountData.value;
  } else if (newValue === 1) {
    option_column.value.series[0].data = priceData.value;
  }
});


const options = [
  {
    value: 0,
    label: '数量占比',
  },
  {
    value: 1,
    label: '金额占比',
  },
]

</script>

<style scoped>
.card{
  width: 695px;
  height: 340px;
  margin-bottom: 10px;
  margin-right: 15px;
  --el-card-padding: 10px
}
.card-header{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.my-chart{
  height: 300px;
}
</style>