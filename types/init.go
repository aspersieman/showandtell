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

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}
