package runner_test

import (
	"github.com/zaksoup/yoloist/runner"
	"encoding/json"

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
type outputData struct {
	Args []string
}

var _ = Describe("Runner", func() {
	var (
		cmdRunner  runner.Runner
		fakeParams []string
		expectedOutput outputData
		err error
	)

	BeforeEach(func() {
		fakeParams = []string{
			"param1",
			"param2",
			"--option1",
			"--option2",
		}

		expectedOutput = outputData{
			Args: []string{
				"param1",
				"param2",
				"--option1",
				"--option2",
			},
		}

		cmdRunner = runner.Runner{
			Path: pathToFakeProcess,
		}
	})

	Describe("Run", func() {
		It("calls the process with args", func() {
			Expect(cmdRunner.Run()).To(Succeed())

			var output outputData
			bytes := cmdRunner.Stdout

			err = json.Unmarshal(bytes, &output)
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(Equal(expectedOutput))
		})
	})
})
