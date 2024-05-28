<template>
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
      placeholder="在当前页搜索"
      prefix-icon="Search"
      @input="searchChange"
      clearable
  />
</template>

<script setup>
//对外事件列表
import {ref} from "vue";

const emit = defineEmits(["add", "download", "upload", "search"]);

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

function searchChange(){
  emit("search", search.value)
}
</script>

<style scoped>
</style>