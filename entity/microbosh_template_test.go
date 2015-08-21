package entity_test

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("microbosh template", func() {

		var test = entity.NewMicroBOSH("deployment-microbosh",
			entity.NewNetWork("vip", "netid"),
			entity.NewResources("16384", "flavor_100", "zone2"),
			entity.NewCloudProperties("auth_url", "username", "apikey", "tenant", "defaultkeyname", "privatekey", "ebsurl", "ebskey", "ebssercetkey"))

		var mt = entity.NewMicroBOSHTemplate(test)
		_, err := mt.CreateMicroBOSHYaml(entity.MicroBOSHTemplateTextV2, "/home/ubuntu/temp/microbosh.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
})
