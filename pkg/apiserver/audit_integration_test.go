// +build integration

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

package apiserver_test

import (
	"time"

	"github.com/appvia/kore/pkg/apiclient"
	"github.com/appvia/kore/pkg/apiclient/models"
	"github.com/appvia/kore/pkg/apiclient/operations"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func findEvent(list []*models.V1AuditEvent, user string, operation string, after time.Time) *models.V1AuditEvent {

	for _, item := range list {
		evtTime, err := time.Parse(time.RFC3339, item.Spec.CreatedAt)
		Expect(err).ToNot(HaveOccurred())
		if item.Spec.User == user && item.Spec.Operation == operation && evtTime.Unix() > after.Unix() {
			return item
		}
	}
	return nil
}

var _ = Describe("GET /audit (ListAuditEvents)", func() {
	var api *apiclient.Kore

	BeforeEach(func() {
		api = getApi()
	})

	When("called anonymously", func() {
		It("should return 401", func() {
			_, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuthAnon())
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&operations.ListAuditEventsUnauthorized{}))
		})
	})

	When("called as a non-admin", func() {
		It("should return 403", func() {
			_, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuth(TestUserTeam1))
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&operations.ListAuditEventsForbidden{}))
		})
	})

	When("called as admin", func() {
		It("should return a list of audit events", func() {
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			Expect(*&resp.Payload.Items).To(BeAssignableToTypeOf([]*models.V1AuditEvent{}))
		})
	})

	When("called without a since parameter", func() {
		It("should return a list of all audit events from the last 60 minutes by default", func() {
			// Get the audit events for the default period (60m)
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			sinceDuration, err := time.ParseDuration("60m")
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			expectedStartTime := time.Now().Add(-sinceDuration)
			for _, i := range resp.Payload.Items {
				Expect(time.Parse(time.RFC3339, i.Spec.CreatedAt)).To(BeTemporally(">=", expectedStartTime))
			}
		})
	})

	When("called with a since parameter", func() {
		It("should return a list of all audit events for the period specified", func() {
			since := "1m"
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams().WithSince(&since), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			sinceDuration, err := time.ParseDuration(since)
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			expectedStartTime := time.Now().Add(-sinceDuration)
			for _, i := range resp.Payload.Items {
				Expect(time.Parse(time.RFC3339, i.Spec.CreatedAt)).To(BeTemporally(">=", expectedStartTime))
			}
		})
	})

	When("audit events exist", func() {
		It("should include events which are not tied to a specific team", func() {
			// Do any action which should cause an audit event to be raised.
			_, err := api.Operations.ListUsers(operations.NewListUsersParams(), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}

			since := "1m"
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams().WithSince(&since), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}

			// Check items contains the event relating to the above.
			item := findEvent(resp.Payload.Items, getTestUser(TestUserAdmin).Username, "ListUsers", time.Now().Add(-time.Duration(2)*time.Second))
			Expect(item).ToNot(BeNil())
		})

		It("should include events which are tied to a specific team", func() {
			// Do any action which should cause a TEAM audit event to be raised.
			_, err := api.Operations.ListTeamMembers(operations.NewListTeamMembersParams().WithTeam(getTestTeam(TestTeam1).Name), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}

			since := "1m"
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams().WithSince(&since), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}

			// Check items contains the event relating to the above.
			item := findEvent(resp.Payload.Items, getTestUser(TestUserAdmin).Username, "ListTeamMembers", time.Now().Add(time.Duration(-5)*time.Second))
			Expect(item).ToNot(BeNil())
		})
	})
})
