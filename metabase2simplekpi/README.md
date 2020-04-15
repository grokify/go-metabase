# Metabase to SimpleKPI

## Synopsis

```go
    m2sCfg := metabase2simplekpi.Config{
		MetabaseConfig:  metabase.Config{...},
		SimplekpiConfig: simplekpiutil.Config{...},
		SimplekpiUserID: 111,
		LoadSimpleKpi:   true,
		Datasets:        []DatasetInfo{},
	}

	err = m2sCfg.InitClients()
	if err != nil {
		log.Fatal(errors.Wrap(err, "E_MS2CFG_InitClients"))
	}

	err = m2sCfg.ExecWriteCSVs()
	if err != nil {
		log.Fatal(errors.Wrap(err, "E_MS2CFG_ExecWriteCSVs"))
	}
```