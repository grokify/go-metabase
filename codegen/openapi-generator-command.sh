#perl -p -i -e 's/\s+\[Beta\]\s*$/\n/g' swagger_spec.yaml

if ! command -v java &> /dev/null
then
    echo "You must install Java before run this script"
    exit
fi

if [ ! -f openapi-generator-cli.jar ]; then
  curl --output openapi-generator-cli.jar -L https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.6.0/openapi-generator-cli-6.6.0.jar
fi

#GITHUB_PATH="$(git config --get remote.origin.url | cut -d : -f 2)"
GITHUB_USER=grokify #"$(echo "${GITHUB_PATH}" | cut -d / -f 1)"
GITHUB_REPO_NAME=go-metabase #"$(echo "${GITHUB_PATH}" | cut -d / -f 2 | sed 's/.git//')"

java -jar openapi-generator-cli.jar generate \
  -i swagger_spec.yaml \
  -g go \
  -o metabase \
  --git-repo-id "${GITHUB_REPO_NAME}" \
  --git-user-id "${GITHUB_USER}" \
  --additional-properties=packageName=metabase
# perl -p -i -e 's/(\[\]\[\]\S+)/[][]interface{}/g' metabase/model_dataset_query_results_data.go
echo -e "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> metabase/client.go
gofmt -s -w metabase/*.go

rm -rf metabase/test
rm -f metabase/go.*
rm -f metabase/git_push.sh
rm -rf ../metabase
mv metabase ../
