type SocketStore = {
  isConnected: boolean
  message: string
  reconnectError: boolean
  heartBeatInterval: number
  heartBeatTimer: number
}

type socketType = {
  $connect: () => void
}

interface SocketMessage {
  socket_id: number
  data: any
}

export type { SocketStore, socketType, SocketMessage }
