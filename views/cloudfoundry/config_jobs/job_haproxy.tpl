{{with .CloudFoundryJobs.haproxy }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}