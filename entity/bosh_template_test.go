package entity_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("b o s h template", func() {

		var test = NewSimpleBOSH("deployment-bosh")
		var bt = NewBOSHTemplate(test)
		_, err := bt.CreateBOSHYaml("/home/ubuntu/temp/bosh.yml")

		if err != nil {
			GinkgoT().Error(err)
		}
	})
})
