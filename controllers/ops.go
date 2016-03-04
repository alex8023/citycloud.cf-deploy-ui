package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"strings"
	"time"
)

type OpsController struct {
	BaseController
}

type OpsMonitorController struct {
	BaseController
}

type OpsMonitorRestController struct {
	BaseController
}

func (this *OpsController) Get() {
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "ops"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion
	jobs := this.LoadJobMonitor()
	this.Data["JobMonitorStruct"] = jobs
	this.TplNames = "ops/index.tpl"
}

func (this *OpsMonitorController) Get() {
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "ops"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion

	agentId := this.GetString("agent_id")

	agentId = this.CheckAgentId(agentId)

	this.Data["AgentId"] = agentId

	this.TplNames = "ops/index_monitor.tpl"
}

func (this *OpsMonitorRestController) Get() {
	agentId := this.GetString("agent_id")
	natsServerIp := mi.NetWork.Vip
	if iaasVersion == vsphereVersion {
		natsServerIp = vsphereMicro.VsphereNetWork.Ip
	}
	natsServerIp = strings.Trim(natsServerIp, " ")

	//TODO 暂时直接从nats请求信息，需要持久话请求到的健康信息和负载信息
	monitorStr, err := utils.GetMonitor(natsServerIp, agentId, time.Duration(2))

	result := entity.ResponseMessage{}

	if err != nil {
		result.Code = utils.ResponseCodeFailed
		result.Data = fmt.Errorf("Request AgentId %s monitor info error %s", agentId, err)
		logger.Error("Request AgentId %s monitor info error %s", agentId, err)

	} else {
		result.Code = utils.ResponseCodeSuccess
		result.Data = monitorStr
	}
	this.Data["json"] = &result
	this.ServeJson(false)
}

func (this *OpsMonitorController) CheckAgentId(agentId string) string {
	return agentId
}

func (this *OpsController) LoadJobMonitor() map[string]entity.JobMonitorStruct {
	var jobs map[string]entity.JobMonitorStruct = map[string]entity.JobMonitorStruct{}
	var errors = make([]error, 0)
	haproxyJob := entity.JobMonitorStruct{}
	haproxyJob.JobName = utils.Job_Haproxy
	err := haproxyJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Haproxy, err))
	}
	jobs[utils.Job_Haproxy] = haproxyJob

	gorouterJob := entity.JobMonitorStruct{}
	gorouterJob.JobName = utils.Job_Gorouter
	err = gorouterJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Gorouter, err))
	}
	jobs[utils.Job_Gorouter] = gorouterJob

	postgresJob := entity.JobMonitorStruct{}
	postgresJob.JobName = utils.Job_Postgres
	err = postgresJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Postgres, err))
	}
	jobs[utils.Job_Postgres] = postgresJob

	nfsJob := entity.JobMonitorStruct{}
	nfsJob.JobName = utils.Job_NFS
	err = nfsJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_NFS, err))
	}
	jobs[utils.Job_NFS] = nfsJob

	natsJob := entity.JobMonitorStruct{}
	natsJob.JobName = utils.Job_NATS
	err = natsJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_NATS, err))
	}
	jobs[utils.Job_NATS] = natsJob

	syslog_aggregatorJob := entity.JobMonitorStruct{}
	syslog_aggregatorJob.JobName = utils.Job_Syslog
	err = syslog_aggregatorJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Syslog, err))
	}
	jobs[utils.Job_Syslog] = syslog_aggregatorJob

	etcdJob := entity.JobMonitorStruct{}
	etcdJob.JobName = utils.Job_Etcd
	err = etcdJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Etcd, err))
	}
	jobs[utils.Job_Etcd] = etcdJob

	loggregatorJob := entity.JobMonitorStruct{}
	loggregatorJob.JobName = utils.Job_Loggregator
	err = loggregatorJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Loggregator, err))
	}
	jobs[utils.Job_Loggregator] = loggregatorJob

	loggregator_trafficJob := entity.JobMonitorStruct{}
	loggregator_trafficJob.JobName = utils.Job_Loggregator_Traffic
	err = loggregator_trafficJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Loggregator_Traffic, err))
	}
	jobs[utils.Job_Loggregator_Traffic] = loggregator_trafficJob

	uaaJob := entity.JobMonitorStruct{}
	uaaJob.JobName = utils.Job_Uaa
	err = uaaJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Uaa, err))
	}
	jobs[utils.Job_Uaa] = uaaJob

	cloud_controller_ngJob := entity.JobMonitorStruct{}
	cloud_controller_ngJob.JobName = utils.Job_Cloud_Controller_NG
	err = cloud_controller_ngJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Cloud_Controller_NG, err))
	}
	jobs[utils.Job_Cloud_Controller_NG] = cloud_controller_ngJob

	cloud_controller_workerJob := entity.JobMonitorStruct{}
	cloud_controller_workerJob.JobName = utils.Job_Cloud_Controller_Worker
	err = cloud_controller_workerJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Cloud_Controller_Worker, err))
	}
	jobs[utils.Job_Cloud_Controller_Worker] = cloud_controller_workerJob

	cloud_controller_clockJob := entity.JobMonitorStruct{}
	cloud_controller_clockJob.JobName = utils.Job_Cloud_Controller_Clock
	err = cloud_controller_clockJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Cloud_Controller_Clock, err))
	}
	jobs[utils.Job_Cloud_Controller_Clock] = cloud_controller_clockJob

	hm9000Job := entity.JobMonitorStruct{}
	hm9000Job.JobName = utils.Job_Hm9000
	err = hm9000Job.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Hm9000, err))
	}
	jobs[utils.Job_Hm9000] = hm9000Job

	statsJob := entity.JobMonitorStruct{}
	statsJob.JobName = utils.Job_Stats
	err = statsJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Stats, err))
	}
	jobs[utils.Job_Stats] = statsJob

	dea_nextJob := entity.JobMonitorStruct{}
	dea_nextJob.JobName = utils.Job_Dea_Next
	err = dea_nextJob.Load()
	if err != nil {
		errors = append(errors, fmt.Errorf("Load JobMonitorStruct %s err %v", utils.Job_Dea_Next, err))
	}
	jobs[utils.Job_Dea_Next] = dea_nextJob
	if len(errors) != 0 {
		this.Data["MessageErr"] = fmt.Sprintf("%v", errors)
	}
	return jobs
}

