/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashString returns a md5 of the string
func HashString(v string) string {
	return string(HashBytes([]byte(v)))
}

// HashBytes returns a md5 of the bytes
func HashBytes(v []byte) []byte {
	w := md5.New()
	_, _ = w.Write(v)

	return []byte(hex.EncodeToString(w.Sum(nil)))
}
