{{with .CloudFoundryJobs.syslog_aggregator }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}