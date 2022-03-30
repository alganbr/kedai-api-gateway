package main

import (
	"github.com/alganbr/kedai-api-gateway/cmd"
	"os"
)

// @title           Kedai API Gateway
// @version         1.0
// @description     This is Kedai API Gateway.

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
