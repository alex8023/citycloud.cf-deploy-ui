package entity_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("microbosh template", func() {

		var test = NewMicroBOSH("deployment-microbosh",
			NewNetWork("vip", "netid"),
			NewResources("16384", "flavor_100", "zone2"),
			NewCloudProperties("auth_url", "username", "apikey", "tenant", "defaultkeyname", "privatekey", "ebsurl", "ebskey", "ebssercetkey"))

		var mt = NewMicroBOSHTemplate(test)
		_, err := mt.CreateMicroBOSHYaml(MicroBOSHTemplateTextV2, "/home/ubuntu/temp/microbosh.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
})
