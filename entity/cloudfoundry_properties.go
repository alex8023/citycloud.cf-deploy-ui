package entity

const (
	CloudFoundryPropertiesTemplate string = `{{define "properties"}}
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
        memory_limit: {{with .Properties.JobProperties}}{{if .cc_quota_definitions_memory_limit}}{{.cc_quota_definitions_memory_limit.Value}}{{else}}524288{{end}}{{end}}
        total_services: {{with .Properties.JobProperties}}{{if .cc_quota_definitions_total_services}}{{.cc_quota_definitions_total_services.Value}}{{else}}1000{{end}}{{end}}
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
{{end}}`
)
