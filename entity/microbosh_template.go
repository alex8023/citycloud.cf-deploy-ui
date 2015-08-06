package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

const (
	MicroBOSHTemplateTextV2 string = `
---
name: {{.Name}}
{{with .Network}}
network:
  type: manual
  vip: {{.Vip}}
  cloud_properties:
    net_id: {{.NetId}}
{{end}}

{{with .Resources}}
resources:
  persistent_disk: {{.PersistentDisk}}
  cloud_properties:
    instance_type: {{.InstanceType}}
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
{{end}}
{{with .CloudProperties}}
cloud:
  plugin: openstack
  properties:
    openstack:
      auth_url: {{.AuthUrl}}
      tenant: {{.Tenant}}
      username: {{.UserName}}
      api_key: {{.ApiKey}}
      default_key_name: {{.DefaultKeyName}}
      private_key: {{.PrivateKey}}
      default_security_groups: ["default"]
      cci_ebs_url: {{.CciEbsUrl}}
      cci_ebs_accesskey: {{.CciEbsAccesskey}}
      cci_ebs_secretkey: {{.CciEbsSecretkey}}
      cci_connection_options: 
      cci_ebs_api_key: 
      cci_ebs_username:
{{end}}

apply_spec:
  properties:
    director:
      max_threads: 1
    hm:
      resurrector_enabled: true
    ntp: # This example uses the North American NTP servers. Edit for your region.
      - 0.asia.pool.ntp.org
      - 1.asia.pool.ntp.org
`

	MicroBOSHTemplateTextV3 string = `
---
name: {{.Name}}
{{with .Network}}
network:
  type: manual
  vip: {{.Vip}}
  cloud_properties:
    net_id: {{.NetId}}
{{end}}

{{with .Resources}}
resources:
  persistent_disk: {{.PersistentDisk}}
  cloud_properties:
    instance_type: {{.InstanceType}}
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
{{end}}
{{with .CloudProperties}}
cloud:
  plugin: openstack
  properties:
    openstack:
      auth_url: {{.AuthUrl}}
      tenant: {{.Tenant}}
      username: {{.UserName}}
      api_key: {{.ApiKey}}
      default_key_name: {{.DefaultKeyName}}
      private_key: {{.PrivateKey}}
      default_security_groups: ["default"]
{{end}}

apply_spec:
  properties:
    director:
      max_threads: 1
    hm:
      resurrector_enabled: true
    ntp: # This example uses the North American NTP servers. Edit for your region.
      - 0.asia.pool.ntp.org
      - 1.asia.pool.ntp.org
`
)

type MicroBOSHTemplate struct {
	microbosh MicroBOSH
}

func NewMicroBOSHTemplate(microbosh MicroBOSH) (mt MicroBOSHTemplate) {
	mt.microbosh = microbosh
	return
}

func (mt MicroBOSHTemplate) CreateMicroBOSHYaml(template, path string) (bool, error) {
	logger.Debug("Create MicroBOSH deployment file : %s", path)

	return utils.CreateYmlFile("microbosh", template, path, mt.microbosh)
}
