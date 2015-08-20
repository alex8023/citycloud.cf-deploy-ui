{{with .CloudFoundryJobs.loggregator }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}