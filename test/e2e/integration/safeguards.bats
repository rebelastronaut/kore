#!/usr/bin/env bats
#
# Copyright 2020 Appvia Ltd <info@appvia.io>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
load helper

setup() {
  ${KORE} get cluster ${CLUSTER} | grep -i deleting && skip || true
  ${KORE} get cluster ${CLUSTER} | grep -i pending && skip || true
}

@test "We should not be able to delete the team if clusters exist" {
  runit "${KORE} delete teams ${TEAM} || true"
  [[ "$status" -eq 0 ]]
}

@test "We should not be able to delete the cluster if a namespace exists" {
  runit "${KORE} create namespace -c ${CLUSTER} ingress -t ${TEAM}"
  [[ "$status" -eq 0 ]]
  runit "${KORE} delete teams ${TEAM} 2>&1 | grep 'the following objects need to be deleted first'"
  [[ "$status" -eq 0 ]]
  runit "${KORE} delete namespaceclaims ${CLUSTER}-ingress -t ${TEAM}"
  [[ "$status" -eq 0 ]]
  retry 4 "${KORE} get namespaceclaims ${CLUSTER}-ingress -t ${TEAM} || true"
  [[ "$status" -eq 0 ]]
}

#@test "We should not be able to create a gke cluster with an disallowed paramater" {
#  ${KORE} get allocation gkecredentials-gke -t ${TEAM} || skip
#
#  runit "${KORE} create cluster ${CLUSTER} -p gke-development -a gkecredentials-gke --param '{\"enableIstio\":\"true\"}' || true"
#  [[ "$status" -eq 0 ]]
#}
