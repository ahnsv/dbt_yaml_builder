package dbt_yaml_builder

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "language for the greeting",
				EnvVars: []string{"APP_LANG"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "add_model",
				Aliases: []string{"m"},
				Usage:   "모델 추가",
				Action:  ModelCommand,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
