{{with .CloudFoundryJobs.cloud_controller_worker }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}