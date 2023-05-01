<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { useAuthenticationStore } from '@/stores/authentication'

const authenticationStore = useAuthenticationStore()
const router = useRouter()

const username = ref('')
const password = ref('')
const message = ref('')

const login = () => {
  message.value = ''
  if (username.value === '' || password.value === '') {
    alert('Email and password are required')
  } else {
    authenticationStore.login(username.value, password.value).then(() => {
      if (authenticationStore.isLoggedIn) {
        router.push('/')
      } else {
        message.value = 'Invalid email or password'
      }
    })
  }
}
</script>

<template>
  <div class="flex login-container flex-row">
    <section class="left-section flex h-full w-full">&nbsp;</section>
    <section class="login-section bg-gray-50/50 dark:bg-gray-900/50">
      <div
        class="flex-col items-center justify-center self-end px-6 py-8 mx-auto md:h-screen lg:py-0"
      >
        <div
          class="w-full bg-white rounded-lg shadow dark:border md:mt-4 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700"
        >
          <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1
              class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white"
            >
              Log in
            </h1>
            <div
              v-if="message != ''"
              class="p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400"
              role="alert"
            >
              {{ message }}
            </div>
            <div>
              <label
                for="email"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Your email</label
              >
              <input
                v-model="username"
                type="email"
                name="email"
                id="email"
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="name@company.com"
                required
              />
            </div>
            <div>
              <label
                for="password"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Password</label
              >
              <input
                v-model="password"
                type="password"
                name="password"
                id="password"
                placeholder="••••••••"
                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />
            </div>
            <button
              @click="login()"
              type="submit"
              class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Sign in
            </button>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.login-container {
  background: url('@/assets/core-values.jpg') no-repeat center center;
  background-size: cover;
}

.login-section {
  min-width: 25rem;
}

@media not all and (min-width: 640px) {
  .login-section {
    min-width: 100%;
  }

  .left-section {
    width: 0%;
  }
}
</style>
