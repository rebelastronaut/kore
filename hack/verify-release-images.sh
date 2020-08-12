#!/bin/bash
#
# Copyright (C) 2020 Appvia Ltd <info@appvia.io>
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License
# as published by the Free Software Foundation; either version 2
# of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
set -o errexit
set -o nounset
set -o pipefail

BRANCH=$(git rev-parse --abbrev-ref HEAD)
FORCE_IMAGE_PUSH=${FORCE_IMAGE_PUSH:-false}
VERSION=${VERSION:-latest}

# Used by CI to push without confirmation
if [[ ${FORCE_IMAGE_PUSH} == true ]]; then
  echo "Force image push detected, skipping checks"
  exit 0
fi

if [[ "${VERSION}" == "latest" ]]; then
  if [[ ${BRANCH} != "master" ]]; then
    echo "Refusing to push latest on none master branch"
    exit 1
  fi

  echo "Are you REALLY sure you want to push image latest? (y/n)"
  read -n1 choice
  if [[ ! "${choice}" =~ ^[Yy]$ ]]; then
    echo "exitting without pushing images"
    exit 1
  fi
fi
