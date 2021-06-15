package metabase2simplekpi

import (
	"github.com/grokify/go-metabase/metabaseutil"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/timeseries"
)

func QueryMetabase(cfg Config, dsInfo DatasetInfo) (*SqlResponse, error) {
	resp, err := metabaseutil.QuerySQLHttp(
		cfg.MetabaseHttpClient,
		cfg.MetabaseConfig.BaseURL,
		dsInfo.MetabaseQuery.DatabaseID,
		dsInfo.MetabaseQuery.NativeSQL())
	if err != nil {
		return &SqlResponse{}, err
	}
	return HTTPResponseToSqlResponse(resp)
}

func QueryMetabaseTS(cfg Config, dsInfo DatasetInfo) (timeseries.TimeSeries, error) {
	ts := timeseries.TimeSeries{}
	sqlResp, err := QueryMetabase(cfg, dsInfo)
	if err != nil {
		return ts, err
	}
	return SqlResponseToTS(
		dsInfo.KpiName,
		sqlResp,
		dsInfo.MetabaseQuery.ColIdxCount,
		dsInfo.MetabaseQuery.ColIdxTime)
}

func UpdateSimpleKpi(cfg Config, dsInfo DatasetInfo) []error {
	sts, err := QueryMetabaseTS(cfg, dsInfo)
	if err != nil {
		return []error{err}
	}
	return UpdateSimpleKpiTS(cfg, dsInfo, sts)
}

func UpdateSimpleKpiSqlResponse(cfg Config, dsInfo DatasetInfo, sqlResp *SqlResponse) []error {
	ts, err := SqlResponseToTS(dsInfo.KpiName, sqlResp,
		dsInfo.MetabaseQuery.ColIdxCount,
		dsInfo.MetabaseQuery.ColIdxTime)
	if err != nil {
		return []error{err}
	}
	return UpdateSimpleKpiTS(cfg, dsInfo, ts)
}

func UpdateSimpleKpiTS(cfg Config, dsInfo DatasetInfo, ts timeseries.TimeSeries) []error {
	_, resps, err := simplekpiutil.UpsertKpiEntriesStaticTimeSeries(
		cfg.SimplekpiApiClient,
		int64(cfg.SimplekpiUserID),
		int64(dsInfo.SimplekpiKpiId),
		ts)
	if err != nil {
		return []error{err}
	}
	return simplekpiutil.KpiEntryResponseErrors(resps)
}
