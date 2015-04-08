package main

import (
	"fmt"
	"log"
	"path"

	"github.com/codegangsta/cli"
)

var cmdSearch = func(c *cli.Context) {
	client, err := getClient(c)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Search(c.String("query"), 1, 100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of Repositories: %d\n", res.NumberOfResults)
	for _, r := range res.Results {
		fmt.Printf(" -  Name: %s\n    Tags: %d\n    Layers: %d\n",
			path.Join(r.Namespace, r.Repository), len(r.Tags), len(r.Layers),
		)
	}

}
