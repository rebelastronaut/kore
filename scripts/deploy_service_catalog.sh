#!/usr/bin/env sh

set -eu

err_exit() {
  echo "Error: $1"
  exit 1
}

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." >/dev/null 2>&1 && pwd )"
CHARTS_DIR="${PROJECT_DIR}/charts"
CHARTS="kore-costs"

SERVICE_CATALOG_GS_BUCKET=${SERVICE_CATALOG_GS_BUCKET:-kore-service-catalog}
SERVICE_CATALOG_URL=${SERVICE_CATALOG_URL:-https://${SERVICE_CATALOG_GS_BUCKET}.storage.googleapis.com}

[ -n "${SERVICE_CATALOG_URL}" ] || err_exit "SERVICE_CATALOG_URL must be set"
[ -n "${SERVICE_CATALOG_GS_BUCKET}" ] || err_exit "SERVICE_CATALOG_GS_BUCKET must be set"

cd "${CHARTS_DIR}"

rm -rf index.yaml *.tgz

gsutil cp gs://kore-service-catalog/index.yaml .

for chart in ${CHARTS}; do
  helm package "${chart}"
done

helm repo index . --url "${SERVICE_CATALOG_URL}" --merge index.yaml

gsutil -h "Cache-Control:private, max-age=0, no-transform" cp index.yaml gs://kore-service-catalog
gsutil -h "Cache-Control:private, max-age=0, no-transform" cp *.tgz gs://kore-service-catalog
