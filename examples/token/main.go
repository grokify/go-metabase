package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/type/stringsutil"
	"github.com/grokify/oauth2more/metabase"
	mo "github.com/grokify/oauth2more/metabase"
	"github.com/jessevdk/go-flags"
)

type mbOptions struct {
	EnvPath   string `short:"e" long:"env_path" description:"Env File" required:"false"`
	Site      string `short:"s" long:"site" description:"Your Site" required:"false"`
	Username  string `short:"u" long:"username" description:"Your username" required:"false"`
	Password  string `short:"p" long:"password" description:"Your password" required:"false"`
	SessionId string `short:"t" long:"token" description:"Your token" required:"false"`
}

func (opts *mbOptions) Inflate(mbCfg mo.Config) mo.Config {
	if len(opts.Site) > 0 {
		mbCfg.BaseUrl = opts.Site
	}
	if len(opts.Username) > 0 {
		mbCfg.Username = opts.Username
	}
	if len(opts.Password) > 0 {
		mbCfg.Password = opts.Password
	}
	if len(opts.SessionId) > 0 {
		mbCfg.SessionId = opts.SessionId
	}
	return mbCfg
}

func main() {
	opts := mbOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	read, err := config.LoadDotEnvSkipEmptyInfo(
		opts.EnvPath,
		os.Getenv("ENV_PATH"),
		".env")
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
	mbCfg = opts.Inflate(mbCfg)

	fmtutil.PrintJSON(mbCfg)

	_, authInfo, err := metabase.NewClientPassword(
		mbCfg.BaseUrl, mbCfg.Username, mbCfg.Password, mbCfg.TlsSkipVerify)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s\n", authInfo.Id)

	fmt.Println("DONE")
}
