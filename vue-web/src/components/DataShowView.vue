<template>
  <my-table
      ref="myTable"
      :table-col-list="tableColList"
      :default-data="state.dataArray"
      :search-data="searchData"
      :add-data-template="addForm"
      :edit-data-template="editForm"
      :loading="state.isLoading"
      :add-f-k-list="state.addFKList"
      :edit-f-k-list="state.editFKList"
      :show-f-k-list="state.showFKList"
      :key-data="keyData"
      @add="add"
      @upload="upload"
      @del="del"
      @edit="edit"
  >
  </my-table>
</template>

<script setup>
import axios from "axios";
import {ElMessage, ElNotification} from "element-plus";
import MyTable from "@/components/MyTable.vue";
import {onMounted, reactive} from "vue";

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
  state.dataArray = await getData()
  await getFKList()
})

//获取外键列表
async function getFKList() {
  const addFormItem = prop.addForm.item
  let addUrl = ''
  let editUrl = ''
  let showUrl = ''
  for (const i in addFormItem) {
    if (addFormItem[i].isFK === true) {
      // state.addFKList.push({
      //   name: addFormItem[i].FKData.property,
      //   data: await getData(addFormItem[i].FKData.url)
      // })
      addUrl = addFormItem[i].FKData.url
      state.addFKList = await getData(addUrl)
    }
  }
  const editFormItem = prop.editForm.item
  for (const i in editFormItem) {
    if (editFormItem[i].isFK === true) {
      // state.editFKList.push({
      //   name: editFormItem[i].FKData.property,
      //   data: await getData(editFormItem[i].FKData.url)
      // })
      editUrl = editFormItem[i].FKData.url
      if(editUrl === addUrl) {
        state.editFKList = state.addFKList
      }
      else{
        state.editFKList = await getData(editUrl)
      }
    }
  }
  for (const i in prop.tableColList){
    if(prop.tableColList[i].isFK === true){
      showUrl = prop.tableColList[i].FKData.url
      if(showUrl === addUrl) {
        state.showFKList = state.addFKList
      }
      else if(showUrl === editUrl) {
        state.showFKList = state.editFKList
      }
      else{
        state.showFKList = await getData(showUrl)
      }
    }
  }
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
  isLoading: true,  //数据是否正在加载
  dataArray: [],  //接收用户对象的列表
  addFKList: [],  //添加窗口外键列表
  editFKList: [],  //编辑窗口外键列表
  showFKList: [],  //显示映射外键列表
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
        state.dataArray = await getData()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！")
        console.error("deleteData:", error)
      })
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
        state.dataArray = await getData()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！该数据是否已被添加？")
        console.error("addData:", error)
      })
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
    state.dataArray = await getData()
  })
      .catch( error => {
        ElMessage.error("网络请求出错了！")
        console.error("updateData:", error)
      })
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
          for(const item in message.data.detail){
            if(message.data.detail[item].code === 402){
              errMessage += (message.data.detail[item].message + "\n")
              ElNotification({
                title: '失败信息',
                message: message.data.detail[item].message,
                type: 'error',
                duration: 30000,
              })
            }
          }
        }
        console.log("uploadData:", message)
        state.dataArray = await getData()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！数据是否已被添加过？")
        console.error("uploadData:", error)
      })
  state.isLoading = false
}
</script>

<style scoped>

</style>