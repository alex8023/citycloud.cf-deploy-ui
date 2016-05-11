package entity

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"time"
)

type NatsResult struct {
	Value Result `json:"value"`
}

type Result struct {
	JobState string `json:"job_state"`
	Index    int    `json:"index"`
	AgentId  string `json:"agent_id"`
	Job      Job    `json:"job"`
}

type Job struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

type Monitor struct {
	Id       int64
	Value    string `orm:"size(3200)"`
	AgentId  string
	JobName  string
	Index    int
	JobState string
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func (monitor *Monitor) Load() error {
	err := orm.NewOrm().Read(monitor, "Id")
	if err != nil {
		beego.Error("Read Monitor error : %s", err)
	}
	return err
}

func (monitor *Monitor) LoadByAgentId() error {
	err := orm.NewOrm().QueryTable(new(Monitor)).Filter("AgentId", monitor.AgentId).OrderBy("Updated").Limit(1).One(monitor)
	if err != nil {
		beego.Error("Read Monitor error : %s", err)
	}
	return err
}

func (monitor *Monitor) Save() error {
	_, err := orm.NewOrm().Insert(monitor)
	if err != nil {
		beego.Error("Insert Monitor error : %s", err)
	}
	return err
}

func (monitor *Monitor) Update() error {
	_, err := orm.NewOrm().Update(monitor)
	if err != nil {
		beego.Error("Update Monitor error %s ", err)
	}
	return err
}

func (monitor *Monitor) Delete() error {
	_, err := orm.NewOrm().Delete(monitor)
	if err != nil {
		beego.Error("Delete Monitor error %s ", err)
	}
	return err
}

func (monitor *Monitor) DeleteByAgentId() error {
	_, err := orm.NewOrm().QueryTable(new(Monitor)).Filter("AgentId", monitor.AgentId).Delete()
	if err != nil {
		beego.Error("Delete Monitor error %s ", err)
	}
	return err
}

//同一个agent的结果保存20条记录，滚动覆盖记录
func (monitor *Monitor) SaveOrUpdate() error {
	count, err := orm.NewOrm().QueryTable(new(Monitor)).Filter("AgentId", monitor.AgentId).Count()

	if err != nil {
		beego.Error("Count Monitor error : %s", err)
		return err
	}

	if count < utils.Monitor_Agent_Size {
		return monitor.Save()
	}

	var _monitor Monitor
	//选取同一个agent_id 的最早update的记录
	err = orm.NewOrm().QueryTable(new(Monitor)).Filter("AgentId", monitor.AgentId).OrderBy("Updated").Limit(1).One(&_monitor)

	if err == orm.ErrNoRows {
		return monitor.Save()
	}

	if err == nil {
		monitor.Id = _monitor.Id
		return monitor.Update()
	}

	return err
}

type AgentVm struct {
	Id      int64
	AgentId string
}

func (agentVm *AgentVm) Load() error {
	err := orm.NewOrm().Read(agentVm, "Id")
	if err != nil {
		beego.Error("Read AgentVm error : %s", err)
	}
	return err
}

func (agentVm *AgentVm) Save() error {
	_, err := orm.NewOrm().Insert(agentVm)
	if err != nil {
		beego.Error("Insert AgentVm error : %s", err)
	}
	return err
}

func (agentVm *AgentVm) Update() error {
	_, err := orm.NewOrm().Update(agentVm)
	if err != nil {
		beego.Error("Update AgentVm error %s ", err)
	}
	return err
}

type JobMonitorStruct struct {
	JobName          string
	CloudFoundryJobs CloudFoundryJobs
	Monitors         []Monitor
}

func (jobMonitorStruct *JobMonitorStruct) Load() error {
	if jobMonitorStruct.JobName != "" {
		job := CloudFoundryJobs{}
		job.JobName = jobMonitorStruct.JobName
		job.Load()
		jobMonitorStruct.CloudFoundryJobs = job

		monitors := make([]Monitor, 0)

		for i := 0; i < job.Instances; i++ {
			var monitor Monitor
			err := orm.NewOrm().QueryTable(new(Monitor)).Filter("JobName", job.JobName).Filter("Index", i).OrderBy("Updated").Limit(1).One(&monitor)
			if err != nil {
				return err
			}
			monitors = append(monitors, monitor)
		}
		jobMonitorStruct.Monitors = monitors
	}
	return nil
}

func init() {
	orm.RegisterModel(new(Monitor), new(AgentVm))
}
