import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { ApiLogin, ApiLogout, Token, User } from '@/models/models'
import { query } from '@/utils/api'

export const useAuthenticationStore = defineStore('authenticationStore', () => {
  const user = ref<User | null>(null)
  const token = ref<Token | null>(null)
  const isLoggedIn = ref<boolean>(false)

  function init() {
    getToken()
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
      token.value = {
        accessToken: tokens?.accessToken,
        refreshToken: tokens?.refreshToken,
        expiresIn: tokens?.expiresIn,
        decoded: tokens?.decoded
      }
      user.value = {
        email: tokens?.decoded?.claims.email,
        name: tokens?.decoded?.claims.name,
        username: tokens?.decoded?.claims.preferred_username,
        realm_access: tokens?.decoded?.claims.realm_access,
        resource_access: tokens?.decoded?.claims.resource_access
      }

      isLoggedIn.value = true
    }
  }
  async function logout() {
    const body: ApiLogout = {
      refresh_token: token.value?.refreshToken
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
    const response = await fetch(query(`/api/v1/auth/token`))
    const data = await response.json()
    const tokens = data?.data
    token.value = {
      accessToken: tokens?.accessToken,
      refreshToken: tokens?.refreshToken,
      expiresIn: tokens?.expiresIn,
      decoded: tokens?.decoded
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
