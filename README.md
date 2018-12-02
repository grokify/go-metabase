# Metabase Go SDK

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-metabase.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-metabase
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-metabase
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-metabase
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-metabase
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-metabase/blob/master/LICENSE

Go SDK built using OpenAPI 3.0 spec and OpenAPI Generator.

Documentation:

* https://github.com/grokify/go-metabase/blob/master/client/README.md

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

References:

* Metabase API Documentation:
  * https://github.com/metabase/metabase/blob/master/docs/api-documentation.md
* Metabase OpenAPI 3.0 Spec:
  * https://github.com/grokify/go-metabase/blob/master/codegen/swagger_spec.yaml