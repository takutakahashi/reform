package exec

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	commandList,
	commandAdd,
	commandDelete,
}

var commandList = cli.Command{
	Name:        "list",
	usage:       "tbd",
	Description: "listing repository",
	Action:      doList,
}

func doList(c *cli.Context) {}

var commandAdd = cli.Command{
	Name:        "add",
	usage:       "tbd",
	Description: "adding repository",
	Action:      doAdd,
}

func doAdd(c *cli.Context) {}

var commandDelete = cli.Command{
	Name:        "delete",
	usage:       "tbd",
	Description: "deleting repository",
	Action:      doDelete,
}

func doDelete(c *cli.Context) {}
