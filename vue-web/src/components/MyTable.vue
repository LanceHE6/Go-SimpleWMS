<template>
  <el-container
    v-loading="loading"
  >
    <el-header
      style="height: 20px"
    >

      <table-header
          :operations="operations"
          :has-submit-page="hasSubmitPage"
          :large="large"
          @add="add"
          @download="download"
          @upload="upload"
          @search="searchChange"
          @print="print"
          @refresh="refresh"
      />

    </el-header>

    <el-main>
      <table-body
        ref="myTable"
        :default-data="defaultData"
        :table-col-list="tableColList"
        :key-data="keyData"
        :search-data="searchData"
        :show-f-k-map="showFKMap"
        :operations="operations"
        :height="height"
        :large="large"
        :search="search"
        @edit="edit"
        @delete="confirmDel"
        @upload-img="uploadImg"
        @view-files="viewFiles"
        @selection-change="handleSelectionChange"
      />
    </el-main>

    <el-footer
        style="margin-top: 5px"
    >
      <el-pagination
          v-model:current-page="currentPage"
          :page-count="pageCount"
          background
          layout="prev, pager, next, jumper"
          @current-change="pageChange"
      />
    </el-footer>
  </el-container>

  <el-dialog
      v-if="!hasSubmitPage && operations.edit"
      v-model="editFormVisible"
      title="编辑"
      width="500"
      center
  >

    <el-form :model="editForm.data" :rules="editForm.rules" ref="myEditForm" label-position="top" :class="editDialogClass" status-icon>

      <el-form-item
        v-for="item in editForm.item"
        :label="item.label"
        :prop="item.prop"
      >
        <el-input
          v-if="item.isInput"
          v-model.trim="editForm.data[item.dataName]"
          :type="item.type"
          :show-password="item.type === 'password'"
          autocomplete="off"
        />

        <el-select
            v-if="item.isSelect"
            v-model.trim="editForm.data[item.dataName]"
            placeholder="请选择"
        >
          <el-option
              v-for="i in item.selectOptions"
              :label="i.label"
              :value="i.value"
          />
        </el-select>

        <el-select
            v-if="item.isFK"
            v-model.trim="editForm.data[item.dataName]"
            placeholder="请选择"
        >
          <el-option
              v-for="i in Array.from(editFKMap.get(item.FKData.property))"
              :disabled="'status' in i && i['status'] === 0"
              :label="i[item.FKData.label]"
              :value="i[item.FKData.property]"
          />
        </el-select>

      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="editFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitEditForm(myEditForm)">
          确定
        </el-button>
      </div>
    </template>

  </el-dialog>
  <el-dialog
      v-if="!hasSubmitPage && operations.add"
      v-model="addFormVisible"
      title="添加"
      width="500"
      center
  >

    <el-form :model="addForm.data" :rules="addForm.rules" ref="myAddForm" label-position="top" :class="addDialogClass" status-icon>
      <el-form-item
          v-for="item in addForm.item"
          :label="item.label"
          :prop="item.prop"
      >
        <el-input
            v-if="item.isInput"
            v-model.trim="addForm.data[item.dataName]"
            :type="item.type"
            :show-password="item.type === 'password'"
            autocomplete="off"
        />

        <el-select
            v-if="item.isSelect"
            v-model.trim="addForm.data[item.dataName]"
            placeholder="请选择"
        >
          <el-option
              v-for="i in item.selectOptions"
              :label="i.label"
              :value="i.value"
          />
        </el-select>

        <el-select
            v-if="item.isFK"
            v-model.trim="addForm.data[item.dataName]"
            placeholder="请选择"
        >
          <el-option
              v-for="i in Array.from(addFKMap.get(item.FKData.property))"
              :disabled="'status' in i && i['status'] === 0"
              :label="i[item.FKData.label]"
              :value="i[item.FKData.property]"
          />
        </el-select>

      </el-form-item>

    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="addFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitAddForm(myAddForm)">
          确定
        </el-button>
      </div>
    </template>

  </el-dialog>
  <el-dialog
      v-if="!hasSubmitPage && operations.upload"
      v-model="uploadFormVisible"
      title="批量导入"
      width="700"
      @closed="uploadDialogClosed"
      center
  >
    <el-upload
        ref="myUploadForm"
        action=""
        accept="application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
        :auto-upload="false"
        :http-request="excelToJson"
        :limit="1"
        drag
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        拖拽文件到此处或 <em>点击此处上传</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          <el-text type="info">支持xls或xlsx类型文件, </el-text>
          <el-button type="primary" @click="downloadTemplate" text>点击此处下载提交模版</el-button>
          <el-text type="info">, 输入格式参考导出表格中的格式 </el-text>
        </div>
      </template>
    </el-upload>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="uploadFormVisible = false">取消</el-button>
        <el-button type="primary" @click="submitUploadData()">
          上传
        </el-button>
      </div>
    </template>

  </el-dialog>
  <el-dialog
      v-if="operations.uploadImg"
      v-model="uploadImgVisible"
      title="上传图片"
      width="700"
      @opened="uploadImgDialogOpen"
      center
  >
    <el-upload
        ref="myUploadImgForm"
        class="avatar-uploader"
        accept="image/jpeg, image/png, image/gif, image/bmp, image/x-bmp, image/webp, image/tiff, image/x-tiff, image/svg+xml"
        list-type="picture-card"
        :auto-upload="false"
        :limit="5"
        :on-preview="handlePictureCardPreview"
        :on-change="handleImgChange"
        :on-remove="(file) => handleRemove(file, imgList)"
        :on-exceed="uploadExceed"
    >
      <el-icon><Plus /></el-icon>
      <template #tip>
        <div class="el-upload__tip">
          <el-text type="info">支持jpg、png、svg、webp等图片类型文件，</el-text>
          <el-text type="warning">最多上传五张图片。</el-text>
        </div>
      </template>
    </el-upload>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="uploadImgVisible = false">取消</el-button>
        <el-button type="primary" @click="submitUploadImgData()">
          上传
        </el-button>
      </div>
    </template>

  </el-dialog>
  <el-dialog
      v-model="viewFileVisible"
      title="附件列表"
      width="800"
      center
  >
    <div>
      <el-button
        type="primary"
        icon="Upload"
        style="margin-bottom: 10px"
        :disabled="!operations.uploadFile"
        @click="uploadFileVisible = true"
        plain
      >
        上传附件
      </el-button>
    </div>
    <table-body
      :table-col-list="[
          {property: 'path', label: '文件名', sortable: false, operable: true, operationEvent: previewFile},
      ]"
      :default-data="viewFileList"
      :operations="{
        edit: false,
        delete: false,
      }"
      height="50vh"
    />
  </el-dialog>
  <el-dialog
      v-if="operations.uploadFile"
      v-model="uploadFileVisible"
      title="上传附件"
      width="700"
      @opened="uploadFileDialogOpen"
      center
  >
    <el-upload
        ref="myUploadFileForm"
        accept="text/plain, text/markdown, application/msword, application/vnd.ms-excel, application/vnd.ms-powerpoint, application/pdf,
        application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.openxmlformats-officedocument.presentationml.presentation,
        application/vnd.openxmlformats-officedocument.wordprocessingml.document"
        :auto-upload="false"
        :limit="5"
        :on-change="handleFileChange"
        :on-remove="handleRemove"
        :on-exceed="uploadExceed"
        drag
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        拖拽文件到此处或 <em>点击此处上传</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          <el-text type="info">支持txt、doc、xls、ppt、pdf等类型文件，</el-text>
          <el-text type="warning">最多上传五个文件。</el-text>
        </div>
      </template>
    </el-upload>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="uploadFileVisible = false">取消</el-button>
        <el-button type="primary" @click="submitUploadFileData()">
          上传
        </el-button>
      </div>
    </template>

  </el-dialog>

  <el-dialog v-model="imgDialogVisible" class="img-dialog">
    <el-image :src="dialogImageUrl" alt="Preview Image" :fit="'contain'" class="preview-img"/>
  </el-dialog>
