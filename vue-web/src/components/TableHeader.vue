<template>
  <div class="my-table-header">
    <el-button v-if="(operations.add || hasSubmitPage) && permission >= 2" type="primary" icon="Plus" @click="add">
      添加
    </el-button>

    <el-button v-if="operations.upload" type="primary" icon="Upload" @click="upload" plain>
      导入
    </el-button>


    <el-button v-if="operations.download" type="primary" icon="Download" @click="download" plain>
      导出
    </el-button>

    <el-button v-if="operations.print" type="warning" icon="Printer" @click="print" plain>
      打印
    </el-button>

    <span class="my-table-header" v-if="!large">
      <el-input
          v-model="search"
          style="width: 240px; margin-left: 20px"
          placeholder="在当前页搜索"
          prefix-icon="Search"
          @input="searchChange"
          clearable
      />
    </span>
    <span class="my-table-header" v-else>
      <el-input
          v-model="search"
          style="width: 240px; margin-left: 20px; margin-right: 10px;"
          placeholder="关键字搜索"
          clearable
      />
      <el-button type="success" icon="Search" @click="searchChange">
        搜索
      </el-button>
    </span>

    <el-button type="primary" icon="Refresh" @click="refresh" plain style="margin-left: 10px">
      刷新
    </el-button>
  </div>
</template>

<script setup>
//对外事件列表
import {ref} from "vue";

const prop = defineProps({
  operations:{
    type: Object,
    default: () => null,
    description: '表格支持的操作, 包含增删查改以及导入导出等等(通常无需手动配置)'
  },
  hasSubmitPage:{
    type: Boolean,
    default: () => false,
    description: '是否需要单独的提交页面(由父组件传入, 无需手动配置)'
  },
  large:{
    type: Boolean,
    default: () => false,
    description: '是否为后端分页(由父组件传入, 无需手动配置)'
  },
  permission:{
    type: Number,
    default: () => 0,
    description: '用户权限等级'
  }
})

const emit = defineEmits(["add", "download", "upload", "search", "print", "refresh"]);

//搜索栏文字
const search = ref('')

function add(){
  emit("add");
}

function download(){
  emit("download");
}

function upload(){
  emit("upload");
}

function print(){
  emit("print");
}

function searchChange(){
  emit("search", search.value)
}

function refresh(){
  emit("refresh")
}
</script>

<style scoped>
.my-table-header {
  display: flex;

  /* 防止flex子项换行 */
  flex-wrap: nowrap; /* 仅当使用flex布局时 */
}
</style>