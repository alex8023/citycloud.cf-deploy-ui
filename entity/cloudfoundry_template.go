package entity

import (
	utils "github.com/citycloud/citycloud.cf-deploy-ui/utils"
)
const CloudFoundryTemplateText string = `
---
name: {{.Name}}


`

type CloudFoundryTemplate struct {
	cloudfoundry CloudFoundry
}

func NewCloudFoundryTemplate(cloudfoundry CloudFoundry)(cft CloudFoundryTemplate){
	cft.cloudfoundry = cloudfoundry
	return
}

func (cft CloudFoundryTemplate)CreateCloudFoundryYaml(path string) (bool,error) {
	return utils.CreateYmlFile("cloudfoundry",CloudFoundryTemplateText,path,cft.cloudfoundry)
}