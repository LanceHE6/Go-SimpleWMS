<template>
<el-container
  v-loading="state.isLoading"
>
  <el-header
      class="main-header"
  >
    <el-form
      class="header"
      :model="form.data"
      :rules="form.rules"
    >
      <el-form-item
        v-for="item in form.item"
        :label="item.label"
        :prop="item.prop"
        :class="item.isLong ? 'form-long-item' : 'form-item'"
      >
        <el-input
            v-if="item.isInput === true"
            v-model.trim="form.data[item.dataName]"
            autocomplete="off"
        />

        <el-select
            v-if="item.isMapping === true"
            v-model.trim="form.data[item.dataName]"
            placeholder="请选择"
            filterable
        >
          <el-option
              v-for="i in item.selectOptions"
              :label="i.label"
              :value="i.value"
          />
        </el-select>

        <el-select
            v-if="item.isFK === true && hasFKData(item.FKData.property)"
            v-model.trim="form.data[item.dataName]"
            placeholder="请选择"
            filterable
        >
          <el-option
              v-for="i in Array.from(FKMap.get(item.FKData.property))"
              :disabled="'status' in i && i['status'] === 0"
              :label="i[item.FKData.label]"
              :value="i[item.FKData.property]"
          />
        </el-select>
      </el-form-item>
    </el-form>
  </el-header>

  <el-divider class="my-divider"/>

  <el-main
    class="main-body"
  >
    <div
      class="my-table-header"
    >
      <el-button
        type="primary"
        icon="Plus"
        @click="addItem"
        plain
      >
        添加
      </el-button>
      <el-button
          type="danger"
          icon="Delete"
          @click="removeItem"
          plain
      >
        删除
      </el-button>
    </div>
    <el-table
      :data="tableData"
      :border="true"
      :stripe="true"
      height="40vh"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />

      <el-table-column
          v-for="item in mainTableColList"
          :property="item.property"
          :label="item.label"
          :width="item.width"
          :sortable="item.sortable"
          :formatter="mapping(item.property)">
        <template #default="scope" v-if="item.isImage">
          <div style="display: flex; align-items: center">
            <el-image
                class="table-col-img"
                v-if="scope.row.images && scope.row.images.length > 0"
                :src="`${axios.defaults.baseURL}/${scope.row.images[0].path}`"
                fit="cover"
                :preview-src-list="scope.row.images.map(imgObj => axios.defaults.baseURL + '/' + imgObj.path)"
                preview-teleported
            >
              <template #error>
                <div class="error-image-slot">
                  <el-icon><Picture /></el-icon>
                </div>
              </template>
            </el-image>

            <div class="error-image-slot" v-else>
              <el-icon><Plus /></el-icon>
            </div>

          </div>
        </template>
      </el-table-column>
    </el-table>
  </el-main>

  <el-divider/>

  <el-footer>
    footer
  </el-footer>
</el-container>

<el-dialog
    v-model="addFormVisible"
    title="添加货品"
    width="1230"
    @open="handleAddDialogOpen"
    center
>

  <DataShowView
    ref="myDataShowView"
    :table-col-list="addTableColList"
    :key-data="'gid'"
    :urls="urls"
    height="40vh"
    large
  />
  <div>
    <el-text type="primary" style="margin-left: 20px">
      注意：重复的选择将会被忽略！
    </el-text>
  </div>
  <template #footer>
    <div class="dialog-footer">
      <el-button @click="addFormVisible = false">取消</el-button>
      <el-button type="primary" @click="submitAddForm">
        确定
      </el-button>
    </div>
  </template>

</el-dialog>
</template>

<script setup>

import {computed, markRaw, onMounted, reactive, ref} from "vue";
import axios from "axios";
import {ElMessage, ElMessageBox} from "element-plus";
import {Plus, Picture, Delete} from "@element-plus/icons-vue";
import DataShowView from "@/components/DataShowView.vue";
import {getObjectArrayUnion, getObjectArrayDifference} from "@/utils/arrayUtil"

const hasFKData = computed(() => (key) => FKMap.value.has(key));
const tableData = ref([])  //表格数据
const addTableData = ref([])  //添加窗口表格数据
const myDataShowView = ref(null)  //数据展示窗口子组件
let multipleSelection = []  //用户选择的主窗口元素

const state =  reactive({
  isLoading: true,  //数据是否正在加载
  dialogIsLoading: true,  //添加窗口是否正在加载
  itemCount: 0,  //已选元素数量
})

const prop = defineProps({
  submitForm: {
    type: Object,
    default: () => {},
    description: '核心表单'
  },
  mainTableColList:{
    type: Array,
    default: () => [],
    description: '主页面表头'
  },
  addTableColList:{
    type: Array,
    default: () => [],
    description: '添加页面表头'
  },
  urls:{
    type: Object,
    default: () => {},
    description: '请求url'
  }
});

const form = ref(prop.submitForm)

const addFormVisible = ref(false)

