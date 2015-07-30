package entity

const (
	CloudFoundryResourcePoolTemplate string = `{{define "resourcepool"}}{{with .ResourcesPools}}
resource_pools:{{range .}}
- cloud_properties:
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
    instance_type: {{.InstanceType}}
  name: {{.Name}}
  network: {{.DefaultNetWork}}
  size: {{.Size}}
  stemcell:
    name: bosh-openstack-kvm-ubuntu-lucid-go_agent
    version: 2719
{{end}}{{end}}{{end}}`
)
