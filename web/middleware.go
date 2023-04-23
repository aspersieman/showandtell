package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
	types "bitbucket.org/envirovisionsolutions/showandtell/types"
)

type KeycloakMiddleware struct {
	keycloak *auth.Keycloak
}

func NewMiddleware(keycloak *auth.Keycloak) *KeycloakMiddleware {
	return &KeycloakMiddleware{keycloak: keycloak}
}

func ExtractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func (auth *KeycloakMiddleware) VerifyToken(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.ApiResponse{Data: "Authorization header missing"})
	}

	token = ExtractBearerToken(token)

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.ApiResponse{Data: "Bearer Token missing"})
	}

	// call Keycloak API to verify the access token
	result, err := auth.keycloak.Gocloak.RetrospectToken(context.Background(), token, auth.keycloak.ClientId, auth.keycloak.ClientSecret, auth.keycloak.Realm)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.ApiResponse{Data: fmt.Sprintf("Invalid or malformed token: %s", err.Error())})
	}

	jwt, _, err := auth.keycloak.Gocloak.DecodeAccessToken(context.Background(), token, auth.keycloak.Realm)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.ApiResponse{Data: fmt.Sprintf("Invalid or malformed token: %s", err.Error())})
	}

	jwtj, _ := json.Marshal(jwt)
	log.Printf("token: %v\n", string(jwtj))

	// check if the token isn't expired and valid
	if !*result.Active {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.ApiResponse{Data: "Invalid or expired Token"})
	}
	return ctx.Next()
}
