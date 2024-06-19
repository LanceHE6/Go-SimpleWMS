<template>
  <el-table
      ref="myTable"
      :data="filterTableData"
      :border="true"
      :stripe="true"
      :height="height"
      @selection-change="handleSelectionChange"
      style="width: 100%"
  >
    <el-table-column type="selection" width="55" fixed />
    <el-table-column type="index" label="序号" align="center" header-align="center" width="55" fixed/>

    <el-table-column
        v-if="operations.edit || operations.del"
        align="center"
        header-align="center"
        width="135"
        label="操作"
        fixed
    >
      <template #default="scope">
        <el-button
            v-if="operations.edit"
            size="small"
            type="success"
            @click="edit(scope.row)"
            plain
        >编辑</el-button>
        <el-button
            v-if="operations.del"
            size="small"
            type="warning"
            @click="del(scope.row)"
            plain
        >删除</el-button>
      </template>
    </el-table-column>

    <el-table-column
        v-for="item in tableColList"
        align="center"
        header-align="center"
        :type="item.isExpand ? 'expand' : 'default'"
        :property="item.property"
        :label="item.label"
        :width="item.width"
        :sortable="item.sortable"
        :formatter="(row) => itemMapping(row, item.property, item)"
        :fixed="item.isFixed"
    >
      <template #default="scope" v-if="item.operable">
          <span class="blue-font-color" @click="item.operationEvent(scope.row[keyData], scope.row[item.property])">
            {{ itemMapping(scope.row, item.property, item) }}
          </span>
      </template>
      <template #default="scope" v-if="item.isImage">
        <div style="display: flex; align-items: center; justify-content: center">
          <el-image
              class="table-col-img"
              v-if="isArrNotEmpty(scope.row, item)"
              :src="`${axios.defaults.baseURL}/${getArrData(scope.row, item)[0].path}`"
              fit="cover"
              :preview-src-list="getArrData(scope.row, item).map(imgObj => axios.defaults.baseURL + '/' + imgObj.path)"
              preview-teleported
          >
            <template #error>
              <div
                  class="error-image-slot"
              >
                <el-icon><Picture /></el-icon>
              </div>
            </template>
          </el-image>

          <div
              v-else
              :class="['error-image-slot', operations.uploadImg ? 'is_upload_img' : '']"
              @click="uploadImg(getObjKeyData(scope.row, keyData))">
            <el-icon v-if="operations.uploadImg"><Plus /></el-icon>
            <el-icon v-else><Picture /></el-icon>
          </div>

          <el-button
              v-if="isArrNotEmpty(scope.row, item) && operations.uploadImg"
              type="success"
              icon="Edit"
              @click="uploadImg(getObjKeyData(scope.row, keyData))"
              circle
              plain
              style="margin-left: 10px"
          />
        </div>
      </template>

      <template #default="scope" v-if="item.isInput">
        <div style="display: flex; align-items: center">
          <el-input
              v-if="item.type === 'number'"
              v-model.number="scope.row[item.property]"
              :type="item.type"
          />
          <el-input
              v-else
              v-model="scope.row[item.property]"
              :type="item.type"
          />
        </div>
      </template>

      <template #default="scope" v-if="item.isExpand">
        <div class="child_table_div">
          <el-table
              :data="scope.row[item.property]"
              :stripe="true"
              max-height="30vh"
          >
            <el-table-column
                v-for="child_item in item.children"
                align="center"
                header-align="center"
                :label="child_item.label"
                :prop="child_item.property"
                :width="child_item.width"
                :sortable="child_item.sortable"
                :formatter="(row) => itemMapping(row, child_item.property, child_item)"
                :fixed="child_item.isFixed"
            >
              <template #default="prop" v-if="child_item.isImage">
                <div style="display: flex; align-items: center; justify-content: center">
                  <el-image
                      class="table-col-img"
                      v-if="isArrNotEmpty(prop.row, child_item)"
                      :src="`${axios.defaults.baseURL}/${getArrData(prop.row,child_item)[0].path}`"
                      fit="cover"
                      :preview-src-list="getArrData(prop.row, child_item).map(imgObj => axios.defaults.baseURL + '/' + imgObj.path)"
                      preview-teleported
                  >
                    <template #error>
                      <div
                          class="error-image-slot"
                      >
                        <el-icon><Picture /></el-icon>
                      </div>
                    </template>
                  </el-image>
                  <div
                      v-else
                      class="error-image-slot"
                  >
                    <el-icon><Picture /></el-icon>
                  </div>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup>

