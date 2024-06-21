<template>
<el-container
  v-loading="state.isLoading"
>
  <el-header
      class="main-header"
  >
    <el-form
      ref="myForm"
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
            :disabled="item.isKey && tableData.length > 0"
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
            :disabled="item.isKey && tableData.length > 0"
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

        <el-date-picker
            v-if="item.isDate"
            v-model="form.data[item.dataName]"
            type="datetime"
            placeholder="选择日期和时间"
            :shortcuts="shortcuts"
            value-format="YYYY-MM-DD HH:mm:ss"
        />
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
    <table-body
      :default-data="tableData"
      :table-col-list="mainTableColList"
      :show-f-k-map="mainFKMap"
      :operations="operations"
      height="50vh"
      @selection-change="handleSelectionChange"
    />
  </el-main>

  <el-divider class="my-divider"/>

  <el-footer
    class="main-footer"
  >
    <el-button
      type="primary"
      @click="saveCheck"
    >
    保存
    </el-button>

    <el-button
      type="warning"
      plain
      @click="cancel"
    >
    取消
    </el-button>
  </el-footer>
</el-container>

<el-dialog
    v-model="addFormVisible"
    title="添加货品"
    width="1300"
    @open="handleAddDialogOpen"
    center
>

  <DataShowView
    ref="myDataShowView"
    :table-col-list="addTableColList"
    :key-data="keyData"
    :urls="urls"
    :extra-params="state.extraParams"
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
import {Delete, Picture} from "@element-plus/icons-vue";
import DataShowView from "@/components/DataShowView.vue";
import {getObjectArrayDifference, getObjectArrayUnionByKey} from "@/utils/arrayUtil"
import {axiosGet, axiosPost} from "@/utils/axiosUtil.js";
import TableBody from "@/components/TableBody.vue";
import EventBus from "@/utils/eventBus.js";

const hasFKData = computed(() => (key) => FKMap.value.has(key));
const tableData = ref([])  //表格数据
const myDataShowView = ref(null)  //数据展示窗口子组件
const myForm = ref(null)  //首部表单
let multipleSelection = []  //用户选择的主窗口元素

//对外事件列表
const emit = defineEmits(["addTab", "removeTab"]);

const operations = {
  add: false,
  del: false,
  edit: false,
  upload: false,
  download: false,
  print: false,
  uploadImg: false,
}

const state =  reactive({
  isLoading: true,  //数据是否正在加载
  itemCount: 0,  //已选元素数量
  extraParams: {} //getData的额外参数
})

const prop = defineProps({
  keyData:{
    type: [String, Array],
    default: () => undefined,
    description: '数据主键, 如有多层则用列表封装, 格式如: ["key", "p1", "p2", ...]'
  },
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
  },
  extraParams:{
    type: Array,
    default: () => [],
    description: 'getData请求的额外参数'
  }
});

//日期选择器的便携选项
const shortcuts = [
  {
    text: '今天',
    value: new Date(),
  },
  {
    text: '昨天',
    value: () => {
      const date = new Date()
      date.setDate(date.getDate() - 1)
      return date
    },
  },
  {
    text: '前天',
    value: () => {
      const date = new Date()
      date.setDate(date.getDate() - 2)
      return date
    },
  },
  {
    text: '大前天',
    value: () => {
      const date = new Date()
      date.setDate(date.getDate() - 3)
      return date
    },
  }
]

//提交表单及其配置
const form = ref(prop.submitForm)

//添加窗口是否可见
const addFormVisible = ref(false)

//提交表单外键
const FKMap = ref(new Map())
//主表格外键
const mainFKMap = ref(new Map())

onMounted(async () => {
  state.isLoading = true

  await getAllFKList()

  state.isLoading = false
})

async function getDataAndSet(url, propName, stateMap, FKDataMap) {
  if (!FKDataMap.has(propName)) {
    const data = await getData(url, undefined, `getFK-${propName}`);
    stateMap.set(propName, data);
    FKDataMap.set(propName, data);
  } else {
    stateMap.set(propName, FKDataMap.get(propName));
  }
}

async function getFKList(form, targetFKMap, FKDataMap) {
  for (const item of (form || [])) {
    let currentItem = item
    //多层对象，先解开
    while (currentItem.isParent){
      currentItem = currentItem['child']
    }
    if (currentItem.isFK) {
      await getDataAndSet(currentItem.FKData.url, currentItem.FKData.property, targetFKMap, FKDataMap);
    }
  }
}

