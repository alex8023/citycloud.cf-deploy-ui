package utils

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

func CreateYmlFile(templateName, templateText, path string, data interface{}) (bool, error) {

	templateYml := template.Must(template.New(templateName).Parse(templateText))

	buffer := bytes.NewBuffer([]byte{})

	err := templateYml.Execute(buffer, data)
	if err != nil {
		return false, err
	}

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)

	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false, err
	}

	defer file.Close()

	_, err = file.Write(buffer.Bytes())

	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

//将多个模板文件整合到一个文件输出，顺序以输入的模板文件名为顺序进行整合
func CreateYmlFileFromFile(path string, data interface{}, templatePath ...string) (bool, error) {

	templateYml := template.Must(template.ParseFiles(templatePath...))

	buffer := bytes.NewBuffer([]byte{})
	for _, item := range templatePath {

		err := templateYml.ExecuteTemplate(buffer, filepath.Base(item), data)
		if err != nil {
			return false, err
		}
	}
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)

	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false, err
	}

	defer file.Close()

	_, err = file.Write(buffer.Bytes())

	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

//将多个模板文件整合到一个文件输出，顺序以输入的模板文件名为顺序进行整合
func ParseTemplateFile2File(path string, data interface{}, templatePath ...string) (bool, error) {

	templateresult, parseerr := template.ParseFiles(templatePath...)
	if parseerr != nil {
		return false, parseerr
	}

	templateYml := template.Must(templateresult, parseerr)

	buffer := bytes.NewBuffer([]byte{})
	for _, item := range templatePath {

		err := templateYml.ExecuteTemplate(buffer, filepath.Base(item), data)
		if err != nil {
			return false, err
		}
	}
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)

	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false, err
	}

	defer file.Close()

	_, err = file.Write(buffer.Bytes())

	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}
	return true, nil
}
