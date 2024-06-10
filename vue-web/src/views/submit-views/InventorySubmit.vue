<template>
  <p>Debug Mode: {{ route.query.debug === 'true' ? 'Yes' : 'No' }}</p>
</template>

<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()
const debug = route.query.debug

</script>

<style scoped>

</style>