package utils

import (
	"testing"
)

type String struct{
	Name string
	Age int
}

func TestTemplate(t *testing.T){
	var templates = `{{.Name}}`
	var path = "/home/ubuntu/abc.yml"
	data := &String{Name:"lucy",Age:20}
	CreateYmlFile("helloworld",templates,path,data)
}