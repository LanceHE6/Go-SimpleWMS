<template>
  <el-card
      class="card"
  >
    <template #header>
      <div class="card-header">
        <el-text><b>{{'出入库分析'}}</b></el-text>

      </div>
    </template>
  </el-card>
</template>

<script setup>

</script>

<style scoped>
.card{
  width: 695px;
  height: 340px;
  margin-bottom: 10px;
  margin-right: 15px;
  --el-card-padding: 10px
}
.card-header{
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>