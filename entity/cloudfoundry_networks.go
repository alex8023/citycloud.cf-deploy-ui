package entity

const (
	CloudFoundryNetworksTemplateV2 string = `{{define "networks"}}{{with .NetWorks}}
networks:
- name: {{.public.Name}}
  type: {{.public.NetType}}
  cloud_properties: {}
- name: {{.private.Name}}
  subnets:
  - cloud_properties: 
      net_id: {{.private.NetId}}
    range: {{.private.Cidr}}
    dns: [{{.private.PowerDns}},{{.private.Dns}}]
    reserved:
    - {{.private.ReservedIp}}
    static:
    - {{.private.StaticIp}}
{{end}}{{end}}`

	CloudFoundryNetworksTemplateV3 string = `{{define "networks"}}{{with .NetWorks}}
networks:
- name: {{.private.Name}}
  type: {{.private.NetType}}
  subnets:
  - cloud_properties: 
      net_id: {{.private.NetId}}
    range: {{.private.Cidr}}
    dns: [{{.private.PowerDns}},{{.private.Dns}}]
    reserved:
    - {{.private.ReservedIp}}
    static:
    - {{.private.StaticIp}}
- name: {{.public.Name}}
  type: {{.public.NetType}}
  subnets:
  - cloud_properties: 
      net_id: {{.public.NetId}}  
    range: {{.public.Cidr}}
    static:
    - {{.public.StaticIp}}
{{end}}{{end}}`
)
