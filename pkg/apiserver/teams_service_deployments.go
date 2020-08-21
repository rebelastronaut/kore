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

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/apiserver/filters"
	"github.com/appvia/kore/pkg/apiserver/params"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/kore/authentication"

	restful "github.com/emicklei/go-restful"
)

func (u teamHandler) addServiceDeploymentRoutes(ws *restful.WebService) {
	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/servicedeployments")).To(u.listServiceDeployments).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Lists all service deployments for a team").
			Operation("ListServiceDeployments").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "List of all service deployments for a team", servicesv1.ServiceDeploymentList{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/servicedeployments/{name}")).To(u.getServiceDeployment).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Returns a service deployment").
			Operation("GetServiceDeployment").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name of the service deployment")).
			Returns(http.StatusNotFound, "the service deployment with the given name doesn't exist", nil).
			Returns(http.StatusOK, "The requested service deployment details", servicesv1.ServiceDeployment{}),
	)
	ws.Route(
		withAllErrors(ws.PUT("/{team}/servicedeployments/{name}")).To(u.updateServiceDeployment).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Filter(u.readonlyServiceDeploymentFilter).
			Doc("Creates or updates a service deployment").
			Operation("UpdateServiceDeployment").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the service deployment")).
			Reads(servicesv1.Service{}, "The definition for the service deployment").
			Returns(http.StatusOK, "The service deployment details", servicesv1.ServiceDeployment{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.DELETE("/{team}/servicedeployments/{name}")).To(u.deleteServiceDeployment).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Filter(u.readonlyServiceDeploymentFilter).
			Doc("Deletes a service deployment").
			Operation("DeleteServiceDeployment").
			Param(ws.PathParameter("name", "Is the name of the service deployment")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(params.DeleteCascade()).
			Returns(http.StatusNotFound, "the service deployment with the given name doesn't exist", nil).
			Returns(http.StatusOK, "Contains the former service deployment definition from the kore", servicesv1.ServiceDeployment{}),
	)
}

func (u teamHandler) readonlyServiceDeploymentFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	handleErrors(req, resp, func() error {
		name := req.PathParameter("name")
		team := req.PathParameter("team")

		serviceDeployment, err := u.Teams().Team(team).ServiceDeployments().Get(req.Request.Context(), name)
		if err != nil && err != kore.ErrNotFound {
			return err
		}

		if serviceDeployment != nil && serviceDeployment.Annotations[kore.AnnotationReadOnly] == "true" {
			resp.WriteHeader(http.StatusForbidden)
			return nil
		}

		// @step: continue with the chain
		chain.ProcessFilter(req, resp)
		return nil
	})
}

// listServiceDeployments returns all the servicedeployments from a team
func (u teamHandler) listServiceDeployments(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team := req.PathParameter("team")

		var list *servicesv1.ServiceDeploymentList
		var err error

		user := authentication.MustGetIdentity(req.Request.Context())
		if user.IsGlobalAdmin() {
			list, err = u.Teams().Team(team).ServiceDeployments().List(req.Request.Context())
		} else {
			list, err = u.Teams().Team(team).ServiceDeployments().List(req.Request.Context(), func(sd servicesv1.ServiceDeployment) bool {
				return sd.Annotations[kore.AnnotationSystem] != "true"
			})
		}
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, list)
	})
}

// getServiceDeployment returns a service deployment from a team
func (u teamHandler) getServiceDeployment(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		name := req.PathParameter("name")
		team := req.PathParameter("team")

		serviceDeployment, err := u.Teams().Team(team).ServiceDeployments().Get(req.Request.Context(), name)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, serviceDeployment)
	})
}

// updateServiceDeployment is responsible for creating or updating a service deployment
func (u teamHandler) updateServiceDeployment(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team := req.PathParameter("team")

		serviceDeployment := &servicesv1.ServiceDeployment{}
		if err := req.ReadEntity(serviceDeployment); err != nil {
			return err
		}

		if serviceDeployment.Annotations[kore.AnnotationReadOnly] != "" {
			writeError(req, resp, fmt.Errorf("setting %q annotation is not allowed", kore.AnnotationReadOnly), http.StatusForbidden)
			return nil
		}

		user := authentication.MustGetIdentity(req.Request.Context())

		if err := u.Teams().Team(team).ServiceDeployments().Update(req.Request.Context(), serviceDeployment, user.IsGlobalAdmin()); err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, serviceDeployment)
	})
}

// deleteServiceDeployment is responsible for deleting a service deployment from a team
func (u teamHandler) deleteServiceDeployment(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		ctx := req.Request.Context()
		name := req.PathParameter("name")
		team := req.PathParameter("team")

		object, err := u.Teams().Team(team).ServiceDeployments().Delete(ctx, name, parseDeleteOpts(req)...)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, object)
	})
}
