package runner_test

import (
	"github.com/zaksoup/yoloist/runner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
type Runner struct{
	Path string
	Params []string
}
*/

/*
thing = Runner{
	Path: "brew",
	Params: []string,
}
*/

var _ = Describe("Runner", func() {
	var (
		cmdRunner  runner.Runner
		fakeParams []string
	)

	BeforeEach(func() {
		var err error

		fakeParams = []string{
			"param1",
			"param2",
			"--option1",
			"--option2",
		}

		runner = runner.Runner{
			Path: pathToFakeProcess,
		}
	})

	Describe("Run", func() {
		It("calls the process with args", func() {
			Expect(runner.Run()).To(Succeed())

			//Expect stdout to marshal into the struct we think it should be
		})
	})
})
