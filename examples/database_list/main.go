package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/go-metabase/metabaseutil"
	mo "github.com/grokify/goauth/metabase"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/strconv/strconvutil"
)

type optionsDbList struct {
	Verbose []bool `short:"v" long:"verbose" description:"Verbose - Include Tables" required:"false"`
}

func main() {
	loaded, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH")}, 1)
	if err != nil {
		log.Fatal(errorsutil.Wrap(err, "E_LOAD_ENV"))
	}
	fmtutil.PrintJSON(loaded)

	opts := optionsDbList{}
	_, err = flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	clientCfg := mo.Config{
		BaseURL:       os.Getenv("METABASE_BASE_URL"),
		SessionID:     os.Getenv("METABASE_SESSION_ID"),
		Username:      os.Getenv("METABASE_USERNAME"),
		Password:      os.Getenv("METABASE_PASSWORD"),
		TLSSkipVerify: strconvutil.MustParseBool(os.Getenv("METABASE_TLS_SKIP_VERIFY")),
	}
	err = clientCfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	apiClient, _, err := metabaseutil.NewApiClient(clientCfg)
	if err != nil {
		log.Fatal(err)
	}

	verbose := false
	if len(opts.Verbose) > 0 {
		verbose = true
	}
	err = printDatabaseList(apiClient, verbose)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}

func printDatabaseList(apiClient *metabase.APIClient, verbose bool) error {

	request := apiClient.DatabaseApi.ListDatabases(context.Background())
	request.IncludeTables(true)

	info, resp, err := apiClient.DatabaseApi.ListDatabasesExecute(request)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Status Code [%v]", resp.StatusCode)
	}

	for _, db := range info {
		fmt.Printf("DB_ID [%v] DB_NAME [%v]\n", db.Id, db.Name)
		fmtutil.PrintJSON(db)
		if verbose {
			for _, tb := range db.Tables {
				fmt.Printf("DB_ID [%v] DB_NAME [%v] TB_ID [%v] TB_NAME [%v] [%v]\n",
					db.Id, db.Name, tb.Id, tb.Name, tb)
				fmtutil.PrintJSON(tb)
			}
		}
	}
	return nil
}
