#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [ ! -f ./demo.env ] ; then
    echo "You could copy and edit the file:"
    echo "    cp ./hack/compose/demo.env.tmpl ./demo.env"
    exit 1
fi
source ./demo.env

if [[ ${KUBE_CONFIG_FILE:-} != '' ]]; then
    KUBE_CONFIG_FLAG="--kubeconfig ${KUBE_CONFIG_FILE}"
else
    KUBE_CONFIG_FLAG=''
    KUBE_API_SERVER=${KUBE_API_SERVER:-http://127.0.0.1:8080}
fi

export \
    KORE_IDP_CLIENT_ID \
    KORE_IDP_CLIENT_SECRET \
    KORE_IDP_SERVER_URL \
    KORE_IDP_USER_CLAIMS \
    KORE_IDP_CLIENT_SCOPES \
    KORE_FEATURE_GATES \
    CLOUD_INFO_URL

./bin/kore-apiserver \
    ${KUBE_CONFIG_FLAG} \
    --listen localhost:10080 \
    --verbose \
    --admin-pass password \
    --admin-token password \
    --api-public-url http://localhost:10080 \
    --ui-public-url http://localhost:3000 \
    --kore-authentication-plugin basicauth \
    --kore-authentication-plugin admintoken \
    --kore-authentication-plugin openid \
    --kore-authentication-plugin localjwt \
    --local-jwt-public-key MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAIG6XiNhkwDETU2zk0tGlI0DKlbEJcN4jxwJBqhd3neReLDnqg9SBgKepdy9Nxw5LAd1gNoBkLvdFJg9SbHlM0sCAwEAAQ== \
    --certificate-authority	hack/ca/ca.pem \
    --certificate-authority-key hack/ca/ca-key.pem \
    --users-db-url 'root:pass@tcp(localhost:3306)/kore?parseTime=true'
