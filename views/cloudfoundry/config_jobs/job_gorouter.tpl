{{with .CloudFoundryJobs.gorouter }}
{{template "cloudfoundry/config_jobs/basic.tpl" .}}
{{end}}