package entity

import (
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

const BOSHTemplateText string = `
---
name: {{.Name}}
director_uuid: {{.Uuid}}

release:
  name: bosh
  version: 0+dev.22

compilation:
  workers: 3
  network: default
  reuse_compilation_vms: true
  cloud_properties:
    instance_type: {{.InstanceType}} # CHANGE or flavor_94
    availability_zone: zone2 # 指定区域zone2

update:
  canaries: 2
  canary_watch_time: 3000-120000
  update_watch_time: 3000-120000
  max_in_flight: 4

networks:
  - name: floating
    type: vip
    cloud_properties: {}
  - name: default
    type: dynamic
    cloud_properties: {}

resource_pools:
  - name: common
    network: default
    size: 8
    stemcell:
      name: bosh-openstack-kvm-ubuntu-lucid-go_agent
      version: 2719
    cloud_properties:
      instance_type: {{.InstanceType}} # CHANGE
      availability_zone: zone2

jobs:
  - name: nats
    template: nats
    instances: 1
    resource_pool: common
    networks:
      - name: default
        default: [dns, gateway]

  - name: redis
    template: redis
    instances: 1
    resource_pool: common
    networks:
      - name: default
        default: [dns, gateway]

  - name: postgres
    template: postgres
    instances: 1
    resource_pool: common
    persistent_disk: 16384
    networks:
      - name: default
        default: [dns, gateway]

  - name: powerdns
    template: powerdns
    instances: 1
    resource_pool: common
    networks:
      - name: default
        default: [dns, gateway]
      - name: floating
        static_ips:
          - {{.PowerDns}} # CHANGE

  - name: blobstore
    template: blobstore
    instances: 1
    resource_pool: common
    persistent_disk: 51200
    networks:
      - name: default
        default: [dns, gateway]

  - name: director
    template: director
    instances: 1
    resource_pool: common
    persistent_disk: 16384
    networks:
      - name: default
        default: [dns, gateway]
      - name: floating
        static_ips:
          - {{.Director}} # CHANGE

  - name: registry
    template: registry
    instances: 1
    resource_pool: common
    networks:
      - name: default
        default: [dns, gateway]

  - name: health_monitor
    template: health_monitor
    instances: 1
    resource_pool: common
    networks:
      - name: default
        default: [dns, gateway]

properties:
  nats:
    address: 0.nats.default.{{.Name}}.microbosh
    user: nats
    password: nats

  redis:
    address: 0.redis.default.{{.Name}}.microbosh
    password: redis

  postgres: &bosh_db
    host: 0.postgres.default.{{.Name}}.microbosh
    user: postgres
    password: postgres
    database: bosh

  dns:
    address: {{.PowerDns}}  # CHANGE
    db: *bosh_db
    recursor: {{.MicroBOSHIP}} # CHANGE<microbosh_ip_address>

  blobstore:
    address: 0.blobstore.default.{{.Name}}.microbosh
    agent:
      user: agent
      password: agent
    director:
      user: director
      password: director

  director:
    name: bosh
    address: 0.director.default.{{.Name}}.microbosh
    db: *bosh_db

  registry:
    address: 0.registry.default.{{.Name}}.microbosh
    db: *bosh_db
    http:
      user: registry
      password: registry

  hm:
    http:
      user: hm
      password: hm
    director_account:
      user: admin
      password: admin
    resurrector_enabled: true

  ntp:
    - 0.asia.pool.ntp.org
    - 1.asia.pool.ntp.org
{{with .CloudProperties}}
  openstack:
    auth_url: {{.AuthUrl}}
    tenant: {{.Tenant}}
    username: {{.UserName}}
    api_key: {{.ApiKey}}
    default_key_name: {{.DefaultKeyName}}
    private_key: {{.PrivateKey}}
    default_security_groups: [bosh]
    cci_ebs_url: {{.CciEbsUrl}}
    cci_ebs_accesskey: {{.CciEbsAccesskey}}
    cci_ebs_secretkey: {{.CciEbsSecretkey}}
    cci_connection_options: 
    cci_ebs_api_key: 
    cci_ebs_username:
{{end}}
`

type BOSHTemplate struct {
	bosh BOSH
}

func NewBOSHTemplate(bosh BOSH) (bt BOSHTemplate) {
	bt.bosh = bosh
	return
}

func (bt BOSHTemplate) CreateBOSHYaml(path string) (bool, error) {
	logger.Debug("Create BOSH deployment file : %s", path)
	return utils.CreateYmlFile("bosh", BOSHTemplateText, path, bt.bosh)
}