//获取外键列表
async function getAllFKList() {
  const mainFormItem = prop.mainTableColList  //主表格外键数据
  const formItem = prop.submitForm && prop.submitForm.item  //提交表单外键数据
  const FKDataMap = new Map()  //当前已获取外键, 避免重复请求

  // 主窗口外键获取
  await getFKList(mainFormItem, mainFKMap.value, FKDataMap)

  // 提交窗口外键获取
  await getFKList(formItem, FKMap.value, FKDataMap)


  //子表格显示外键映射
  for (const item of mainFormItem) {
    if ('children' in item) {
      await getFKList(item.children, mainFKMap.value, FKDataMap)
    }
  }
}

/**
 * getData()
 * 获取数据的请求
 * 结果：dataArray中包含一到多个data对象
 * */

const getData = async (url, params = {}, name = 'getData') => {
  const result = await axiosGet({url: url, params: params, name: name})
  if(result){
    if (result && result.data && result.data.rows) {
      return result.data.rows;
    }
  }
  else{
    return undefined
  }
}

/**
 * addData()
 * 新增数据 param: newData对象
 * */
const addData=async (data) => {
  state.isLoading = true
  console.log("add", data)
  const result = axiosPost({url: prop.urls['addData'], data: data, name: 'addData'})
  if(result){
    ElMessage.success("数据添加成功！")
  }
  state.isLoading = false
}

const submitAddForm = () => {
  tableData.value = getObjectArrayUnionByKey(tableData.value, myDataShowView.value.getMultipleSelection(), prop.keyData)
  //将用户需要输入的内容放入数据列表中
  for(const item of prop.mainTableColList){
    if(item.isInput){
      for(const data of tableData.value){
        if(!(item.property in data)){  //如果数据不存在该属性, 则为其添加此属性
          data[item.property] = ''
        }
      }
    }
  }
  addFormVisible.value = false
}

//用户选择主窗口中的表项时回调此方法
const handleSelectionChange = (val) => {
  multipleSelection = val
  console.log("change", val)
}

//用户打开新增窗口时回调此方法
const handleAddDialogOpen = async () => {
  await myDataShowView.value.update(1)
  myDataShowView.value.clearSelection()
}

const addItem = async () => {
  const params = {}
  if (prop.extraParams.length > 0) {
    for (const item of prop.extraParams) {
      if (!prop.submitForm.data[item.dataName]) {
        ElMessage.error(`请先选择${item.label}`)
        return
      } else {
        params[item.prop] = prop.submitForm.data[item.dataName]
      }
    }
    state.extraParams = params
  }
  addFormVisible.value = true
}

//用户点击删除按钮
const removeItem = () => {
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

//用户点击最下方的保存按钮, 执行检查步骤
const saveCheck = async () => {
  //检查上方表单中的必填项以及规则
  await myForm.value.validate(async (valid) => {
    if (valid) {
      //检查用户是否添加了表项
      if (tableData.value.length === 0) {
        await ElMessageBox.confirm(
            `请添加表项`,
            '注意',
            {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning',
            }
        )
        return
      }
      //检查表格中的所有必填项
      for (const item of prop.mainTableColList) {
        if (item.required) {
          for (const data of tableData.value) {
            if (!data[item.property]) {
              await ElMessageBox.confirm(
                  `请输入${item.label}`,
                  '注意',
                  {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning',
                  }
              )
              return
            }
          }
        }
      }
      //检查通过
      await save()
    }
  })
}

//保存
const save = async () => {
  for (const i of prop.submitForm.listDataTemplate) {

    const objList = tableData.value.map(item => {
      // 初始化一个空对象来存储新属性
      const newObj = {};
      // 遍历children数组来动态设置新对象的属性
      i.children.forEach(child => {
        // 使用child.dataName从原始对象中提取值，并使用child.prop作为新对象的属性名
        if(Array.isArray(child.dataName)){
          let val = item
          for(const i of child.dataName){
            val = val[i]
          }
          newObj[child.prop] = val
        }
        else{
          newObj[child.prop] = item[child.dataName];
        }
      });

      // 返回构建好的新对象
      return newObj;
    })
    prop.submitForm.data[i.prop] = JSON.stringify(objList)
  }
  console.log("data", prop.submitForm.data)
  await addData(prop.submitForm.data)
  setTimeout(() => {
    EventBus.emit("refresh", 1)
  }, 500);
  emit("removeTab")
}

//用户点击最下方的取消按钮
const cancel = () => {
  ElMessageBox.confirm(
      '数据未保存，你确定要退出吗？',
      '注意',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
  .then(() => {
    emit("removeTab")
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
.main-footer{
  display: flex;
  justify-content: center; /* 水平分布，两端对齐 */
}
.main-footer .el-button{
  margin-top: 10px;
}
.form-item {
  text-align: justify;
  width: 250px;
  padding-right: 30px;
}
.form-long-item {
  text-align: justify;
  width: 420px;
  padding-right: 30px;
}
</style>