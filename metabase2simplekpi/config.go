package metabase2simplekpi

// $ go run sql.go -d 3 -q abc -s 'https://data.corp.ringcentral.com/pla-prod' -t 'ab33281e-d9c1-46e5-8c3b-63ea9e4a77fb'

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

func (cfg *Config) InitMetabase() error {
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

func (cfg *Config) InitSimplekpi() error {
	if cfg.SimplekpiApiClient != nil {
		return nil
	}
	client, err := simplekpiutil.NewClient(
		cfg.SimplekpiConfig.Site,
		cfg.SimplekpiConfig.Username,
		cfg.SimplekpiConfig.Token)
	if err != nil {
		return err
	}
	cfg.SimplekpiApiClient = client
	return nil
}
