<template>
  <!--  根据路由展示页面-->
  <router-view/>
</template>

<script setup>
import Zoom from '@/utils/Zoom.js'
import {onMounted} from "vue";

onMounted(() => {
  new Zoom().init()
})


</script>
