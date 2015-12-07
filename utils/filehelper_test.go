package utils_test

import (
	. "github.com/citycloud/citycloud.cf-deploy-ui/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filehelper", func() {
	It("test detech file", func() {
		home, filename, isdir, err := Detechfile("/home/ubuntu/deploy")

		Expect(home).To(Equal("/home/ubuntu"))
		Expect(filename).To(Equal("deploy"))
		Expect(isdir).To(Equal(true))
		Expect(err == nil).To(Equal(true))
	})
})
