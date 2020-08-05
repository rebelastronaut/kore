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

@test "We should be able to create a user" {
  runit "${KORE} create user local@appvia.io -E"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to delete a existing profie" {
  runit "${KORE} profiles delete local-user --force"
  [[ "$status" -eq 0 ]]
}

@test "We should be to associate a identity with the user" {
  runit "echo -n ${LOCAL_USER_PASS} | ${KORE} create identity basicauth -u local@appvia.io --password -"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to create a local profile with this user" {
  runit "echo -n ${LOCAL_USER_PASS} | ${KORE} profile configure local-user --account basicauth -a http://localhost:10080 --user local@appvia.io --password -"
  [[ "$status" -eq 0 ]]
  runit "${KORE} profiles use local-user"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to authenticate with the user" {
  runit "${KORE} --profile local-user get teams"
  [[ "$status" -eq 0 ]]
}

@test "We should see the basicauth identity" {
  runit "${KORE} --profile local-user get identity | grep basicauth"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to add the user to the e2e team" {
  runit "${KORE} --profile local create member -u local@appvia.io -t ${TEAM}"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to access to the cluster using the local user" {
  runit "${KORE} -t ${TEAM} kubeconfig"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to authorize without error" {
  runit "${KORE} -t ${TEAM} alpha authorize >/dev/null"
  [[ "$status" -eq 0 ]]
  runit "file ${HOME}/.kore/authorized"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to see the pods in the default namespace" {
  runit "${KUBECTL} --context ${CLUSTER} get po"
  [[ "$status" -eq 0 ]]
}

@test "We should remove the local user from the team" {
  runit "${KORE} delete member -u local@appvia.io -t ${TEAM}"
  [[ "$status" -eq 0 ]]
}

@test "We should no longer have access to the cluster" {
  runit "${KUBECTL} --context ${CLUSTER} get po || true"
  [[ "$status" -eq 0 ]]
}

@test "We should be able to switch back to the local profile" {
  runit "${KORE} profiles use local"
  [[ "$status" -eq 0 ]]
}
