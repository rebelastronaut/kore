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

package controllers

import (
	"fmt"
)

type ReconcileError struct {
	Err      error
	Critical bool
}

func NewReconcileError(err error, critical bool) *ReconcileError {
	return &ReconcileError{
		Err:      err,
		Critical: critical,
	}
}

func (r *ReconcileError) Error() string {
	return r.Err.Error()
}

func (r *ReconcileError) Wrap(message string) *ReconcileError {
	if r != nil && r.Err != nil {
		r.Err = fmt.Errorf("%s: %w", message, r.Err)
	}
	return r
}

func (r *ReconcileError) Wrapf(format string, args ...interface{}) *ReconcileError {
	if r != nil && r.Err != nil {
		r.Err = fmt.Errorf(format, append(args, r.Err)...)
	}
	return r
}
