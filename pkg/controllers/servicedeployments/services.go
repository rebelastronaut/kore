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

package servicedeployments

import (
	"fmt"

	"github.com/appvia/kore/pkg/controllers/helpers"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	orgv1 "github.com/appvia/kore/pkg/apis/org/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type servicesComponent struct {
	serviceDeployment *servicesv1.ServiceDeployment
}

func newServicesComponent(serviceDeployment *servicesv1.ServiceDeployment) *servicesComponent {
	return &servicesComponent{
		serviceDeployment: serviceDeployment,
	}
}

func (c *servicesComponent) Reconcile(ctx kore.Context) (reconcile.Result, error) {
	plan, err := ctx.Kore().Plans().Get(ctx, c.serviceDeployment.Spec.Plan)
	if err != nil {
		return reconcile.Result{}, err
	}

	clusterSelector := c.serviceDeployment.Spec.ClusterSelector

	var teamFilters []func(orgv1.Team) bool

	if len(clusterSelector.Teams) > 0 {
		teamFilters = append(teamFilters, func(t orgv1.Team) bool {
			return utils.Contains(t.Name, clusterSelector.Teams)
		})
	}

	var clusterFilters []func(cluster clustersv1.Cluster) bool

	if len(clusterSelector.Kinds) > 0 {
		clusterFilters = append(clusterFilters, func(c clustersv1.Cluster) bool {
			return utils.Contains(c.Spec.Kind, clusterSelector.Kinds)
		})
	}

	if len(clusterSelector.MatchLabels) > 0 || len(clusterSelector.MatchExpressions) > 0 {
		selector, err := v1.LabelSelectorAsSelector(&clusterSelector.LabelSelector)
		if err != nil {
			return reconcile.Result{}, controllers.NewCriticalError(fmt.Errorf("cluster selector is invalid: %w", err))
		}

		plans, err := ctx.Kore().Plans().List(ctx, func(p configv1.Plan) bool {
			return selector.Matches(labels.Set(p.Labels))
		})
		if err != nil {
			return reconcile.Result{}, fmt.Errorf("failed to list plans: %w", err)
		}

		var planNames []string
		for _, plan := range plans.Items {
			planNames = append(planNames, plan.Name)
		}

		clusterFilters = append(clusterFilters, func(c clustersv1.Cluster) bool {
			return utils.Contains(c.Spec.Plan, planNames) || selector.Matches(labels.Set(c.Labels))
		})
	}

	teams, err := ctx.Kore().Teams().List(ctx, teamFilters...)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to list teams: %w", err)
	}

	var services []servicesv1.Service

	for _, team := range teams.Items {
		clusters, err := ctx.Kore().Teams().Team(team.Name).Clusters().List(ctx, clusterFilters...)
		if err != nil {
			return reconcile.Result{}, fmt.Errorf("failed to list clusters for team %s: %w", team.Name, err)
		}

		for _, cluster := range clusters.Items {
			services = append(services, c.createService(&cluster, plan))
		}
	}

	return helpers.EnsureServices(ctx, services, c.serviceDeployment, &c.serviceDeployment.Status.Components)
}

func (c *servicesComponent) Delete(ctx kore.Context) (reconcile.Result, error) {
	clusterSelector := c.serviceDeployment.Spec.ClusterSelector

	var teamFilters []func(orgv1.Team) bool

	if len(clusterSelector.Teams) > 0 {
		teamFilters = append(teamFilters, func(t orgv1.Team) bool {
			return utils.Contains(t.Name, clusterSelector.Teams)
		})
	}

	teams, err := ctx.Kore().Teams().List(ctx, teamFilters...)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to list teams: %w", err)
	}

	for _, team := range teams.Items {
		res, err := helpers.DeleteServices(ctx, team.Name, c.serviceDeployment, &c.serviceDeployment.Status.Components)
		if err != nil || res.Requeue || res.RequeueAfter > 0 {
			return res, err
		}
	}

	return reconcile.Result{}, nil
}

func (c *servicesComponent) createService(cluster *clustersv1.Cluster, plan *configv1.Plan) servicesv1.Service {
	serviceName := c.serviceDeployment.Spec.ServiceName
	if serviceName == "" {
		serviceName = c.serviceDeployment.Name
	}

	return servicesv1.Service{
		TypeMeta: v1.TypeMeta{
			Kind:       "Service",
			APIVersion: servicesv1.GroupVersion.String(),
		},
		ObjectMeta: v1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", cluster.Name, serviceName),
		},
		Spec: servicesv1.ServiceSpec{
			Kind:              plan.Kind,
			Plan:              plan.Name,
			Cluster:           corev1.MustGetOwnershipFromObject(cluster),
			ClusterNamespace:  c.serviceDeployment.Spec.ClusterNamespace,
			Configuration:     c.serviceDeployment.Spec.Configuration.DeepCopy(),
			ConfigurationFrom: c.serviceDeployment.Spec.ConfigurationFrom.DeepCopy(),
		},
	}
}
