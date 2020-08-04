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

package allocations

import (
	"context"
	"time"

	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"

	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	orgv1 "github.com/appvia/kore/pkg/apis/org/v1"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type acCtrl struct {
	kore.Interface
	// mgr is the manager
	mgr manager.Manager
	// stopCh is the stop channel
	stopCh chan struct{}
}

func init() {
	controllers.Register(&acCtrl{})
}

// Name returns the name of the controller
func (a acCtrl) Name() string {
	return finalizerName
}

// Run is called when the controller is started
func (a *acCtrl) Run(ctx context.Context, cfg *rest.Config, hi kore.Interface) error {
	a.Interface = hi

	// @step: create the manager for the controller
	mgr, err := manager.New(cfg, controllers.DefaultManagerOptions(a))
	if err != nil {
		log.WithError(err).Error("failed to create the manager")

		return err
	}

	// @step: set the controller manager
	a.mgr = mgr

	// @step: create the controller
	ctrl, err := controller.New(a.Name(), mgr, controllers.DefaultControllerOptions(a))
	if err != nil {
		log.WithError(err).Error("failed to create the controller")

		return err
	}

	// @step: setup watches for the resources
	if err := ctrl.Watch(&source.Kind{Type: &configv1.Allocation{}},
		&handler.EnqueueRequestForObject{},
		&predicate.GenerationChangedPredicate{}); err != nil {

		log.WithError(err).Error("failed to create watcher on resource")

		return err
	}

	// @step: we need to setup a watch for teams and requeue all allocations
	// which as allocated to AllTeams
	err = ctrl.Watch(&source.Kind{Type: &orgv1.Team{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: handler.ToRequestsFunc(func(o handler.MapObject) []reconcile.Request {

			items := &configv1.AllocationList{}
			if err := a.mgr.GetClient().List(ctx, items, client.InNamespace("")); err != nil {
				log.WithError(err).Error("failed to force reconcilation of allocations on team change")

				return []reconcile.Request{}
			}

			// @step: build a request for all allocations which reference all team scope
			requests := make([]reconcile.Request, 0)

			for _, a := range items.Items {
				if utils.Contains(configv1.AllTeams, a.Spec.Teams) {
					requests = append(requests, reconcile.Request{
						NamespacedName: types.NamespacedName{
							Namespace: a.GetNamespace(),
							Name:      a.GetName(),
						},
					})
				}
			}

			return requests
		}),
	})
	if err != nil {
		return err
	}

	go func() {
		log.Info("starting the controller loop")

		for {
			a.stopCh = make(chan struct{})

			if err := mgr.Start(a.stopCh); err != nil {
				log.WithError(err).Error("failed to start the controller")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// @step: use a routine to catch the stop channel
	go func() {
		<-ctx.Done()
		log.WithFields(log.Fields{
			"controller": a.Name(),
		}).Info("stopping the controller")

		close(a.stopCh)
	}()

	return nil
}

// Stop is responsible for calling a halt on the controller
func (a acCtrl) Stop(context.Context) error {
	log.WithFields(log.Fields{
		"controller": a.Name(),
	}).Info("attempting to stop the controller")

	return nil
}
