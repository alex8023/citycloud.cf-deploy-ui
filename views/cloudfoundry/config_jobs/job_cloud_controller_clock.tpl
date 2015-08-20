{{with .CloudFoundryJobs.cloud_controller_clock }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}