</template>

<script setup>
import {markRaw, ref, watch} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {Delete, Plus, UploadFilled} from "@element-plus/icons-vue";
import * as XLSX from "xlsx";
import TableHeader from "@/components/TableHeader.vue";
import {editObjKeyData, getObjKeyData} from "@/utils/objectUtil.js";
import TableBody from "@/components/TableBody.vue";
import axios from "axios";
import EventBus from "@/utils/eventBus.js";

const prop = defineProps({
  large:{
    type: Boolean,
    default: () => false,
    description: '是否启用大数据表格(由后端分页的数据)'
  },
  pageCount:{
    type: Number,
    default: () => 1,
    description: '数据总页数'
  },
  keyData:{
    type: [String, Array],
    default: () => undefined,
    description: '数据主键, 如有多层则用列表封装, 格式如: ["key", "p1", "p2", ...]'
  },
  loading:{
    type: Boolean,
    default: () => true,
    description: '组件是否正在加载'
  },
  defaultData: {
    type: Array,
    default: () => [],
    description: '默认table数据列表'
  },
  tableColList:{
    type: Array,
    default: () => [],
    description: 'table显示列, 用于定义该table的首列元素'+
        '\n列表内对象格式为：{property: "属性名", label: "显示名", sortable: "是否能排序", width: "宽度"}'
  },
  searchData:{
    type: String,
    default: () => '',
    description: 'table搜索框的搜索依据, 即根据什么属性做搜索'
  },
  addDataTemplate:{
    type: Object,
    default: () => {},
    description: '新增数据模版, 包含数据模版以及约束模版'+
        '\n对象格式为：{data: "数据模版对象", rules: "约束模版对象"}'
  },
  editDataTemplate:{
    type: Object,
    default: () => {},
    description: '编辑数据模版, 包含数据模版以及约束模版'+
        '\n对象格式为：{data: "数据模版对象", rules: "约束模版对象"}'
  },
  addFKMap:{
    type: Map,
    default: () => null,
    description: '添加窗口外键数据'+
        '\n其中key为外键名, data为外键对象'
  },
  editFKMap:{
    type: Map,
    default: () => null,
    description: '编辑窗口外键数据'+
        '\n其中key为外键名, data为外键对象'
  },
  showFKMap:{
    type: Map,
    default: () => null,
    description: '显示外键数据'+
        '\n其中key为外键名, data为外键对象'
  },
  operations:{
    type: Object,
    default: () => null,
    description: '表格支持的操作, 包含增删查改以及导入导出等等(通常无需手动配置)'
  },
  hasSubmitPage:{
    type: Boolean,
    default: () => false,
    description: '是否需要提交表单的页面, 如果为false则使用简单的窗口来提交表单'
  },
  height:{
    type: String,
    default: () => '60vh',
    description: '表格高度'
  }
});
//对外事件列表
const emit = defineEmits([
  "add", "download", "upload",
  "edit", "del", "update",
  "search", "refresh", "uploadImg",
  "selectionChange", "uploadFile"
]);

