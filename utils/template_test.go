package utils_test

import (
	"fmt"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Describe("Testing Template with Ginkgo", func() {
	It("single file can be parsed", func() {
		p := Person{
			Name: "ginkgo",
			Age:  "20",
		}

		fileDir, _ := os.Getwd()

		// test.yml
		//{{.Name}}
		//{{.Age}}
		templatefile := filepath.Join(fileDir, "test.yml")
		templatefile1 := filepath.Join(fileDir, "test1.yml")
		bytes, _ := ioutil.ReadFile(templatefile)
		fmt.Println(string(bytes))
		bytes1, _ := ioutil.ReadFile(templatefile1)
		fmt.Println(string(bytes1))

		result, errors := utils.CreateYmlFileFromFile("/home/ubuntu/temp/test.yml", p, templatefile, templatefile1)
		bytes2, _ := ioutil.ReadFile("/home/ubuntu/temp/test.yml")
		fmt.Println(string(bytes2))

		fmt.Println(result)
		fmt.Println(errors)
	})

	It("Test ...string for args []string", func() {
		//		str := []string{"hello", "kity", "hi", "lucy"}

		Tarry("hello", "kity", "hi", "lucy")
	})

	FIt("Test parse from map", func() {
		mmap := make(map[string]string)
		mmap["mysql_ip"] = "10.162.2.146"
		mmap["mysql_port"] = "3306"
		mmap["mysql_user"] = "root"
		mmap["mysql_passwd"] = "123456"
		templatesfile := "/home/ubuntu/customservice/cfweb/templates/app-data.xml"

		utils.ParseTemplateFile2File("/home/ubuntu/temp/app-data.xml", mmap, templatesfile)
	})
})

type Person struct {
	Name string
	Age  string
}

func Tarry(str ...string) {
	fmt.Println(str)
}
