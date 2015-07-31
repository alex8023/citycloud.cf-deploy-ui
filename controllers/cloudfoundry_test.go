package controllers

import (
	"fmt"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"reflect"
	"testing"
)

func TestCommitData(testing *testing.T) {
	var cloud = cf
	AssertType(cloud)
	AssertType(cloud.CloudFoundryJobs)
	AssertType(cloud.CloudFoundryProperties)
	AssertType(cloud.Compilation)
	AssertType(cloud.NetWorks)
	AssertType(cloud.ResourcesPools)
	AssertType("abc")
}

func TestCommitData2(testing *testing.T) {

}

func AssertType(a interface{}) {
	// 使用类型断言
	switch a.(type) {
	case map[string]entity.CloudFoundryJobs:
		fmt.Println("==========================")
		fmt.Println("CloudFoundryJobs")
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a).Kind())
	case entity.CloudFoundryProperties:
		fmt.Println("==========================")
		fmt.Println("CloudFoundryProperties")
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a).Kind())
	case entity.Compilation:
		fmt.Println("==========================")
		fmt.Println("Compilation")
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a).Kind())
	case entity.CloudFoundry:
		fmt.Println("==========================")
		fmt.Println("CloudFoundry")
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a).Kind())
	case map[string]entity.NetWorks:
		fmt.Println("==========================")
		fmt.Println("NetWorks")
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a).Kind())
	case []entity.ResourcesPools:
		fmt.Println("==========================")
		fmt.Println("ResourcesPools")
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(reflect.TypeOf(a).Kind())
	default:
		fmt.Printf("Unknow Type %s\n", reflect.TypeOf(a))
	}

}
