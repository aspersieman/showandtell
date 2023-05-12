<script setup lang="ts">
import { RouterView } from 'vue-router'
import { onMounted, onBeforeUnmount } from 'vue'
import { initFlowbite } from 'flowbite'
import { getCurrentInstance } from 'vue'
import type { ComponentInternalInstance } from 'vue'

import type { SocketMessage } from '@/models/models'
import { useAuthenticationStore } from '@/stores/authentication'
import { useScheduleStore } from '@/stores/schedules'
import NavBar from './components/NavBar.vue'
import FooterBar from './components/FooterBar.vue'

const scheduleStore = useScheduleStore()
const authenticationStore = useAuthenticationStore()

const { proxy } = getCurrentInstance() as ComponentInternalInstance

const listenSocket = () => {
  proxy.$socket.onmessage = (event) => {
    let eventData: SocketMessage = {
      socket_id: 0,
      data: {}
    }
    try {
      eventData = JSON.parse(event.data)
    } catch (e) {
      console.log(e)
    }

    console.log('Message from server ', eventData)
    /*
    console.log('Message from server ', socketId, event.data, isRunning.value)
    if (eventData.socket_id == socketId && !isRunning.value) {
      time.value = eventData?.data?.time
      progress.value = eventData?.data?.progress
    }
    */
  }
}
const initSockets = () => {
  setTimeout(() => {
    if (proxy == null) return
    proxy.$connect()
    listenSocket()
  }, 100)
}

onMounted(() => {
  initSockets()
  initFlowbite()
  authenticationStore.init()
  scheduleStore.getSchedulesHome()
})
onBeforeUnmount(() => {
  proxy?.$disconnect()
})
</script>

<template>
  <div class="flex flex-col h-screen justify-between">
    <NavBar />
    <RouterView v-slot="{ Component }">
      <Transition name="fade" mode="out-in">
        <component :is="Component" />
      </Transition>
    </RouterView>
    <FooterBar />
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
