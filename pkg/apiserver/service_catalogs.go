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
	"fmt"
	"net/http"

	"github.com/appvia/kore/pkg/utils/validation"

	"github.com/appvia/kore/pkg/apiserver/params"

	"github.com/appvia/kore/pkg/apiserver/filters"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterHandler(&serviceCatalogsHandler{})
}

type serviceCatalogsHandler struct {
	kore.Interface
	// DefaultHandler implements default features
	DefaultHandler
}

// Register is called by the api server on registration
func (p *serviceCatalogsHandler) Register(i kore.Interface, builder utils.PathBuilder) (*restful.WebService, error) {
	path := builder.Add("servicecatalogs")

	log.WithFields(log.Fields{
		"path": path.Base(),
	}).Info("registering the servicecatalogs webservice")

	p.Interface = i

	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path(path.Base())

	ws.Route(
		withAllNonValidationErrors(ws.GET("")).To(p.listServiceCatalogs).
			Doc("Returns all the available service catalogs").
			Operation("ListServiceCatalogs").
			Returns(http.StatusOK, "A list of service catalogs", servicesv1.ServiceCatalogList{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{name}")).To(p.getServiceCatalog).
			Doc("Returns a specific service catalog").
			Operation("GetServiceCatalog").
			Param(ws.PathParameter("name", "The name of the service catalog you wish to retrieve")).
			Returns(http.StatusNotFound, "the service catalog with the given name doesn't exist", nil).
			Returns(http.StatusOK, "Contains the service catalog definition", servicesv1.ServiceCatalog{}),
	)

	ws.Route(
		withAllErrors(ws.PUT("/{name}")).To(p.updateServiceCatalog).
			Filter(filters.Admin).
			Filter(p.readOnlyServiceCatalogFilter).
			Doc("Creates or updates a service catalog").
			Operation("UpdateServiceCatalog").
			Param(ws.PathParameter("name", "The name of the service catalog you wish to create or update")).
			Reads(servicesv1.ServiceCatalog{}, "The specification for the service catalog you are creating or updating").
			Returns(http.StatusOK, "Contains the service catalog definition", servicesv1.ServiceCatalog{}),
	)

	ws.Route(
		withAllErrors(ws.DELETE("/{name}")).To(p.deleteServiceCatalog).
			Filter(filters.Admin).
			Filter(p.readOnlyServiceCatalogFilter).
			Doc("Deletes a service catalog").
			Operation("DeleteServiceCatalog").
			Param(ws.PathParameter("name", "The name of the service catalog you wish to delete")).
			Param(params.DeleteCascade()).
			Returns(http.StatusNotFound, "the service catalog with the given name doesn't exist", nil).
			Returns(http.StatusOK, "Contains the service catalog definition", servicesv1.ServiceCatalog{}),
	)

	return ws, nil
}

func (p *serviceCatalogsHandler) readOnlyServiceCatalogFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	handleErrors(req, resp, func() error {
		name := req.PathParameter("name")

		catalog, err := p.ServiceCatalogs().Get(req.Request.Context(), name)
		if err != nil && err != kore.ErrNotFound {
			return err
		}

		if catalog != nil && catalog.Annotations[kore.AnnotationReadOnly] == "true" {
			resp.WriteHeader(http.StatusForbidden)
			return nil
		}

		// @step: continue with the chain
		chain.ProcessFilter(req, resp)
		return nil
	})
}

// getServiceCatalog returns a specific service catalog
func (p serviceCatalogsHandler) getServiceCatalog(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		catalog, err := p.ServiceCatalogs().Get(req.Request.Context(), req.PathParameter("name"))
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, catalog)
	})
}

// listServiceCatalogs returns all service catalogs in the kore
func (p serviceCatalogsHandler) listServiceCatalogs(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		res, err := p.ServiceCatalogs().List(req.Request.Context())
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, res)
	})
}

// updateServiceCatalog is used to update or create a service catalog in the kore
func (p serviceCatalogsHandler) updateServiceCatalog(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		name := req.PathParameter("name")

		catalog := &servicesv1.ServiceCatalog{}
		if err := req.ReadEntity(catalog); err != nil {
			return err
		}
		if catalog.Name != name {
			return validation.NewError("service catalog failed validation").
				WithFieldError("name", validation.InvalidValue, "doesn't match name given in request path")
		}

		if catalog.Annotations[kore.AnnotationReadOnly] != "" {
			writeError(req, resp, fmt.Errorf("setting %q annotation is not allowed", kore.AnnotationReadOnly), http.StatusForbidden)
			return nil
		}

		if err := p.ServiceCatalogs().Update(req.Request.Context(), catalog); err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, catalog)
	})
}

// deleteServiceCatalog is used to update or create a service catalog in the kore
func (p serviceCatalogsHandler) deleteServiceCatalog(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		name := req.PathParameter("name")

		catalog, err := p.ServiceCatalogs().Delete(req.Request.Context(), name, parseDeleteOpts(req)...)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, catalog)
	})
}

// Name returns the name of the handler
func (p serviceCatalogsHandler) Name() string {
	return "servicecatalogs"
}

// Enabled returns true if the services feature gate is enabled
func (p serviceCatalogsHandler) Enabled() bool {
	return p.Config().IsFeatureGateEnabled(kore.FeatureGateServices)
}
