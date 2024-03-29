<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useAuthenticationStore } from '@/stores/authentication'
import UserAvatar from '@/components/UserAvatar.vue'

const authenticationStore = useAuthenticationStore()
</script>

<template>
  <nav class="bg-white border-gray-200 dark:bg-gray-900">
    <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
      <RouterLink class="flex items-center" to="/">
        <img src="@/assets/show-and-tell.png" class="h-8 mr-3" alt="Show and Tell" />
        <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white"
          >Show & Tell</span
        >
      </RouterLink>
      <div class="flex md:order-2">
        <RouterLink
          v-if="!authenticationStore.isLoggedIn"
          class="text-white mx-1 bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-4 py-2 text-center mr-3 md:mr-0 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
          to="/login"
          >Login
        </RouterLink>
        <user-avatar
          v-else
          :username="authenticationStore.user?.name"
          :email="authenticationStore.user?.email"
        />
        <button
          data-collapse-toggle="navbar-cta"
          type="button"
          class="inline-flex items-center p-2 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
          aria-controls="navbar-cta"
          aria-expanded="false"
        >
          <span class="sr-only">Open main menu</span>
          <svg
            class="w-6 h-6"
            aria-hidden="true"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
              clip-rule="evenodd"
            ></path>
          </svg>
        </button>
      </div>
      <div
        class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1"
        id="navbar-cta"
      >
        <ul
          class="flex flex-col font-medium p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
        >
          <li>
            <RouterLink
              class="nav-link block py-2 pl-3 pr-4 text-gray-900 hover:bg-gray-100 md:hover:bg-transparent md:hover:text-primary-700 md:p-0 md:dark:hover:text-primary-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700"
              to="/"
              >Home
            </RouterLink>
          </li>
          <li>
            <a
              href="https://evsolutions.biz"
              target="_blank"
              class="nav-link block py-2 pl-3 pr-4 text-gray-900 hover:bg-gray-100 md:hover:bg-transparent md:hover:text-primary-700 md:p-0 md:dark:hover:text-primary-500 dark:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700"
              >Website</a
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scope>
.time-picker {
  max-width: 6rem;
}

.nav-link.router-link-active {
  @apply text-primary-700 border-b-2 border-primary-700;
}
</style>
