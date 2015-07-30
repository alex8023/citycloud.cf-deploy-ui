package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

type CloudFoundryTemplate struct {
	cloudfoundry CloudFoundry
}

func NewCloudFoundryTemplate(cloudfoundry CloudFoundry) (cft CloudFoundryTemplate) {
	cft.cloudfoundry = cloudfoundry
	return
}

func (cft CloudFoundryTemplate) CreateCloudFoundryV2Yaml(template, path string) (bool, error) {
	logger.Debug("Create CloudFoundry deployment file : %s", path)
	return utils.CreateYmlFile("cloudfoundry", template, path, cft.cloudfoundry)
}

func (cft CloudFoundryTemplate) CreateCloudFoundryV3Yaml(path string) (bool, error) {
	logger.Debug("Create CloudFoundry deployment file : %s", path)
	return utils.CreateYmlFile("cloudfoundry", CloudFoundryTemplateTextV33, path, cft.cloudfoundry)
}

const (
	CloudFoundryTemplateTextV2 string = CloudFoundryMetaTemplateV2 +
		CloudFoundryUpdateTempalte +
		CloudFoundryCompilationTemplate +
		CloudFoundryNetworksTemplateV2 +
		CloudFoundryResourcePoolTemplate +
		CloudFoundryJobsTemplate +
		CloudFoundryPropertiesTemplate

	CloudFoundryTemplateTextV3 string = CloudFoundryMetaTemplateV3 +
		CloudFoundryUpdateTempalte +
		CloudFoundryCompilationTemplate +
		CloudFoundryNetworksTemplateV3 +
		CloudFoundryResourcePoolTemplate +
		CloudFoundryJobsTemplate +
		CloudFoundryPropertiesTemplate

	TemplateTemplateText string = `{{template "meta" .}}{{template "update" .}}{{template "compilation" .}}{{template "networks" .}}{{template "resourcepool" .}}{{template "jobs" .}}{{template "properties" .}}`

	CloudFoundryTemplateV2 string = CloudFoundryTemplateTextV2 + TemplateTemplateText
	CloudFoundryTemplateV3 string = CloudFoundryTemplateTextV3 + TemplateTemplateText
)
