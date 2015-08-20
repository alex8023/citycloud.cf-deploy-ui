{{with .CloudFoundryJobs.nfs }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}