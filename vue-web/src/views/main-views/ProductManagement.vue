<!--货品管理-->
<template>
  <el-table
      ref="multipleTableRef"
      :data="filterTableData"
      :border="true"
      :stripe="true"
      :default-sort="{ prop: 'wid', order: 'descending' }"
      height="50vh"
      style="width: 100%"
  >
    <el-table-column type="selection" width="55" />

    <el-table-column property="wid" label="仓库ID" width="120" sortable/>

    <el-table-column property="name" label="仓库名" width="120" />

    <el-table-column property="addTime" label="添加时间" width="150" sortable/>

    <el-table-column property="comment" label="备注"/>

    <el-table-column align="right" width="135">
      <template #header>
        <el-input v-model="search" size="small" placeholder="仓库名搜索" />
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
        >编辑</el-button
        >
        <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
        >删除</el-button
        >
      </template>
    </el-table-column>

  </el-table>
</template>

<script setup>
import {computed, reactive, ref} from "vue";

//搜索栏文字
const search = ref('')
const filterTableData = computed(() =>
    tableData.filter(
        (data) =>
            !search.value ||
            data.name.toLowerCase().includes(search.value.toLowerCase())
    )
)

const tableData= [
  {
    wid: '000001',
    name: 'ex-w7',
    addTime: '2024-02-24',
    comment: 'No. 189, Grove St, Los Angeles',
  },
  {
    wid: '000002',
    name: 'w56',
    addTime: '2023-12-21',
    comment: 'No. 189, Grove St, Los Angeles',
  },
  {
    wid: '000003',
    name: 'w102',
    addTime: '2023-11-15',
    comment: 'No. 189, Grove St, Los Angeles',
  },
  {
    wid: '000006',
    name: 'ex-w2',
    addTime: '2024-03-02',
    comment: 'No. 189, Grove St, Los Angeles',
  },
  {
    wid: '000004',
    name: 'w65',
    addTime: '2024-03-14',
    comment: 'No. 189, Grove St, Los Angeles',
  },
  {
    wid: '000007',
    name: 'w2',
    addTime: '2024-03-15',
    comment: 'No. 189, Grove St, Los Angeles',
  },
  {
    wid: '000005',
    name: 'shang',
    addTime: '2024-02-16',
    comment: 'No. 189, Grove St, Los Angeles',
  },
]

</script>

<style scoped>

</style>