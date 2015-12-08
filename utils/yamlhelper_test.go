package utils_test

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"

	"fmt"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Yamlhelper", func() {
	It("Test read key", func() {
		mapp, err := utils.ReadYamlFile("/home/ubuntu/deploy/example/manifest.yml")
		str, err := utils.GetValuesByKey(mapp, "name")
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range str {
			fmt.Println(v)
		}
	})
})
