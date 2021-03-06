package entity

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 组件
type Component struct {
	Id    int64
	Name  string
	Value string
	Sid   int64
}

// 模板
type Template struct {
	Id           int64
	Name         string
	TemplateFile string               //模板路径
	TargetFile   string               //目标路径
	FileType     string               // file type
	Component    map[string]Component `orm:"-"`
	Description  string
	Sid          int64
}

//服务
type Service struct {
	Id          int64
	Name        string
	Description string
	Where       string
}

//联合主键
func (component *Component) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Sid"},
	}
}

func (template *Template) Load() error {
	err := orm.NewOrm().Read(template, "Id")
	if err != nil {
		beego.Error("Read Template error : %s", err)
	}
	return err
}

func (template *Template) Save() error {
	_, err := orm.NewOrm().Insert(template)
	if err != nil {
		beego.Error("Insert Template error : %s", err)
	}
	return err
}

func (template *Template) Update() error {
	_, err := orm.NewOrm().Update(template)
	if err != nil {
		beego.Error("Update Template error %s ", err)
	}
	return err
}

func (template *Template) Delete() error {
	_, err := orm.NewOrm().Delete(template)
	if err != nil {
		beego.Error("Delete Template error %s ", err)
	}
	return err
}

func (component *Component) Load() error {
	err := orm.NewOrm().Read(component, "Id")
	if err != nil {
		beego.Error("Read Component error : %s", err)
	}
	return err
}

func (component *Component) Save() error {
	_, err := orm.NewOrm().Insert(component)
	if err != nil {
		beego.Error("Insert Component error : %s", err)
	}
	return err
}

func (component *Component) Update() error {
	_, err := orm.NewOrm().Update(component)
	if err != nil {
		beego.Error("Update Component error %s ", err)
	}
	return err
}

func (component *Component) Delete() error {
	_, err := orm.NewOrm().Delete(component)
	if err != nil {
		beego.Error("Delete Component error %s ", err)
	}
	return err
}

func LoadComponentList(sid int64) ([]Component, error) {
	var component []Component
	qs := orm.NewOrm().QueryTable(new(Component)).Filter("sid", sid)
	_, err := qs.All(&component)
	if err != nil {
		beego.Error("Load Component List error %s ", err)
		e := fmt.Errorf("Load Component List error %s ", err)
		return component, e
	}
	return component, nil
}

func LoadTemplateList(sid int64) ([]Template, error) {
	var template []Template
	qs := orm.NewOrm().QueryTable(new(Template)).Filter("sid", sid)
	_, err := qs.All(&template)
	if err != nil {
		beego.Error("Load Template List error %s ", err)
		e := fmt.Errorf("Load Template List error %s ", err)
		return template, e
	}

	return template, nil
}

func (service *Service) Load() error {
	err := orm.NewOrm().Read(service, "Id")
	if err != nil {
		beego.Error("Read Service error : %s", err)
	}
	return err
}

func (service *Service) Save() error {
	_, err := orm.NewOrm().Insert(service)
	if err != nil {
		beego.Error("Insert Service error : %s", err)
	}
	return err
}

func (service *Service) Update() error {
	_, err := orm.NewOrm().Update(service)
	if err != nil {
		beego.Error("Update Service error %s ", err)
	}
	return err
}

func (service *Service) Delete() error {
	_, err := orm.NewOrm().Delete(service)
	if err != nil {
		beego.Error("Delete Service error %s ", err)
	}
	return err
}

func LoadServiceList() ([]Service, error) {
	var service []Service
	qs := orm.NewOrm().QueryTable(new(Service)).OrderBy("-id")
	_, err := qs.All(&service)
	if err != nil {
		beego.Error("Load Service List error %s ", err)
		e := fmt.Errorf("Load Service List error %s ", err)
		return service, e
	}
	return service, nil
}

func init() {
	orm.RegisterModel(new(Service), new(Template), new(Component))
}
