package entity

import (
	utils "github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

const micro_bosh string = `
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
      default_security_groups: [bosh]
	  cci_ebs_url: {{.CciEbsUrl}}
	  cci_ebs_accesskey: {{.CciEbsAccesskey}}
	  cci_ebs_secretkey: {{.CciEbsSecretkey}}
{{end}}

`
type MicroBOSHTemplate struct {
	microbosh MicroBOSH
}

func NewMicroBOSHTemplate(microbosh MicroBOSH)(mt MicroBOSHTemplate){
	mt.microbosh = microbosh
	return
}

func (mt MicroBOSHTemplate)CreateMicroBOSHYaml(path string) (bool,error) {
	return utils.CreateYmlFile("microbosh",micro_bosh,path,mt.microbosh)
}

