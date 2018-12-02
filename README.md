# Metabase Go SDK

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]
[![Chat][gitter-svg]][gitter-url]
[![Discuss][forum-svg]][forum-url]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-metabase.svg?branch=master
 [build-status-url]: https://travis-ci.org/grokify/go-metabase
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-metabase
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-metabase
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-url]: https://godoc.org/github.com/grokify/go-metabase
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

* https://github.com/grokify/go-metabase/blob/master/client/README.md

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

* Query Language '98
  * https://github.com/metabase/metabase/wiki/Query-Language-%2798
* Metabase API Documentation:
  * https://github.com/metabase/metabase/blob/master/docs/api-documentation.md
* Metabase OpenAPI 3.0 Spec:
  * https://github.com/grokify/go-metabase/blob/master/codegen/swagger_spec.yaml

### Example Queries

* https://github.com/metabase/metabase/issues/2683
* https://github.com/metabase/metabase/issues/5635