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
  <NavBar />
  <RouterView />
  <FooterBar />
</template>

<style scoped></style>
