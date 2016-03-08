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

func (this *OpsController) Post() {
	agentId := this.GetString("agent_id")
	monitorStr, err := getMonitorByAgentId(agentId)
	if err != nil {
		this.Data["AgentError"] = fmt.Sprintf("Request AgentId [%s] monitor info error: %s", agentId, err)
		logger.Error("Request AgentId [%s] monitor info error: %s", agentId, err)
	} else {
		err = saveOrUpdateMonitor(monitorStr)
		this.Data["AgentError"] = fmt.Sprintf("Save error: %s", err)
		logger.Error("saveOrUpdateMonitor monitorStr [%s] Error %s", monitorStr, err)
	}

	this.Get()
}

func (this *OpsMonitorController) Get() {
	this.Layout = "index.tpl"
	this.Data["NavBarFocus"] = "ops"
	this.Data["IaaSVersion"] = iaasVersion
	this.Data["DefaultVersion"] = defaultVersion

	agentId := this.GetString("agent_id")

	agentId = this.CheckAgentId(agentId)

	this.Data["AgentId"] = agentId

	if agentId != "" {
		monitor := entity.Monitor{}
		monitor.AgentId = agentId
		err := monitor.LoadByAgentId()
		if err != nil {
			this.Data["MessageErr"] = fmt.Sprintf("Load Agent Monitor info err %s", err)
		}
		this.Data["Monitor"] = monitor
	}

	this.TplNames = "ops/index_monitor.tpl"
}

func (this *OpsMonitorRestController) Get() {
	agentId := this.GetString("agent_id")

	//TODO 暂时直接从nats请求信息，需要持久话健康信息和负载信息
	monitorStr, err := getMonitorByAgentId(agentId)

	result := entity.ResponseMessage{}

	if err != nil {
		result.Code = utils.ResponseCodeFailed
		result.Data = fmt.Sprintf("Request AgentId [%s] monitor info error: %s", agentId, err)
		logger.Error("Request AgentId %s monitor info error %s", agentId, err)

	} else {
		result.Code = utils.ResponseCodeSuccess
		result.Data = monitorStr
		err = saveOrUpdateMonitor(monitorStr)
		logger.Error("saveOrUpdateMonitor monitorStr [%s] Error %s", monitorStr, err)
	}
	this.Data["json"] = &result
	this.ServeJson(false)
}

func (this *OpsMonitorController) CheckAgentId(agentId string) string {
	//TODO 验证AgentId是否合法，不合法则返回""
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

func saveOrUpdateMonitor(result string) error {
	var nats entity.NatsResult
	err := json.Unmarshal([]byte(result), &nats)

	if err != nil {
		return err
	}

	monitor := entity.Monitor{}
	monitor.AgentId = nats.Value.AgentId
	monitor.Index = nats.Value.Index
	monitor.JobName = nats.Value.Job.Name
	monitor.JobState = nats.Value.JobState
	monitor.Updated = time.Now()
	monitor.Value = result

	err = monitor.SaveOrUpdate()
	return err
}

func getMonitorByAgentId(agentId string) (string, error) {
	natsServerIp := mi.NetWork.Vip

	if iaasVersion == vsphereVersion {
		natsServerIp = vsphereMicro.VsphereNetWork.Ip
	}

	natsServerIp = strings.Trim(natsServerIp, " ")

	//TODO 暂时直接从nats请求信息，需要持久话健康信息和负载信息 // 3s超时
	monitorStr, err := utils.GetMonitor(natsServerIp, agentId, time.Duration(3)*time.Second)
	return monitorStr, err
}
