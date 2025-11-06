package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/jessevdk/go-flags"

	"github.com/1024pix/go-metabase/metabase"
	"github.com/1024pix/go-metabase/metabaseutil"
	mo "github.com/grokify/goauth/metabase"
)

type Options struct {
	Env     string `short:"e" long:"envfile" description:"Env File" required:"false"`
	Id      int32  `short:"i" long:"databaseId" description:"The Id of database to get metadata" required:"true"`
	Verbose bool   `short:"v" long:"verbose" description:"Verbose - Include Hidden" required:"false"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	read, err := config.LoadDotEnv([]string{
		opts.Env, os.Getenv("ENV_PATH"), ".env"}, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmtutil.PrintJSON(read)

	clientCfg := mo.Config{
		BaseURL:       os.Getenv("METABASE_BASE_URL"),
		Username:      os.Getenv("METABASE_USERNAME"),
		Password:      os.Getenv("METABASE_PASSWORD"),
		SessionID:     os.Getenv("METABASE_SESSION_ID"),
		TLSSkipVerify: stringsutil.ToBool(os.Getenv("METABASE_TLS_SKIP_VERIFY"))}

	err = clientCfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	apiClient, _, err := metabaseutil.NewApiClient(clientCfg)
	if err != nil {
		log.Fatal(err)
	}

	err = printDatabase(apiClient, opts.Verbose, opts.Id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}

func printDatabase(apiClient *metabase.APIClient, verbose bool, dbId int32) error {

	request := apiClient.DatabaseApi.GetDatabase(context.Background(), dbId)
	request = request.Include("tables")

	db, resp, err := apiClient.DatabaseApi.GetDatabaseExecute(request)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Status Code [%v]", resp.StatusCode)
	}

	fmt.Printf("DB_ID [%v] DB_NAME [%v]\n", db.Id, *db.Name)
	fmtutil.PrintJSON(db)

	return nil
}
