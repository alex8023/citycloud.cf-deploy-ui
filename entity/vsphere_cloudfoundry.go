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
	VsphereCompilation     VsphereCompilation
	NetWorks               map[string]NetWorks
	VsphereResourcesPools  []VsphereResourcesPool
	CloudFoundryJobs       map[string]CloudFoundryJobs
	Properties             Properties
	Release                Release
	Stemcells              Stemcells
	CloudFoundryProperties CloudFoundryProperties
}

func LoadVsphereResourcePool(vsphereResourcesPool []VsphereResourcesPool) ([]VsphereResourcesPool, error) {
	qs := orm.NewOrm().QueryTable(new(VsphereResourcesPool))

	cnt, _ := qs.Count()
	if cnt == 0 {
		for index, value := range vsphereResourcesPool {
			value.Load()
			vsphereResourcesPool[index] = value
		}
		return vsphereResourcesPool, nil
	}

	var res []VsphereResourcesPool
	_, err := qs.All(&res)

	for index, value := range res {
		vsphereResource := VsphereResource{}

		vsphereResource.Id = value.Vid
		vsphereResource.Load()

		value.VsphereResource = vsphereResource
		res[index] = value
	}

	return res, err
}

func UpdateVsphereResourcePool(vsphereResourcesPool []VsphereResourcesPool) ([]VsphereResourcesPool, error) {
	qs := orm.NewOrm().QueryTable(new(VsphereResourcesPool))
	cond := orm.NewCondition()

	cond1 := cond.And("id__isnull", false)
	qs.SetCond(cond1).Delete()

	insert, _ := qs.PrepareInsert()

	for index, values := range vsphereResourcesPool {
		id, err := insert.Insert(&values)

		if err != nil {
			logger.Error("Insert VsphereResourcesPool error :%s", err)
		} else {
			values.Id = id
			vsphereResourcesPool[index] = values
		}

	}

	insert.Close()

	return vsphereResourcesPool, nil
}

func LoadVsphereResource(vsphereResources []VsphereResource) ([]VsphereResource, error) {
	qs := orm.NewOrm().QueryTable(new(VsphereResource))

	cnt, _ := qs.Count()
	if cnt == 0 {
		for index, value := range vsphereResources {
			value.Load()
			vsphereResources[index] = value
		}
		return vsphereResources, nil
	}

	var res []VsphereResource
	_, err := qs.All(&res)
	return res, err
}

func init() {
	orm.RegisterModel(new(VsphereResourcesPool), new(VsphereCompilation))
}
