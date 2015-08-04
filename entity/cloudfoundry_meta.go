package entity

const (
	CloudFoundryMetaTemplateV2 string = `{{define "meta"}}
---
name: {{.CloudFoundryProperties.Name}}
director_uuid: {{.CloudFoundryProperties.Uuid}}
releases:
- name: {{.Release.Name}}
  version: {{.Release.Version}}
meta:
  environment: {{.CloudFoundryProperties.Name}}
  stemcell:
    name: {{.Stemcells.Name}}
    version: {{.Stemcells.Version}}
  openstack:
    auth_url: {{.CloudFoundryProperties.CloudProperties.AuthUrl}}
    tenant: {{.CloudFoundryProperties.CloudProperties.Tenant}}
    username: {{.CloudFoundryProperties.CloudProperties.UserName}}
    api_key: {{.CloudFoundryProperties.CloudProperties.ApiKey}}
    security_groups: ["default"]
    default_key_name: {{.CloudFoundryProperties.CloudProperties.DefaultKeyName}}
    private_key: {{.CloudFoundryProperties.CloudProperties.PrivateKey}}
    cci_ebs_url: {{.CloudFoundryProperties.CloudProperties.CciEbsUrl}}
    cci_ebs_accesskey: {{.CloudFoundryProperties.CloudProperties.CciEbsAccesskey}}
    cci_ebs_secretkey: {{.CloudFoundryProperties.CloudProperties.CciEbsSecretkey}}
    cci_connection_options: 
    cci_ebs_api_key: 
    cci_ebs_username:
{{end}}`

	CloudFoundryMetaTemplateV3 string = `
{{define "meta"}}---
name: {{.CloudFoundryProperties.Name}}
director_uuid: {{.CloudFoundryProperties.Uuid}}
releases:
- name: {{.Release.Name}}
  version: {{.Release.Version}}
meta:
  environment: {{.CloudFoundryProperties.Name}}
  stemcell:
    name: {{.Stemcells.Name}}
    version: {{.Stemcells.Version}}
  openstack:
    auth_url: {{.CloudFoundryProperties.CloudProperties.AuthUrl}}
    tenant: {{.CloudFoundryProperties.CloudProperties.Tenant}}
    username: {{.CloudFoundryProperties.CloudProperties.UserName}}
    api_key: {{.CloudFoundryProperties.CloudProperties.ApiKey}}
    security_groups: ["default"]
    default_key_name: {{.CloudFoundryProperties.CloudProperties.DefaultKeyName}}
    private_key: {{.CloudFoundryProperties.CloudProperties.PrivateKey}}
{{end}}`
)
