package entity

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
)

type Deploy struct {
	Id  int64
	Sid int64
	Ops string
}

func (deploy *Deploy) Load() error {
	err := orm.NewOrm().Read(deploy, "Id")
	if err != nil {
		logger.Error("Read Deploy error : %s", err)
	}
	return err
}

func (deploy *Deploy) Save() error {
	_, err := orm.NewOrm().Insert(deploy)
	if err != nil {
		logger.Error("Insert Deploy error : %s", err)
	}
	return err
}

func (deploy *Deploy) Update() error {
	_, err := orm.NewOrm().Update(deploy)
	if err != nil {
		logger.Error("Update Deploy error %s ", err)
	}
	return err
}

func (deploy *Deploy) Delete() error {
	_, err := orm.NewOrm().Delete(deploy)
	if err != nil {
		logger.Error("Delete Deploy error %s ", err)
	}
	return err
}

// sid = deploy.id
type Operation struct {
	Id      int64
	Start   string
	Restart string
	Stop    string
	Sid     int64
}

func (operation *Operation) Load() error {
	err := orm.NewOrm().Read(operation, "Id")
	if err != nil {
		logger.Error("Read Operation error : %s", err)
	}
	return err
}

func (operation *Operation) Save() error {
	_, err := orm.NewOrm().Insert(operation)
	if err != nil {
		logger.Error("Insert Operation error : %s", err)
	}
	return err
}

func (operation *Operation) Update() error {
	_, err := orm.NewOrm().Update(operation)
	if err != nil {
		logger.Error("Update Operation error %s ", err)
	}
	return err
}

func (operation *Operation) Delete() error {
	_, err := orm.NewOrm().Delete(operation)
	if err != nil {
		logger.Error("Delete Operation error %s ", err)
	}
	return err
}

func (operation *Operation) LoadBySid() error {
	err := orm.NewOrm().QueryTable(new(Operation)).Filter("sid", operation.Sid).One(operation)
	if err != nil {
		logger.Error("Read Operation error : %s", err)
		return err
	}
	return nil
}

// sid = deploy.id
type OnPaaS struct {
	Id     int64
	Sid    int64
	Api    string
	User   string
	Passwd string
	Org    string
	Space  string
}

func (onPaaS *OnPaaS) Load() error {
	err := orm.NewOrm().Read(onPaaS, "Id")
	if err != nil {
		logger.Error("Read OnPaaS error : %s", err)
	}
	return err
}

func (onPaaS *OnPaaS) LoadBySid() error {
	err := orm.NewOrm().QueryTable(new(OnPaaS)).Filter("sid", onPaaS.Sid).One(onPaaS)
	if err != nil {
		logger.Error("Read OnPaaS error : %s", err)
		return err
	}
	return nil
}

func (onPaaS *OnPaaS) Save() error {
	_, err := orm.NewOrm().Insert(onPaaS)
	if err != nil {
		logger.Error("Insert OnPaaS error : %s", err)
	}
	return err
}

func (onPaaS *OnPaaS) Update() error {
	_, err := orm.NewOrm().Update(onPaaS)
	if err != nil {
		logger.Error("Update OnPaaS error %s ", err)
	}
	return err
}

func (onPaaS *OnPaaS) Delete() error {
	_, err := orm.NewOrm().Delete(onPaaS)
	if err != nil {
		logger.Error("Delete OnPaaS error %s ", err)
	}
	return err
}

func (onPaaS *OnPaaS) TableName() string {
	return "on_paas"
}

// sid = deploy.id
type OnCustom struct {
	Id         int64
	Sid        int64
	Ip         string
	User       string
	Passwd     string
	PrivateKey string
}

func (onCustom *OnCustom) Load() error {
	err := orm.NewOrm().Read(onCustom, "Id")
	if err != nil {
		logger.Error("Read OnCustom error : %s", err)
	}
	return err
}

func (onCustom *OnCustom) LoadBySid() error {
	err := orm.NewOrm().QueryTable(new(OnCustom)).Filter("sid", onCustom.Sid).One(onCustom)
	if err != nil {
		logger.Error("Read OnCustom error : %s,Returned Multi Rows Not One.", err)
		return err
	}
	return nil
}

func (onCustom *OnCustom) Save() error {
	_, err := orm.NewOrm().Insert(onCustom)
	if err != nil {
		logger.Error("Insert OnCustom error : %s", err)
	}
	return err
}

func (onCustom *OnCustom) Update() error {
	_, err := orm.NewOrm().Update(onCustom)
	if err != nil {
		logger.Error("Update OnCustom error %s ", err)
	}
	return err
}

func (onCustom *OnCustom) Delete() error {
	_, err := orm.NewOrm().Delete(onCustom)
	if err != nil {
		logger.Error("Delete OnCustom error %s ", err)
	}
	return err
}

type ServiceDto struct {
	Service  Service
	OnPaaS   OnPaaS
	OnCustom OnCustom
}

func (serviceDto *ServiceDto) Load() error {
	if serviceDto.Service.Id == 0 {
		return fmt.Errorf("Read ServiceDTO %s,Unknow Id", "error")
	}
	if serviceDto.Service.Where == utils.Deploy_On_PaaS {
		onPaaS := OnPaaS{}
		onPaaS.Sid = serviceDto.Service.Id
		err := onPaaS.LoadBySid()
		if err != nil {
			return err
		}
		serviceDto.OnPaaS = onPaaS
	}

	if serviceDto.Service.Where == utils.Deploy_On_Vms {
		onCustom := OnCustom{}
		onCustom.Sid = serviceDto.Service.Id
		err := onCustom.LoadBySid()
		if err != nil {
			return err
		}
		serviceDto.OnCustom = onCustom
	}

	return nil
}

func init() {
	orm.RegisterModel(new(OnCustom), new(OnPaaS), new(Operation), new(Deploy))
}
