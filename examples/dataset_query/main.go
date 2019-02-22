package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"

	"github.com/grokify/go-metabase/metabase"
	mbu "github.com/grokify/go-metabase/util"
)

func main() {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		panic(err)
	}

	apiClient, authResponse, err := mbu.NewApiClientPasswordWithSessionId(
		os.Getenv("METABASE_BASE_URL"),
		os.Getenv("METABASE_USERNAME"),
		os.Getenv("METABASE_PASSWORD"),
		os.Getenv("METABASE_SESSION_ID"))
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(authResponse)

	databaseId := int64(2)
	sourceTableId := int64(518)

	opts := metabase.DatasetQueryJsonQuery{
		Database: databaseId,
		Type:     "query",
		Query: metabase.DatasetQueryDsl{
			SourceTable: sourceTableId,
			Page:        metabase.DatasetQueryDslPage{Page: int64(1), Items: int64(2000)},
		},
		Constraints: metabase.DatasetQueryConstraints{MaxResults: 10000},
	}

	if 1 == 1 {
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
	} else {
		records, err := mbu.GetAllRecords(apiClient, opts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("RECORDS [%v]\n", len(records.Rows))
	}

	fmt.Println("DONE")
}
