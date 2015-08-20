{{with .CloudFoundryJobs.postgres }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}