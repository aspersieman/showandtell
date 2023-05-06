interface ApiLogin {
  username: string
  password: string
}
interface ApiLogout {
  refresh_token: string | undefined
}

interface AccessTokenMethod {
  name: string
  hash: number
}
interface AccessTokenHeader {
  alg: string
  kid: string
  typ: string
}
interface AccessTokenRealmAccess {
  roles: string[]
}
interface AccessTokenResourceAccess {
  resources: string[][]
}
interface AccessTokenClaims {
  acr: string
  allowed_origins: string[]
  aud: string[]
  azp: string
  email: string
  email_verified: boolean
  exp: number
  family_name: string
  given_name: string
  iat: number
  iss: string
  name: string
  preferred_username: string
  realm_access: AccessTokenRealmAccess
  resource_access: AccessTokenResourceAccess
  scope: string
  session_state: string
  sub: string
  typ: string
}
interface DecodedAccessToken {
  raw: string
  method: AccessTokenMethod
  header: AccessTokenHeader
  claims: AccessTokenClaims
  singature: string
  valid: boolean
}
interface Token {
  accessToken: string
  refreshToken: string
  expiresIn: number
  decoded: DecodedAccessToken
}
interface User {
  email: string
  name: string
  username: string
  realm_access: AccessTokenRealmAccess
  resource_access: AccessTokenResourceAccess
}

interface ApiSpeaker {
  id: number
  created_at: Date
  updated_at: Date
  deleted_at: Date | null
  name: string
  topic: string
  description: string
  audience: string
  order: number
  minutes: number
  schedule_id: number
  private: boolean
}

interface SpeakerAdd {
  name: string
  topic: string
  description: string
  audience: string
  order: number
  minutes: number
  schedule_id: number
  private: boolean
}

interface SpeakerUpdate {
  id: number
  name: string
  topic: string
  description: string
  audience: string
  order: number
  minutes: number
  schedule_id: number
  private: boolean
}

interface ApiSchedule {
  id: number
  created_at: Date
  updated_at: Date
  deleted_at: Date | null
  title: string
  description: string
  start_date_time: Date
  end_date_time: Date
  speakers: ApiSpeaker[]
}

interface ScheduleAdd {
  title: string
  description: string
  start_date_time: Date
  end_date_time: Date
  speakers: SpeakerAdd[] | null
}

interface ScheduleUpdate {
  id: number
  title: string
  description: string
  start_date_time: Date
  end_date_time: Date
  speakers: SpeakerAdd[] | null
}

interface ScheduleSearchParams {
  f?: string // start_date_time From
  t?: string // start_date_time To
}

interface SocketMessage {
  socket_id: number
  data: any
}

export type {
  ApiSpeaker,
  SpeakerAdd,
  SpeakerUpdate,
  ApiSchedule,
  ScheduleAdd,
  ScheduleUpdate,
  ScheduleSearchParams,
  ApiLogin,
  ApiLogout,
  Token,
  User,
  SocketMessage
}
