<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
import { initFlowbite } from 'flowbite'

import NavBar from './components/NavBar.vue'
import FooterBar from './components/FooterBar.vue'

onMounted(() => {
  initFlowbite()
})
</script>

<template>
  <div class="flex flex-col h-screen justify-between">
    <NavBar />
    <RouterView />
    <FooterBar />
  </div>
</template>

<style scoped></style>
