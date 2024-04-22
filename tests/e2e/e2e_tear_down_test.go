package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/openshift/rosa/tests/ci/labels"
	"github.com/openshift/rosa/tests/utils/exec/rosacli"
	PH "github.com/openshift/rosa/tests/utils/profilehandler"
)

var _ = Describe("ROSA CLI Test", func() {
	It("Deprovision cluster",
		labels.Critical,
		labels.Destroy,
		func() {
			client := rosacli.NewClient()
			var errs = PH.DestroyClusterByProfile(client, true)
			Expect(errs).To(BeEmpty())
		})
})
