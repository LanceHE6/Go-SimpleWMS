<template>
  <my-table
      ref="myTable"
      :table-col-list="tableColList"
      :default-data="data.warehouseArray"
      :search-data="'name'"
      :add-data-template="addForm"
      :edit-data-template="editForm"
      :loading="data.isLoading"
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

//表头属性列表
const tableColList = [
  {property: "name", label: "仓库名", sortable: false},
  {property: "comment", label: "备注", sortable: false},
  {property: "manager", label: "管理人", sortable: false},
  {property: "status", label: "状态", sortable: false},
  {property: "add_time", label: "创建时间", sortable: true},
  {property: "wid", label: "仓库ID", sortable: true, width: 100},
]

//初始化函数
onMounted(() => {
  warehouseList()
})

//点击子组件的添加按钮, 子组件处理完返回的可提交表单
function add(form){
  warehouseAdd(form)
}

//点击子组件的上传按钮, 子组件处理完返回的可提交表单
function upload(form){
  warehouseUpload(form)
}

//点击子组件的删除按钮
function del(row){
  warehouseDelete(row.uid)
}

//点击子组件的编辑按钮, 子组件处理完返回的可提交表单
function edit(form){
  warehouseUpdate(form)
}

const data =  reactive({
  warehouseArray: [],  //接收用户对象的列表
  isLoading: true  //是否正在加载
})

/**
 * 编辑用户时所用到的表单对象
 * */
const editForm = reactive({
  data : {
    wid:'',
    name:'',
    comment:'',
    manager:'',
    status:'',
  },
  dataNum: 5,
  rules: {
    status:[
      { required: 'true', message: '请选择仓库状态', trigger: 'blur' }
    ],
    name:[
      { required: 'true', message: '请输入仓库名', trigger: 'blur' },
    ],
    manager:[
      { required: 'true', message: '请选择管理人', trigger: 'blur' },
    ]
  },
  item:[
    {label: '仓库名', prop: 'name', dataName: 'name', isInput: true,},
    {label: '是否启用', prop: 'status', dataName: 'status', isSelect: true,
      selectOptions: [
        {label: '禁用', value: 0},
        {label: '启用', value: 1},
      ]},
    {label: '管理人', prop: 'manager', dataName: 'manager', isInput: true},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true},
  ],
  key: 'wid'
})

/**
 * 注册用户时所用到的对象
 * */
const addForm = reactive({
  data :{
    name:'',
    comment:'',
    manager:'',
    status:'',
  },
  dataNum: 4,
  rules: {
    name: [
      { required: 'true', message: '仓库名不能为空', trigger: 'blur' },
    ],
    status:[
      { required: 'true', message: '请选择仓库状态', trigger: 'blur' }
    ],
    manager:[
      { required: 'true', message: '请选择管理人', trigger: 'blur' }
    ]
  }
  ,
  item:[
    {label: '仓库名', prop: 'name', dataName: 'name', isInput: true},
    {label: '备注', prop: 'comment', dataName: 'comment', isInput: true},
    {label: '管理人', prop: 'manager', dataName: 'manager', isInput: true},
    {label: '仓库状态', prop: 'status', dataName: 'status', isSelect: true,
      selectOptions: [
        {label: '禁用', value: 0},
        {label: '启用', value: 1},
      ]},
  ],
})

const token="bearer "+localStorage.getItem("token");

/**
 * warehouseList()
 * 获取仓库的信息的请求
 * 打开网页时自动调用一次
 * 结果：warehouseArray中包含一到多个warehouse对象
 * */

const warehouseList = async () => {
  data.isLoading = true
  await axios.get('/warehouse/list', {
    headers: {
      'Authorization': token
    }
  })
      .then(result => {
        console.log("warehouseList:", result)
        if (result && result.data) {
          for (let i = 0; i < result.data.rows.length; i++){
            result.data.rows[i].add_time = new Date(Number(result.data.rows[i].add_time) * 1000).toLocaleString()
          }
          data.warehouseArray = result.data.rows;
        }
      })
      .catch(error => {
        ElMessage.error("网络请求出错了！")
        console.error("warehouseList:", error)
      })
  data.isLoading = false
}

/**
 * warehouseDelete()
 * 删除仓库的请求，param: wid 类型:String
 * */
const warehouseDelete=async (wid) => {
  data.isLoading = true
  await axios.delete('/warehouse/delete', {
        headers: {
          'Authorization': token
        },
        data: {
          "wid": wid
        }
      }
  )
      .then( async message => {
        ElMessage.success("仓库已被删除！")
        console.log("warehouseDelete:", message)
        await warehouseList()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！")
        console.error("warehouseDelete:", error)
      })
  data.isLoading = false
}

/**
 * warehouseAdd()
 * 新增仓库 param: warehouse对象
 * warehouse={
 *     name:'',
 *     comment:'',
 *     manager:'',
 *     status:0,
 * }
 * */
const warehouseAdd=async (warehouse) => {
  data.isLoading = true
  await axios.post('/warehouse/add', warehouse, {
    headers: {
      'Authorization': token
    }
  })
      .then( async message => {
        ElMessage.success("仓库添加成功！")
        console.log("warehouseAdd:", message)
        await warehouseList()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！仓库名是否已被定义？")
        console.error("warehouseAdd:", error)
      })
  data.isLoading = false
}

/**
 * warehouseUpdate()
 * 修改仓库  param: warehouse对象
 * warehouse={
 *     wid:'',
 *     name:'',
 *     comment:'',
 *     manager:'',
 *     status:0,
 * }
 * */
const warehouseUpdate=async (warehouse) => {
  data.isLoading = true
  await axios.put('/warehouse/update', warehouse, {
        headers: {
          'Authorization': token
        },
      }
  ).then( async message => {
    ElMessage.success("仓库信息修改成功！")
    console.log("warehouseUpdate:", message)
    await warehouseList()
  })
      .catch( error => {
        ElMessage.error("网络请求出错了！")
        console.error("warehouseUpdate:", error)
      })
  data.isLoading = false
}

const warehouseUpload=async (list) => {
  console.log(list)
  data.isLoading = true
  await axios.post('/warehouse/register-batch', {warehouse_list: list}, {
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
        console.log("warehouseUpload:", message)
        await warehouseList()
      })
      .catch( error => {
        ElMessage.error("网络请求出错了！仓库名是否已被定义？")
        console.error("warehouseUpload:", error)
      })
  data.isLoading = false
}
</script>

<style scoped>

</style>