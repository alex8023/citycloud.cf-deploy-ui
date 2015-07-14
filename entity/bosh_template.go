package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

const BOSHTemplateText string = `
---
name: {{.Name}}

`

type BOSHTemplate struct {
	bosh BOSH
}

func NewBOSHTemplate(bosh BOSH) (bt BOSHTemplate) {
	bt.bosh = bosh
	return
}

func (bt BOSHTemplate) CreateBOSHYaml(path string) (bool, error) {
	logger.Debug("Create BOSH deployment file : %s", path)
	return utils.CreateYmlFile("bosh", BOSHTemplateText, path, bt.bosh)
}
