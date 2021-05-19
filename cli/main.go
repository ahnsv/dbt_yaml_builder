package main

import (
	"log"
	"os"

	"github.com/ahnsv/dbt_yaml_builder"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dbt_yaml_builder",
		Usage: "dbt 생산성 향상을 위한 cli 툴",
		Commands: []*cli.Command{
			{
				Name:    "add_model",
				Aliases: []string{"model"},
				Usage:   "모델 추가",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
					},
					&cli.StringFlag{
						Name:    "description",
						Value:   "",
						Aliases: []string{"d"},
					},
					&cli.StringSliceFlag{
						Name:    "columns",
						Value:   cli.NewStringSlice(),
						Aliases: []string{"c"},
					},
				},
				Action: dbt_yaml_builder.AddModelActionHandler,
			},
			{
				Name:    "delete_model",
				Aliases: []string{"dm"},
				Usage:   "모델 삭제",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
					},
				},
				Action: dbt_yaml_builder.DeleteModelActionHandler,
			},
		},
	}
	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
