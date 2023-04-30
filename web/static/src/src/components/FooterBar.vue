<script setup lang="ts">
const year = new Date().getFullYear()
</script>
<template>
  <footer class="bg-white rounded-lg shadow m-4 dark:bg-gray-800">
    <div class="w-full mx-auto max-w-screen-xl p-4 md:flex md:items-center md:justify-between">
      <div class="whitespace-nowrap">
        <span class="text-sm text-gray-500 sm:text-center dark:text-gray-400">
          <img src=@/assets/icon.png alt="EVSolutions" class="h-10 float-left" /> © {{ year }}
          <a href="https://evsolutions.biz/" class="hover:underline"> EVSolutions™ </a>
        </span>
      </div>
      <ul
        class="flex flex-wrap items-center justify-end mt-3 text-sm font-medium text-gray-500 dark:text-gray-400 sm:mt-0">
        <li>
          <a href="/swagger" target="_blank" class="mr-4 hover:underline md:mr-6">Documentation</a>
        </li>
      </ul>
    </div>
  </footer>
</template>
