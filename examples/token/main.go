package main

import (
	"fmt"
	"log"

	"github.com/grokify/oauth2more/metabase"
	"github.com/jessevdk/go-flags"
)

type optionsToken struct {
	Site     string `short:"s" long:"site" description:"Your Site" required:"true"`
	Username string `short:"u" long:"username" description:"Your username" required:"true"`
	Password string `short:"p" long:"password" description:"Your password" required:"true"`
}

func main() {
	opts := optionsToken{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	_, authInfo, err := metabase.NewClientPassword(opts.Site, opts.Username, opts.Password, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s\n", authInfo.Id)

	fmt.Println("DONE")
}