//提交表单外键
const FKMap = ref(new Map())
//主表格外键
const mainFKMap = ref(new Map())
onMounted(async () => {
  state.dialogIsLoading = true
  state.isLoading = true
  await getFKList()
  addTableData.value = await getData(prop.urls.getData)
  state.isLoading = false
  state.dialogIsLoading = false
})

//获取外键列表
async function getFKList() {
  const mainFormItem = prop.mainTableColList  //主表格外键数据
  const formItem = prop.submitForm && prop.submitForm.item  //提交表单外键数据
  const urlList = []  //外键url列表
  const FKDataMap = new Map()  //当前已获取外键, 避免重复请求

  //提交表单外键获取
  for (const i in formItem) {
    if (formItem[i].isFK === true) {
      const url = formItem[i].FKData.url  //外键url
      const property = formItem[i].FKData.property  //外键名

      //如果外键url不在列表中, 说明还没有请求过该外键
      if (!urlList.includes(url)) {
        urlList.push(url)
        const data = await getData(url)
        FKMap.value.set(property, data)
        FKDataMap.set(property, data)
      }
    }
  }

  //主表格外键获取
  for (const i in mainFormItem) {
    if (mainFormItem[i].isFK === true) {
      const addUrl = mainFormItem[i].FKData.url
      const addProp = mainFormItem[i].FKData.property

      //如果外键url不在列表中, 说明还没有请求过该外键
      if(!urlList.includes(addUrl)){
        urlList.push(addUrl)
        const data = await getData(addUrl)
        mainFKMap.value.set(addProp, data)
        FKDataMap.set(addProp, data)
      }
      //如果已经请求过该外键, 则在已获取外键中取值, 无需反复请求
      else{
        mainFKMap.value.set(addProp, FKDataMap.get(addProp))
      }

    }
  }
}

const token="bearer "+localStorage.getItem("token");

/**
 * getData()
 * 获取数据的请求
 * 结果：dataArray中包含一到多个data对象
 * */

const getData = async (url, params = {}) => {
  let resultList = []
  url = '/api' + url
  await axios.get(url, {
    headers: {
      'Authorization': token
    },
    params: params
  })
      .then(result => {
        console.log("submitViewGetData:", result)
        if (result && result.data && result.data.rows) {
          for (let i = 0; i < result.data.rows.length; i++){
            result.data.rows[i].created_at = new Date(result.data.rows[i].created_at).toLocaleString()
          }
          resultList = result.data.rows;
        }
      })
      .catch(error => {
        ElMessage.error("网络请求出错了！")
        console.error("submitViewGetData:", error)
      })
  return resultList
}

// 数据显示转换(映射)
function mapping(property){
  return (row) => {
    //遍历每一个属性
    for(const i of prop.mainTableColList){
      let item = i
      if(item.property === property){
        //多层对象, 一层层解开
        while(item.isParent){
          row = row[item.property]
          item = item['child']
        }
        //外键映射
        if(item.isFK){
          //从外键map中获取对应的外键表
          const fkList = mainFKMap.value.get(item.FKData.property)
          console.log("666", item.FKData.property, mainFKMap.value)
          for(const item2 of fkList){
            // 映射
            if(row[property] === item2[item.FKData.property]){
              return item2[item.FKData.label]
            }
          }
        }
        //普通映射
        else if(item.isMapping){
          //从映射表获取映射对象
          for(const j of item.mappingList){
            const item2 = j
            // 映射
            if(row[property] === item2.value){
              return item2.label
            }
          }
        }

        //不需要映射或者没有匹配的映射则返回原值
        return row[item.property]
      }
    }
  }
}

const submitAddForm = () => {
  tableData.value = getObjectArrayUnion(tableData.value, myDataShowView.value.getMultipleSelection())
  addFormVisible.value = false
}

//用户选择主窗口中的表项时回调此方法
const handleSelectionChange = (val) => {
  multipleSelection = val
}

//用户打开新增窗口时回调此方法
const handleAddDialogOpen = () =>{
  myDataShowView.value.clearSelection()
}

function addItem(){
  addFormVisible.value = true
}

function removeItem(){
  ElMessageBox.confirm(
      '你确定要删除选中的行吗？',
      '注意',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        icon: markRaw(Delete),
        type: 'warning',
      }
  )
  .then(() => {
    tableData.value = getObjectArrayDifference(tableData.value, multipleSelection)
  })
}

</script>

<style scoped>
.main-header{
  width: 100%;
  height: auto;
}
.header{
  display: flex;
  flex-wrap: wrap; /* 允许容器内子元素换行 */
  justify-content: flex-start; /* 水平分布，两端对齐 */
}
.main-body{
  padding: 5px
}
.my-table-header{
  margin-bottom: 10px;
}
.my-divider{
  margin-top: 10px;
  margin-bottom: 5px;
}
.form-item {
  text-align: justify;
  width: 225px;
  padding-right: 35px;
}
.form-long-item {
  text-align: justify;
  width: 745px;
  padding-right: 35px;
}
.table-col-img{
  height: 50px;
  width: 50px;
}

.error-image-slot{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 50px;
  height: 50px;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 15px;
}
</style>