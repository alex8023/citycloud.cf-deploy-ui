package entity

import (
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/logger"
)

type VsphereCompilation struct {
	Id               int64
	Workers          int
	DefaultNetWork   string
	Vid              int64            // vsphere resources id
	VsphereResources VsphereResources `orm:"-"`
}

type VsphereResourcesPools struct {
	Id               int64
	Name             string
	Size             int
	DefaultNetWork   string
	Vid              int64            // vsphere resources id
	VsphereResources VsphereResources `orm:"-"`
}
