package dbt_yaml_builder

import (
	"log"

	"github.com/thoas/go-funk"
	"gopkg.in/yaml.v2"
)

func AddNewSourceTable(src DbtSourceYaml, sourceName string, tableEntry SourceTable) DbtSourceYaml {
	nextYaml := DbtSourceYaml(src)
	index := funk.IndexOf(nextYaml.Sources, func(source Source) bool {
		return source.Name == sourceName
	})
	if index == -1 {
		log.Fatalf("sourceName: [%v]을 찾을 수 없습니다", sourceName)
	}
	nextYaml.Sources[index].SourceTables = append(nextYaml.Sources[index].SourceTables, tableEntry)
	_, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatal("nextYaml을 파싱할 수 없습니다")
	}
	return nextYaml
}

func DeleteSourceTable(src DbtSourceYaml, sourceName string, tableName string) DbtSourceYaml {
	return src
}

func AddModel(src DbtSchemaYaml, modelEntry Models) DbtSchemaYaml {
	return src
}

func DeleteModel(src DbtSchemaYaml, modelName string) DbtSchemaYaml {
	return src
}
