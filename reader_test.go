package dbt_yaml_builder

import (
	"fmt"
	"testing"
)

func Test_ReadDbtSourceYaml(t *testing.T) {
	yaml := ReadDbtSourceYaml("./tests/dbt_source_gitlab.yaml")
	fmt.Print(yaml)
}
