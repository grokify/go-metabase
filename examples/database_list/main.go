package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/go-metabase/metabaseutil"
)

type optionsDbList struct {
	Site    string `short:"s" long:"site" description:"Your Site" required:"true"`
	Token   string `short:"t" long:"token" description:"Your Token" required:"true"`
	Verbose []bool `short:"v" long:"verbose" description:"Verbose - Include Tables" required:"false"`
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
		fmt.Printf("DB_ID [%v] DB_NAME [%v]\n",
			db.Id, db.Name)
		if verbose {
			for _, tb := range db.Tables {
				fmt.Printf("DB_ID [%v] DB_NAME [%v] TB_ID [%v] TB_NAME [%v]\n",
					db.Id, db.Name, tb.Id, tb.Name)
			}
		}
	}
	return nil
}

func main() {
	opts := optionsDbList{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := metabaseutil.NewApiClientSessionId(
		opts.Site, opts.Token, true)

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
