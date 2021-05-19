package dbt_yaml_builder

import funk "github.com/thoas/go-funk"

func AddNewSourceTable(src DbtSourceYaml, sourceName string, tableEntry SourceTables) DbtSourceYaml {
	sources := src.Sources
	selectedSource := funk.Find(sources, func(source Sources) bool {
		return source.Name == sourceName
	})
	nextSource := append(selectedSource.SourceTables, tableEntry) // FIXME
	return nextSource
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
