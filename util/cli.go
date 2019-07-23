package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/grokify/go-metabase/metabase"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	mo "github.com/grokify/oauth2more/metabase"
)

type AppConfig struct {
	Config    string `short:"c" long:"config" description:"Config (cli,env)" required:"true"`
	URL       string `short:"u" long:"url" description:"The base URL" required:"false"`
	Username  string `short:"n" long:"username" description:"The username" required:"false"`
	Password  string `short:"p" long:"password" description:"The password" required:"false"`
	SessionId string `short:"s" long:"sessionid" description:"The sessionId" required:"false"`
}

func (opts *AppConfig) validate() error {
	err := opts.validateCLI()
	if err != nil {
		return err
	}
	return opts.loadEnv()
}

func (opts *AppConfig) validateCLI() error {
	opts.Config = strings.TrimSpace(opts.Config)
	if opts.Config == "cli" {
		missingFields := []string{}
		opts.URL = strings.TrimSpace(opts.URL)
		opts.Username = strings.TrimSpace(opts.Username)
		opts.Password = strings.TrimSpace(opts.Password)
		opts.SessionId = strings.TrimSpace(opts.SessionId)
		if len(opts.URL) == 0 {
			missingFields = append(missingFields, "url")
		}
		if len(opts.SessionId) == 0 {
			if len(opts.Username) == 0 {
				missingFields = append(missingFields, "username")
			}
			if len(opts.Password) == 0 {
				missingFields = append(missingFields, "password")
			}
		}
		if len(missingFields) > 0 && len(missingFields) < 3 {
			msg := strings.Join(missingFields, ",")
			return fmt.Errorf("Missing Fields: [%s]", msg)
		}
	} else if opts.Config != "env" {
		return fmt.Errorf("ConfigLocation is not in (cli|env) [%s]", opts.Config)
	}
	return nil
}

func (opts *AppConfig) loadEnv() error {
	if opts.Config == "env" {
		err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
		if err != nil {
			return err
		}
		if len(opts.URL) == 0 {
			opts.URL = os.Getenv("METABASE_BASE_URL")
		}
		if len(opts.Username) == 0 {
			opts.Username = os.Getenv("METABASE_USERNAME")
		}
		if len(opts.Password) == 0 {
			opts.Password = os.Getenv("METABASE_PASSWORD")
		}
		if len(opts.SessionId) == 0 {
			opts.SessionId = os.Getenv("METABASE_SESSION_ID")
		}
	}
	return nil
}

func (opts *AppConfig) GetClient() (*metabase.APIClient, *mo.AuthResponse, error) {
	err := opts.validate()
	if err != nil {
		return nil, nil, err
	}
	fmtutil.PrintJSON(opts)
	return NewApiClientPasswordWithSessionId(
		opts.URL, opts.Username, opts.Password, opts.SessionId, true)
}
