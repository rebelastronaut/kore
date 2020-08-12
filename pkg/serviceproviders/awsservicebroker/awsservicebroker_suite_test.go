package awsservicebroker_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAwsservicebroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Awsservicebroker Suite")
}
