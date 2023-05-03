<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import ProfileAvatar from 'vue-profile-avatar'
import { useAuthenticationStore } from '@/stores/authentication'

const props = defineProps({
  username: String,
  email: String
})

const authenticationStore = useAuthenticationStore()
const router = useRouter()
const dropdown = ref(false)

const logout = () => {
  authenticationStore.logout().then(() => {
    router.push('/')
  })
}
const toggleDropdown = () => {
  dropdown.value = !dropdown.value
}
</script>

<template>
  <div class="relative">
    <ProfileAvatar
      @click="toggleDropdown()"
      class="w-10 h-10 rounded-full cursor-pointer"
      :username="props.username"
    >
    </ProfileAvatar>

    <div
      v-if="dropdown"
      class="absolute z-10 -left-32 bg-white divide-y divide-gray-100 rounded-lg shadow drop-shadow-md w-44 dark:bg-gray-700 dark:divide-gray-600"
    >
      <div class="px-4 py-3 text-sm text-gray-900 dark:text-white">
        <div class="truncate">{{ props.username }}</div>
        <div class="font-medium truncate">{{ props.email }}</div>
      </div>
      <ul class="py-2 text-sm text-gray-700 dark:text-gray-200" aria-labelledby="avatarButton">
        <li>
          <RouterLink
            to="/schedules"
            class="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
            >Schedules</RouterLink
          >
        </li>
      </ul>
      <div class="py-1">
        <a
          @click="logout()"
          href="#"
          class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 dark:text-gray-200 dark:hover:text-white"
          >Log out</a
        >
      </div>
    </div>
  </div>
</template>