//表格当前已选内容
let multipleSelection = []
const getMultipleSelection = () => multipleSelection
//清空表格选择
const clearSelection = () => myTable.value.clearSelection()
//暴露函数，可供父组件调用
defineExpose({
  getMultipleSelection,
  clearSelection
});

//上传图片列表
const imgList = ref([])
//上传附件列表
const fileList = ref([])

//表格高度
//const tableHeight = document.documentElement.clientHeight * 0.6
//const colHeight = ref(tableHeight / 10)

const dialogImageUrl = ref('')
const imgDialogVisible = ref(false)

//表格数据列表
const tableData = ref(prop.defaultData)
//编辑表单数据模版
const editForm = ref(prop.editDataTemplate)
//添加表单数据模版
const addForm = ref(prop.addDataTemplate)

//当前页数
const currentPage = ref(1)

//显示名表头
const tableHead = prop.tableColList.map( item => {
  return item.label
})

//表格
const myTable = ref(null)

//编辑表单
const myEditForm = ref(null)
//编辑表单是否可见
let editFormVisible = ref(false)

//新增表单
const myAddForm = ref(null)
//新增表单是否可见
let addFormVisible = ref(false)

//上传窗口组件
const myUploadForm = ref(null)
//上传表单是否可见
let uploadFormVisible = ref(false)

//上传图片窗口组件
const myUploadImgForm = ref(null)
//上传图片窗口是否可见
let uploadImgVisible = ref(false)
//上传图片的对应列id
let uploadImgId = ref('')

//查看文件窗口组件
const myViewFileForm = ref(null)
//查看文件窗口是否可见
let viewFileVisible = ref(false)
//文件列表
let viewFileList = ref([])

//上传文件窗口是否可见
let uploadFileVisible = ref(false)
//上传文件窗口组件
const myUploadFileForm = ref(null)
//上传文件的对应列id
let uploadFileId = ref('')

//监听数据变化并实时更新表格
watch(() => prop.defaultData, (newValue) => {
  tableData.value = newValue;
});

//搜索栏文字
const search = ref('')

//输入元素超过4个的窗口分列显示
const DIALOG_COL = 4
let addDialogClass = prop.operations.add
    ? ref(prop.addDataTemplate.dataNum > DIALOG_COL ? 'multi-column' : 'single-column')
    : null
let editDialogClass = prop.operations.edit
    ? ref(prop.editDataTemplate.dataNum > DIALOG_COL ? 'multi-column' : 'single-column')
    : null


//搜索
function searchChange(s){
  search.value = s
  if(prop.large){
    emit("search", search.value);
  }
}

function pageChange(){
  emit("update", currentPage.value);
}

