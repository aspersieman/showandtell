package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
)

func EnvironmentCreate() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	file := filepath.Join(cwd, ".env")
	if _, err := os.Stat(file); err != nil {
		fmt.Printf("Creating new environment file '%s'\n", file)
		f, err := os.Create(file)
		if err != nil {
			return err
		}

		content := `
APP_NAME="Show and Tell"
APP_PORT=8020
APP_ENV=production
APP_WEB_CACHE_LENGTH_MINUTES=2

VIRTUAL_HOST=showandtell.local
NGINX_PORT=80

KEYCLOAK_BASE_URL=https://auth.evsolutions.co.za/auth/
KEYCLOAK_REALM=evs
KEYCLOAK_CLIENT_ID=showandtell
KEYCLOAK_CLIENT_SECRET=
KEYCLOAK_USER_ROLE=showandtell-user
KEYCLOAK_ADMIN_GROUPS="/Administrator,/EVS Management"`
		_, err = f.WriteString(content + "\n")
		if err != nil {
			return err
		}
		f.Sync()
	} else {
		fmt.Printf("Environment file '%s' already exists\n", file)
	}
	return nil
}

func EnvironmentPrint() {
	appEnv := utils.EnvGet("APP_ENV", "production")
	fmt.Printf("INFO: Application Environment: %s\n", appEnv)
}

func EnvironmentValidate() {
	if _, err := os.Stat(".env"); err != nil {
		fmt.Println("ERROR: '.env' file not found")
		os.Exit(1)
	}
}
