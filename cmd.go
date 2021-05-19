package dbt_yaml_builder

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func ModelCommand(c *cli.Context) error {
	path := c.Args().Get(0)
	yaml := ReadDbtSchemaYaml(path)
	fmt.Printf("%v", yaml)

	return nil
}
