package metabase2simplekpi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gotilla/encoding/jsonutil"
	"github.com/grokify/gotilla/fmt/fmtutil"
	mo "github.com/grokify/oauth2more/metabase"
	"github.com/pkg/errors"
)

// Config runs a set of Metabase to SimpleKPI queries.
// Call Config.InitClients() and then Config.Exec()
// or Config.ExecWriteCSVs().
type Config struct {
	MetabaseConfig     mo.Config
	SimplekpiConfig    simplekpiutil.Config
	SimplekpiUserID    int64
	LoadSimpleKpi      bool
	Datasets           []DatasetInfo
	MetabaseHttpClient *http.Client
	SimplekpiApiClient *simplekpi.APIClient
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

func (cfg *Config) Exec() error {
	return ExecConfig(*cfg, func(ds DatasetInfo, sr *SqlResponse) error {
		return nil
	})
}

func (cfg *Config) ExecFunc(funcSqlResp func(ds DatasetInfo, sr *SqlResponse) error) error {
	return ExecConfig(*cfg, funcSqlResp)
}

func (cfg *Config) ExecWriteCSVs() error {
	return ExecConfig(*cfg, WriteSQLResponseCSV)
}

func ExecConfig(m2sCfg Config, funcSqlResp func(ds DatasetInfo, sr *SqlResponse) error) error {
	for i, ds := range m2sCfg.Datasets {
		fmt.Printf("MB2SK_PROCESSING [%v/%v][%v][%v]\n", i, len(m2sCfg.Datasets), ds.KpiName, ds.MetabaseQueryExec)
		err := ExecDataset(m2sCfg, ds, funcSqlResp)
		if err != nil {
			return err
		}
	}
	return nil
}

func ExecDataset(m2sCfg Config, ds DatasetInfo, funcSqlResp func(ds DatasetInfo, sr *SqlResponse) error) error {
	if !ds.MetabaseQueryExec {
		return nil
	}
	sr, err := QueryMetabase(m2sCfg, ds)
	if err != nil {
		fmtutil.PrintJSON(ds)
		fmt.Printf("NATIVE_SQL: %s\n", ds.NativeSQL())
		log.Fatal(errors.Wrap(err, fmt.Sprintf("E_QUERY_MB [%v]", ds.KpiName)))
	}

	if ds.SimplekpiUpdateExec {
		errs := UpdateSimpleKpiSqlResponse(m2sCfg, ds, sr)
		if len(errs) > 0 {
			bytes, err := ErrorsToJSON(errs)
			if err != nil {
				return errors.Wrap(err, "metabase2simplekpi.ExecDataset")
			}
			return fmt.Errorf(string(bytes))
		}
	}

	if funcSqlResp != nil {
		err := funcSqlResp(ds, sr)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteSQLResponseCSV(ds DatasetInfo, sr *SqlResponse) error {
	tbl, err := SqlResponseToTable(
		sr, int64(ds.SimplekpiKpiId),
		ds.MetabaseQueryColIdxCount,
		ds.MetabaseQueryColIdxDate)
	if err != nil {
		return errors.Wrap(err, "metabase2simplekpi.WriteSQLResponseCSV")
	}
	filename := ds.KpiName + ".csv"
	err = tbl.WriteCSV(filename)
	if err != nil {
		return errors.Wrap(err, "metabase2simplekpi.WriteSQLResponseCSV")
	}
	fmt.Printf("SUCCESS [%v]\n", filename)
	return nil
}

func ErrorsToJSON(errs []error) ([]byte, error) {
	errStrings := []string{}
	for _, err := range errs {
		errStrings = append(errStrings, err.Error())
	}
	return jsonutil.MarshalSimple(errStrings, "", "")
}
