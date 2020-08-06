#!/bin/sh

set -euo pipefail

ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../.." >/dev/null 2>&1 && pwd )"

for idpvar in KORE_IDP_CLIENT_ID KORE_IDP_CLIENT_SECRET KORE_IDP_SERVER_URL KORE_IDP_USER_CLAIMS KORE_IDP_CLIENT_SCOPES; do
  export ${idpvar}=$(kubectl --context kind-kore -n kore get secret kore-idp -o json | jq -r ".data.${idpvar}" | base64 --decode)
done

export KORE_API_TOKEN=$(kubectl --context kind-kore -n kore get secret kore-api -o json | jq -r ".data.KORE_ADMIN_TOKEN" | base64 --decode)

export KORE_FEATURE_GATES="services=true,application_services=true,monitoring_services=true"
export KORE_UI_SHOW_PROTOTYPES=true

KORE_API_URL=$(kubectl --context kind-kore -n kore get deployment kore-apiserver -o=jsonpath="{.spec.template.spec.containers[?(.name=='kore-apiserver')].env[?(.name == 'KORE_API_PUBLIC_URL')].value}")

export KORE_API_PUBLIC_URL=${KORE_API_URL}
export KORE_API_URL=${KORE_API_URL}/api/v1alpha1

set +e
CA_CERT=$( kubectl --context kind-kore get secret -n kore kore-api-tls -o=jsonpath='{.data.ca\.cert}' | base64 -D)
set -e

if [ -n "${CA_CERT}" ]; then
  echo "${CA_CERT}" > "${ROOT_DIR}/ui/node_extra_ca_certs.pem"
  export NODE_EXTRA_CA_CERTS="${ROOT_DIR}/ui/node_extra_ca_certs.pem"
fi

exec "$@"