function print(){

}

function refresh(){
  emit("refresh");
}

function add(){
  //如果用窗口提交表单则显示窗口
  if(!prop.hasSubmitPage){
    addFormVisible.value = true
  }
  //否则直接向父组件提交事件, 让父组件处理
  else{
    emit("add")
  }
}

async function submitAddForm(form) {
  if (!form) return
  await form.validate((valid) => {
    if (valid) {
      //类型转换
      for (const i in prop.addDataTemplate.dataType) {
        if (i in addForm.value.data) {
          if (prop.addDataTemplate.dataType[i] === "String") {
            addForm.value.data[i] = addForm.value.data[i].toString()
          }
          else if (prop.addDataTemplate.dataType[i] === "Int") {
            addForm.value.data[i] = addForm.value.data[i] !== '' ? parseInt(addForm.value.data[i]) : 0
          }
          else if (prop.addDataTemplate.dataType[i] === "Float") {
            addForm.value.data[i] = addForm.value.data[i] !== '' ? parseFloat(addForm.value.data[i]) : 0.0
          }
          else{
            console.log('unknown dataType', i)
          }
        }
      }
      //提交窗口数据给父组件
      emit("add", addForm.value.data);
      //清空原窗口
      for (const item in addForm.value.item) {
        addForm.value.data[addForm.value.item[item].dataName] = ''
      }
      //关闭窗口
      addFormVisible.value = false
    }
  })
}

//点击下载表格按钮
function download(){
  const workbook = XLSX.utils.book_new();
  const worksheet = XLSX.utils.json_to_sheet(tableData.value)
  console.log(tableData.value)
  XLSX.utils.book_append_sheet(workbook, worksheet, "sheet1")
  //XLSX.utils.sheet_add_aoa(worksheet, [tableHead], { origin: "A1" });
  worksheet["!cols"] = new Array(tableHead.length).fill({ wch: 25 });
  XLSX.writeFileXLSX(workbook, "导出表格.xlsx")
  emit("download");
}
function upload(){
  uploadFormVisible.value = true
}

function uploadImg(id){
  if(prop.operations.uploadImg === true){
    uploadImgVisible.value = true
    uploadImgId.value = id
  }
}

function viewFiles(data, id){
  if(prop.operations.uploadFile === true){
    viewFileVisible.value = true
    viewFileList.value = data
    uploadFileId.value = id
  }
}

function submitUploadData(){
  myUploadForm.value.submit()
  uploadFormVisible.value = false
}

function submitUploadImgData(){
  myUploadImgForm.value.clearFiles()
  emit("uploadImg",uploadImgId.value ,imgList.value)
  uploadImgVisible.value = false
}

function submitUploadFileData(){
  console.log("2333")
  myUploadFileForm.value.clearFiles()
  emit("uploadFile",uploadFileId.value ,fileList.value)
  viewFileVisible.value = false
  uploadFileVisible.value = false
  setTimeout(() => {
    EventBus.emit("refresh", 1)
  }, 500);
}
//提交编辑表单
async function submitEditForm(form){
  if (!form) return
  await form.validate((valid) => {
    if (valid) {
      //类型转换
      for (const i in prop.editDataTemplate.dataType) {
        if (editForm.value.data[i]) {
          if (prop.editDataTemplate.dataType[i] === "String") {
            editForm.value.data[i] = editForm.value.data[i].toString()
          }
          else if (prop.editDataTemplate.dataType[i] === "Int") {
            editForm.value.data[i] = editForm.value.data[i] !== '' ? parseInt(editForm.value.data[i]) : 0
          }
          else if (prop.editDataTemplate.dataType[i] === "Float") {
            editForm.value.data[i] = editForm.value.data[i] !== '' ? parseFloat(editForm.value.data[i]) : 0.0
          }
        }
      }
      //提交窗口数据给父组件
      emit("edit", editForm.value.data);
      //清空原窗口
      for(const item in editForm.value.item){
        editForm.value.data[editForm.value.item[item].dataName] = ''
      }
      //关闭窗口
      editFormVisible.value = false
    }
  })
}
function edit(row){
  console.log("edit", row)
  if(typeof prop.keyData === "string"){
    editForm.value.data[prop.keyData] = row[prop.keyData]
  }
  else{
    editForm.value.data = editObjKeyData(editForm.value.data, getObjKeyData(row, prop.keyData), prop.keyData)
  }
  for(const item in editForm.value.item){
    const i = editForm.value.item[item].dataName
    editForm.value.data[i] = row[i]
  }
  editFormVisible.value = true
}

