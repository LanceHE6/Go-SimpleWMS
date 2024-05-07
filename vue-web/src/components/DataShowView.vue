<template>
  <my-table
      ref="myTable"
      :table-col-list="tableColList"
      :default-data="state.currentDataArray"
      :search-data="searchData"
      :add-data-template="addForm"
      :edit-data-template="editForm"
      :loading="state.isLoading"
      :add-f-k-map="state.addFKMap"
      :edit-f-k-map="state.editFKMap"
      :show-f-k-map="state.showFKMap"
      :key-data="keyData"
      :page-count="state.pageCount"
      @add="add"
      @upload="upload"
      @del="del"
      @edit="edit"
      @update="update"
  >
  </my-table>
</template>

<script setup>
import axios from "axios";
import {ElMessage, ElNotification} from "element-plus";
import {onMounted, reactive} from "vue";
import MyTable from "@/components/MyTable.vue";

const PAGE_SIZE = 10 //每页展示多少个数据

const prop = defineProps({
  tableColList:{
    type: Array,
    default: () => [],
    description: '表头属性列表'
  },
  keyData:{
    type: String,
    default: () => '',
    description: '主键'
  },
  searchData:{
    type: String,
    default: () => '',
    description: '搜索框的搜索依据'
  },
  addForm:{
    type: Object,
    default: () => {},
    description: '添加界面数据'
  },
  editForm:{
    type: Object,
    default: () => {},
    description: '编辑界面数据'
  },
  deleteDataBody:{
    type: Object,
    default: () => {},
    description: '删除请求体'
  },
  urls:{
    type: Object,
    default: () => {},
    description: '请求url'
  }
});

//初始化函数
onMounted(async () => {
  await getFKList()
  state.allDataArray = await getData()
  update(1)
})

//获取外键列表
async function getFKList() {
  const addFormItem = prop.addForm.item  //添加窗口外键数据
  const editFormItem = prop.editForm.item  //编辑窗口外键数据
  const tableColItem = prop.tableColList  //表格显示外键映射
  const urlList = []  //外键url列表
  const FKDataMap = new Map()  //当前已获取外键, 避免重复请求

  //添加窗口外键获取
  for (const i in addFormItem) {
    if (addFormItem[i].isFK === true) {
      const addUrl = addFormItem[i].FKData.url  //外键url
      const addProp = addFormItem[i].FKData.property  //外键名

      //如果外键url不在列表中, 说明还没有请求过该外键
      if(!urlList.includes(addUrl)){
        urlList.push(addUrl)
        const data = await getData(addUrl)
        state.addFKMap.set(addProp, data)
        FKDataMap.set(addProp, data)
      }
      //第一个for循环不需要else判断

    }
  }

  //编辑窗口外键获取
  for (const i in editFormItem) {
    if (editFormItem[i].isFK === true) {
      const editUrl = editFormItem[i].FKData.url
      const editProp = editFormItem[i].FKData.property

      //如果外键url不在列表中, 说明还没有请求过该外键
      if(!urlList.includes(editUrl)){
        urlList.push(editUrl)
        const data = await getData(editUrl)
        state.editFKMap.set(editProp, data)
        FKDataMap.set(editProp, data)
      }
      //如果已经请求过该外键, 则在已获取外键中取值, 无需反复请求
      else{
        state.editFKMap.set(editProp, FKDataMap.get(editProp))
      }

    }
  }

  //表格显示外键映射
  for (const i in tableColItem){
    if (tableColItem[i].isFK === true) {
      const showUrl = tableColItem[i].FKData.url
      const showProp = tableColItem[i].FKData.property

      if(!urlList.includes(showUrl)){
        urlList.push(showUrl)
        const data = await getData(showUrl)
        state.showFKMap.set(showProp, data)
        FKDataMap.set(showProp, data)
      }
      //如果已经请求过该外键, 则在已获取外键中取值, 无需反复请求
      else{
        state.showFKMap.set(showProp, FKDataMap.get(showProp))
      }

    }
  }
}

//页数更新时更新数据
function update(currentPage){
  state.currentPage = currentPage
  //获取数据总数
  const dataSize = state.allDataArray.length

  // 计算总页数
  state.pageCount = Math.ceil(dataSize / PAGE_SIZE);

  //判断数据更新后当前页数是否大于总页数
  if(currentPage > state.pageCount) {
    currentPage = state.pageCount
  }

  // 计算当前页的第一个元素和最后一个元素在 allDataArray 中的索引
  const startIndex = (currentPage - 1) * PAGE_SIZE;
  const endIndex = Math.min(startIndex + PAGE_SIZE, dataSize);

  // 使用这些索引来获取对应的元素
  state.currentDataArray = state.allDataArray.slice(startIndex, endIndex);
}

