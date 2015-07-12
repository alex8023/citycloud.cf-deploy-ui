package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
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
	logger.Debug("Create CloudFoundry deployment file : %s",path)
	return utils.CreateYmlFile("cloudfoundry",CloudFoundryTemplateText,path,cft.cloudfoundry)
}