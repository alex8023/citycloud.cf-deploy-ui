package entity

const (
	CloudFoundryCompilationTemplate string = `{{define "compilation"}}{{with .Compilation}}
compilation:
  cloud_properties:
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
    instance_type: {{.InstanceType}}
  network: {{.DefaultNetWork}}
  reuse_compilation_vms: true
  workers: {{.Workers}}
{{end}}{{end}}`
)
