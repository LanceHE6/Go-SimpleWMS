<template>
  <el-container
      v-loading="state.isLoading"
      class="main-body"
  >
    <el-collapse
        v-model="activeNames"
    >
      <el-collapse-item
          name="1">
        <template #title>
          <el-icon class="header-icon">
            <price-tag />
          </el-icon>
          <el-text size="large">仓库概况</el-text>
        </template>
        <div>
          <el-button
              type="primary"
              size="small"
              text
              @click="warehouseCardLimitChange"
              :disabled="state.warehouseCardDataList.length === 0"
          >
            {{ state.warehouseCardBtnStr }}
          </el-button>
        </div>
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
          name="2">
        <template #title>
          <el-icon class="header-icon">
            <box />
          </el-icon>
          <el-text size="large">货品相关</el-text>
        </template>
        <div
            class="card-body"
        >
          <goods-type-card
            v-loading="state.goodsTypeIsLoading"
            :option-amount-data="state.optionAmountData"
            :option-price-data="state.optionPriceData"
          />
          <inventory-card

          />
        </div>
      </el-collapse-item>
      <el-collapse-item
          name="3">
        <template #title>
          <el-icon class="header-icon">
            <folder-add />
          </el-icon>
          <el-text size="large">入库相关</el-text>
        </template>
        <div
            v-loading="false"
            class="card-body"
        >

        </div>
      </el-collapse-item>
      <el-collapse-item
          name="4">
        <template #title>
          <el-icon class="header-icon">
            <folder-remove />
          </el-icon>
          <el-text size="large">出库相关</el-text>
        </template>
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
import {axiosGet} from "@/utils/axiosUtil.js";
import {Box, FolderAdd, FolderRemove, PriceTag} from "@element-plus/icons-vue";
import GoodsTypeCard from "@/components/cards/GoodsTypeCard.vue";
import InventoryCard from "@/components/cards/InventoryCard.vue";

// 使用 ref 创建响应式数据
const activeNames = ref(['1', '2', '3', '4'])  // 初始化时的活动界面

const state =  reactive({
  selectWid: '0',  //选中的仓库id(如果是所有仓库则为0)
  isLoading: false,  //主页面是否正在加载
  goodsTypeIsLoading: false,  //类型分布卡片是否正在加载
  warehouseCardDataList:[],  //仓库信息卡片数据列表
  hasExtraWarehouse: false,  //是否还有仓库卡片未显示
  warehouseCardBtnStr: '显示所有仓库',  //仓库概况上方的按钮文字
  warehouseCardLimit: 3,  //显示多少个仓库卡片
  optionAmountData: [],  //图表中的数据(数量占比)
  optionPriceData: [], //图表中的数据(金额占比)
})

//获取出入库表请求参数
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

