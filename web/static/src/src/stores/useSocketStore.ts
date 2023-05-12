import { defineStore } from 'pinia'
import { store } from '@/stores/sockets'
import main from '@/main'
import type { SocketStore } from '@/types/sockets'

export const useSocketStore = defineStore({
  id: 'socketStore',
  state: (): SocketStore => ({
    isConnected: false,
    message: '',
    reconnectError: false,
    heartBeatInterval: 50000,
    heartBeatTimer: 0
  }),
  actions: {
    SOCKET_ONOPEN(event: any) {
      console.log('Websocket connection successful ')
      main.config.globalProperties.$socket = event.currentTarget
      this.isConnected = true
      this.heartBeatTimer = window.setInterval(() => {
        const message = 'HeartBeat'
        this.isConnected &&
          main.config.globalProperties.$socket.sendObj({
            code: 200,
            msg: message
          })
      }, this.heartBeatInterval)
    },
    SOCKET_ONCLOSE(event: any) {
      this.isConnected = false
      window.clearInterval(this.heartBeatTimer)
      this.heartBeatTimer = 0
      console.log('Websocket connection close: ' + new Date())
      console.log(event)
    },
    SOCKET_ONERROR(event: any) {
      console.error(event)
    },
    SOCKET_ONMESSAGE(message: any) {
      this.message = message
    },
    SOCKET_RECONNECT(count: any) {
      console.info('Websocket connection reconnect...', count)
    },
    SOCKET_RECONNECT_ERROR() {
      this.reconnectError = true
    }
  }
})

// Need to be used outside the setup
export function useSocketStoreWithOut() {
  return useSocketStore(store)
}
