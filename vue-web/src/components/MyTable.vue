<template>
  <el-container>
    <el-header
      style="height: 20px"
    >
      <el-button type="primary" icon="Plus" @click="add">
        添加
      </el-button>

      <el-button type="primary" icon="Download" @click="download" plain>
        导入
      </el-button>

      <el-button type="primary" icon="Upload" @click="upload" plain>
        导出
      </el-button>

      <el-input
          v-model="search"
          style="width: 240px; margin-left: 20px"
          placeholder="按昵称搜索"
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
</template>

<script setup>
import {computed, markRaw, ref, watch} from "vue";
import {ElMessageBox} from "element-plus";
import {Delete} from "@element-plus/icons-vue";

const prop = defineProps({
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
  }
});

//对外事件列表
const emit = defineEmits(["add", "download", "upload", "edit", "del"]);

//表格数据列表
const tableData= ref(prop.defaultData)

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
            data.nickname.toLowerCase().includes(search.value.toLowerCase())
    )
)

function add(){
  emit("add");
}
function download(){
  emit("download");
}
function upload(){
  emit("upload");
}
function edit(row){
  emit("edit", row);
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

</script>

<style scoped>
</style>