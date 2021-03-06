# Go API client for metabase

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./metabase"
```

## Documentation for API Endpoints

All URIs are relative to *http://example.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DatabaseApi* | [**ListDatabases**](docs/DatabaseApi.md#listdatabases) | **Get** /api/database | List Databases
*DatasetApi* | [**QueryDatabase**](docs/DatasetApi.md#querydatabase) | **Post** /api/dataset | Execute a query


## Documentation For Models

 - [Database](docs/Database.md)
 - [DatabaseDetails](docs/DatabaseDetails.md)
 - [DatabaseTable](docs/DatabaseTable.md)
 - [DatasetQueryConstraints](docs/DatasetQueryConstraints.md)
 - [DatasetQueryDsl](docs/DatasetQueryDsl.md)
 - [DatasetQueryDslPage](docs/DatasetQueryDslPage.md)
 - [DatasetQueryJsonQuery](docs/DatasetQueryJsonQuery.md)
 - [DatasetQueryNative](docs/DatasetQueryNative.md)
 - [DatasetQueryOpts](docs/DatasetQueryOpts.md)
 - [DatasetQueryResults](docs/DatasetQueryResults.md)
 - [DatasetQueryResultsCol](docs/DatasetQueryResultsCol.md)
 - [DatasetQueryResultsColFingerprint](docs/DatasetQueryResultsColFingerprint.md)
 - [DatasetQueryResultsColFingerprintGlobal](docs/DatasetQueryResultsColFingerprintGlobal.md)
 - [DatasetQueryResultsColFingerprintType](docs/DatasetQueryResultsColFingerprintType.md)
 - [DatasetQueryResultsColTarget](docs/DatasetQueryResultsColTarget.md)
 - [DatasetQueryResultsData](docs/DatasetQueryResultsData.md)
 - [DatasetQueryResultsMetadata](docs/DatasetQueryResultsMetadata.md)
 - [DatasetQueryResultsMetadataColumn](docs/DatasetQueryResultsMetadataColumn.md)
 - [DatasetQueryResultsNativeForm](docs/DatasetQueryResultsNativeForm.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



