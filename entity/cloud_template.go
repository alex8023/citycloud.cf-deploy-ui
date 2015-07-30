package entity

const (
	CloudFoundryTemplateTextV22 string = `
---
name: {{with .CloudFoundryProperties}}{{.Name}}{{end}}
director_uuid: {{with .CloudFoundryProperties}}{{.Uuid}}{{end}}
releases:
- name: cf2
  version: 170
meta:
  environment: {{with .CloudFoundryProperties}}{{.Name}}{{end}}
  stemcell:
    name: bosh-openstack-kvm-ubuntu-lucid-go_agent
    version: 2719
  openstack:
    auth_url: http://10.212.1.2:5000/v2.0
    tenant: Project_paas
    username: paas
    api_key: paasdunhuang
    security_groups: ["default"]
    default_key_name: paasdunhuang
    private_key: ~/bosh-workspace/keys/paasdunhuang.private
    cci_ebs_url: http://10.212.1.2:8070/EBS
    cci_ebs_accesskey: 73baf6b238dd42d2a602272ea01246c5 
    cci_ebs_secretkey: b28ce2b591c84f328763f59db3653144
    cci_connection_options: 
    cci_ebs_api_key: 
    cci_ebs_username:
{{with .Compilation}}
compilation:
  cloud_properties:
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
    instance_type: {{.InstanceType}}
  network: {{.DefaultNetWork}}
  reuse_compilation_vms: true
  workers: {{.Workers}}
{{end}}

update:
  canaries: 1
  canary_watch_time: 30000 - 90000
  update_watch_time: 30000 - 90000
  max_in_flight: 4
  max_errors: 4
{{with .NetWorks}}
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
{{end}}

{{with .ResourcesPools}}
resource_pools:{{range .}}
- cloud_properties:
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
    instance_type: {{.InstanceType}}
  name: {{.Name}}
  network: {{.DefaultNetWork}}
  size: {{.Size}}
  stemcell:
    name: bosh-openstack-kvm-ubuntu-lucid-go_agent
    version: 2719
{{end}}{{end}}

{{with .CloudFoundryJobs}}
jobs:
- name: {{.haproxy.Name}}
  instances: {{.haproxy.Instances}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    default: [dns,gateway]
    static_ips:{{range .haproxy.StaticIp}}
    - {{.}}{{end}}
  - name: {{$.NetWorks.public.Name}}
    static_ips:
    - {{$.NetWorks.public.StaticIp}}
  release: cf2
  resource_pool: {{.haproxy.ResourcesPool}}
  template: haproxy


- name: {{.gorouter.Name}}
  instances: {{.gorouter.Instances}}
  resource_pool: {{.gorouter.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .gorouter.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: gorouter
    release: cf2

- name: {{.postgres.Name}}
  instances: {{.postgres.Instances}}
  resource_pool: {{.postgres.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .postgres.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: postgres
    release: cf2
  persistent_disk: 30720

- name: {{.nfs.Name}}
  instances: {{.nfs.Instances}}
  resource_pool: {{.nfs.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .nfs.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: debian_nfs_server
    release: cf2
  persistent_disk: 102400

- name: {{.nats.Name}}
  instances: {{.nats.Instances}}
  resource_pool: {{.nats.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .nats.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: nats
    release: cf2
  - name: nats_stream_forwarder
    release: cf2

- name: {{.syslog_aggregator.Name}}
  instances: {{.syslog_aggregator.Instances}}
  resource_pool: {{.syslog_aggregator.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .syslog_aggregator.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: syslog_aggregator
    release: cf2

- name: {{.etcd.Name}}
  instances: {{.etcd.Instances}}
  resource_pool: {{.etcd.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .etcd.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: etcd
    release: cf2
  - name: etcd_metrics_server
    release: cf2

- name: {{.loggregator.Name}}
  instances: {{.loggregator.Instances}}
  resource_pool: {{.loggregator.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .loggregator.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: loggregator
    release: cf2

- name: {{.loggregator_traffic.Name}}
  instances: {{.loggregator_traffic.Instances}}
  resource_pool: {{.loggregator_traffic.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .loggregator_traffic.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: loggregator_trafficcontroller
    release: cf2

- name: {{.uaa.Name}}
  instances: {{.uaa.Instances}}
  resource_pool: {{.uaa.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .uaa.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: uaa
    release: cf2


- name: {{.cloud_controller_ng.Name}}
  instances: {{.cloud_controller_ng.Instances}}
  resource_pool: {{.cloud_controller_ng.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_ng.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_ng
    release: cf2

- name: {{.cloud_controller_worker.Name}}
  instances: {{.cloud_controller_worker.Instances}}
  resource_pool: {{.cloud_controller_worker.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_worker.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_worker
    release: cf2
	
- name: {{.cloud_controller_clock.Name}}
  instances: {{.cloud_controller_clock.Instances}}
  resource_pool: {{.cloud_controller_clock.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_clock.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_clock


- name: {{.hm9000.Name}}
  instances: {{.hm9000.Instances}}
  resource_pool: {{.hm9000.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .hm9000.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: hm9000
    release: cf2

- name: {{.stats.Name}}
  instances: {{.stats.Instances}}
  resource_pool: {{.stats.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .stats.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: collector
    release: cf2


- name: {{.dea_next.Name}}
  instances: {{.dea_next.Instances}}
  resource_pool: {{.dea_next.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .dea_next.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: dea_next
    release: cf2
  - name: dea_logging_agent
    release: cf2
  persistent_disk: 102400
{{end}}

properties:
  system_domain: {{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  system_domain_organization: {{with .CloudFoundryProperties}}{{.SystemDomainOrg}}{{end}}
  support_address: http://support.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  domain: {{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  app_domains:
  - {{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  description: CCI PaaS v2 sponsored by Pivotal
  syslog_aggregator:
    address: {{range .CloudFoundryJobs.syslog_aggregator.StaticIp}}{{.}}{{end}}
    port: 80
  etcd:
    machines:{{range .CloudFoundryJobs.etcd.StaticIp}}
    - {{.}}{{end}}
  etcd_metrics_server:
    etcd:
      machine: 127.0.0.1
      port: 4001
    nats:
      machines:{{range .CloudFoundryJobs.nats.StaticIp}}
      - {{.}}{{end}}
      port: 4222
      username: nats
      password: c1oudc0w
    status:
      username: etcd
      password: c1oudc0w
  ha_proxy:
    ssl_pem: ! '-----BEGIN CERTIFICATE-----

      MIICLzCCAZgCCQCSoIG3LoeSMTANBgkqhkiG9w0BAQUFADBcMQswCQYDVQQGEwJV

      UzELMAkGA1UECBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xEDAOBgNVBAoT

      B1Bpdm90YWwxFjAUBgNVBAsTDUNsb3VkIEZvdW5kcnkwHhcNMTMxMDE3MjMxNzM5

      WhcNMTQxMDE3MjMxNzM5WjBcMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExFjAU

      BgNVBAcTDVNhbiBGcmFuY2lzY28xEDAOBgNVBAoTB1Bpdm90YWwxFjAUBgNVBAsT

      DUNsb3VkIEZvdW5kcnkwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBANqyjJMH

      FDbJ4XM2vLRxz6i82Gd2Y2dnAKYgu67FgheHGQJSv38lvn8JrAOBXu2QJgP8sJa+

      XqUWUTMo9BWvhvedQLojtcpLdULc0FhlIXn8bThGlTQyrSg9QJobhQZGziFVrdQM

      UZy4X+s6A2Szb9iOg3TYibnnfnuvdZli83eRAgMBAAEwDQYJKoZIhvcNAQEFBQAD

      gYEAbFu65WT+czpiJqHhNz5AnGYA8kieVlV7KhcljLtsU4Sxut5Vq9uXFBE09KCg

      YkyZ9KfzXArTeRCfcbm5xq12I+1nf6q0QjT1y3P6ztax0rpyb0i+4GWvA+ruMhfx

      n4QM1MkiJVYBGKkXFZDd1zsdR3pY4vm1uiMY75IvJQfgb08=

      -----END CERTIFICATE-----

      -----BEGIN RSA PRIVATE KEY-----

      MIICXQIBAAKBgQDasoyTBxQ2yeFzNry0cc+ovNhndmNnZwCmILuuxYIXhxkCUr9/

      Jb5/CawDgV7tkCYD/LCWvl6lFlEzKPQVr4b3nUC6I7XKS3VC3NBYZSF5/G04RpU0

      Mq0oPUCaG4UGRs4hVa3UDFGcuF/rOgNks2/YjoN02Im55357r3WZYvN3kQIDAQAB

      AoGAa88G81fTBCtDA1vhbIaKWuE1QNOgrxGcxUhvnPlqZxTHJFkMY66EmPV4oYW9

      +RhNVTvVBYq092boAnNW1/Xebvrk1SnBDkrLntnGPmExkatOkPTFFlNXfePu6qOJ

      ULwYg8rKRwpvLoQXxbzMDXApPBifBNWGHVneGuHLpwPEQgECQQD0IJOecEyorrCR

      6+wmJBS/IwqQO0Ooj7kZEg65MHi9exVe+XFvP0lW2NAUsDuBLz79hjslSqIJjRG8

      c6q36oqhAkEA5VVeEip+T4DV12I5A5maGexVMeC92K7EGU/H8YhltxVZ/RtNngdT

      3r19WeDbDF7R5CJy2f7AYullk3S6fkk28QJBALdEiW1k4rezRMyW7tYHOifvN2vl

      gbpWAt/GRZVSxSGB+B4vZq/cM8NlynftgQ5PGJucnGQ3bgN7irgNoTimc2ECQFMX

      QBMy5DroAhKcmu2r/IKB90gwFnjosVI+bsIbWkcgbE9hUhj8rK2aWE11Q8hSnpea

      x6QmQgxUZiIr+9n/qvECQQDiDIeSmo2mRYsaqr4CQ3Ak+EDjsm9XTpXHuqm+xgFO

      iDIeQCWd3/twqdDTR7FaDE7Q0i559u7A1yLumUn8caLF

      -----END RSA PRIVATE KEY-----'
  networks:
    apps: {{$.NetWorks.private.Name}}
  ssl:
    skip_cert_verify: true
  nats:
    address: {{range .CloudFoundryJobs.nats.StaticIp}}{{.}}{{end}}
    user: nats
    password: c1oudc0w
    port: 4222
    machines:{{range .CloudFoundryJobs.nats.StaticIp}}
    - {{.}}{{end}}
    authorization_timeout: 30
    debug: true
    trace: true
    monitor_port: 0
    prof_port: 0
  router:
    servers:
      z1:{{range .CloudFoundryJobs.gorouter.StaticIp}}
      - {{.}}{{end}}
      z2: []
    endpoint_timeout: 120
    status:
      user: router
      password: c1oudc0w
  loggregator_endpoint:
    host: {{range .CloudFoundryJobs.loggregator.StaticIp}}{{.}}{{end}}
    shared_secret: c1oudc0w
    use_ssl: false
    port: 3456
  loggregator:
    servers:
      z1:{{range .CloudFoundryJobs.loggregator.StaticIp}}
      - {{.}}{{end}}
    status:
      user: loggregator
      password: c1oudc0w
      port: 5768
  traffic_controller:
    zone: z1
    status:
      user: trafic_controller
      password: c1oudc0w
      port: 6789
  collector:
    datadog_api_key: collector
    datadog_application_key: c1oudc0w
    deployment_name: cf-release
    use_datadog: false
    use_tsdb: false
  disk_quota_enabled: true
  debian_nfs_server:
    no_root_squash: true
  nfs_server:
    no_root_squash: true
    address: {{range .CloudFoundryJobs.nfs.StaticIp}}{{.}}{{end}}
    network: {{.NetWorks.private.Cidr}}
    share: /var/vcap/store
  directory_server_protocol: http
  dea_next:
    advertise_interval_in_seconds: 5
    allow_networks: []
    default_health_check_timeout: 60
    deny_networks: []
    directory_server_protocol: http
    disk_mb: 16384
    disk_overcommit_factor: 2
    evacuation_bail_out_time_in_seconds: 0
    heartbeat_interval_in_seconds: 10
    instance_disk_inode_limit: 200000
    kernel_network_tuning_enabled: false
    logging_level: debug
    memory_mb: 8192
    memory_overcommit_factor: 3
    staging_disk_inode_limit: 200000
    staging_disk_limit_mb: 4096
    staging_memory_limit_mb: 1024
  dea_logging_agent:
    status:
      user: admin
      password: c1oudc0w
  ccdb: &21120980
    address: {{range .CloudFoundryJobs.postgres.StaticIp}}{{.}}{{end}}
    databases:
    - citext: true
      name: ccdb
      tag: cc
    db_scheme: postgres
    port: 5524
    roles:
    - name: ccadmin
      password: c1oudc0w
      tag: admin
  ccdb_ng: *21120980
  cc: &21119160
    bulk_api_user: bulk_api
    bulk_api_password: c1oudc0w
    srv_api_uri: http://api.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
    db_logging_level: debug2
    ccng.logging_level: debug2
    db_encryption_key: c1oudc0w
    quota_definitions:
      default:
        memory_limit: 524288
        total_services: 1000
        non_basic_services_allowed: true
        total_routes: 1000
        trial_db_allowed: true
    hm9000_noop: false
    staging_upload_user: uploaduser
    staging_upload_password: c1oudc0w
    install_buildpacks:
    - name: java_buildpack
      package: buildpack_java
    - name: ruby_buildpack
      package: buildpack_ruby
    - name: nodejs_buildpack
      package: buildpack_nodejs
  ccng: *21119160
  databases:
    db_scheme: postgres
    address: {{range .CloudFoundryJobs.postgres.StaticIp}}{{.}}{{end}}
    port: 5524
    roles:
    - tag: admin
      name: ccadmin
      password: c1oudc0w
    - tag: admin
      name: uaaadmin
      password: c1oudc0w
    databases:
    - tag: cc
      name: ccdb
      citext: true
    - tag: uaa
      name: uaadb
      citext: true
  db: databases
  login:
    enabled: false
  uaa:
    admin:
      client_secret: c1oudc0w
    batch:
      password: batch
      username: c1oudc0w
    catalina_opts: -Xmx512m -XX:MaxPermSize=256m
    cc:
      client_secret: c1oudc0w
    client:
      autoapprove:
      - cf
    clients:
      cf:
        override: true
        authorized-grant-types: password,implicit,refresh_token
        authorities: uaa.none
        scope: cloud_controller.read,cloud_controller.write,openid,password.write,cloud_controller.admin,scim.read,scim.write
        access-token-validity: 7200
        refresh-token-validity: 1209600
      app-direct:
        access-token-validity: 1209600
        authorities: app_direct_invoice.write
        authorized-grant-types: authorization_code,client_credentials,password,refresh_token,implicit
        override: true
        redirect-uri: http://console.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
        refresh-token-validity: 1209600
        secret: c1oudc0w
      cc_service_broker_client: 
      developer_console:
        access-token-validity: 1209600
        authorities: scim.write,scim.read,cloud_controller.read,cloud_controller.write,password.write,uaa.admin,uaa.resource,cloud_controller.admin,billing.admin
        authorized-grant-types: authorization_code,client_credentials
        override: true
        redirect-uri: http://console.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}/oauth/callback
        refresh-token-validity: 1209600
        scope: openid,cloud_controller.read,cloud_controller.write,password.write,console.admin,console.support
        secret: c1oudc0w
      login:
        authorities: oauth.login
        authorized-grant-types: authorization_code,client_credentials,refresh_token
        override: true
        redirect-uri: http://login.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
        scope: openid,oauth.approvals
        secret: c1oudc0w
      servicesmgmt:
        authorities: uaa.resource,oauth.service,clients.read,clients.write,clients.secret
        authorized-grant-types: authorization_code,client_credentials,password,implicit
        autoapprove: true
        override: true
        redirect-uri: http://servicesmgmt.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}/auth/cloudfoundry/callback
        scope: openid,cloud_controller.read,cloud_controller.write
        secret: c1oudc0w
      space-mail:
        access-token-validity: 1209600
        authorities: scim.read,scim.write,cloud_controller.admin
        authorized-grant-types: client_credentials
        override: true
        refresh-token-validity: 1209600
        secret: c1oudc0w
      support-services:
        access-token-validity: 1209600
        authorities: portal.users.read
        authorized-grant-types: authorization_code,client_credentials
        redirect-uri: http://support-signon.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
        refresh-token-validity: 1209600
        scope: scim.write,scim.read,openid,cloud_controller.read,cloud_controller.write
        secret: c1oudc0w
    jwt:
      signing_key: ! '-----BEGIN RSA PRIVATE KEY-----

        MIICXAIBAAKBgQDHFr+KICms+tuT1OXJwhCUmR2dKVy7psa8xzElSyzqx7oJyfJ1

        JZyOzToj9T5SfTIq396agbHJWVfYphNahvZ/7uMXqHxf+ZH9BL1gk9Y6kCnbM5R6

        0gfwjyW1/dQPjOzn9N394zd2FJoFHwdq9Qs0wBugspULZVNRxq7veq/fzwIDAQAB

        AoGBAJ8dRTQFhIllbHx4GLbpTQsWXJ6w4hZvskJKCLM/o8R4n+0W45pQ1xEiYKdA

        Z/DRcnjltylRImBD8XuLL8iYOQSZXNMb1h3g5/UGbUXLmCgQLOUUlnYt34QOQm+0

        KvUqfMSFBbKMsYBAoQmNdTHBaz3dZa8ON9hh/f5TT8u0OWNRAkEA5opzsIXv+52J

        duc1VGyX3SwlxiE2dStW8wZqGiuLH142n6MKnkLU4ctNLiclw6BZePXFZYIK+AkE

        xQ+k16je5QJBAN0TIKMPWIbbHVr5rkdUqOyezlFFWYOwnMmw/BKa1d3zp54VP/P8

        +5aQ2d4sMoKEOfdWH7UqMe3FszfYFvSu5KMCQFMYeFaaEEP7Jn8rGzfQ5HQd44ek

        lQJqmq6CE2BXbY/i34FuvPcKU70HEEygY6Y9d8J3o6zQ0K9SYNu+pcXt4lkCQA3h

        jJQQe5uEGJTExqed7jllQ0khFJzLMx0K6tj0NeeIzAaGCQz13oo2sCdeGRHO4aDh

        HH6Qlq/6UOV5wP8+GAcCQFgRCcB+hrje8hfEEefHcFpyKH+5g1Eu1k0mLrxK2zd+

        4SlotYRHgPCEubokb2S1zfZDWIXW3HmggnGgM949TlY=

        -----END RSA PRIVATE KEY-----'
      verification_key: ! '-----BEGIN PUBLIC KEY-----

        MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDHFr+KICms+tuT1OXJwhCUmR2d

        KVy7psa8xzElSyzqx7oJyfJ1JZyOzToj9T5SfTIq396agbHJWVfYphNahvZ/7uMX

        qHxf+ZH9BL1gk9Y6kCnbM5R60gfwjyW1/dQPjOzn9N394zd2FJoFHwdq9Qs0wBug

        spULZVNRxq7veq/fzwIDAQAB

        -----END PUBLIC KEY-----'
    no_ssl: true
    scim:
      users:
      - admin|admin|scim.write,scim.read,openid,cloud_controller.admin
    url: http://uaa.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  uaadb:
    address: {{range .CloudFoundryJobs.postgres.StaticIp}}{{.}}{{end}}
    databases:
    - citext: true
      name: uaadb
      tag: uaa
    db_scheme: postgresql
    port: 5524
    roles:
    - name: uaaadmin
      password: c1oudc0w
      tag: admin
`
	CloudFoundryTemplateTextV33 string = `
---
name: {{with .CloudFoundryProperties}}{{.Name}}{{end}}
director_uuid: {{with .CloudFoundryProperties}}{{.Uuid}}{{end}}
releases:
- name: cf2
  version: 170
meta:
  environment: {{with .CloudFoundryProperties}}{{.Name}}{{end}}
  stemcell:
    name: bosh-openstack-kvm-ubuntu-lucid-go_agent
    version: 2719
  openstack:
    auth_url: http://10.212.1.2:5000/v2.0
    tenant: Project_paas
    username: paas
    api_key: paasdunhuang
    security_groups: ["default"]
    default_key_name: paasdunhuang
    private_key: ~/bosh-workspace/keys/paasdunhuang.private
    cci_ebs_url: http://10.212.1.2:8070/EBS
    cci_ebs_accesskey: 73baf6b238dd42d2a602272ea01246c5 
    cci_ebs_secretkey: b28ce2b591c84f328763f59db3653144
    cci_connection_options: 
    cci_ebs_api_key: 
    cci_ebs_username:
{{with .Compilation}}
compilation:
  cloud_properties:
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
    instance_type: {{.InstanceType}}
  network: {{.DefaultNetWork}}
  reuse_compilation_vms: true
  workers: {{.Workers}}
{{end}}

update:
  canaries: 1
  canary_watch_time: 30000 - 90000
  update_watch_time: 30000 - 90000
  max_in_flight: 4
  max_errors: 4

{{with .NetWorks}}
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
{{end}}

{{with .ResourcesPools}}
resource_pools:{{range .}}
- cloud_properties:
    {{if .AvailabilityZone}}availability_zone: {{.AvailabilityZone}}{{end}}
    instance_type: {{.InstanceType}}
  name: {{.Name}}
  network: {{.DefaultNetWork}}
  size: {{.Size}}
  stemcell:
    name: bosh-openstack-kvm-ubuntu-lucid-go_agent
    version: 2719
{{end}}{{end}}

{{with .CloudFoundryJobs}}
jobs:
- name: {{.haproxy.Name}}
  instances: {{.haproxy.Instances}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    default: [dns,gateway]
    static_ips:{{range .haproxy.StaticIp}}
    - {{.}}{{end}}
  - name: {{$.NetWorks.public.Name}}
    static_ips:
    - {{$.NetWorks.public.StaticIp}}
  release: cf2
  resource_pool: {{.haproxy.ResourcesPool}}
  template: haproxy


- name: {{.gorouter.Name}}
  instances: {{.gorouter.Instances}}
  resource_pool: {{.gorouter.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .gorouter.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: gorouter
    release: cf2

- name: {{.postgres.Name}}
  instances: {{.postgres.Instances}}
  resource_pool: {{.postgres.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .postgres.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: postgres
    release: cf2
  persistent_disk: 30720

- name: {{.nfs.Name}}
  instances: {{.nfs.Instances}}
  resource_pool: {{.nfs.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .nfs.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: debian_nfs_server
    release: cf2
  persistent_disk: 102400

- name: {{.nats.Name}}
  instances: {{.nats.Instances}}
  resource_pool: {{.nats.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .nats.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: nats
    release: cf2
  - name: nats_stream_forwarder
    release: cf2

- name: {{.syslog_aggregator.Name}}
  instances: {{.syslog_aggregator.Instances}}
  resource_pool: {{.syslog_aggregator.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .syslog_aggregator.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: syslog_aggregator
    release: cf2

- name: {{.etcd.Name}}
  instances: {{.etcd.Instances}}
  resource_pool: {{.etcd.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .etcd.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: etcd
    release: cf2
  - name: etcd_metrics_server
    release: cf2

- name: {{.loggregator.Name}}
  instances: {{.loggregator.Instances}}
  resource_pool: {{.loggregator.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .loggregator.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: loggregator
    release: cf2

- name: {{.loggregator_traffic.Name}}
  instances: {{.loggregator_traffic.Instances}}
  resource_pool: {{.loggregator_traffic.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .loggregator_traffic.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: loggregator_trafficcontroller
    release: cf2

- name: {{.uaa.Name}}
  instances: {{.uaa.Instances}}
  resource_pool: {{.uaa.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .uaa.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: uaa
    release: cf2


- name: {{.cloud_controller_ng.Name}}
  instances: {{.cloud_controller_ng.Instances}}
  resource_pool: {{.cloud_controller_ng.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_ng.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_ng
    release: cf2

- name: {{.cloud_controller_worker.Name}}
  instances: {{.cloud_controller_worker.Instances}}
  resource_pool: {{.cloud_controller_worker.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_worker.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_worker
    release: cf2
	
- name: {{.cloud_controller_clock.Name}}
  instances: {{.cloud_controller_clock.Instances}}
  resource_pool: {{.cloud_controller_clock.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_clock.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_clock


- name: {{.hm9000.Name}}
  instances: {{.hm9000.Instances}}
  resource_pool: {{.hm9000.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .hm9000.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: hm9000
    release: cf2

- name: {{.stats.Name}}
  instances: {{.stats.Instances}}
  resource_pool: {{.stats.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .stats.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: collector
    release: cf2


- name: {{.dea_next.Name}}
  instances: {{.dea_next.Instances}}
  resource_pool: {{.dea_next.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .dea_next.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: dea_next
    release: cf2
  - name: dea_logging_agent
    release: cf2
  persistent_disk: 102400
{{end}}

properties:
  system_domain: {{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  system_domain_organization: {{with .CloudFoundryProperties}}{{.SystemDomainOrg}}{{end}}
  support_address: http://support.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  domain: {{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  app_domains:
  - {{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  description: CCI PaaS v2 sponsored by Pivotal
  syslog_aggregator:
    address: {{range .CloudFoundryJobs.syslog_aggregator.StaticIp}}{{.}}{{end}}
    port: 80
  etcd:
    machines:{{range .CloudFoundryJobs.etcd.StaticIp}}
    - {{.}}{{end}}
  etcd_metrics_server:
    etcd:
      machine: 127.0.0.1
      port: 4001
    nats:
      machines:{{range .CloudFoundryJobs.nats.StaticIp}}
      - {{.}}{{end}}
      port: 4222
      username: nats
      password: c1oudc0w
    status:
      username: etcd
      password: c1oudc0w
  ha_proxy:
    ssl_pem: ! '-----BEGIN CERTIFICATE-----

      MIICLzCCAZgCCQCSoIG3LoeSMTANBgkqhkiG9w0BAQUFADBcMQswCQYDVQQGEwJV

      UzELMAkGA1UECBMCQ0ExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xEDAOBgNVBAoT

      B1Bpdm90YWwxFjAUBgNVBAsTDUNsb3VkIEZvdW5kcnkwHhcNMTMxMDE3MjMxNzM5

      WhcNMTQxMDE3MjMxNzM5WjBcMQswCQYDVQQGEwJVUzELMAkGA1UECBMCQ0ExFjAU

      BgNVBAcTDVNhbiBGcmFuY2lzY28xEDAOBgNVBAoTB1Bpdm90YWwxFjAUBgNVBAsT

      DUNsb3VkIEZvdW5kcnkwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBANqyjJMH

      FDbJ4XM2vLRxz6i82Gd2Y2dnAKYgu67FgheHGQJSv38lvn8JrAOBXu2QJgP8sJa+

      XqUWUTMo9BWvhvedQLojtcpLdULc0FhlIXn8bThGlTQyrSg9QJobhQZGziFVrdQM

      UZy4X+s6A2Szb9iOg3TYibnnfnuvdZli83eRAgMBAAEwDQYJKoZIhvcNAQEFBQAD

      gYEAbFu65WT+czpiJqHhNz5AnGYA8kieVlV7KhcljLtsU4Sxut5Vq9uXFBE09KCg

      YkyZ9KfzXArTeRCfcbm5xq12I+1nf6q0QjT1y3P6ztax0rpyb0i+4GWvA+ruMhfx

      n4QM1MkiJVYBGKkXFZDd1zsdR3pY4vm1uiMY75IvJQfgb08=

      -----END CERTIFICATE-----

      -----BEGIN RSA PRIVATE KEY-----

      MIICXQIBAAKBgQDasoyTBxQ2yeFzNry0cc+ovNhndmNnZwCmILuuxYIXhxkCUr9/

      Jb5/CawDgV7tkCYD/LCWvl6lFlEzKPQVr4b3nUC6I7XKS3VC3NBYZSF5/G04RpU0

      Mq0oPUCaG4UGRs4hVa3UDFGcuF/rOgNks2/YjoN02Im55357r3WZYvN3kQIDAQAB

      AoGAa88G81fTBCtDA1vhbIaKWuE1QNOgrxGcxUhvnPlqZxTHJFkMY66EmPV4oYW9

      +RhNVTvVBYq092boAnNW1/Xebvrk1SnBDkrLntnGPmExkatOkPTFFlNXfePu6qOJ

      ULwYg8rKRwpvLoQXxbzMDXApPBifBNWGHVneGuHLpwPEQgECQQD0IJOecEyorrCR

      6+wmJBS/IwqQO0Ooj7kZEg65MHi9exVe+XFvP0lW2NAUsDuBLz79hjslSqIJjRG8

      c6q36oqhAkEA5VVeEip+T4DV12I5A5maGexVMeC92K7EGU/H8YhltxVZ/RtNngdT

      3r19WeDbDF7R5CJy2f7AYullk3S6fkk28QJBALdEiW1k4rezRMyW7tYHOifvN2vl

      gbpWAt/GRZVSxSGB+B4vZq/cM8NlynftgQ5PGJucnGQ3bgN7irgNoTimc2ECQFMX

      QBMy5DroAhKcmu2r/IKB90gwFnjosVI+bsIbWkcgbE9hUhj8rK2aWE11Q8hSnpea

      x6QmQgxUZiIr+9n/qvECQQDiDIeSmo2mRYsaqr4CQ3Ak+EDjsm9XTpXHuqm+xgFO

      iDIeQCWd3/twqdDTR7FaDE7Q0i559u7A1yLumUn8caLF

      -----END RSA PRIVATE KEY-----'
  networks:
    apps: {{$.NetWorks.private.Name}}
  ssl:
    skip_cert_verify: true
  nats:
    address: {{range .CloudFoundryJobs.nats.StaticIp}}{{.}}{{end}}
    user: nats
    password: c1oudc0w
    port: 4222
    machines:{{range .CloudFoundryJobs.nats.StaticIp}}
    - {{.}}{{end}}
    authorization_timeout: 30
    debug: true
    trace: true
    monitor_port: 0
    prof_port: 0
  router:
    servers:
      z1:{{range .CloudFoundryJobs.gorouter.StaticIp}}
      - {{.}}{{end}}
      z2: []
    endpoint_timeout: 120
    status:
      user: router
      password: c1oudc0w
  loggregator_endpoint:
    host: {{range .CloudFoundryJobs.loggregator.StaticIp}}{{.}}{{end}}
    shared_secret: c1oudc0w
    use_ssl: false
    port: 3456
  loggregator:
    servers:
      z1:{{range .CloudFoundryJobs.loggregator.StaticIp}}
      - {{.}}{{end}}
    status:
      user: loggregator
      password: c1oudc0w
      port: 5768
  traffic_controller:
    zone: z1
    status:
      user: trafic_controller
      password: c1oudc0w
      port: 6789
  collector:
    datadog_api_key: collector
    datadog_application_key: c1oudc0w
    deployment_name: cf-release
    use_datadog: false
    use_tsdb: false
  disk_quota_enabled: true
  debian_nfs_server:
    no_root_squash: true
  nfs_server:
    no_root_squash: true
    address: {{range .CloudFoundryJobs.nfs.StaticIp}}{{.}}{{end}}
    network: {{.NetWorks.private.Cidr}}
    share: /var/vcap/store
  directory_server_protocol: http
  dea_next:
    advertise_interval_in_seconds: 5
    allow_networks: []
    default_health_check_timeout: 60
    deny_networks: []
    directory_server_protocol: http
    disk_mb: 16384
    disk_overcommit_factor: 2
    evacuation_bail_out_time_in_seconds: 0
    heartbeat_interval_in_seconds: 10
    instance_disk_inode_limit: 200000
    kernel_network_tuning_enabled: false
    logging_level: debug
    memory_mb: 8192
    memory_overcommit_factor: 3
    staging_disk_inode_limit: 200000
    staging_disk_limit_mb: 4096
    staging_memory_limit_mb: 1024
  dea_logging_agent:
    status:
      user: admin
      password: c1oudc0w
  ccdb: &21120980
    address: {{range .CloudFoundryJobs.postgres.StaticIp}}{{.}}{{end}}
    databases:
    - citext: true
      name: ccdb
      tag: cc
    db_scheme: postgres
    port: 5524
    roles:
    - name: ccadmin
      password: c1oudc0w
      tag: admin
  ccdb_ng: *21120980
  cc: &21119160
    bulk_api_user: bulk_api
    bulk_api_password: c1oudc0w
    srv_api_uri: http://api.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
    db_logging_level: debug2
    ccng.logging_level: debug2
    db_encryption_key: c1oudc0w
    quota_definitions:
      default:
        memory_limit: 524288
        total_services: 1000
        non_basic_services_allowed: true
        total_routes: 1000
        trial_db_allowed: true
    hm9000_noop: false
    staging_upload_user: uploaduser
    staging_upload_password: c1oudc0w
    install_buildpacks:
    - name: java_buildpack
      package: buildpack_java
    - name: ruby_buildpack
      package: buildpack_ruby
    - name: nodejs_buildpack
      package: buildpack_nodejs
  ccng: *21119160
  databases:
    db_scheme: postgres
    address: {{range .CloudFoundryJobs.postgres.StaticIp}}{{.}}{{end}}
    port: 5524
    roles:
    - tag: admin
      name: ccadmin
      password: c1oudc0w
    - tag: admin
      name: uaaadmin
      password: c1oudc0w
    databases:
    - tag: cc
      name: ccdb
      citext: true
    - tag: uaa
      name: uaadb
      citext: true
  db: databases
  login:
    enabled: false
  uaa:
    admin:
      client_secret: c1oudc0w
    batch:
      password: batch
      username: c1oudc0w
    catalina_opts: -Xmx512m -XX:MaxPermSize=256m
    cc:
      client_secret: c1oudc0w
    client:
      autoapprove:
      - cf
    clients:
      cf:
        override: true
        authorized-grant-types: password,implicit,refresh_token
        authorities: uaa.none
        scope: cloud_controller.read,cloud_controller.write,openid,password.write,cloud_controller.admin,scim.read,scim.write
        access-token-validity: 7200
        refresh-token-validity: 1209600
      app-direct:
        access-token-validity: 1209600
        authorities: app_direct_invoice.write
        authorized-grant-types: authorization_code,client_credentials,password,refresh_token,implicit
        override: true
        redirect-uri: http://console.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
        refresh-token-validity: 1209600
        secret: c1oudc0w
      cc_service_broker_client: 
      developer_console:
        access-token-validity: 1209600
        authorities: scim.write,scim.read,cloud_controller.read,cloud_controller.write,password.write,uaa.admin,uaa.resource,cloud_controller.admin,billing.admin
        authorized-grant-types: authorization_code,client_credentials
        override: true
        redirect-uri: http://console.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}/oauth/callback
        refresh-token-validity: 1209600
        scope: openid,cloud_controller.read,cloud_controller.write,password.write,console.admin,console.support
        secret: c1oudc0w
      login:
        authorities: oauth.login
        authorized-grant-types: authorization_code,client_credentials,refresh_token
        override: true
        redirect-uri: http://login.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
        scope: openid,oauth.approvals
        secret: c1oudc0w
      servicesmgmt:
        authorities: uaa.resource,oauth.service,clients.read,clients.write,clients.secret
        authorized-grant-types: authorization_code,client_credentials,password,implicit
        autoapprove: true
        override: true
        redirect-uri: http://servicesmgmt.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}/auth/cloudfoundry/callback
        scope: openid,cloud_controller.read,cloud_controller.write
        secret: c1oudc0w
      space-mail:
        access-token-validity: 1209600
        authorities: scim.read,scim.write,cloud_controller.admin
        authorized-grant-types: client_credentials
        override: true
        refresh-token-validity: 1209600
        secret: c1oudc0w
      support-services:
        access-token-validity: 1209600
        authorities: portal.users.read
        authorized-grant-types: authorization_code,client_credentials
        redirect-uri: http://support-signon.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
        refresh-token-validity: 1209600
        scope: scim.write,scim.read,openid,cloud_controller.read,cloud_controller.write
        secret: c1oudc0w
    jwt:
      signing_key: ! '-----BEGIN RSA PRIVATE KEY-----

        MIICXAIBAAKBgQDHFr+KICms+tuT1OXJwhCUmR2dKVy7psa8xzElSyzqx7oJyfJ1

        JZyOzToj9T5SfTIq396agbHJWVfYphNahvZ/7uMXqHxf+ZH9BL1gk9Y6kCnbM5R6

        0gfwjyW1/dQPjOzn9N394zd2FJoFHwdq9Qs0wBugspULZVNRxq7veq/fzwIDAQAB

        AoGBAJ8dRTQFhIllbHx4GLbpTQsWXJ6w4hZvskJKCLM/o8R4n+0W45pQ1xEiYKdA

        Z/DRcnjltylRImBD8XuLL8iYOQSZXNMb1h3g5/UGbUXLmCgQLOUUlnYt34QOQm+0

        KvUqfMSFBbKMsYBAoQmNdTHBaz3dZa8ON9hh/f5TT8u0OWNRAkEA5opzsIXv+52J

        duc1VGyX3SwlxiE2dStW8wZqGiuLH142n6MKnkLU4ctNLiclw6BZePXFZYIK+AkE

        xQ+k16je5QJBAN0TIKMPWIbbHVr5rkdUqOyezlFFWYOwnMmw/BKa1d3zp54VP/P8

        +5aQ2d4sMoKEOfdWH7UqMe3FszfYFvSu5KMCQFMYeFaaEEP7Jn8rGzfQ5HQd44ek

        lQJqmq6CE2BXbY/i34FuvPcKU70HEEygY6Y9d8J3o6zQ0K9SYNu+pcXt4lkCQA3h

        jJQQe5uEGJTExqed7jllQ0khFJzLMx0K6tj0NeeIzAaGCQz13oo2sCdeGRHO4aDh

        HH6Qlq/6UOV5wP8+GAcCQFgRCcB+hrje8hfEEefHcFpyKH+5g1Eu1k0mLrxK2zd+

        4SlotYRHgPCEubokb2S1zfZDWIXW3HmggnGgM949TlY=

        -----END RSA PRIVATE KEY-----'
      verification_key: ! '-----BEGIN PUBLIC KEY-----

        MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDHFr+KICms+tuT1OXJwhCUmR2d

        KVy7psa8xzElSyzqx7oJyfJ1JZyOzToj9T5SfTIq396agbHJWVfYphNahvZ/7uMX

        qHxf+ZH9BL1gk9Y6kCnbM5R60gfwjyW1/dQPjOzn9N394zd2FJoFHwdq9Qs0wBug

        spULZVNRxq7veq/fzwIDAQAB

        -----END PUBLIC KEY-----'
    no_ssl: true
    scim:
      users:
      - admin|admin|scim.write,scim.read,openid,cloud_controller.admin
    url: http://uaa.{{with .CloudFoundryProperties}}{{.SystemDomain}}{{end}}
  uaadb:
    address: {{range .CloudFoundryJobs.postgres.StaticIp}}{{.}}{{end}}
    databases:
    - citext: true
      name: uaadb
      tag: uaa
    db_scheme: postgresql
    port: 5524
    roles:
    - name: uaaadmin
      password: c1oudc0w
      tag: admin
`
)