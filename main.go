package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	registry "github.com/shipyard/shipyard/registry/v1"
)

func getClient(c *cli.Context) (*registry.RegistryClient, error) {
	return registry.NewRegistryClient(c.GlobalString("url"), nil)
}

func main() {
	app := cli.NewApp()
	app.Name = "registry-manager"
	app.Usage = "docker registry management"
	app.Version = "0.0.1"
	app.Author = ""
	app.Email = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "registry url",
			Value: "http://localhost:5000",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "search",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "query, q",
					Usage: "query to search",
					Value: "",
				},
			},
			Action: cmdSearch,
		},
		{
			Name: "delete-repo",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of repo",
					Value: "",
				},
			},
			Action: cmdDeleteRepo,
		},
		{
			Name: "delete-tag",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of repo",
					Value: "",
				},
				cli.StringFlag{
					Name:  "tag, t",
					Usage: "name of tag",
					Value: "latest",
				},
			},
			Action: cmdDeleteTag,
		},
		{
			Name: "inspect-repo",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of repo",
					Value: "",
				},
			},
			Action: cmdInspectRepo,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
