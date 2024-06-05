<template>
  <el-card
      class="card"
  >
    <template #header>
      <div class="card-header">
        <el-text>{{name}}</el-text>
        <el-text type="info">共{{typeQuantity}}种货物</el-text>
      </div>
    </template>
    <div class="card-main">
      <div class="main-title">
        <el-text>
          <el-icon style="vertical-align: middle">
            <Sunny />
          </el-icon>
          今日概况
        </el-text>
      </div>
      <div class="main-body">
        <div class="main-content">
          <el-tag size="small" type="success" style="width: 40px">入库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="eqv" />
            <text class="text">笔</text>
          </div>
          <text class="text">金额 {{entryPriceQuantity}}</text>
        </div>
        <el-divider direction="vertical" style="height: 60px"/>
        <div class="main-content">
          <el-tag size="small" type="warning" style="width: 40px">出库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="oqv" />
            <text class="text">笔</text>
          </div>
          <text class="text">金额 {{outPriceQuantity}}</text>
        </div>
      </div>
      <el-divider style="margin:5px"/>
      <div class="main-title">
        <el-text>
          <el-icon style="vertical-align: middle">
            <Moon />
          </el-icon>
          昨日概况
        </el-text>
      </div>
      <div class="main-body">
        <div class="main-content">
          <el-tag size="small" type="success" style="width: 40px">入库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="y_eqv" />
            <text class="text">笔</text>
          </div>
          <text class="text">金额 {{yEntryPriceQuantity}}</text>
        </div>
        <el-divider direction="vertical" style="height: 60px"/>
        <div class="main-content">
          <el-tag size="small" type="warning" style="width: 40px">出库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="y_oqv" />
            <text class="text">笔</text>
          </div>
          <text class="text">金额 {{yOutPriceQuantity}}</text>
        </div>
      </div>
    </div>

  </el-card>
</template>

<script setup>
import {Moon, Sunny} from "@element-plus/icons-vue";

const prop = defineProps({
  name: {
    type: String,
    default: () => 'Null',
    description: '仓库名'
  },
  typeQuantity: {
    type: Number,
    default: () => 0,
    description: '货品种类数量'
  },
  entryQuantity: {
    type: Number,
    default: () => 0,
    description: '今日入库数'
  },
  outQuantity: {
    type: Number,
    default: () => 0,
    description: '今日出库数'
  },
  entryPriceQuantity: {
    type: Number,
    default: () => 0,
    description: '今日入库总价'
  },
  outPriceQuantity: {
    type: Number,
    default: () => 0,
    description: '今日出库总价'
  },
  yEntryQuantity: {
    type: Number,
    default: () => 0,
    description: '昨日入库数'
  },
  yOutQuantity: {
    type: Number,
    default: () => 0,
    description: '昨日出库数'
  },
  yEntryPriceQuantity: {
    type: Number,
    default: () => 0,
    description: '昨日入库总价'
  },
  yOutPriceQuantity: {
    type: Number,
    default: () => 0,
    description: '昨日出库总价'
  },
});
import { useTransition } from '@vueuse/core'
import {ref} from "vue";

const eq = ref(0)
const eqv = useTransition(eq, {
  duration: 420,
})
eq.value = prop.entryQuantity

const oq = ref(0)
const oqv = useTransition(oq, {
  duration: 340,
})
oq.value = prop.outQuantity

const y_eq = ref(0)
const y_eqv = useTransition(y_eq, {
  duration: 500,
})
y_eq.value = prop.yEntryQuantity

const y_oq = ref(0)
const y_oqv = useTransition(y_oq, {
  duration: 450,
})
y_oq.value = prop.yOutQuantity

</script>

<style scoped>
.card{
  width: 340px;
  height: 300px;
  margin-bottom: 10px;
  margin-right: 15px;
  --el-card-padding: 10px
}
.card-header{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.card-main{
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 子元素在父容器中垂直分布 */
}
.main-title{
  margin-bottom: 5px;
}
.main-body{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.main-content{
  width: 50%;
  display: flex;
  flex-direction: column;
  padding: 5px
}
.number-body{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.number{
  margin-right: 10px;
}
.text{
  font-size: 12px;
}
</style>