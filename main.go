package main

import (
	"embed"
	"flag"
	"fmt"

	cmd "bitbucket.org/envirovisionsolutions/showandtell/cmd"
	database "bitbucket.org/envirovisionsolutions/showandtell/database"
	types "bitbucket.org/envirovisionsolutions/showandtell/types"
	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
)

//go:embed web/static/src/dist/index.html
var EmbedIndex embed.FS

//go:embed web/static/src/dist/* web/static/src/dist/assets/*
var EmbedDirStatic embed.FS

//go:embed web/static/img/favicon.ico
var EmbedFavIcon embed.FS

//go:embed web/static/VERSION
var EmbedVersion embed.FS

// @title Show and Tell
// @version 1.0
// @description EVS Show and Tell collaboration system
// @contact.name Support
// @contact.email support@evsolutions.biz
// @host localhost:8021
// @BasePath /api/v1
func main() {
	var serve bool
	var migrate bool
	var seed bool
	var env string
	flag.BoolVar(&serve, "serve", false, "Serve the API as well as the API documentation")
	flag.BoolVar(&migrate, "migrate", false, "Migrate database")
	flag.BoolVar(&seed, "seed", false, "Seed database")
	flag.StringVar(&env, "env", "", "Environment. [create|print]")
	flag.Parse()
	if env != "" {
		if env == "create" {
			err := cmd.EnvironmentCreate()
			if err != nil {
				fmt.Printf("ERROR: Could not create .env file: %s\n", err)
			}
		} else if env == "print" {
			cmd.EnvironmentPrint()
		}
	} else if serve {
		cmd.EnvironmentValidate()
		utils.CmdContext = types.CmdContext{
			IndexFile: EmbedIndex,
			StaticDir: EmbedDirStatic,
			FavIcon:   EmbedFavIcon,
			Version:   EmbedVersion,
		}
		utils.Setup()
		cmd.Serve()
	} else if migrate {
		cmd.EnvironmentValidate()
		database.RunMigrations(true)
	} else if seed {
		cmd.EnvironmentValidate()
		database.ConnectDb()
		database.RunSeeder()
	} else {
		flag.PrintDefaults()
	}
}
