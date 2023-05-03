package utils

import (
	"strings"

	types "bitbucket.org/envirovisionsolutions/showandtell/types"
)

func GetAppVersion() types.Version {
	version, _ := CmdContext.StaticDir.ReadFile("web/static/VERSION")
	latest := strings.TrimSuffix(string(version), "\n")
	supported := make([]string, 0)
	v := types.Version{
		LatestVersion:     latest,
		SupportedVersions: supported,
	}
	return v
}