function confirmDel(row){
  ElMessageBox.confirm(
      '你确定要删除' + row[prop.searchData] + '吗？数据被删除后无法恢复！',
      '注意',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        icon: markRaw(Delete),
        type: 'warning',
      }
  )
  .then(() => {
    del(row)
  })
}
function del(row){
  emit("del", row);
}

function excelToJson(e){
  let file = e.file // 文件信息
  if (!file) {
    // 没有文件
    return false
  } else if (!/\.(xls|xlsx)$/.test(file.name.toLowerCase())) {
    // 格式根据自己需求定义
    ElMessage.error('上传格式不正确，请上传xls或者xlsx格式')
    return false
  }

  const fileReader = new FileReader()
  fileReader.onload = (ev) => {
    try {
      const data = ev.target.result
      const workbook = XLSX.read(data, {
        type: 'binary' // 以字符编码的方式解析
      })
      const excelName = workbook.SheetNames[0] // 取第一张表
      const exl = XLSX.utils.sheet_to_json(workbook.Sheets[excelName]) // 生成json表格内容
      //console.log("excelToJson:", exl)
      // 将 JSON 数据上传给服务器
      let dataList = []
      for (const item in exl) {
        const data = exl[item]
        console.log("data:", data)
        for (const i in prop.addDataTemplate.dataType){
          if(data[i]) {
            if (prop.addDataTemplate.dataType[i] === "String") {
              data[i] = data[i].toString()
            }
            else if(prop.addDataTemplate.dataType[i] === "Int"){
              data[i] = data[i] !== '' ? parseInt(data[i]) : 0
            }
            else if(prop.addDataTemplate.dataType[i] === "Float"){
              data[i] = data[i] !== '' ? parseFloat(data[i]) : 0.0
            }
          }
          dataList.push(data)
        }
      }
      dataList = Array.from(new Set(dataList))
      emit("upload", dataList)
    } catch (error) {
      console.error("excelToJson:", error)
      return false
    }
  }
  fileReader.readAsBinaryString(file)
}

function downloadTemplate(){
  const workbook = XLSX.utils.book_new();
  const worksheet = XLSX.utils.json_to_sheet([addForm.value.data])
  XLSX.utils.book_append_sheet(workbook, worksheet, "sheet1")
  worksheet["!cols"] = new Array(addForm.value.dataNum).fill({ wch: 15 });
  XLSX.writeFileXLSX(workbook, "提交模版.xlsx")
}

function uploadDialogClosed(){
  myUploadForm.value.clearFiles()
}

function previewFile(_, val){
  const path = `${axios.defaults.baseURL}/${val}`
  window.open(path, '_blank');
}

const handleImgChange = (uploadFile) => {
  if (!/\.(jfif|pjpeg|jpeg|pjp|jpg|png|gif|bmp|webp|tif|tiff|svgz|svg)$/.test(uploadFile.name.toLowerCase())) {
    // 格式根据自己需求定义
    ElMessage.error('上传格式不正确，请上传支持的图片格式')
    myUploadImgForm.value.handleRemove(uploadFile)
    return false
  }
  imgList.value.push(uploadFile.raw)
}

const handleFileChange = (uploadFile) => {
  if (!/\.(txt|doc|dot|docx|xls|xlsx|ppt|pptx|pdf|md)$/.test(uploadFile.name.toLowerCase())) {
    // 格式根据自己需求定义
    ElMessage.error('上传格式不正确，请上传支持的文件格式')
    myUploadFileForm.value.handleRemove(uploadFile)
    return false
  }
  fileList.value.push(uploadFile.raw)
}

const uploadImgDialogOpen = () => {
  myUploadImgForm.value.clearFiles()
  imgList.value.length = 0  //清空图片数组
}

const uploadFileDialogOpen = () => {
  myUploadFileForm.value.clearFiles()
  fileList.value.length = 0  //清空附件数组
}

//上传图片或附件超限时调用此函数
const uploadExceed = () => {
  ElMessage.error("上传量已达上限！")
}

const handlePictureCardPreview = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url
  imgDialogVisible.value = true
}

const handleRemove = (f, list) => {
  const index = list.indexOf(f.raw)
  if(index !== -1){
    list.splice(index, 1)
  }
}

const handleSelectionChange = (val) => {
  multipleSelection = val
  emit("selectionChange", val)
}

</script>

<style scoped>
.single-column {
  display: grid;
  grid-template-columns: 1fr;
  gap: 10px;
}

.multi-column {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 10px;
}

.preview-img{
  max-height: 70vh;
  width: auto;
}

.img-dialog{
  width: auto;
  height: auto;
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>