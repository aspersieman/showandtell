<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
import { initFlowbite } from 'flowbite'

import NavBar from './components/NavBar.vue'

onMounted(() => {
  initFlowbite()
})
</script>

<template>
  <NavBar />

  <RouterView />
</template>

<style scoped></style>
