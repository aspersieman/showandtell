import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import Countdown from 'vue3-flip-countdown'
import router from './router'
import 'flowbite'

import './assets/base.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Countdown)

app.mount('#app')
