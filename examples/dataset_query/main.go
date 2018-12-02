package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	"github.com/grokify/go-metabase/client"
	mo "github.com/grokify/oauth2more/metabase"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	serverURL := os.Getenv("METABASE_BASE_URL")

	var httpClient *http.Client

	if len(os.Getenv("METABASE_SESSION_ID")) > 0 {
		httpClient = mo.NewClientId(os.Getenv("METABASE_SESSION_ID"), true)
	} else {
		httpClient2, res, err := mo.NewClientPassword(
			serverURL,
			os.Getenv("METABASE_USERNAME"),
			os.Getenv("METABASE_PASSWORD"),
			true)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(res)
		httpClient = httpClient2
	}

	apiConfig := metabase.NewConfiguration()
	apiConfig.BasePath = serverURL
	apiConfig.HTTPClient = httpClient
	apiClient := metabase.NewAPIClient(apiConfig)

	databaseId := int64(2)
	sourceTableId := int64(518)

	opts := metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "query",
		Query: metabase.DatasetQueryDsl{
			SourceTable: sourceTableId,
			Page:        metabase.DatasetQueryDslPage{Page: int64(5), Items: int64(2000)},
		},
		Constraints: metabase.DatasetQueryConstraints{MaxResults: 10000},
	}

	info, resp, err := apiClient.DatasetApi.QueryDatabase(
		context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("Status Code [%v]", resp.StatusCode))
	}

	fmtutil.PrintJSON(info)
	fmtutil.PrintJSON(info.JsonQuery)
	fmt.Printf("ROWS [%v]\n", len(info.Data.Rows))
	if len(info.Data.Rows) > 0 {
		fmtutil.PrintJSON(info.Data.Rows[0])
	}
	fmt.Println("DONE")
}
