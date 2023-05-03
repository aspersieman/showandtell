import type { Token } from '@/models/models'

export function getApiBaseUrl(mode: string = 'http') {
  console.log('getApiBaseUrl: ', import.meta.env)
  const protocol = location.protocol
  if (mode === 'http') {
    return `${protocol}//${import.meta.env.VITE_VIRTUAL_HOST}`
  }
  if (mode === 'ws') {
    return `ws://${import.meta.env.VITE_VIRTUAL_HOST}`
  }
  return `${protocol}//${import.meta.env.VITE_VIRTUAL_HOST}`
}

export function query(url: string, params: Object | never = {}) {
  const queryUrl = new URL(getApiBaseUrl() + url)
  Object.entries(params).forEach(([key, value]) => {
    queryUrl.searchParams.set(key, value)
  })

  const queryString = queryUrl.toString()
  return queryString
}

export function getTokenLocal(): Token {
  return JSON.parse(localStorage.getItem('token') || '{}')
}

export function headers(): HeadersInit {
  const token = getTokenLocal()
  const headers: HeadersInit = new Headers()
  if (token?.accessToken) {
    headers.set('Content-Type', 'application/json')
    headers.set('Accept', 'application/json')
    headers.set('Authorization', `Bearer ${token?.accessToken}`)
  } else {
    headers.set('Content-Type', 'application/json')
    headers.set('Accept', 'application/json')
  }
  return headers
}
