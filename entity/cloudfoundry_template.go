package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"os"
	"strings"
)

type CloudFoundryTemplate struct {
	CloudFoundry CloudFoundry
}

func NewCloudFoundryTemplate(cloudfoundry CloudFoundry) (cft CloudFoundryTemplate) {
	cft.CloudFoundry = cloudfoundry
	return
}

func (cft CloudFoundryTemplate) CreateCloudFoundryYaml(template, path string) (bool, error) {
	logger.Debug("Create CloudFoundry deployment file : %s", path)
	return utils.CreateYmlFile("cloudfoundry", template, path, cft.CloudFoundry)
}

func (cft CloudFoundryTemplate) CreateCloudFoundryV3Yaml(path string) (bool, error) {
	logger.Debug("Create CloudFoundry deployment file : %s", path)
	return utils.CreateYmlFile("cloudfoundry", CloudFoundryTemplateV3, path, cft.CloudFoundry)
}

func (cft CloudFoundryTemplate) CreateDefaultCloudFoundryYamlFile(version, path string) (bool, error) {

	return cft.CreateCloudFoundryYamlFile(version, path, workDir+CloudFoundryMetaTemplateV2File, workDir+CloudFoundryUpdateTempalteFile, workDir+CloudFoundryCompilationTemplateFile, workDir+CloudFoundryNetworksTemplateV2File, workDir+CloudFoundryResourcePoolTemplateFile, workDir+CloudFoundryJobsTemplateFile, workDir+CloudFoundryPropertiesTemplateFile)
}

func (cft CloudFoundryTemplate) CreateDefaultCloudFoundryYamlFileWithAppPath(version, path string, appPath string) (bool, error) {
	if !strings.HasSuffix(appPath, "/") {
		appPath = appPath + "/"
	}
	return cft.CreateCloudFoundryYamlFile(version, path, appPath+CloudFoundryMetaTemplateV2File, appPath+CloudFoundryUpdateTempalteFile, appPath+CloudFoundryCompilationTemplateFile, appPath+CloudFoundryNetworksTemplateV2File, appPath+CloudFoundryResourcePoolTemplateFile, appPath+CloudFoundryJobsTemplateFile, appPath+CloudFoundryPropertiesTemplateFile)
}

func (cft CloudFoundryTemplate) CreateCloudFoundryYamlFile(version, path string, template ...string) (bool, error) {
	logger.Debug("Create CloudFoundry deployment file : %s", path)
	if version == "CCI-IaaS3.0" {
		return utils.CreateYmlFileFromFile(path, cft.CloudFoundry, template...)
	}

	return utils.CreateYmlFileFromFile(path, cft.CloudFoundry, template...)

}

func (cft CloudFoundryTemplate) CreateCloudFoundryV3YamlFile(path string) (bool, error) {
	logger.Debug("Create CloudFoundry deployment file : %s", path)
	return utils.CreateYmlFileFromFile(path, cft.CloudFoundry, workDir+CloudFoundryMetaTemplateV3File, workDir+CloudFoundryUpdateTempalteFile, workDir+CloudFoundryCompilationTemplateFile, workDir+CloudFoundryNetworksTemplateV3File, workDir+CloudFoundryResourcePoolTemplateFile, workDir+CloudFoundryJobsTemplateFile, workDir+CloudFoundryPropertiesTemplateFile)

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

	// template file path
	CloudFoundryCompilationTemplateFile  = "yaml/cloudfoundry/compilation.yml"
	CloudFoundryJobsTemplateFile         = "yaml/cloudfoundry/jobs.yml"
	CloudFoundryMetaTemplateV2File       = "yaml/cloudfoundry/meta_v2.yml"
	CloudFoundryMetaTemplateV3File       = "yaml/cloudfoundry/meta_v3.yml"
	CloudFoundryNetworksTemplateV2File   = "yaml/cloudfoundry/networks_v2.yml"
	CloudFoundryNetworksTemplateV3File   = "yaml/cloudfoundry/networks_v3.yml"
	CloudFoundryPropertiesTemplateFile   = "yaml/cloudfoundry/properties.yml"
	CloudFoundryResourcePoolTemplateFile = "yaml/cloudfoundry/resourcepool.yml"
	CloudFoundryUpdateTempalteFile       = "yaml/cloudfoundry/update.yml"
)

func init() {
	workDir, _ = os.Getwd()
	end := strings.LastIndex(workDir, "/")
	rs := []rune(workDir)
	workDir = string(rs[:end+1])
}

var workDir = ""
