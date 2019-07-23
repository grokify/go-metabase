package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-metabase/metabase"
	mbu "github.com/grokify/go-metabase/util"
)

func printDatabaseList(apiClient *metabase.APIClient) error {
	opts := metabase.ListDatabasesOpts{
		IncludeTables: optional.NewBool(true)}

	info, resp, err := apiClient.DatabaseApi.ListDatabases(
		context.Background(), &opts)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Status Code [%v]", resp.StatusCode)
	}
	//fmtutil.PrintJSON(info)

	for _, db := range info {
		for _, tb := range db.Tables {
			fmt.Printf("DB_ID [%v] DB_NAME [%v] TB_ID [%v] TB_NAME [%v]\n",
				db.Id, db.Name, tb.Id, tb.Name)
		}
	}
	return nil
}

func main() {
	opts := mbu.AppConfig{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	apiClient, authResponse, err := opts.GetClient()
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(authResponse)

	err = printDatabaseList(apiClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
