<template>
  <el-card
      class="card"
      :class="{ 'card-selected': isSelected }"
  >
    <template #header>
      <div class="card-header" @click="onSelected">
        <span>
          <el-avatar :icon="isSelected ? 'Select' : 'SemiSelect'"
                     style="margin-right: 10px; width: 15px; height: 15px; --el-avatar-icon-size: 10px;"/>
          <el-text><b>{{name}}</b></el-text>
        </span>
        <el-text type="info">共{{typeQuantity}}种货物</el-text>
      </div>
    </template>
    <div class="card-main">
      <div class="card-title">
        <el-text>
          <el-icon style="vertical-align: middle">
            <Sunny />
          </el-icon>
          今日概况
        </el-text>
      </div>
      <div class="card-body">
        <div class="card-content">
          <el-tag size="small" :type="isSelected ? 'primary' : 'success'" style="width: 40px">入库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="eqv" />
            <el-text class="item-text">笔</el-text>
          </div>
          <el-text class="item-text">金额 {{entryPriceQuantity}}</el-text>
        </div>
        <el-divider direction="vertical" style="height: 60px"/>
        <div class="card-content">
          <el-tag size="small" :type="isSelected ? 'primary' : 'warning'" style="width: 40px">出库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="oqv" />
            <el-text class="item-text">笔</el-text>
          </div>
          <el-text class="item-text">金额 {{outPriceQuantity}}</el-text>
        </div>
      </div>
      <el-divider style="margin:5px"/>
      <div class="card-title">
        <el-text>
          <el-icon style="vertical-align: middle">
            <Moon />
          </el-icon>
          昨日概况
        </el-text>
      </div>
      <div class="card-body">
        <div class="card-content">
          <el-tag size="small" :type="isSelected ? 'primary' : 'success'" style="width: 40px">入库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="y_eqv" />
            <el-text class="item-text">笔</el-text>
          </div>
          <el-text class="item-text">金额 {{yEntryPriceQuantity}}</el-text>
        </div>
        <el-divider direction="vertical" style="height: 60px"/>
        <div class="card-content">
          <el-tag size="small" :type="isSelected ? 'primary' : 'warning'" style="width: 40px">出库</el-tag>
          <div class="number-body">
            <el-statistic class="number" :value="y_oqv" />
            <el-text class="item-text">笔</el-text>
          </div>
          <el-text class="item-text">金额 {{yOutPriceQuantity}}</el-text>
        </div>
      </div>
    </div>

  </el-card>
</template>

<script setup>
import {Moon, Sunny} from "@element-plus/icons-vue";
import { useTransition } from '@vueuse/core'
import {ref, watch} from "vue";

const emit = defineEmits(["select"]);

const prop = defineProps({
  wid: {
    type: String,
    default: () => '0',
    description: '唯一标识'
  },
  selected: {
    type: Boolean,
    default: () => false,
    description: '是否被选中'
  },
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


const isSelected = ref(prop.selected)
watch(() => prop.selected, (newValue) => {
  isSelected.value = newValue;
});

const onSelected = () =>{
  //必须选中一个, 所以不能点击自身来取消掉选中状态
  if(!isSelected.value) {
    isSelected.value = !isSelected.value
    emit("select", prop.wid);
  }
}

const eq = ref(0)
const eqv = useTransition(eq, {
  duration: 800,
})
eq.value = prop.entryQuantity

const oq = ref(0)
const oqv = useTransition(oq, {
  duration: 720,
})
oq.value = prop.outQuantity

const y_eq = ref(0)
const y_eqv = useTransition(y_eq, {
  duration: 760,
})
y_eq.value = prop.yEntryQuantity

const y_oq = ref(0)
const y_oqv = useTransition(y_oq, {
  duration: 700,
})
y_oq.value = prop.yOutQuantity

</script>

<style scoped>
.card{
  width: 340px;
  height: 310px;
  margin-bottom: 10px;
  margin-right: 15px;
  --el-card-padding: 10px
}
.card-header{
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer
}
.card-main{
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 子元素在父容器中垂直分布 */
}
.card-title{
  margin-bottom: 5px;
}
.card-body{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.card-content{
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
.item-text{
  width: 100%;
  font-size: 12px;
  color: black;
}

.card-selected{
  background-color: cornflowerblue;
}
.card-selected .el-text{
  color: white;
}
.card-selected .el-statistic{
  --el-statistic-content-color: white;
}

.card-selected .el-avatar{
  --el-avatar-bg-color: orange;
}
</style>