package upgrade

import (
	"github.com/konflux-ci/e2e-tests/pkg/framework"
	"github.com/konflux-ci/e2e-tests/tests/upgrade/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = framework.UpgradeSuiteDescribe("Create users and check their state", Label("upgrade-cleanup"), func() {
	defer GinkgoRecover()

	var fw *framework.Framework

	BeforeAll(func() {
		fw, _ = utils.PrepareForUpgradeTests()
	})

	It("Delete AppStudioProvisionedUser", func() {
		_, err := fw.SandboxController.DeleteUserSignup(utils.AppStudioProvisionedUser)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Delete AppStudioDeactivatedUser", func() {
		_, err := fw.SandboxController.DeleteUserSignup(utils.DeactivatedUser)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Delete AppStudioBannedUser", func() {
		_, err := fw.SandboxController.DeleteUserSignup(utils.BannedUser)
		Expect(err).NotTo(HaveOccurred())
	})

})
