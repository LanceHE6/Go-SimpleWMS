<template>
  <el-container
    v-loading="loading"
  >
    <el-header
      style="height: 20px"
    >
      <el-button type="primary" icon="Plus" @click="add">
        添加
      </el-button>

      <el-button type="primary" icon="Download" @click="upload" plain>
        导入
      </el-button>


      <el-button type="primary" icon="Upload" @click="download" plain>
        导出
      </el-button>

      <el-input
          v-model="search"
          style="width: 240px; margin-left: 20px"
          placeholder="搜索"
          prefix-icon="Search"
          clearable
      />

    </el-header>

    <el-main>
      <el-table
          :data="filterTableData"
          :border="true"
          :stripe="true"
          height="60vh"
          style="width: 100%"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column
          v-for="item in tableColList"
          :property="item.property"
          :label="item.label"
          :width="item.width"
          :sortable="item.sortable"
        >

        </el-table-column>

        <el-table-column width="135" label="操作">
          <template #default="scope">
            <el-button
                size="small"
                type="success"
                @click="edit(scope.row)"
                plain
            >编辑</el-button>
            <el-button
                size="small"
                type="warning"
                @click="confirmDel(scope.row)"
                plain
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>

  <el-dialog
      v-model="editFormVisible"
      title="编辑"
      width="500"
      center
  >

    <el-form :model="editForm.data" :rules="editForm.rules" ref="myEditForm" label-position="top" status-icon>

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
      v-model="addFormVisible"
      title="添加"
      width="500"
      center
  >

    <el-form :model="addForm.data" :rules="addForm.rules" ref="myAddForm" label-position="top" status-icon>
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
      v-model="uploadFormVisible"
      title="批量导入"
      width="700"
      center
  >
    <el-upload
        ref="myUploadForm"
        action=""
        accept="application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
        :auto-upload="false"
        :http-request="excelToJson"
        limit="1"
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
</template>

<script setup>
import {computed, markRaw, ref, watch} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {Delete, UploadFilled} from "@element-plus/icons-vue";
import * as XLSX from "xlsx";

const prop = defineProps({
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
  }
});

//对外事件列表
const emit = defineEmits(["add", "download", "upload", "edit", "del"]);

//表格数据列表
const tableData = ref(prop.defaultData)
//编辑表单数据模版
const editForm = ref(prop.editDataTemplate)
//添加表单数据模版
const addForm = ref(prop.addDataTemplate)

//显示名表头
const tableHead = prop.tableColList.map( item => {
  return item.label
})

//编辑表单
const myEditForm = ref(null)
//编辑表单是否可见
const editFormVisible = ref(false)

//新增表单
const myAddForm = ref(null)
//新增表单是否可见
const addFormVisible = ref(false)

//上传窗口组件
const myUploadForm = ref(null)
//上传表单是否可见
const uploadFormVisible = ref(false)


//监听数据变化并实时更新表格
watch(() => prop.defaultData, (newValue) => {
  tableData.value = newValue;
});


//搜索栏文字
const search = ref('')

//筛选函数, 根据搜索框来做筛选
const filterTableData = computed(() =>
    tableData.value.filter(
        (data) =>
            !search.value ||
            data[prop.searchData].toLowerCase().includes(search.value.toLowerCase())
    )
)

function add(){
  addFormVisible.value = true
}

async function submitAddForm(form) {
  if (!form) return
  await form.validate((valid) => {
    if (valid) {
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

function submitUploadData(){
  myUploadForm.value.submit()
  uploadFormVisible.value = false
}

//提交编辑表单
async function submitEditForm(form){
  if (!form) return
  await form.validate((valid) => {
    if (valid) {
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
  editForm.value.data[editForm.value.key] = row[editForm.value.key]
  for(const item in editForm.value.item){
    editForm.value.data[editForm.value.item[item].dataName] = row[editForm.value.item[item].dataName]
  }
  editFormVisible.value = true
}

function confirmDel(row){
  ElMessageBox.confirm(
      '你确定要删除用户' + row.nickname + '吗？用户被删除后无法恢复！',
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
      console.log("excelToJson:", exl)
      // 将 JSON 数据上传给服务器
      let userList = []
      for (const item in exl) {
        const user = exl[item]
        user.phone = user.phone.toString()
        user.password = user.password.toString()
        userList.push(user)
      }
      userList = Array.from(new Set(userList))
      emit("upload", userList)
      // document.getElementsByName('file')[0].value = '' // 根据自己需求，可重置上传value为空，允许重复上传同一文件
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

</script>

<style scoped>
</style>