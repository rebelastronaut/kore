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
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/appvia/kore/pkg/apiserver/params"

	"github.com/appvia/kore/pkg/apiserver/filters"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	eks "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	gcp "github.com/appvia/kore/pkg/apis/gcp/v1alpha1"
	gke "github.com/appvia/kore/pkg/apis/gke/v1alpha1"
	orgv1 "github.com/appvia/kore/pkg/apis/org/v1"
	securityv1 "github.com/appvia/kore/pkg/apis/security/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/apiserver/types"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/kore/assets"
	"github.com/appvia/kore/pkg/utils"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterHandler(&teamHandler{})
}

type teamHandler struct {
	kore.Interface
	// DefaultHandler implements default features
	DefaultHandler
}

// Register is called by the api server to register the service
func (u *teamHandler) Register(i kore.Interface, builder utils.PathBuilder) (*restful.WebService, error) {
	u.Interface = i
	path := builder.Path("teams")

	log.WithFields(log.Fields{
		"path": path,
	}).Info("registering the teams webservice with container")

	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path(path)

	ws.Filter(func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		team := req.PathParameter("team")
		if team == "" {
			chain.ProcessFilter(req, resp)
			return
		}

		// Team resource endpoints do this check themselves
		if strings.HasSuffix(req.Request.RequestURI, fmt.Sprintf("teams/%s", team)) {
			chain.ProcessFilter(req, resp)
			return
		}

		exists, err := u.Teams().Exists(context.Background(), team)
		if err != nil {
			handleError(req, resp, err)
			return
		}
		if !exists {
			writeError(req, resp, fmt.Errorf("team %q does not exist", team), http.StatusNotFound)
			return
		}

		chain.ProcessFilter(req, resp)
	})

	ws.Route(
		ws.PUT("/invitation/{token}").To(u.invitationSubmit).
			Doc("Used to verify and handle the team invitation generated links").
			Operation("InvitationSubmit").
			Param(ws.PathParameter("token", "The generated base64 invitation token which was provided from the team")).
			// As there's no body, need to explicitly say we consume any MIME type. Arguably a go-restful bug:
			Consumes(restful.MIME_JSON, "*/*").
			Returns(http.StatusOK, "Indicates the generated link is valid and the user has been granted access", types.TeamInvitationResponse{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("").To(u.listTeams).
			Doc("Returns all the teams in the kore").
			Operation("ListTeams").
			Returns(http.StatusOK, "A list of all the teams in the kore", orgv1.TeamList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}").To(u.findTeam).
			Doc("Return information related to the specific team in the kore").
			Operation("GetTeam").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the team definition from the kore", orgv1.Team{}).
			Returns(http.StatusNotFound, "Team does not exist", nil).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		withAllErrors(ws.PUT("/{team}")).To(u.updateTeam).
			Doc("Used to create or update a team in the kore").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Operation("UpdateTeam").
			Reads(orgv1.Team{}, "Contains the definition for a team in the kore").
			Returns(http.StatusOK, "Contains the team definition from the kore", orgv1.Team{}).
			Returns(http.StatusNotFound, "Team does not exist", nil).
			Returns(http.StatusNotModified, "Indicates the request was processed but no changes applied", orgv1.Team{}),
	)

	ws.Route(
		ws.DELETE("/{team}").To(u.deleteTeam).
			Doc("Used to delete a team from the kore").
			Operation("RemoveTeam").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(params.DeleteCascade()).
			Returns(http.StatusOK, "Contains the former team definition from the kore", orgv1.Team{}).
			Returns(http.StatusNotFound, "Team does not exist", nil).
			Returns(http.StatusNotAcceptable, "Indicates you cannot delete the team for one or more reasons", Error{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Team Audit Events

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/audits")).To(u.findTeamAudit).
			Doc("Used to return a collection of events against the team").
			Operation("ListTeamAudit").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.QueryParameter("since", "The duration to retrieve from the audit log").DefaultValue("60m")).
			Returns(http.StatusOK, "A collection of audit events against the team", orgv1.AuditEventList{}).
			Returns(http.StatusNotFound, "Team does not exist", nil),
	)

	// Team Members

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/members")).To(u.findTeamMembers).
			Doc("Returns a list of user memberships in the team").
			Operation("ListTeamMembers").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains a collection of team memberships for this team", List{}).
			Returns(http.StatusNotFound, "Team does not exist", nil),
	)

	ws.Route(
		withAllErrors(ws.PUT("/{team}/members/{user}")).To(u.addTeamMember).
			Doc("Used to add a user to the team via membership").
			Operation("AddTeamMember").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("user", "Is the user you are adding to the team")).
			// As there's no body, need to explicitly say we consume any MIME type. Arguably a go-restful bug:
			Consumes(restful.MIME_JSON, "*/*").
			Returns(http.StatusOK, "The user has been successfully added to the team", orgv1.TeamMember{}).
			Returns(http.StatusNotFound, "Team does not exist", nil),
	)

	ws.Route(
		ws.DELETE("/{team}/members/{user}").To(u.removeTeamMember).
			Doc("Used to remove team membership from the team").
			Operation("RemoveTeamMember").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("user", "Is the user you are removing from the team")).
			Returns(http.StatusOK, "The user has been successfully removed from the team", orgv1.TeamMember{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Team Invitations

	ws.Route(
		ws.GET("/{team}/invites/user").To(u.listInvites).
			Doc("Used to return a list of all the users whom have pending invitations").
			Operation("ListInvites").
			Param(ws.PathParameter("team", "The name of the team you are pulling the invitations for")).
			Returns(http.StatusOK, "A list of users whom have an invitation for the team", orgv1.TeamInvitationList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/invites/user/{user}").To(u.inviteUser).
			Doc("Used to create an invitation for the team").
			Operation("InviteUser").
			Param(ws.PathParameter("team", "The name of the team you are creating an invition")).
			Param(ws.PathParameter("user", "The name of the username of the user the invitation is for")).
			Param(ws.QueryParameter("expire", "The expiration of the generated link").DefaultValue("1h")).
			Returns(http.StatusOK, "Indicates the team invitation for the user has been successful", nil).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/invites/user/{user}").To(u.removeInvite).
			Doc("Used to remove a user invitation for the team").
			Operation("RemoveInvite").
			Param(ws.PathParameter("team", "The name of the team you are deleting the invitation")).
			Param(ws.PathParameter("user", "The username of the user whos invitation you are removing")).
			Returns(http.StatusOK, "Indicates the team invitation has been successful removed", nil).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Invitation Links

	ws.Route(
		ws.GET("/{team}/invites/generate").To(u.inviteLink).
			Doc("Used to generate a link which provides automatic membership of the team").
			Operation("GenerateInviteLink").
			Param(ws.PathParameter("team", "The name of the team you are creating an invition link")).
			Param(ws.QueryParameter("expire", "The expiration of the generated link").DefaultValue("1h")).
			Returns(http.StatusOK, "A generated URI which can be used to join a team", "").
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/invites/generate/{user}").To(u.inviteLinkByUser).
			Doc("Used to generate for a specific user to join a team").
			Operation("GenerateInviteLinkForUser").
			Param(ws.PathParameter("team", "The name of the team you are creating an invition link")).
			Param(ws.PathParameter("user", "The username of the user the link should be limited for")).
			Returns(http.StatusOK, "A generated URI which users can use to join the team", "").
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Team Allocations

	ws.Route(
		ws.GET("/{team}/allocations").To(u.findAllocations).
			Doc("Used to return a list of all the allocations in the team").
			Operation("ListAllocations").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.QueryParameter("assigned", "Retrieves all allocations which have been assigned to you")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.AllocationList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.GET("/{team}/allocations/{name}").To(u.findAllocation).
			Doc("Used to return an allocation within the team").
			Operation("GetAllocation").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the allocation you wish to return")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.Allocation{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.PUT("/{team}/allocations/{name}").To(u.updateAllocation).
			Doc("Used to create/update an allocation within the team.").
			Operation("UpdateAllocation").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the allocation you wish to update")).
			Reads(configv1.Allocation{}, "The definition of the Allocation").
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.Allocation{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.DELETE("/{team}/allocations/{name}").To(u.deleteAllocation).
			Doc("Remove an allocation from a team").
			Operation("RemoveAllocation").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the allocation you wish to delete")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.Allocation{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Team Namespaces

	ws.Route(
		ws.GET("/{team}/namespaceclaims").To(u.findNamespaces).
			Doc("Used to return all namespaces for the team").
			Operation("ListNamespaces").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former definition from the kore", clustersv1.NamespaceClaimList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/namespaceclaims/{name}").To(u.findNamespace).
			Doc("Used to return the details of a namespace within a team").
			Operation("GetNamespace").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the namespace claim you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.NamespaceClaim{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/namespaceclaims/{name}").To(u.updateNamespace).
			Doc("Used to create or update the details of a namespace within a team").
			Operation("UpdateNamespace").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the namespace claim you are acting upon")).
			Reads(clustersv1.NamespaceClaim{}, "The definition for namespace claim").
			Returns(http.StatusOK, "Contains the definition from the kore", clustersv1.NamespaceClaim{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/namespaceclaims/{name}").To(u.deleteNamespace).
			Doc("Used to remove a namespace from a team").
			Operation("RemoveNamespace").
			Param(ws.PathParameter("name", "Is name the of the namespace claim you are acting upon")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(params.DeleteCascade()).
			Returns(http.StatusOK, "Contains the former definition from the kore", clustersv1.NamespaceClaim{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Secrets is used to provision a secret in the team

	ws.Route(
		ws.GET("/{team}/secrets").To(u.findTeamSecrets).
			Operation("ListTeamSecrets").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Used to return all the secrets within the team").
			Returns(http.StatusOK, "Contains the definition for the resource", configv1.Secret{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/secrets/{name}").To(u.findTeamSecret).
			Operation("GetTeamSecret").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the secert in the name")).
			Doc("Used to retrieve the secret from the team").
			Returns(http.StatusOK, "Contains the definition for the resource", configv1.Secret{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/secrets/{name}").To(u.updateTeamSecret).
			Operation("UpdateTeamSecret").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of secret you are creating / updating")).
			Doc("Used to update the secret in the team").
			Reads(configv1.Secret{}, "The definition for the secret you are creating or updating").
			Returns(http.StatusOK, "Contains updated definition of the secret", configv1.Secret{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/secrets/{name}").To(u.deleteTeamSecret).
			Operation("DeleteTeamSecret").
			Param(ws.PathParameter("name", "Is name the of the secret you are acting upon")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Used to delete the secret from team").
			Returns(http.StatusOK, "Contains the former definition of the secret", configv1.Secret{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Team Kubernetes clusters

	ws.Route(
		ws.GET("/{team}/kubernetes").To(u.listKubernetes).
			Doc("Lists all Kubernetes objects available for a team").
			Operation("ListKubernetes").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.KubernetesList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/kubernetes/{name}").To(u.getKubernetes).
			Doc("returns a specific Kubernetes object").
			Operation("GetKubernetes").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes object you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.Kubernetes{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// Team Clusters

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/clusters")).To(u.listClusters).
			Doc("Lists all clusters for a team").
			Operation("ListClusters").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "List of all clusters for a team", clustersv1.ClusterList{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/clusters/{name}")).To(u.getCluster).
			Doc("Returns a cluster").
			Operation("GetCluster").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes cluster you are acting upon")).
			Returns(http.StatusNotFound, "the cluster with the given name doesn't exist", nil).
			Returns(http.StatusOK, "The requested cluster details", clustersv1.Cluster{}),
	)
	ws.Route(
		withAllErrors(ws.PUT("/{team}/clusters/{name}")).To(u.updateCluster).
			Doc("Creates or updates a cluster").
			Operation("UpdateCluster").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the cluster")).
			Reads(clustersv1.Cluster{}, "The definition for kubernetes cluster").
			Returns(http.StatusOK, "The cluster details", clustersv1.Cluster{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.DELETE("/{team}/clusters/{name}")).To(u.deleteCluster).
			Doc("Deletes a cluster").
			Operation("RemoveCluster").
			Param(ws.PathParameter("name", "Is the name of the cluster")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(params.DeleteCascade()).
			Returns(http.StatusNotFound, "the cluster with the given name doesn't exist", nil).
			Returns(http.StatusOK, "Contains the former cluster definition from the kore", clustersv1.Cluster{}),
	)

	// Team Cloud Providers

	// GKE Clusters

	ws.Route(
		ws.GET("/{team}/gkes").To(u.findGKEs).
			Doc("Returns a list of Google Container Engine clusters which the team has access").
			Operation("ListGKEs").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKEList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/gkes/{name}").To(u.findGKE).
			Doc("Returns a specific Google Container Engine cluster to which the team has access").
			Operation("GetGKE").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKE{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// GKE Credentials - @TODO these all need to be autogenerated

	ws.Route(
		ws.GET("/{team}/gkecredentials").To(u.findGKECredientalss).
			Doc("Returns a list of GKE Credentials to which the team has access").
			Operation("ListGKECredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentialsList{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/gkecredentials/{name}").To(u.findGKECredientals).
			Doc("Returns a specific GKE Credential to which the team has access").
			Operation("GetGKECredential").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentials{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/gkecredentials/{name}").To(u.updateGKECredientals).
			Doc("Creates or updates a specific GKE Credential to which the team has access").
			Operation("UpdateGKECredential").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Reads(gke.GKECredentials{}, "The definition for GKE Credentials").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentials{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/gkecredentials/{name}").To(u.deleteGKECredientals).
			Doc("Deletes a specific GKE Credential from the team").
			Operation("RemoveGKECredential").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentials{}).
			Returns(http.StatusInternalServerError, "A generic API error containing the cause of the error", Error{}),
	)

	// GCP Project Claims

	ws.Route(
		ws.GET("/{team}/projectclaims").To(u.findProjectClaims).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Is the used tor return a list of Google Container Engine clusters which thhe team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gcp.ProjectClaimList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/projectclaims/{name}").To(u.findProjectClaim).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the resource you are acting on")).
			Doc("Is the used tor return a list of Google Container Engine clusters which thhe team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gcp.ProjectClaim{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// GCP Organization

	ws.Route(
		ws.GET("/{team}/organizations").To(u.findOrganizations).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Is the used tor return a list of gcp organizations").
			Operation("ListGCPOrganizations").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gcp.OrganizationList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/organizations/{name}").To(u.findOrganization).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the resource you are acting on")).
			Doc("Is the used tor return a specific gcp organization").
			Operation("GetGCPOrganization").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gcp.Organization{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.PUT("/{team}/organizations/{name}").To(u.updateOrganization).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the resource you are acting on")).
			Operation("UpdateGCPOrganization").
			Reads(gcp.Organization{}, "The definition for GCP organization").
			Doc("Is used to provision or update a gcp organization").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gcp.Organization{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/organizations/{name}").To(u.deleteOrganization).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the resource you are acting on")).
			Doc("Is used to delete a managed gcp organization").
			Operation("DeleteGCPOrganization").
			Returns(http.StatusOK, "Contains the former team definition from the kore", gcp.Organization{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// EKS clusters

	ws.Route(
		ws.GET("/{team}/eks").To(u.findEKSs).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Is the used to return a list of Amazon EKS clusters which thhe team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/eks/{name}").To(u.findEKS).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS cluster you are acting upon")).
			Doc("Is the used to return a EKS cluster which the team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKS{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// EKS Nodegroups
	ws.Route(
		ws.GET("/{team}/eksnodegroups").To(u.findEKSNodeGroups).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Is the used tor return a list of Amazon EKS clusters which the team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSNodeGroupList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/eksnodegroups/{name}").To(u.findEKSNodeGroup).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the EKS nodegroup")).
			Doc("Is the used to return a EKS cluster which the team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSNodeGroup{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// EKS Credentials - @TODO these all need to be autogenerated

	ws.Route(
		ws.GET("/{team}/ekscredentials").To(u.listEKSCredentials).
			Operation("ListEKSCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Is the used tor return a list of Amazon EKS credentials which thhe team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSCredentialsList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/ekscredentials/{name}").To(u.getEKSCredentials).
			Operation("GetEKSCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS Credentials you are acting upon")).
			Doc("Is the used tor return a list of EKS Credentials which the team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSCredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/ekscredentials/{name}").To(u.updateEKSCredentials).
			Operation("UpdateEKSCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS credentials you are acting upon")).
			Reads(eks.EKSCredentials{}, "The definition for EKS Credentials").
			Doc("Is used to provision or update a EKS credentials in the kore").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSCredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/ekscredentials/{name}").To(u.deleteEKSCredentials).
			Operation("DeleteEKSCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS credentials you are acting upon")).
			Doc("Is used to delete a EKS credentials from the kore").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSCredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// EKSVPC EKS dependencies

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/eksvpcs").To(u.findEKSVPCs).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Is the used to return a list of Amazon EKS VPC which thhe team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSVPCList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
		),
	)

	ws.Route(
		ws.GET("/{team}/eksvpcs/{name}").To(u.findEKSVPC).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS VPC you are acting upon")).
			Doc("Is the used to return a EKS VPC which the team has access").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSVPC{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/eksvpcs/{name}").To(u.updateEKSVPC).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS VPC you are acting upon")).
			Doc("Is used to provision or update a EKS VPC in the kore").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSVPC{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/eksvpcs/{name}").To(u.deleteEKSVPC).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the EKS VPC you are acting upon")).
			Doc("Is used to delete a EKS VPC from the kore").
			Returns(http.StatusOK, "Contains the former team definition from the kore", eks.EKSVPC{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team-specific plan details
	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/plans/{plan}")).To(u.getTeamPlanDetails).
			Operation("GetTeamPlanDetails").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("plan", "Is name the of the plan you're interested in")).
			Doc("Returns the plan, the JSON schema of the plan, and what what parameters are allowed to be edited by this team when using the plan").
			Returns(http.StatusOK, "Contains details of the plan", TeamPlan{}).
			Returns(http.StatusNotFound, "Team or plan doesn't exist", nil),
	)

	// Team services

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/services")).To(u.listServices).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Lists all services for a team").
			Operation("ListServices").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "List of all services for a team", servicesv1.ServiceList{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/services/{name}")).To(u.getService).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Returns a service").
			Operation("GetService").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name of the service")).
			Returns(http.StatusNotFound, "the service with the given name doesn't exist", nil).
			Returns(http.StatusOK, "The requested service details", servicesv1.Service{}),
	)
	ws.Route(
		withAllErrors(ws.PUT("/{team}/services/{name}")).To(u.updateService).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Filter(u.readonlyServiceFilter).
			Doc("Creates or updates a service").
			Operation("UpdateService").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the service")).
			Reads(servicesv1.Service{}, "The definition for the service").
			Returns(http.StatusOK, "The service details", servicesv1.Service{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.DELETE("/{team}/services/{name}")).To(u.deleteService).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Filter(u.readonlyServiceFilter).
			Doc("Deletes a service").
			Operation("DeleteService").
			Param(ws.PathParameter("name", "Is the name of the service")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(params.DeleteCascade()).
			Returns(http.StatusNotFound, "the service with the given name doesn't exist", nil).
			Returns(http.StatusOK, "Contains the former service definition from the kore", servicesv1.Service{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/servicecredentials")).To(u.listServiceCredentials).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Lists all service credentials for a team").
			Operation("ListServiceCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.QueryParameter("cluster", "Is the name of the cluster you are filtering for")).
			Param(ws.QueryParameter("service", "Is the name of the service you are filtering for")).
			Returns(http.StatusOK, "List of all service credentials for a team", servicesv1.ServiceCredentials{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/servicecredentials/{name}")).To(u.getServiceCredentials).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Returns the requsted service credentials").
			Operation("GetServiceCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name of the service credentials")).
			Returns(http.StatusNotFound, "the service credentials with the given name doesn't exist", nil).
			Returns(http.StatusOK, "The requested service crendential details", servicesv1.ServiceCredentials{}),
	)

	ws.Route(
		withAllErrors(ws.PUT("/{team}/servicecredentials/{name}")).To(u.updateServiceCredentials).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Creates or updates service credentials").
			Operation("UpdateServiceCredentials").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the service credentials")).
			Reads(servicesv1.ServiceCredentials{}, "The definition for the service credentials").
			Returns(http.StatusOK, "The service credentail details", servicesv1.ServiceCredentials{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.DELETE("/{team}/servicecredentials/{name}")).To(u.deleteServiceCredentials).
			Filter(filters.FeatureGateFilter(u.Config(), kore.FeatureGateServices)).
			Doc("Deletes the given service credentials").
			Operation("DeleteServiceCredentials").
			Param(ws.PathParameter("name", "Is the name of the service credentials")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(params.DeleteCascade()).
			Returns(http.StatusNotFound, "the service credentials with the given name doesn't exist", nil).
			Returns(http.StatusOK, "Contains the former service credentials definition", servicesv1.ServiceCredentials{}),
	)

	// Team Security
	ws.Route(
		withAllNonValidationErrors(ws.GET("/{team}/security")).To(u.getTeamSecurityOverview).
			Doc("Returns an overview of the security posture for resources owned by this team").
			Operation("GetTeamSecurityOverview").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "The requested security overview", securityv1.SecurityOverview{}),
	)

	return ws, nil
}

// Name returns the name of the handler
func (u teamHandler) Name() string {
	return "teams"
}

// findTeamAudit returns the audit log for a team
func (u teamHandler) findTeamAudit(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team := req.PathParameter("team")

		since := req.QueryParameter("since")
		if since == "" {
			since = "60m"
		}
		tm, err := time.ParseDuration(since)
		if err != nil {
			return err
		}

		list, err := u.Audit().AuditEventsTeam(req.Request.Context(), team, tm)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, list)
	})
}

// Teams Management

// deleteTeam is responsible for deleting a team from the kore
func (u teamHandler) deleteTeam(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		err := u.Teams().Delete(req.Request.Context(), req.PathParameter("team"), parseDeleteOpts(req)...)
		if err != nil {
			return err
		}
		resp.WriteHeader(http.StatusOK)

		return nil
	})
}

// findTeam returns a specific team
func (u teamHandler) findTeam(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team, err := u.Teams().Get(req.Request.Context(), req.PathParameter("team"))
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, team)
	})
}

// listTeams returns all the teams in the kore
func (u teamHandler) listTeams(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		teams, err := u.Teams().List(req.Request.Context())
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, teams)
	})
}

// updateTeam is responsible for updating for creating a team in the kore
func (u teamHandler) updateTeam(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team := &orgv1.Team{}
		if err := req.ReadEntity(team); err != nil {
			return err
		}
		team, err := u.Teams().Update(req.Request.Context(), team)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, team)
	})
}

func (u teamHandler) getTeamPlanDetails(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		plan, err := u.Plans().Get(req.Request.Context(), req.PathParameter("plan"))
		if err != nil {
			return err
		}

		schema, err := assets.GetClusterSchema(plan.Spec.Kind)
		if err != nil {
			writeError(req, resp, err, http.StatusNotFound)
			return nil
		}

		editableParams, err := u.Plans().GetEditablePlanParams(req.Request.Context(), req.PathParameter("team"), plan.Spec.Kind)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, TeamPlan{
			Plan:           plan.Spec,
			Schema:         schema,
			EditableParams: editableParams,
		})
	})
}
