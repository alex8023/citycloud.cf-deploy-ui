{{with .CloudFoundryJobs.nats }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}