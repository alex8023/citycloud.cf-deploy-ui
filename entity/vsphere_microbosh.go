package entity

import (
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
)

type VsphereMicro struct {
	Id                int64
	Name              string
	VsphereResourceId int64
	VsphereNetWork    VsphereNetWork  `orm:"-"`
	VsphereResource   VsphereResource `orm:"-"`
	VsphereVcenter    VsphereVcenter  `orm:"-"`
}

func NewVsphereMicro(
	name string,
	vsphereNetWork VsphereNetWork,
	vsphereResource VsphereResource,
	vsphereVcenter VsphereVcenter,
) (vsphereMicro VsphereMicro) {
	vsphereMicro.Name = name
	vsphereMicro.VsphereNetWork = vsphereNetWork
	vsphereMicro.VsphereResource = vsphereResource
	vsphereMicro.VsphereVcenter = vsphereVcenter
	return
}

func (vsphereMicro *VsphereMicro) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(vsphereMicro).One(vsphereMicro, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(vsphereMicro, "Id")
	if errors != nil {
		logger.Error("Read VsphereMicro error : %s", errors)
		_, err := orm.NewOrm().Insert(vsphereMicro)
		if err != nil {
			logger.Error("Insert VsphereMicro error %s ", err)
		}
	}

	return nil
}

func (vsphereMicro *VsphereMicro) Update() error {
	_, err := orm.NewOrm().Update(vsphereMicro)
	if err != nil {
		logger.Error("Update VsphereMicro error %s ", err)
	}
	return err
}

type VsphereNetWork struct {
	Id       int64
	Ip       string
	NetMask  string
	GateWay  string
	Dns      string
	VlanName string
}

func NewVsphereNetWork(
	ip string,
	netMask string,
	gateWay string,
	dns string,
	vlanName string,
) (vsphereNetWork VsphereNetWork) {
	vsphereNetWork.Ip = ip
	vsphereNetWork.NetMask = netMask
	vsphereNetWork.GateWay = gateWay
	vsphereNetWork.Dns = dns
	vsphereNetWork.VlanName = vlanName
	return
}

func (vsphereNetWork *VsphereNetWork) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(vsphereNetWork).One(vsphereNetWork, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(vsphereNetWork, "Id")
	if errors != nil {
		logger.Error("Read VsphereNetWork error : %s", errors)
		_, err := orm.NewOrm().Insert(vsphereNetWork)
		if err != nil {
			logger.Error("Insert VsphereNetWork error %s ", err)
		}
	}

	return nil
}

func (vsphereNetWork *VsphereNetWork) Update() error {
	_, err := orm.NewOrm().Update(vsphereNetWork)
	if err != nil {
		logger.Error("Update VsphereNetWork error %s ", err)
	}
	return err
}

type VsphereResource struct {
	Id             int64
	PersistentDisk string
	Ram            string
	Disk           string
	Cpu            string
}

func NewVsphereResource(
	persistentDisk string,
	ram string,
	disk string,
	cpu string,
) (vsphereResource VsphereResource) {
	vsphereResource.PersistentDisk = persistentDisk
	vsphereResource.Ram = ram
	vsphereResource.Disk = disk
	vsphereResource.Cpu = cpu
	return
}

func (vsphereResource *VsphereResource) Load() error {

	errors := orm.NewOrm().Read(vsphereResource, "Id")
	if errors != nil {
		logger.Error("Read VsphereResource error : %s", errors)
	}

	return errors
}

func (vsphereResource *VsphereResource) Update() error {
	_, err := orm.NewOrm().Update(vsphereResource)
	if err != nil {
		logger.Error("Update VsphereResource error %s ", err)
	}
	return err
}

type VsphereVcenter struct {
	Id             int64
	UserName       string
	Passwd         string
	Host           string
	DataCenterName string
	VmFolder       string
	TemplateFolder string
	DiskPath       string
	DataStore      string
	ClustersName   string
}

func NewVsphereVcenter(
	userName string,
	passwd string,
	host string,
	dataCenterName string,
	vmFolder string,
	templateFolder string,
	diskPath string,
	dataStore string,
	clustersName string,
) (vsphereVcenter VsphereVcenter) {
	vsphereVcenter.UserName = userName
	vsphereVcenter.Passwd = passwd
	vsphereVcenter.Host = host
	vsphereVcenter.DataCenterName = dataCenterName
	vsphereVcenter.VmFolder = vmFolder
	vsphereVcenter.TemplateFolder = templateFolder
	vsphereVcenter.DiskPath = diskPath
	vsphereVcenter.DataStore = dataStore
	vsphereVcenter.ClustersName = clustersName
	return
}
func (vsphereVcenter *VsphereVcenter) Load() error {
	queryOneErr := orm.NewOrm().QueryTable(vsphereVcenter).One(vsphereVcenter, "Id")
	if queryOneErr != nil {
		logger.Error("Query One failed %s", queryOneErr)
	}

	errors := orm.NewOrm().Read(vsphereVcenter, "Id")
	if errors != nil {
		logger.Error("Read VsphereVcenter error : %s", errors)
		_, err := orm.NewOrm().Insert(vsphereVcenter)
		if err != nil {
			logger.Error("Insert VsphereVcenter error %s ", err)
		}
	}

	return nil
}

func (vsphereVcenter *VsphereVcenter) Update() error {
	_, err := orm.NewOrm().Update(vsphereVcenter)
	if err != nil {
		logger.Error("Update VsphereVcenter error %s ", err)
	}
	return err
}

func init() {
	orm.RegisterModel(new(VsphereNetWork), new(VsphereResource), new(VsphereVcenter), new(VsphereMicro))
}
