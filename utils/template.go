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
