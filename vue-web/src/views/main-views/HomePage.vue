<template>
  <el-container
      v-loading="state.isLoading"
      class="main-body"
  >
    <el-collapse
        v-model="activeNames"
    >
      <el-collapse-item
          title="仓库概况"
          name="1">
        <div
            v-loading="state.warehouseCardDataList.length === 0"
            class="card-body"
        >
          <el-empty
              v-if="state.warehouseCardDataList.length === 0"
              description="暂无数据"
              style="width: 100%"
          />
          <warehouse-card
              v-for="item in state.warehouseCardDataList"
              :wid="item.wid"
              :selected="item.selected"
              :name="item.name"
              :typeQuantity ="item.typeQuantity"
              :entry-quantity="item.entryQuantity"
              :out-quantity="item.outQuantity"
              :entry-price-quantity="item.entryPriceQuantity"
              :out-price-quantity="item.outPriceQuantity"
              :y-entry-quantity="item.yEntryQuantity"
              :y-out-quantity="item.yOutQuantity"
              :y-entry-price-quantity="item.yEntryPriceQuantity"
              :y-out-price-quantity="item.yOutPriceQuantity"
              @select="onWarehouseCardSelect"
          />
        </div>
      </el-collapse-item>
      <el-collapse-item
          title="货品相关"
          name="2">
        <div
            v-loading="false"
            class="card-body"
        >

        </div>
      </el-collapse-item>
    </el-collapse>
  </el-container>
</template>

<script setup>
import {computed, onMounted, reactive, ref} from 'vue';
import WarehouseCard from "@/components/cards/WarehouseCard.vue";
import axios from "axios";
import {ElMessage} from "element-plus";

// 使用 ref 创建响应式数据
const data = ref(getRandomData());
const activeNames = ref(['1','2'])  // 初始化时的活动界面

const state =  reactive({
  selectWid: '0',  //选中的仓库id(如果是所有仓库则为0)
  isLoading: false,  //数据是否正在加载
  warehouseCardDataList:[],  //仓库信息卡片数据列表
  hasExtraWarehouse: false,  //是否还有仓库卡片未显示
})

function getInvParams(wid = '', type = -1, date = 0){
  const w = wid === '' ? {} : {warehouse : wid}
  const t = type === -1 ? {} : {type : type}
  return {
    ...w,
    ...t,
    page: -1,
    created_at: getDate(date)
  }
}

// 初始化函数
onMounted(async () => {
  data.value = getRandomData();

  const warehouseList = (await getData("/warehouse/list")).rows
  const warehouseCardDataList = []  //仓库信息卡片数据列表缓存

  // 并发执行多个请求
  const typeQuantityPromise = getData("/goods/search");
  const entryPromise = getData("/inv/search", getInvParams(undefined,1,undefined));
  const outPromise = getData("/inv/search", getInvParams(undefined,2,undefined));
  const yEntryPromise = getData("/inv/search", getInvParams(undefined,1,-1));
  const yOutPromise = getData("/inv/search", getInvParams(undefined,2,-1));

  const [typeQuantity, entry, out, yEntry, yOut] = await Promise.all([
    typeQuantityPromise,
    entryPromise,
    outPromise,
    yEntryPromise,
    yOutPromise
  ]);


  //今日入库总价
  let entryPrice = 0
  for(const item of entry.rows){
    entryPrice += item['goods']['quantity'] * item['goods']['unit_price']
  }

  //今日出库总价
  let outPrice = 0
  for(const item of out.rows){
    outPrice += item['goods']['quantity'] * item['goods']['unit_price']
  }

  //昨日入库总价
  let yEntryPrice = 0
  for(const item of yEntry.rows){
    yEntryPrice += item['goods']['quantity'] * item['goods']['unit_price']
  }

  //昨日出库总价
  let yOutPrice = 0
  for(const item of yOut.rows){
    yOutPrice += item['goods']['quantity'] * item['goods']['unit_price']
  }

  warehouseCardDataList.push({
    wid: '0',
    selected: true,
    name: '所有仓库',
    typeQuantity: typeQuantity.total,
    entryQuantity: entry.total,
    outQuantity: out.total,
    entryPriceQuantity: entryPrice,
    outPriceQuantity: outPrice,
    yEntryQuantity: yEntry.total,
    yOutQuantity: yOut.total,
    yEntryPriceQuantity: yEntryPrice,
    yOutPriceQuantity: yOutPrice,
  })

  let i = 1  //限制显示个数
  for(const item of warehouseList){
    if(i > 3){
      break
    }
    else{
      i++
    }
    const wid = item['wid']
    const itemTypeQuantityPromise = getData("/goods/search",{
      warehouse: wid
    })
    const itemEntryPromise = getData("/inv/search", getInvParams(wid,1,undefined))
    const itemOutPromise = getData("/inv/search", getInvParams(wid,2,undefined))
    const yItemEntryPromise = getData("/inv/search", getInvParams(wid,1,-1))
    const yItemOutPromise = getData("/inv/search", getInvParams(wid,2,-1))

    const [itemTypeQuantity, itemEntry, itemOut, yItemEntry, yItemOut] = await Promise.all([
      itemTypeQuantityPromise,
      itemEntryPromise,
      itemOutPromise,
      yItemEntryPromise,
      yItemOutPromise
    ]);

    console.log("entry", itemEntry)

    //今日入库总价
    let entryPrice = 0
    for(const j in itemEntry.rows){
      entryPrice += itemEntry.rows[j]['goods']['quantity'] * itemEntry.rows[j]['goods']['unit_price']
    }

    //今日出库总价
    let outPrice = 0
    for(const j in itemOut.rows){
      outPrice += itemOut.rows[j]['goods']['quantity'] * itemOut.rows[j]['goods']['unit_price']
    }

    //昨日入库总价
    let yEntryPrice = 0
    for(const j in yItemEntry.rows){
      yEntryPrice += yItemEntry.rows[j]['goods']['quantity'] * yItemEntry.rows[j]['goods']['unit_price']
    }

    //昨日出库总价
    let yOutPrice = 0
    for(const j in yItemOut.rows){
      yOutPrice += yItemOut.rows[j]['goods']['quantity'] * yItemOut.rows[j]['goods']['unit_price']
    }

    warehouseCardDataList.push({
      wid: wid,
      selected: false,
      name: item.name,
      typeQuantity: itemTypeQuantity.total,
      entryQuantity: itemEntry.total,
      outQuantity: itemOut.total,
      entryPriceQuantity: entryPrice,
      outPriceQuantity: outPrice,
      yEntryQuantity: yItemEntry.total,
      yOutQuantity: yItemOut.total,
      yEntryPriceQuantity: yEntryPrice,
      yOutPriceQuantity: yOutPrice,
    })
  }
  state.warehouseCardDataList = warehouseCardDataList
  console.log("total", state.warehouseCardDataList)
});

