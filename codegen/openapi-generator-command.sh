#perl -p -i -e 's/\s+\[Beta\]\s*$/\n/g' swagger_spec.yaml
java -jar openapi-generator-cli.jar generate -i swagger_spec.yaml -g go -o metabase -D packageName=metabase
perl -p -i -e 's/(\[\]\[\]\S+)/[][]interface{}/g' metabase/model_dataset_query_results_data.go
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> metabase/client.go
gofmt -s -w metabase/*.go
