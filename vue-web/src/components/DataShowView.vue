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
      :operations="state.operations"
      :large="large"
      :has-submit-page="hasSubmitPage"
      :height="tableHeight"
      @add="add"
      @upload="upload"
      @del="del"
      @edit="edit"
      @update="update"
      @search="startSearch"
      @refresh="initialize"
      @upload-img="uploadImg"
  >
  </my-table>
</template>

<script setup>
import {ElMessage, ElNotification} from "element-plus";
import {onMounted, reactive, ref} from "vue";
import MyTable from "@/components/MyTable.vue";
import {axiosDelete, axiosGet, axiosPost, axiosPut} from "@/utils/axiosUtil.js";
import {editObjKeyData, getObjKeyData} from "@/utils/objectUtil.js";

const PAGE_SIZE = 10 //每页展示多少个数据

const prop = defineProps({
  tableHeight:{
    type: String,
    default: () => '60vh',
    description: '表格高度'
  },
  large:{
    type: Boolean,
    default: () => false,
    description: '是否启用大数据表格(由后端分页的数据)'
  },
  tableColList:{
    type: Array,
    default: () => [],
    description: '表头属性列表'
  },
  keyData:{
    type: [String, Array],
    default: () => undefined,
    description: '数据主键, 如有多层则用列表封装, 格式如: ["key", "p1", "p2", ...]'
  },
  searchData:{
    type: String,
    default: () => '',
    description: '搜索框的搜索依据'
  },
  addForm:{
    type: Object,
    default: () => null,
    description: '添加界面数据'
  },
  editForm:{
    type: Object,
    default: () => null,
    description: '编辑界面数据'
  },
  deleteDataBody:{
    type: Object,
    default: () => null,
    description: '删除请求体'
  },
  urls:{
    type: Object,
    default: () => {},
    description: '请求url'
  },
  extraParams:{
    type: Object,
    default: () => {},
    description: '数据获取请求体中的额外参数'
  },
  upload:{
    type: Boolean,
    default: () => false,
    description: '是否支持上传功能'
  },
  download:{
    type: Boolean,
    default: () => false,
    description: '是否支持下载功能'
  },
  print:{
    type: Boolean,
    default: () => false,
    description: '是否支持打印功能'
  },
  uploadImg:{
    type: Boolean,
    default: () => false,
    description: '是否支持图片上传功能'
  },
  hasSubmitPage:{
    type: Boolean,
    default: () => false,
    description: '是否需要提交表单的页面, 如果为false则使用简单的窗口来提交表单'
  }
});

//表格对象
const myTable = ref(null)

const state =  reactive({
  pageCount: 1, //数据总页数
  currentPage: 1, //当前页数
  searchWord: '',  //后端查询关键词
  isLoading: true,  //数据是否正在加载
  allDataArray: [],  //所有表格展示数据列表
  currentDataArray: [], //当前表格展示数据列表
  addFKMap: new Map(),  //添加窗口外键
  editFKMap: new Map(),  //编辑窗口外键
  showFKMap: new Map(),  //显示映射外键
  operations: {
    add: prop.addForm !== null,
    del: prop.deleteDataBody !== null,
    edit: prop.editForm !== null,
    upload: prop.upload,
    download: prop.download,
    print: prop.print,
    uploadImg: prop.uploadImg
  }
})

//对外事件列表
const emit = defineEmits(["addTab", "removeTab"]);

//获取当前已选项列表
const getMultipleSelection = () => myTable.value.getMultipleSelection()
//清空表格选择
const clearSelection = () => myTable.value.clearSelection()
//暴露函数，可供父组件调用
defineExpose({
  getMultipleSelection,
  clearSelection,
  update
});

//初始化函数
async function initialize(){
  state.isLoading = true
  await getAllFKList()
  if(!prop.large) {
    state.allDataArray = await getData(prop.urls['getData'])
  }
  await update(state.currentPage)
  state.isLoading = false
}
onMounted(initialize)