//点击子组件的添加按钮, 子组件处理完返回的可提交表单
function add(form){
  addData(form)
}

//点击子组件的上传按钮, 子组件处理完返回的可提交表单
function upload(form){
  uploadData(form)
}

//点击子组件的删除按钮
function del(row){
  console.log(JSON.stringify(row))
  prop.deleteDataBody[prop.keyData] = row[prop.keyData]
  deleteData(prop.deleteDataBody)
}

//点击子组件的编辑按钮, 子组件处理完返回的可提交表单
function edit(form){
  updateData(form)
}

const state =  reactive({
  pageCount: 1, //数据总页数
  currentPage: 1, //当前页数
  isLoading: true,  //数据是否正在加载
  allDataArray: [],  //所有表格展示数据列表
  currentDataArray: [], //当前表格展示数据列表
  addFKMap: new Map(),  //添加窗口外键
  editFKMap: new Map(),  //编辑窗口外键
  showFKMap: new Map(),  //显示映射外键
})


const token="bearer "+localStorage.getItem("token");

/**
 * getData()
 * 获取数据的请求
 * 打开网页时自动调用一次
 * 结果：dataArray中包含一到多个data对象
 * */

const getData = async (url = prop.urls.getData) => {
  state.isLoading = true
  let resultList = []
  await axios.get(url, {
    headers: {
      'Authorization': token
    }
  })
      .then(result => {
        console.log("getData:", result)
        if (result && result.data && result.data.rows) {
          for (let i = 0; i < result.data.rows.length; i++){
            result.data.rows[i].created_at = new Date(result.data.rows[i].created_at).toLocaleString()
          }
          resultList = result.data.rows;
        }
      })
      .catch(error => {
        ElMessage.error("网络请求出错了！")
        console.error("getData:", error)
      })
  state.isLoading = false
  return resultList
}

/**
 * deleteData()
 * 删除数据的请求，param: data 类型:Object
 * */
const deleteData=async (data) => {
  state.isLoading = true
  await axios.delete(prop.urls.deleteData, {
        headers: {
          'Authorization': token
        },
        data: data
      }
  )
      .then( async message => {
        ElMessage.success("数据已被删除！")
        console.log("deleteData:", message)
        state.allDataArray = await getData()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！")
        console.error("deleteData:", error)
      })
  update(state.currentPage)
  state.isLoading = false
}

/**
 * addData()
 * 新增数据 param: newData对象
 * */
const addData=async (newData) => {
  console.log("sss",newData)
  state.isLoading = true
  await axios.post(prop.urls.addData, newData, {
    headers: {
      'Authorization': token
    }
  })
      .then( async message => {
        ElMessage.success("数据添加成功！")
        console.log("addData:", message)
        state.allDataArray = await getData()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！该数据是否已被添加？")
        console.error("addData:", error)
      })
  update(state.currentPage)
  state.isLoading = false
}

/**
 * updateData()
 * 修改数据  param: newData对象
 * */
const updateData=async (newData) => {
  state.isLoading = true
  await axios.put(prop.urls.updateData, newData, {
        headers: {
          'Authorization': token
        },
      }
  ).then( async message => {
    ElMessage.success("数据修改成功！")
    console.log("updateData:", message)
    state.allDataArray = await getData()
  })
      .catch( error => {
        ElMessage.error("网络请求出错了！")
        console.error("updateData:", error)
      })
  update(state.currentPage)
  state.isLoading = false
}


/**
 * uploadData()
 * 批量上传数据  param: list
 * */
const uploadData=async (list) => {
  console.log(list)
  state.isLoading = true
  await axios.post(prop.urls.uploadData, {list: list}, {
    headers: {
      'Authorization': token
    }
  })
      .then( async message => {
        if(message.status === 200) {
          ElMessage.success("批量导入成功！")
        }
        else if(message.status === 217){
          ElMessage.warning("部分导入成功")
          let errMessage = ""
          let offset = 0
          for(const item in message.data.detail){
            if(message.data.detail[item].code === 402){
              errMessage += (message.data.detail[item].message + "\n")
              ElNotification({
                title: '失败信息',
                message: message.data.detail[item].message,
                type: 'error',
                duration: 10000,
                offset: offset,
              })
              offset += 100
            }
          }
        }
        console.log("uploadData:", message)
        state.allDataArray = await getData()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！数据是否已被添加过？")
        console.error("uploadData:", error)
      })
  update(state.currentPage)
  state.isLoading = false
}
</script>

<style scoped>

</style>