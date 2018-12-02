package main

import (
	"context"
	"fmt"
	"log"
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
	/*
		httpClient, res, err := mo.NewClientPassword(
			serverURL,
			os.Getenv("METABASE_USERNAME"),
			os.Getenv("METABASE_PASSWORD"),
			true)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(res)
	*/
	id := "831ec6b1-4680-4c60-81e8-d1a6b77201b1"
	httpClient := mo.NewClientId(id, true)

	apiConfig := metabase.NewConfiguration()
	apiConfig.BasePath = serverURL
	apiConfig.HTTPClient = httpClient
	apiClient := metabase.NewAPIClient(apiConfig)

	databaseId := int32(2)
	sourceTableId := int32(518)

	opts := metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "query",
		Query: metabase.DatasetQueryDsl{
			SourceTable: sourceTableId}}

	info, resp, err := apiClient.DatasetApi.QueryDatabase(
		context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("Status Code [%v]", resp.StatusCode))
	}

	info.Constraints.MaxResults = int64(12345)
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}

/*

2018/12/01 16:07:19 json: cannot unmarshal array into Go struct field DatasetQueryResultsMetadata.columns of type metabase.DatasetQueryResultsMetadataColumn


*/