async function getDataAndSet(url, propName, stateMap, FKDataMap) {
  if (!FKDataMap.has(propName)) {
    const data = await getData(url);
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
  const addFormItem = prop.addForm && prop.addForm.item;  // 添加窗口外键数据
  const editFormItem = prop.editForm && prop.editForm.item;  // 编辑窗口外键数据
  const tableColItem = prop.tableColList;  // 表格显示外键映射
  const FKDataMap = new Map();  // 当前已获取外键, 避免重复请求

  // 添加窗口外键获取
  await getFKList(addFormItem, state.addFKMap, FKDataMap)

  // 编辑窗口外键获取
  await getFKList(editFormItem, state.editFKMap, FKDataMap)

  // 表格显示外键映射
  await getFKList(tableColItem, state.showFKMap, FKDataMap)

  //子表格显示外键映射
  for (const item of tableColItem) {
    if ('children' in item) {
      await getFKList(item.children, state.showFKMap, FKDataMap)
    }
  }
}

//页数更新时更新数据
async function update(currentPage) {
  if (!prop.large) {  //前端分页
    //获取数据总数
    const dataSize = state.allDataArray.length

    // 计算总页数
    state.pageCount = Math.max(Math.ceil(dataSize / PAGE_SIZE), 1);

    //判断数据更新后当前页数是否大于总页数
    if (currentPage > state.pageCount) {
      currentPage = state.pageCount
    }
    state.currentPage = currentPage

    // 计算当前页的第一个元素和最后一个元素在 allDataArray 中的索引
    const startIndex = (currentPage - 1) * PAGE_SIZE;
    const endIndex = Math.min(startIndex + PAGE_SIZE, dataSize);

    // 使用这些索引来获取对应的元素
    state.currentDataArray = state.allDataArray.slice(startIndex, endIndex);
  } else {  //后端分页
    state.isLoading = true
    const defaultParams = {
      page: currentPage,
      page_size: PAGE_SIZE,
      keyword: state.searchWord
    }
    state.currentDataArray = await getData(prop.urls['getData'], {...defaultParams, ...prop.extraParams})
    //判断数据更新后当前页数是否大于总页数
    if (currentPage > state.pageCount) {
      currentPage = state.pageCount
    }
    state.currentPage = currentPage
    state.isLoading = false
  }
}

//后端分页时点击查询按钮回调此函数
async function startSearch(s) {
  state.isLoading = true
  state.searchWord = s
  state.currentPage = 1
  const defaultParams = {
    page: state.currentPage,
    page_size: PAGE_SIZE,
    keyword: state.searchWord
  }
  state.currentDataArray = await getData(prop.urls['getData'], {...defaultParams, ...prop.extraParams})
  state.isLoading = false
}

//点击子组件的添加按钮, 子组件处理完返回的可提交表单
function add(form){
  if(!prop.hasSubmitPage){
    addData(form)
  }
  else{
    emit("addTab")
  }
}

//点击子组件的上传按钮, 子组件处理完返回的可提交表单
function upload(form){
  uploadData(form)
}

//点击子组件的删除按钮
function del(row){
  console.log(JSON.stringify(row))
  if(typeof prop.keyData === "string") {
    prop.deleteDataBody[prop.keyData] = row[prop.keyData]
  }
  else{
    prop.deleteDataBody = editObjKeyData(prop.deleteDataBody, getObjKeyData(row, prop.keyData), prop.keyData)
  }
  deleteData(prop.deleteDataBody)
}

//上传图片
async function uploadImg(id, fileList) {

  // 创建一个新的FormData对象
  const formData = new FormData();
  formData.append('goods', id);

  // 遍历文件列表并添加到FormData中
  fileList.forEach((file, _) => {
    formData.append('image', file);
  });


  // 调用uploadImage函数并传入formData
  await uploadImage(formData);
}

//点击子组件的编辑按钮, 子组件处理完返回的可提交表单
function edit(form){
  updateData(form)
}

const token="bearer "+localStorage.getItem("token");

/**
 * getData()
 * 获取数据的请求
 * */

const getData = async (url, params = {}) => {
  const result = await axiosGet({url: url, params: params, name: 'getData'})
  if (result && result.data && result.data.rows) {
    if(prop.large){
      state.pageCount = Math.max(result.data['total_pages'], 1)
    }
    return result.data.rows
  }
  else{
    return undefined
  }
}

/**
 * deleteData()
 * 删除数据的请求
 * */
const deleteData=async (data) => {
  state.isLoading = true
  const result = await axiosDelete({url: prop.urls['deleteData'], data: data, name: 'deleteData'})
  if(result){
    ElMessage.success("数据已被删除！")
    state.allDataArray = await getData(prop.urls['getData'])
    await update(state.currentPage)
  }
  state.isLoading = false
}

/**
 * addData()
 * 新增数据
 * */
const addData=async (data) => {
  state.isLoading = true
  const result = await axiosPost({url: prop.urls['addData'], data: data, name: 'addData'})
  if(result){
    ElMessage.success("数据添加成功！")
    state.allDataArray = await getData(prop.urls['getData'])
    await update(state.currentPage)
  }
  state.isLoading = false
}

/**
 * updateData()
 * 修改数据
 * */
const updateData=async (data) => {
  state.isLoading = true
  const result = await axiosPut({url: prop.urls['updateData'], data: data, name: 'updateData'})
  if(result){
    ElMessage.success("数据修改成功！")
    state.allDataArray = await getData(prop.urls['getData'])
    await update(state.currentPage)
  }
  state.isLoading = false
}


/**
 * uploadData()
 * 批量上传数据
 * */
const uploadData=async (list) => {
  state.isLoading = true
  const result = await axiosPost({url: prop.urls['uploadData'], data: {list: list}, name: 'uploadData'})
  if(result){
    if(result.status === 200) {
      ElMessage.success("批量导入成功！")
    }
    else if(result.status === 217){
      ElMessage.warning("部分导入成功")
      let errMessage = ""
      let offset = 0
      for(const item in result.data.detail){
        if(result.data.detail[item].code === 402){
          errMessage += (result.data.detail[item].message + "\n")
          ElNotification({
            title: '失败信息',
            message: result.data.detail[item].message,
            type: 'error',
            duration: 10000,
            offset: offset,
          })
          offset += 100
        }
      }
    }
    state.allDataArray = await getData(prop.urls['getData'])
    await update(state.currentPage)
  }
  state.isLoading = false
}

/**
 * uploadImage()
 * 上传图片
 * */
const uploadImage=async (data) => {
  state.isLoading = true
  const headers = {
    'Content-Type': 'multipart/form-data'
  }
  const result = axiosPost({url: prop.urls['uploadImage'], data: data, headers: headers, name: 'uploadImage'})
  if(result){
    ElMessage.success("图片上传成功！")
    state.allDataArray = await getData(prop.urls['getData'])
    await update(state.currentPage)
  }
  state.isLoading = false
}
</script>

<style scoped>

</style>