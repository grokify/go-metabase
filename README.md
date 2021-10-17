# Metabase Go SDK

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]
[![Chat][gitter-svg]][gitter-url]
[![Discuss][forum-svg]][forum-url]

 [build-status-svg]: https://github.com/grokify/go-metabase/workflows/go%20build/badge.svg
 [build-status-url]: https://github.com/grokify/go-metabase/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-metabase
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-metabase
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-metabase
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-metabase
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-metabase/blob/master/LICENSE
 [gitter-svg]: https://badges.gitter.im/metabase/metabase.svg
 [gitter-url]: https://gitter.im/metabase/metabase?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge
 [forum-svg]: https://img.shields.io/badge/post-on%20discourse-red.svg
 [forum-url]: http://discourse.metabase.com/

Go SDK built using OpenAPI 3.0 spec and OpenAPI Generator.

## Installation

```
$ go get github.com/grokify/go-metabase/...
```

## Usage

See the examples in the `examples` folder:

https://github.com/grokify/go-metabase/tree/master/examples

## Documentation

Auto-generated documentation is available here:

* https://github.com/grokify/go-metabase/blob/master/metabase/README.md

## Coverage

- [ ] Database
  - [ ] DELETE /api/database/:id
  - [x] GET /api/database/
  - [ ] GET /api/database/:id
  - [ ] GET /api/database/:id/autocomplete_suggestions
  - [ ] GET /api/database/:id/fields
  - [ ] GET /api/database/:id/idfields
  - [ ] GET /api/database/:id/metadata
  - [ ] POST /api/database/
  - [ ] POST /api/database/:id/sync
  - [ ] POST /api/database/sample_dataset
  - [ ] PUT /api/database/:id

- [ ] Dataset
  - [x] POST /api/dataset/
  - [ ] POST /api/dataset/csv
  - [ ] POST /api/dataset/duration
  - [ ] POST /api/dataset/json

## References

* Using the REST API
  * https://github.com/metabase/metabase/wiki/Using-the-REST-API
* Metabase API Documentation:
  * https://github.com/metabase/metabase/blob/master/docs/api-documentation.md
* Metabase OpenAPI 3.0 Spec:
  * https://github.com/grokify/go-metabase/blob/master/codegen/swagger_spec.yaml
* Query Language '98
  * https://github.com/metabase/metabase/wiki/Query-Language-%2798

### Example Queries

* https://github.com/metabase/metabase/issues/2683
* https://github.com/metabase/metabase/issues/5635

* Current User: `curl -XGET 'http://server/api/user/current' -H 'X-Metabase-Session: 38f4939c-ad7f-4cbe-ae54-30946daf8593'`
* Table: `curl -XGET 'https://base_url/api/table/1' -H 'X-Metabase-Session: 11112222-3333-4444-5555-666677778888'`

Describe columns and column types:

`describe mydatabase.mytable;`