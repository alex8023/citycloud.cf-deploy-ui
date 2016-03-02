package entity_test

import (
	"github.com/astaxie/beego/orm"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing Service CRUD", func() {
	It("Service Save", func() {
		var service entity.Service = entity.Service{Name: "mysql-service", Description: "provide mysql store service"}
		err := service.Save()
		Expect(err).ToNot(HaveOccurred())

		err = service.Delete()
		Expect(err).ToNot(HaveOccurred())
	})
	It("Service Update", func() {
		service := entity.Service{Name: "mysql-service", Description: "provide mysql store service"}
		id, err := orm.NewOrm().Insert(&service)
		Expect(err).ToNot(HaveOccurred())
		service1 := entity.Service{}
		service1.Id = id
		errs := service1.Load()
		Expect(errs).ToNot(HaveOccurred())
		service1.Name = "New-mysql-service"
		errs = service1.Update()
		Expect(errs).ToNot(HaveOccurred())
		errs = service1.Load()
		Expect(errs).ToNot(HaveOccurred())

		Expect(service1.Name).To(Equal("New-mysql-service"))

		errs = service1.Delete()
		Expect(errs).ToNot(HaveOccurred())

	})
})

var _ = Describe("Testing Template CRUD", func() {
	It("Template Save", func() {
		var template entity.Template = entity.Template{Name: "template", Description: "description"}
		err := template.Save()
		Expect(err).ToNot(HaveOccurred())

		err = template.Delete()
		Expect(err).ToNot(HaveOccurred())
	})
	It("Template Update", func() {
		template := entity.Template{Name: "template mysql-service", Description: "provide mysql store service"}
		id, err := orm.NewOrm().Insert(&template)
		Expect(err).ToNot(HaveOccurred())
		template1 := entity.Template{}
		template1.Id = id
		errs := template1.Load()
		Expect(errs).ToNot(HaveOccurred())
		template1.Name = "New-template-mysql-service"
		errs = template1.Update()
		Expect(errs).ToNot(HaveOccurred())
		errs = template1.Load()
		Expect(errs).ToNot(HaveOccurred())

		Expect(template1.Name).To(Equal("New-template-mysql-service"))

		errs = template1.Delete()
		Expect(errs).ToNot(HaveOccurred())

	})
})

var _ = Describe("Testing Component CRUD", func() {
	It("Component Save", func() {
		var component entity.Component = entity.Component{Name: "component", Value: "description"}
		err := component.Save()
		Expect(err).ToNot(HaveOccurred())

		err = component.Delete()
		Expect(err).ToNot(HaveOccurred())
	})
	It("Component Update", func() {
		component := entity.Component{Name: "component mysql-service", Value: "provide mysql store service"}
		id, err := orm.NewOrm().Insert(&component)
		Expect(err).ToNot(HaveOccurred())
		component1 := entity.Component{}
		component1.Id = id
		errs := component1.Load()
		Expect(errs).ToNot(HaveOccurred())
		component1.Name = "New-template-mysql-service"
		errs = component1.Update()
		Expect(errs).ToNot(HaveOccurred())
		errs = component1.Load()
		Expect(errs).ToNot(HaveOccurred())

		Expect(component1.Name).To(Equal("New-template-mysql-service"))

		errs = component1.Delete()
		Expect(errs).ToNot(HaveOccurred())

	})
})

func init() {
	mysqlUrl := "root:c1oudc0w@tcp(127.0.0.1:3306)/cf_deploy_ui?charset=utf8&parseTime=true&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", mysqlUrl)

}
