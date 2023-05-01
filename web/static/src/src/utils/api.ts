export function getApiBaseUrl(mode: string = 'http') {
  if (mode === 'http') {
    return 'http://localhost:8021'
  }
  if (mode === 'ws') {
    return 'ws://localhost:8021'
  }
  return 'http://localhost:8021'
}

export function query(url: string, params: Object | never = {}) {
  const queryUrl = new URL(getApiBaseUrl() + url)
  Object.entries(params).forEach(([key, value]) => {
    queryUrl.searchParams.set(key, value)
  })

  const queryString = queryUrl.toString()
  return queryString
}
