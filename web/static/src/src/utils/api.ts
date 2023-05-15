import type { Token } from '@/models/models'

export function getApiBaseUrl(mode: string = 'http') {
  const protocol = location.protocol
  let url = `${protocol}//${import.meta.env.VITE_VIRTUAL_HOST}`
  if (mode === 'http') {
    url = `${protocol}//${import.meta.env.VITE_VIRTUAL_HOST}`
  }
  if (mode === 'ws' && protocol === 'https:') {
    url = `wss://${import.meta.env.VITE_VIRTUAL_HOST}`
  } else if (mode === 'ws' && protocol === 'http:') {
    url = `ws://${import.meta.env.VITE_VIRTUAL_HOST}`
  }
  return url
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
