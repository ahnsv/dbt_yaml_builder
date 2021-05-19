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
	nextYaml := DbtSourceYaml(src)
	sourceIndex := funk.IndexOf(nextYaml.Sources, func(source Source) bool {
		return source.Name == sourceName
	})
	if sourceIndex == -1 {
		log.Fatalf("sourceName: [%v]을 찾을 수 없습니다", sourceName)
	}
	tableIndex := funk.IndexOf(nextYaml.Sources[sourceIndex].SourceTables, func(table SourceTable) bool {
		return table.Name == tableName
	})
	nextYaml.Sources[sourceIndex].SourceTables = append(nextYaml.Sources[sourceIndex].SourceTables[:tableIndex], nextYaml.Sources[sourceIndex].SourceTables[tableIndex+1:]...)
	_, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatal("nextYaml을 파싱할 수 없습니다")
	}
	return nextYaml
}

func AddModel(src DbtSchemaYaml, modelEntry Model) DbtSchemaYaml {
	nextYaml := DbtSchemaYaml(src)
	sameModelNameExists := funk.IndexOf(nextYaml.Models, func(model Model) bool {
		return model.Name == modelEntry.Name
	}) != -1
	if sameModelNameExists != false {
		log.Fatalf("같은 이름은 Model이 이미 존재합니다 [%v]", modelEntry.Name)
	}
	nextYaml.Models = append(nextYaml.Models, modelEntry)
	_, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatal("nextYaml을 파싱할 수 없습니다")
	}
	return nextYaml
}

func DeleteModel(src DbtSchemaYaml, modelName string) DbtSchemaYaml {
	nextYaml := DbtSchemaYaml(src)
	modelIndex := funk.IndexOf(nextYaml.Models, func(model Model) bool {
		return model.Name == modelName
	})
	if modelIndex == -1 {
		log.Fatalf("같은 이름의 Model을 찾을 수 없습니다 [%v]", modelName)
	}
	nextYaml.Models = append(nextYaml.Models[:modelIndex], nextYaml.Models[modelIndex+1:]...)
	_, err := yaml.Marshal(&nextYaml)
	if err != nil {
		log.Fatal("nextYaml을 파싱할 수 없습니다")
	}
	return nextYaml
}
