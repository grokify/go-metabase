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
	Env string `short:"e" long:"envfile" description:"Env File" required:"false"`
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

	//err = clientCfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	apiClient, _, err := metabaseutil.NewApiClient(clientCfg)
	if err != nil {
		log.Fatal(err)
	}

	err = printPermissionsGraph(apiClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}

func printPermissionsGraph(apiClient *metabase.APIClient) error {

	request := apiClient.PermissionsApi.GetPermissionsGraph(context.Background())

	peg, resp, err := apiClient.PermissionsApi.GetPermissionsGraphExecute(request)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("Status Code [%v]", resp.StatusCode)
	}

	fmt.Printf("REVISION [%d]\n", *peg.Revision)
	fmt.Printf("%+v\n", peg.Groups)
	return nil
}
