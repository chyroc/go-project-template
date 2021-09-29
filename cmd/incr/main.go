package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/chyroc/go-project-template"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "incr",
		Usage: "incr int",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			i, err := strconv.ParseInt(c.Args().First(), 10, 64)
			if err != nil {
				return err
			}
			fmt.Println(go_project_template.Incr(int(i)))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
