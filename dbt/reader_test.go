package dbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadDbtSourceYaml(t *testing.T) {
	yaml := ReadDbtSourceYaml("./tests/dbt_source_gitlab.yaml")
	firstSourceTableName := yaml.Sources[0].SourceTables[0].Name
	assert.Equal(t, firstSourceTableName, "events")
}

func Test_ReadDbtSchemaYaml(t *testing.T) {
	yaml := ReadDbtSchemaYaml("./tests/dbt_schema_gitlab.yaml")
	mart_arr_columns := yaml.Models[0].Columns
	assert.GreaterOrEqual(t, len(mart_arr_columns), 9)
}
