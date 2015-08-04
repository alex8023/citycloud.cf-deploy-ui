package entity

const (
	CloudFoundryJobsTemplate string = `{{define "jobs"}}{{with .CloudFoundryJobs}}
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
    - {{$.CloudFoundryProperties.FloatingIp}}
  release: {{$.Release.Name}}
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
    release: {{$.Release.Name}}

- name: {{.postgres.Name}}
  instances: {{.postgres.Instances}}
  resource_pool: {{.postgres.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .postgres.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: postgres
    release: {{$.Release.Name}}
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
    release: {{$.Release.Name}}
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
    release: {{$.Release.Name}}
  - name: nats_stream_forwarder
    release: {{$.Release.Name}}

- name: {{.syslog_aggregator.Name}}
  instances: {{.syslog_aggregator.Instances}}
  resource_pool: {{.syslog_aggregator.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .syslog_aggregator.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: syslog_aggregator
    release: {{$.Release.Name}}

- name: {{.etcd.Name}}
  instances: {{.etcd.Instances}}
  resource_pool: {{.etcd.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .etcd.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: etcd
    release: {{$.Release.Name}}
  - name: etcd_metrics_server
    release: {{$.Release.Name}}

- name: {{.loggregator.Name}}
  instances: {{.loggregator.Instances}}
  resource_pool: {{.loggregator.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .loggregator.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: loggregator
    release: {{$.Release.Name}}

- name: {{.loggregator_traffic.Name}}
  instances: {{.loggregator_traffic.Instances}}
  resource_pool: {{.loggregator_traffic.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .loggregator_traffic.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: loggregator_trafficcontroller
    release: {{$.Release.Name}}

- name: {{.uaa.Name}}
  instances: {{.uaa.Instances}}
  resource_pool: {{.uaa.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .uaa.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: uaa
    release: {{$.Release.Name}}

- name: {{.cloud_controller_ng.Name}}
  instances: {{.cloud_controller_ng.Instances}}
  resource_pool: {{.cloud_controller_ng.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_ng.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_ng
    release: {{$.Release.Name}}

- name: {{.cloud_controller_worker.Name}}
  instances: {{.cloud_controller_worker.Instances}}
  resource_pool: {{.cloud_controller_worker.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_worker.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_worker
    release: {{$.Release.Name}}

- name: {{.cloud_controller_clock.Name}}
  instances: {{.cloud_controller_clock.Instances}}
  resource_pool: {{.cloud_controller_clock.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .cloud_controller_clock.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: cloud_controller_clock
    release: {{$.Release.Name}}

- name: {{.hm9000.Name}}
  instances: {{.hm9000.Instances}}
  resource_pool: {{.hm9000.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .hm9000.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: hm9000
    release: {{$.Release.Name}}

- name: {{.stats.Name}}
  instances: {{.stats.Instances}}
  resource_pool: {{.stats.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .stats.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: collector
    release: {{$.Release.Name}}

- name: {{.dea_next.Name}}
  instances: {{.dea_next.Instances}}
  resource_pool: {{.dea_next.ResourcesPool}}
  networks:
  - name: {{$.NetWorks.private.Name}}
    static_ips:{{range .dea_next.StaticIp}}
    - {{.}}{{end}}
  templates:
  - name: dea_next
    release: {{$.Release.Name}}
  - name: dea_logging_agent
    release: {{$.Release.Name}}
  persistent_disk: 102400
{{end}}{{end}}`
)
