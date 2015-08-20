{{with .CloudFoundryJobs.uaa }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}