package auth

import (
	"github.com/Nerzal/gocloak/v13"

	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
)

type Keycloak struct {
	Gocloak      *gocloak.GoCloak
	ClientId     string
	ClientSecret string
	Realm        string
}

func NewKeycloak() *Keycloak {
	baseUrl := utils.EnvGet("KEYCLOAK_BASE_URL", "")
	clientId := utils.EnvGet("KEYCLOAK_CLIENT_ID", "")
	clientSecret := utils.EnvGet("KEYCLOAK_CLIENT_SECRET", "")
	realm := utils.EnvGet("KEYCLOAK_REALM", "")
	return &Keycloak{
		Gocloak:      gocloak.NewClient(baseUrl),
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Realm:        realm,
	}
}
