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

package indexer

// Interface is the contract to the indexer
type Interface interface {
	// Delete remove the id from the index
	Delete(string) error
	// DeleleByQuery removes all the documents which match the query
	DeleteByQuery(interface{}) (int, error)
	// DeleleByQueryRaw removes all the documents which the raw query
	DeleteByQueryRaw(string) (int, error)
	// Index is responsible is add a document the index
	Index(string, interface{}) error
	// Query is responsible for searching the index
	Query(interface{}) ([]string, error)
	// QueryRaw is responsible for searching the index with a raw query
	QueryRaw(string) ([]string, error)
	// Size returns the size of the index
	Size() (uint64, error)
}
