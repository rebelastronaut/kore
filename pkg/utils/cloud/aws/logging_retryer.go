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

package aws

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
)

const maxRetries = 13

// LoggingRetryer adds some logging when we are retrying, so we have some idea what is happening
// Right now it is very basic - e.g. it only logs when we retry (so doesn't log when we fail due to too many retries)
// It was copied from k8s.io/kops/upup/pkg/fi/cloudup/awsup/logging_retryer.go; the original version used glog, and
// didn't export the constructor
type LoggingRetryer struct {
	client.DefaultRetryer
}

var _ request.Retryer = &LoggingRetryer{}

func newLoggingRetryer() *LoggingRetryer {
	return &LoggingRetryer{
		client.DefaultRetryer{NumMaxRetries: maxRetries},
	}
}

// RetryRules extends on DefaultRetryer.RetryRules
func (l LoggingRetryer) RetryRules(r *request.Request) time.Duration {
	duration := l.DefaultRetryer.RetryRules(r)

	service := r.ClientInfo.ServiceName
	name := "?"
	if r.Operation != nil {
		name = r.Operation.Name
	}
	methodDescription := service + "/" + name

	var errorDescription string
	if r.Error != nil {
		// We could check aws error Code & Message, but we expect them to be in the string
		errorDescription = fmt.Sprintf("%v", r.Error)
	} else {
		errorDescription = fmt.Sprintf("%d %s", r.HTTPResponse.StatusCode, r.HTTPResponse.Status)
	}

	log.Warningf("retryable error (%s) from %s - will retry after delay of %v", errorDescription, methodDescription, duration)

	return duration
}
