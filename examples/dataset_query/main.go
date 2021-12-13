package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-metabase/metabase"
	mbu "github.com/grokify/go-metabase/metabaseutil"
)

func main() {
	appCfg := mbu.AppConfig{}
	_, err := flags.Parse(&appCfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient, authResponse, err := appCfg.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(authResponse)

	databaseId := int64(3)
	sourceTableId := int64(656)

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
