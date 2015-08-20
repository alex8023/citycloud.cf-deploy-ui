{{with .CloudFoundryJobs.cloud_controller_ng }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}