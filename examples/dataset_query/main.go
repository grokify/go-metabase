package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grokify/gotilla/fmt/fmtutil"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/go-metabase/util"
	mo "github.com/grokify/oauth2more/metabase"
)

func main() {
	cfg := mo.InitConfig{
		LoadEnv:              true,
		EnvPath:              "ENV_PATH",
		EnvMetabaseBaseUrl:   "METABASE_BASE_URL",
		EnvMetabaseSessionId: "METABASE_SESSION_ID",
		EnvMetabaseUsername:  "METABASE_USERNAME",
		EnvMetabasePassword:  "METABASE_PASSWORD",
		TlsSkipVerify:        true}

	apiClient, _, err := util.NewApiClientEnv(cfg)
	if err != nil {
		log.Fatal(err)
	}

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

	if 1 == 0 {
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
		records, err := util.GetAllRecords(apiClient, opts)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("RECORDS [%v]\n", len(records.Rows))
	}

	fmt.Println("DONE")
}
