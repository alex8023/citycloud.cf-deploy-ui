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