func Save() {
	var result = `{
	    "value": {
	        "properties": {
	            "logging": {
	                "max_log_file_size": ""
	            }
	        },
	        "job": {
	            "name": "postgres",
	            "release": "",
	            "template": "postgres",
	            "version": "5",
	            "sha1": "0228430151570ea1e6cbb0c26c79b5ceaa83ef4d",
	            "blobstore_id": "f26d465e-55fb-4efa-9b6f-3c03e3c06bba",
	            "templates": [
	                {
	                    "name": "postgres",
	                    "version": "5",
	                    "sha1": "0228430151570ea1e6cbb0c26c79b5ceaa83ef4d",
	                    "blobstore_id": "f26d465e-55fb-4efa-9b6f-3c03e3c06bba"
	                }
	            ]
	        },
	        "packages": {
	            "common": {
	                "name": "common",
	                "version": "6.1",
	                "sha1": "41bbc068713cc4a2de2433a3b8d8f03ce399c442",
	                "blobstore_id": "17e30dc6-434b-4d5a-5cda-6203395d0333"
	            },
	            "postgres": {
	                "name": "postgres",
	                "version": "5.1",
	                "sha1": "92a88bb74c02c9fe8459590ef695dbc6f4d413f0",
	                "blobstore_id": "5780870b-09f3-46d8-6bb2-b96f52ed3529"
	            }
	        },
	        "configuration_hash": " ",
	        "networks": {
	            "cf12": {
	                "cloud_properties": {
	                    "net_id": "53f22403-b386-4158-a30f-07f0b8cc2ea7"
	                },
	                "default": [
	                    "dns",
	                    "gateway"
	                ],
	                "dns": [
	                    "192.168.138.11",
	                    "10.10.245.59"
	                ],
	                "dns_record_name": "0.postgres.cf12.cf-release.microbosh",
	                "ip": "192.168.138.22",
	                "netmask": "255.255.255.0"
	            }
	        },
	        "resource_pool": {
	            "cloud_properties": {
	                "availability_zone": "bigdata",
	                "instance_type": "flavor_737"
	            },
	            "name": "normal",
	            "stemcell": {
	                "name": "bosh-openstack-kvm-ubuntu-lucid-go_agent",
	                "version": "2719"
	            }
	        },
	        "deployment": "cf-release",
	        "index": 0,
	        "persistent_disk": 30720,
	        "persistent_disk_pool": {
	            "cloud_properties": {},
	            "disk_size": 30720,
	            "name": "e629ad1c-48ea-44ca-82fe-02a023c50819"
	        },
	        "rendered_templates_archive": {
	            "sha1": "36f90aaa68ef2ad08e3e8a3b2e2d01490e8ad02a",
	            "blobstore_id": "a43cbd05-4917-47dc-968d-f316063d078b"
	        },
	        "agent_id": "6b3733ee-7b6f-4d65-9b68-d3d08ddba267",
	        "bosh_protocol": "1",
	        "job_state": "running",
	        "vitals": {
	            "cpu": {
	                "sys": "0.0",
	                "user": "0.1",
	                "wait": "0.0"
	            },
	            "disk": {
	                "ephemeral": {
	                    "inode_percent": "0",
	                    "percent": "1"
	                },
	                "persistent": {
	                    "inode_percent": "0",
	                    "percent": "1"
	                },
	                "system": {
	                    "inode_percent": "32",
	                    "percent": "40"
	                }
	            },
	            "load": [
	                "0.02",
	                "0.08",
	                "0.12"
	            ],
	            "mem": {
	                "kb": "164720",
	                "percent": "8"
	            },
	            "swap": {
	                "kb": "0",
	                "percent": "0"
	            }
	        },
	        "vm": {
	            "name": "vm-447d7acb-4a60-4932-b8fc-d2dc155eeccc"
	        },
	        "ntp": {
	            "offset": "0.134190",
	            "timestamp": "22 Feb 11:15:07"
	        }
	    }
	}`
	var nats entity.NatsResult
	json.Unmarshal([]byte(result), &nats)

	monitor := entity.Monitor{}
	monitor.AgentId = nats.Value.AgentId
	monitor.Index = nats.Value.Index
	monitor.JobName = nats.Value.Job.Name
	monitor.JobState = nats.Value.JobState
	monitor.Updated = time.Now()
	monitor.Value = result
	monitor.Save()
}
