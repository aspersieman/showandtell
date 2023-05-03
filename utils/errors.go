package utils

import (
	"encoding/json"
	"log"
	"net/http"

	types "bitbucket.org/envirovisionsolutions/showandtell/types"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("ERROR: %e", err)
	}
}

func CheckHttpErr(err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		log.Printf("ERROR: Cannot encode response. %s", err.Error())
		http.Error(w, http.StatusText(statusCode), statusCode)
		err := json.NewEncoder(w).Encode(types.ApiResponse{
			Data: http.StatusText(statusCode),
		})
		CheckErr(err)
	}
}
