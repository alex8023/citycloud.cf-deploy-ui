package utils

//resolve import cycle not allowed
const (
	Job_Haproxy                 = "haproxy"
	Job_Gorouter                = "gorouter"
	Job_Postgres                = "postgres"
	Job_NFS                     = "nfs"
	Job_NATS                    = "nats"
	Job_Syslog                  = "syslog_aggregator"
	Job_Etcd                    = "etcd"
	Job_Loggregator             = "loggregator"
	Job_Loggregator_Traffic     = "loggregator_traffic"
	Job_Uaa                     = "uaa"
	Job_Cloud_Controller_NG     = "cloud_controller_ng"
	Job_Cloud_Controller_Worker = "cloud_controller_worker"
	Job_Cloud_Controller_Clock  = "cloud_controller_clock"
	Job_Hm9000                  = "hm9000"
	Job_Stats                   = "stats"
	Job_Dea_Next                = "dea_next"
)

const (
	Job_Instance = 16
)

const (
	Net_Public  = "public"
	Net_Private = "private"
)

const (
	CloudFoundryProperties = "CloudFoundryProperties"
	NetWorks               = "NetWorks"
	Compilation            = "Compilation"
	ResourcesPools         = "ResourcesPools"
	CloudFoundryJobs       = "CloudFoundryJobs"
	Properties             = "Properties"
	VsphereFlavor          = "vSphereFlavor"
)

//部署类型 PaaS 部署在cci-paas云上，Vms 部署在虚拟机上
const (
	Deploy_On_PaaS = "PaaS"
	Deploy_On_Vms  = "Vms"
)

//文件类型，war 代表程序包 template代表模板文件 exec 代表执行脚本
const (
	FileTypes_War      = "War"
	FileTypes_Template = "Template"
	FileTypes_EXEC     = "Exec"
)

//custom service operation
const (
	Service_Start   = "START"
	Service_Stop    = "STOP"
	Service_Restart = "RESTART"
)

//监控滚动保存的数量
const (
	Monitor_Agent_Size = 20
)

//response error code
const (
	ResponseCodeSuccess = "200"
	ResponseCodeFailed  = "400"
)