//获取仓库卡片信息
async function getWarehouseCardData(){
  const warehouseList = (await getData("/warehouse/list", undefined, 'getWarehouseList')).rows
  const warehouseCardDataList = []  //仓库信息卡片数据列表缓存

  // 并发执行多个请求
  const typeQuantityPromise = getData("/stock/get", undefined, 'getAllQuantity');
  const entryPromise = getData("/inv/search", getInvParams(undefined,1,undefined), 'getAllEntry');
  const outPromise = getData("/inv/search", getInvParams(undefined,2,undefined), 'getAllOut');
  const yEntryPromise = getData("/inv/search", getInvParams(undefined,1,-1), 'getYAllEntry');
  const yOutPromise = getData("/inv/search", getInvParams(undefined,2,-1), 'getYAllOut');

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
    for(const good of item.goods_list){
      entryPrice += good['amount'] * good['goods']['unit_price']
    }
  }

  //今日出库总价
  let outPrice = 0
  for(const item of out.rows){
    for(const good of item.goods_list){
      outPrice += good['amount'] * good['goods']['unit_price']
    }
  }

  //昨日入库总价
  let yEntryPrice = 0
  for(const item of yEntry.rows){
    for(const good of item.goods_list){
      yEntryPrice += good['amount'] * good['goods']['unit_price']
    }
  }

  //昨日出库总价
  let yOutPrice = 0
  for(const item of yOut.rows){
    for(const good of item.goods_list){
      yOutPrice += good['amount'] * good['goods']['unit_price']
    }
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
    if(i > state.warehouseCardLimit){
      break
    }
    else{
      i++
    }
    const wid = item['wid']
    const itemTypeQuantityPromise = getData("/stock/get",{
      warehouse: wid
    }, 'getItemQuatity')
    const itemEntryPromise = getData("/inv/search", getInvParams(wid,1,undefined), 'getItemEntry')
    const itemOutPromise = getData("/inv/search", getInvParams(wid,2,undefined), 'getItemOut')
    const yItemEntryPromise = getData("/inv/search", getInvParams(wid,1,-1), 'getYItemEntry')
    const yItemOutPromise = getData("/inv/search", getInvParams(wid,2,-1), 'getYItemOut')

    const [itemTypeQuantity, itemEntry, itemOut, yItemEntry, yItemOut] = await Promise.all([
      itemTypeQuantityPromise,
      itemEntryPromise,
      itemOutPromise,
      yItemEntryPromise,
      yItemOutPromise
    ]);

    //今日入库总价
    let entryPrice = 0
    for(const item of itemEntry.rows){
      for(const good of item.goods_list){
        entryPrice += good['amount'] * good['goods']['unit_price']
      }
    }

    //今日出库总价
    let outPrice = 0
    for(const item of itemOut.rows){
      for(const good of item.goods_list){
        outPrice += good['amount'] * good['goods']['unit_price']
      }
    }

    //昨日入库总价
    let yEntryPrice = 0
    for(const item of yItemEntry.rows){
      for(const good of item.goods_list){
        yEntryPrice += good['amount'] * good['goods']['unit_price']
      }
    }

    //昨日出库总价
    let yOutPrice = 0
    for(const item of yItemOut.rows){
      for(const good of item.goods_list){
        yOutPrice += good['amount'] * good['goods']['unit_price']
      }
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
}

//获取货品占比卡片信息
async function getGoodsTypeCardData(){
  const params = state.selectWid === '0' ? {} : {
    warehouse: state.selectWid
  }
  const stockDataRow = (await getData('/stock/get', params, 'getGoodsTypeCardData')).rows
  const amountObjList = []
  const priceObjList = []
  for(const item of stockDataRow || []){
    const amountObj = {}
    const priceObj = {}
    amountObj.name = item['goods'].name
    amountObj.value = item['quantity']
    amountObjList.push(amountObj)

    priceObj.name = item['goods'].name
    priceObj.value = item['quantity'] * item['goods'].unit_price
    priceObjList.push(priceObj)
  }
  state.optionAmountData = amountObjList
  state.optionPriceData = priceObjList
}

// 初始化函数
onMounted(async () => {
  await getWarehouseCardData()
  await getGoodsTypeCardData()
});

//点击仓库卡片
const onWarehouseCardSelect = async (wid) => {
  state.selectWid = wid
  for (const item of state.warehouseCardDataList) {
    item.selected = item.wid === wid
  }
  state.goodsTypeIsLoading = true
  await getGoodsTypeCardData()
  state.goodsTypeIsLoading = false
}

//点击显示更多仓库按钮
const warehouseCardLimitChange = async () => {
  state.warehouseCardDataList.length = 0
  if(state.warehouseCardLimit === 3){
    state.warehouseCardLimit = 3000
    state.warehouseCardBtnStr = '显示前3个仓库'
    await getWarehouseCardData()
  }
  else{
    state.warehouseCardLimit = 3
    state.warehouseCardBtnStr = '显示所有仓库'
    await getWarehouseCardData()
  }
  state.selectWid = '0'
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

/**
 * getData()
 * 获取数据的请求
 * */
const getData = async (url, params = {}, name = 'getData') => {
  let resultObj = {}
  const result = await axiosGet({url: url, params: params, name: name})
  if(result){
    resultObj = result.data
  }
  return resultObj
}

//获取时间字符串
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
.main-body {
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 子元素在父容器中垂直分布 */
  width: 100%;
  height: 100%;
  padding: 10px;
}
.card-body{
  display: flex;
  flex-wrap: wrap; /* 允许容器内子元素换行 */
  justify-content: flex-start; /* 水平分布，两端对齐 */
}
.card-body :deep(.el-loading-mask){
  z-index: 20;
}
.header-icon{
  margin-right: 5px;
}

</style>