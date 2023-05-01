import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import 'flowbite'

import './assets/base.css'

const app = createApp(App)

app.use(createPinia())

import { useAuthenticationStore } from '@/stores/authentication'

const authenticationStore = useAuthenticationStore()
const requiresLoggedIn = ['schedules']

router.beforeEach(async (to) => {
  const toRoute = to.name ? String(to.name) : ''
  if (
    // make sure the user is authenticated
    !authenticationStore.isLoggedIn &&
    requiresLoggedIn.includes(toRoute) &&
    to.name !== 'login'
  ) {
    // redirect the user to the login page
    return { name: 'login' }
  }
})
app.use(router)

app.mount('#app')
