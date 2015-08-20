{{with .CloudFoundryJobs.loggregator_traffic }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}