package types

import (
	"embed"
)

// Other
// swagger:model Version
type Version struct {
	// LatestVersion
	// example: 2.0.0
	LatestVersion string `json:"latest_version"`

	// SupportedVersions
	// example: [1.8.0, 1.9.0]
	SupportedVersions []string `json:"supported_versions"`
}

// swagger:model ApiResponse
type ApiResponse struct {
	// Data
	// example: data
	Data interface{} `json:"data"`
}

type CmdContext struct {
	IndexFile embed.FS
	StaticDir embed.FS
	FavIcon   embed.FS
	Version   embed.FS
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenMethod struct {
	Name string `json:"name"`
	Hash int    `json:"hash"`
}
type AccessTokenHeader struct {
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	Typ string `json:"typ"`
}
type AccessTokenRealmAccess struct {
	Roles []string `json:"roles"`
}
type AccessTokenResourceAccess struct {
	Resources [][]string `json:"resources"`
}
type AccessTokenClaims struct {
	Acr               string                    `json:"acr"`
	AllowedOrigins    []string                  `json:"allowed-origins"`
	Aud               []string                  `json:"aud"`
	Azp               string                    `json:"azp"`
	Email             string                    `json:"email"`
	EmailVerified     bool                      `json:"email_verified"`
	Exp               int64                     `json:"exp"`
	FamilyName        string                    `json:"family_name"`
	GivenName         string                    `json:"given_name"`
	Iat               int64                     `json:"iat"`
	Iss               string                    `json:"iss"`
	Name              string                    `json:"name"`
	PreferredUsername string                    `json:"preferred_username"`
	RealmAccess       AccessTokenRealmAccess    `json:"realm_access"`
	ResourceAccess    AccessTokenResourceAccess `json:"resource_access"`
	Scope             string                    `json:"scope"`
	SessionState      string                    `json:"session_state"`
	Sub               string                    `json:"sub"`
	Typ               string                    `json:"typ"`
}
type DecodedAccessToken struct {
	Raw       string            `json:"raw"`
	Method    AccessTokenMethod `json:"method"`
	Header    AccessTokenHeader `json:"header"`
	Claims    AccessTokenClaims `json:"claims"`
	Singature string            `json:"signature"`
	Valid     bool              `json:"valid"`
}
type LoginResponse struct {
	AccessToken  string             `json:"accessToken"`
	RefreshToken string             `json:"refreshToken"`
	ExpiresIn    int                `json:"expiresIn"`
	Decoded      DecodedAccessToken `json:"decoded"`
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
