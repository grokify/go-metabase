package metabase2simplekpi

import (
	"net/http"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	mo "github.com/grokify/oauth2more/metabase"
)

type Config struct {
	MetabaseConfig     mo.Config
	MetabaseHttpClient *http.Client
	SimplekpiConfig    simplekpiutil.Config
	SimplekpiUserID    int64
	SimplekpiApiClient *simplekpi.APIClient
	LoadSimpleKpi      bool
	Datasets           []DatasetInfo
}

type DatasetInfo struct {
	KpiName                 string `json:"kpiName"`
	MetabaseQueryDatabaseId int    `json:"metabaseQueryDatabaseId"`
	MetabaseQueryNativeSQL  string `json:"metabaseQuerySQLNative"`
	SimplekpiKpiId          int    `json:"simplekpiKpiId"`
}

func (cfg *Config) InitClients() error {
	err := cfg.InitMetabaseClient()
	if err != nil {
		return err
	}
	return cfg.InitSimplekpiClient()
}

func (cfg *Config) InitMetabaseClient() error {
	if cfg.MetabaseHttpClient != nil {
		return nil
	}
	httpClient, _, err := mo.NewClient(cfg.MetabaseConfig)
	if err != nil {
		return err
	}
	cfg.MetabaseHttpClient = httpClient
	return nil
}

func (cfg *Config) InitSimplekpiClient() error {
	if cfg.SimplekpiApiClient != nil {
		return nil
	}
	client, err := simplekpiutil.NewApiClient(
		cfg.SimplekpiConfig.Site,
		cfg.SimplekpiConfig.Username,
		cfg.SimplekpiConfig.Token)
	if err != nil {
		return err
	}
	cfg.SimplekpiApiClient = client
	return nil
}