import axios from "axios";
import {Plus, Picture} from "@element-plus/icons-vue";
import {getObjKeyData} from "@/utils/objectUtil.js";
import {computed, ref, watch} from "vue";

const prop = defineProps({
  large:{
    type: Boolean,
    default: () => false,
    description: '是否启用大数据表格(由后端分页的数据)'
  },
  keyData:{
    type: [String, Array],
    default: () => undefined,
    description: '数据主键, 如有多层则用列表封装, 格式如: ["key", "p1", "p2", ...]'
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
  height:{
    type: String,
    default: () => '60vh',
    description: '表格高度'
  },
  search:{
    type: String,
    default: () => '',
    description: '搜索框文字'
  }
});

//对外事件列表
const emit = defineEmits([
  "edit", "delete", "uploadImg", "selectionChange", "operation"
]);

//表格
const myTable = ref(null)
//清空表格选择
const clearSelection = () => myTable.value.clearSelection()

//暴露函数，可供父组件调用
defineExpose({
  clearSelection
});

//表格数据列表
const tableData = ref(prop.defaultData)
//监听数据变化并实时更新表格
watch(() => prop.defaultData, (newValue) => {
  tableData.value = newValue;
});


//筛选函数, 根据搜索框来做筛选(仅对前端分页生效)
const filterTableData = computed(() => {
  // 如果prop.searchData为空字符串，则不进行筛选，直接返回所有数据
  if (prop.searchData === "") {
    return tableData.value
  }

  // 否则，进行筛选
  return tableData.value.filter((data) => {
    // 如果prop.search为空或者prop.large为真，则包含该数据
    if (!prop.search || prop.large) {
      return true
    }

    // 进行正常的搜索比较
    return data[prop.searchData] && data[prop.searchData].toLowerCase().includes(prop.search.toLowerCase());
  })
})

// 数据显示转换(映射)
function itemMapping(row, property, item){
  let currentRow = row
  let currentItem = item
  let currentProp = property
  //多层对象, 一层层解开
  while(currentItem.isParent){
    currentRow = currentRow[currentItem.property]
    currentItem = currentItem['child']
    currentProp = currentItem.property
  }
  //外键映射
  if(currentItem.isFK){
    //从外键map中获取对应的外键表
    const fkList = prop.showFKMap.get(currentItem.FKData.property)
    for(const item2 of fkList){
      // 映射
      if(currentRow[currentProp] === item2[currentItem.FKData.property]){
        return item2[currentItem.FKData.label]
      }
    }
  }
  //普通映射
  if(currentItem.isMapping){
    //从映射表获取映射对象
    for(const j of currentItem.mappingList){
      const item2 = j
      // 映射
      if(currentRow[currentProp] === item2.value){
        return item2.label
      }
    }
  }
  //日期转换
  if(currentItem.isDateFormat){
    currentRow[currentProp] = new Date(currentRow[currentProp]).toLocaleString()
  }
  //不需要映射或者没有匹配的映射则返回原值
  return currentRow[currentItem.property]
}

//判断数据数组内是否有数据
const isArrNotEmpty = (row, item) =>{
  if (row === null || typeof row !== 'object' || Object.keys(row).length === 0) {
    return false;
  }

  //嵌套对象则逐层解开
  while (item && ('isParent' in item) && item.isParent){
    if(!(item.property in row)){
      return false
    }
    row = row[item.property]
    item = item.child
  }

  return Array.isArray(row[item.property]) && row[item.property].length > 0
}

const getArrData = (row, item) =>{
  if (row === null || typeof row !== 'object' || Object.keys(row).length === 0) {
    return [];
  }

  //嵌套对象则逐层解开
  while (item && ('isParent' in item) && item.isParent){
    if(!(item.property in row)){
      return false
    }
    row = row[item.property]
    item = item.child
  }

  return row[item.property]
}

const handleSelectionChange = (val) => {
  emit("selectionChange", val)
}

const edit = (row) =>{
  emit("edit", row)
}

const del = (row) =>{
  emit("delete", row)
}

const uploadImg = (id) =>{
  emit("uploadImg", id)
}
</script>

<style scoped>
.table-col-img{
  height: 50px;
  width: 50px;
}

.error-image-slot{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 50px;
  height: 50px;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 15px;
}
.blue-font-color {
  color: #409eff;
  border-bottom: 1px solid #409eff;
  cursor: pointer;
}
.is_upload_img{
  cursor: pointer;
}
.child_table_div{
  padding: 10px;
}
</style>