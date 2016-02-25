package runner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runner Suite")
}

var _ = BeforeSuite(func() {
	var err error
	pathToFakeProcess, err = gexec.Build("github.com/zaksoup/yoloist/runner/fakes/process")
	Expect(err).NotTo(HaveOccurred())
})
