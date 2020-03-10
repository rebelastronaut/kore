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

package apiserver

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/appvia/kore/pkg/kore"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
func returnNotImplemented(req *restful.Request, wr *restful.Response) {
	wr.WriteHeader(http.StatusNotImplemented)
}
*/

// newList provides an api list type
func newList() *List {
	return &List{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "List",
		},
	}
}

func makeListWithSize(size int) *List {
	l := newList()
	l.Items = make([]string, size)

	return l
}

// handleErrors is a generic wrapper for handling the error from downstream kore brigde
func handleErrors(req *restful.Request, resp *restful.Response, handler func() error) {
	if err := handler(); err != nil {
		code := http.StatusInternalServerError
		switch err {
		case kore.ErrNotFound:
			code = http.StatusNotFound
		case kore.ErrNotAllowed{}:
			code = http.StatusNotAcceptable
		case kore.ErrUnauthorized:
			code = http.StatusForbidden
		case kore.ErrRequestInvalid:
			code = http.StatusBadRequest
		case io.EOF:
			code = http.StatusBadRequest
		}
		if strings.Contains(err.Error(), "record not found") {
			code = http.StatusNotFound
			err = errors.New("resource not found")
		}
		if kerrors.IsNotFound(err) {
			code = http.StatusNotFound
		}

		e := newError(err.Error()).
			WithCode(code).
			WithVerb(req.Request.Method).
			WithURI(req.Request.RequestURI)

		if err := resp.WriteHeaderAndEntity(code, e); err != nil {
			log.WithError(err).Error("failed to respond to request")
		}
	}
}
