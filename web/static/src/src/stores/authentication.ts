import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { ApiLogin, ApiLogout, Token, User } from '@/models/models'
import { getTokenLocal, query } from '@/utils/api'

export const useAuthenticationStore = defineStore('authenticationStore', () => {
  const user = ref<User | null>(null)
  const token = ref<Token | null>(null)
  const isLoggedIn = ref<boolean>(false)

  function init() {
    getToken()
  }

  async function setTokens(tokenData: Token) {
    const tokenRecord = {
      accessToken: tokenData?.accessToken,
      refreshToken: tokenData?.refreshToken,
      expiresIn: tokenData?.expiresIn,
      decoded: tokenData?.decoded
    }
    localStorage.setItem('token', JSON.stringify(tokenRecord))
    token.value = tokenRecord
    const userRecord = {
      email: tokenData?.decoded?.claims.email,
      name: tokenData?.decoded?.claims.name,
      username: tokenData?.decoded?.claims.preferred_username,
      realm_access: tokenData?.decoded?.claims.realm_access,
      resource_access: tokenData?.decoded?.claims.resource_access
    }
    user.value = userRecord
    localStorage.setItem('user', JSON.stringify(userRecord))

    isLoggedIn.value = true
  }

  async function login(username: string, password: string) {
    const body: ApiLogin = {
      username,
      password
    }
    const response = await fetch(query('/api/v1/auth/login', {}), {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    })
    const data = await response.json()
    const tokens = data?.data
    if (response.status === 200) {
      setTokens(tokens)
    }
  }

  async function logout() {
    const t: Token = getTokenLocal()
    const body: ApiLogout = {
      refresh_token: t?.refreshToken
    }
    await fetch(query('/api/v1/auth/logout', {}), {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    })
    token.value = null
    user.value = null

    isLoggedIn.value = false
  }

  async function getToken() {
    const t: Token = getTokenLocal()
    const body: ApiLogout = {
      refresh_token: t?.refreshToken
    }
    const response = await fetch(query(`/api/v1/auth/token`), {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    })
    const data = await response.json()
    const tokens = data?.data
    if (response.status === 200) {
      setTokens(tokens)
    }
  }

  return {
    init,
    login,
    logout,
    user,
    token,
    isLoggedIn,
    getToken
  }
})
