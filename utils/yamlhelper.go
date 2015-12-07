package utils

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetValuesByKey(appdata Map, key string) (str []string, err error) {

	if appdata.Has("applications") {
		appMaps, ok := appdata.Get("applications").([]interface{})
		if !ok {
			return str, errors.New("Expected applications to be a list")
		}
		for _, item := range appMaps {
			mmap := NewMap(item)
			val := stringVal(mmap, key)
			str = append(str, val)
		}
	}
	return
}

func ReadYamlFile(files string) (yamlMap Map, err error) {
	file, err := os.Open(filepath.Clean(files))
	if err != nil {
		return
	}
	defer file.Close()

	manifest, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	mapp := make(map[interface{}]interface{})

	err = yaml.Unmarshal(manifest, &mapp)

	if err != nil {
		return
	}

	if !IsMappable(mapp) || len(mapp) == 0 {
		err = errors.New("Invalid manifest. Expected a map")
		return
	}

	yamlMap = NewMap(mapp)

	return yamlMap, nil
}

func stringVal(yamlMap Map, key string) string {
	val := yamlMap.Get(key)
	if val == nil {
		return ""
	}
	result, ok := val.(string)
	if !ok {
		return ""
	}
	return result
}
