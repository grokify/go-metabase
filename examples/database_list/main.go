package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/go-metabase/metabaseutil"
	mo "github.com/grokify/goauth/metabase"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/strconv/strconvutil"
)

type optionsDbList struct {
	Verbose []bool `short:"v" long:"verbose" description:"Verbose - Include Tables" required:"false"`
}

func main() {
	loaded, err := config.LoadDotEnvSkipEmptyInfo(os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal(errors.Wrap(err, "E_LOAD_ENV"))
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
	opts := metabase.ListDatabasesOpts{
		IncludeTables: optional.NewBool(true)}

	info, resp, err := apiClient.DatabaseApi.ListDatabases(
		context.Background(), &opts)
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
