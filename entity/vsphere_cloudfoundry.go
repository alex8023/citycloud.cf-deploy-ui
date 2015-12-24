package entity

import (
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
)

type VsphereCompilation struct {
	Id              int64
	Workers         int
	DefaultNetWork  string
	Vid             int64           // vsphere resources id
	VsphereResource VsphereResource `orm:"-"`
}

func NewVsphereCompilation(
	workers int,
	defaultNetWork string,
	vid int64,
) (vsphereCompilation VsphereCompilation) {
	vsphereCompilation.Workers = workers
	vsphereCompilation.DefaultNetWork = defaultNetWork
	vsphereCompilation.Vid = vid
	return
}

func (vsphereCompilation *VsphereCompilation) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(vsphereCompilation).One(vsphereCompilation, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(vsphereCompilation, "Id")
	if errors != nil {
		logger.Error("Read VsphereCompilation error : %s", errors)
		_, err := orm.NewOrm().Insert(vsphereCompilation)
		if err != nil {
			logger.Error("Insert VsphereCompilation error %s ", err)
		}
	}

	return nil
}

func (vsphereCompilation *VsphereCompilation) Update() error {
	_, err := orm.NewOrm().Update(vsphereCompilation)
	if err != nil {
		logger.Error("Update VsphereCompilation error %s ", err)
	}
	return err
}

type VsphereResourcesPool struct {
	Id              int64
	Name            string
	Size            int
	DefaultNetWork  string
	Vid             int64           // vsphere resources id
	VsphereResource VsphereResource `orm:"-"`
}

func NewVsphereResourcesPool(
	name string,
	size int,
	defaultNetWork string,
	vid int64,
) (vsphereResourcesPool VsphereResourcesPool) {
	vsphereResourcesPool.Name = name
	vsphereResourcesPool.Size = size
	vsphereResourcesPool.DefaultNetWork = defaultNetWork
	vsphereResourcesPool.Vid = vid
	return
}

func (vsphereResourcesPool *VsphereResourcesPool) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(vsphereResourcesPool).One(vsphereResourcesPool, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(vsphereResourcesPool, "Id")
	if errors != nil {
		logger.Error("Read VsphereResourcesPool error : %s", errors)
		_, err := orm.NewOrm().Insert(vsphereResourcesPool)
		if err != nil {
			logger.Error("Insert VsphereResourcesPool error %s ", err)
		}
	}

	return nil
}

func (vsphereResourcesPool *VsphereResourcesPool) Update() error {
	_, err := orm.NewOrm().Update(vsphereResourcesPool)
	if err != nil {
		logger.Error("Update VsphereResourcesPool error %s ", err)
	}
	return err
}

type VsphereCloudFoundry struct {
	VsphereVcenter        VsphereVcenter
	VsphereCompilation    VsphereCompilation
	NetWorks              map[string]NetWorks
	VsphereResourcesPools []VsphereResourcesPool
	CloudFoundryJobs      map[string]CloudFoundryJobs
	Properties            Properties
	Release               Release
	Stemcells             Stemcells
}

func init() {
	orm.RegisterModel(new(VsphereResourcesPool), new(VsphereCompilation))
}