const onWarehouseCardSelect = (wid) =>{
  state.selectWid = wid
  for(const item of state.warehouseCardDataList){
    item.selected = item.wid === wid
  }
}

// 定义方法
function getRandomData() {
  return [
    {
      time: '2022-9-11',
      value: Math.random() * 100
    },
    {
      time: '2022-9-11',
      value: Math.random() * 100
    },
    {
      time: '2022-9-11',
      value: Math.random() * 100
    },
    {
      time: '2022-9-11',
      value: Math.random() * 100
    }
  ]
}

// 使用 computed 创建计算属性
const option = computed(() => {
  return {
    xAxis: {
      type: 'category',
      data: data.value.map(item => item.time)
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: data.value.map(item => item.value),
        type: 'line'
      }
    ]
  }
});

const token="bearer "+localStorage.getItem("token");

/**
 * getData()
 * 获取数据的请求
 * 结果：result
 * */

const getData = async (url, params = {}) => {
  let resultObj = {}
  await axios.get('/api' + url, {
    headers: {
      'Authorization': token
    },
    params: params
  })
      .then(result => {
        console.log("homepage-getData:", result)
        if (result && result.data && result.data.data && result.data.data.rows) {
          for (let i = 0; i < result.data.data.rows.length; i++){
            result.data.data.rows[i].created_at = new Date(result.data.data.rows[i].created_at).toLocaleString()
          }
          resultObj = result.data.data;
        }
      })
      .catch(error => {
        ElMessage.error("网络请求出错了！")
        console.error("homepage-getData:", error)
      })
  return resultObj
}

const getDate = (extraDay = 0) => {
  // 创建一个Date对象表示当前时间
  const now = new Date();

  // 加上日期偏移量
  now.setDate(now.getDate() + extraDay);

  // 使用getFullYear, getMonth, getDate方法获取各个部分
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0'); // 注意月份是从0开始的，所以需要+1
  const day = String(now.getDate()).padStart(2, '0');


  // 将各个部分组合成所需的字符串格式并返回
  return year + '-' + month + '-' + day
}
</script>

<style scoped>
.chart {
  padding: 10px;
  height: 100%;
  width: 100%;
}
.main-body {
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 子元素在父容器中垂直分布 */
  width: 100%;
  height: 95%;
  padding: 10px;
}
.card-body{
  display: flex;
  flex-wrap: wrap; /* 允许容器内子元素换行 */
  justify-content: flex-start; /* 水平分布，两端对齐 */
}

</style>