{{if eq .IaaSVersion .DefaultVersion}}
	{{template "cloudfoundry/index_networksv3.tpl" .}}
{{else}}
	{{template "cloudfoundry/index_networksv2.tpl" .}}
{{end}}