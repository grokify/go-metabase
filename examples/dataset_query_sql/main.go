package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/type/stringsutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-metabase/metabaseutil"
	mo "github.com/grokify/oauth2more/metabase"
)

type Options struct {
	Env string `short:"e" long:"envfile" description:"Env File" required:"false"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	read, err := config.LoadDotEnvSkipEmptyInfo(
		opts.Env, os.Getenv("ENV_PATH"), ".env")
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(read)

	mbCfg := mo.Config{
		BaseUrl:       os.Getenv("METABASE_BASE_URL"),
		Username:      os.Getenv("METABASE_USERNAME"),
		Password:      os.Getenv("METABASE_PASSWORD"),
		SessionId:     os.Getenv("METABASE_SESSION_ID"),
		TlsSkipVerify: stringsutil.ToBool(os.Getenv("METABASE_TLS_SKIP_VERIFY"))}

	sqlInfo := metabaseutil.SQLInfo{}
	err = json.Unmarshal([]byte(os.Getenv("METABASE_QUERY_SQL_INFO")), &sqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(sqlInfo)

	if 1 == 1 {
		httpClient, authInfo, err := mbCfg.NewClient()
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(authInfo)

		sts, _, err := metabaseutil.QuerySTS(httpClient, mbCfg.BaseUrl, sqlInfo)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(sts)
	}

	if 1 == 0 {
		apiClient, authInfo, err := metabaseutil.NewApiClientConfig(mbCfg)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(authInfo)

		info, resp, err := metabaseutil.QuerySQL(apiClient, sqlInfo.DatabaseID, sqlInfo.SQL)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode >= 300 {
			log.Fatal(fmt.Sprintf("STATUS_CODE [%v]", resp.StatusCode))
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
