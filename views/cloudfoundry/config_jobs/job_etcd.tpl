{{with .CloudFoundryJobs.etcd }}
{{template "cloudfoundry/config_jobs/basic_readonly.tpl" .}}
{{end}}