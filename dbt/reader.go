package dbt

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func ReadDbtSourceYaml(path string) DbtSourceYaml {
	absPath, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(absPath)
	d := DbtSourceYaml{}
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return d
}

func ReadDbtSchemaYaml(path string) DbtSchemaYaml {
	absPath, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(absPath)
	d := DbtSchemaYaml{}
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return d
}

func ReadDbtProjectYaml() {

}
