package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	"github.com/grokify/go-metabase/metabase"
	mo "github.com/grokify/oauth2more/metabase"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	serverURL := os.Getenv("METABASE_BASE_URL")

	httpClient, _, err := mo.NewClientPassword(
		serverURL,
		os.Getenv("METABASE_USERNAME"),
		os.Getenv("METABASE_PASSWORD"),
		true)
	if err != nil {
		log.Fatal(err)
	}

	apiConfig := metabase.NewConfiguration()
	apiConfig.BasePath = serverURL
	apiConfig.HTTPClient = httpClient
	apiClient := metabase.NewAPIClient(apiConfig)

	opts := metabase.ListDatabasesOpts{
		IncludeTables: optional.NewBool(true)}

	info, resp, err := apiClient.DatabaseApi.ListDatabases(
		context.Background(), &opts)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("Status Code [%v]", resp.StatusCode))
	}

	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
