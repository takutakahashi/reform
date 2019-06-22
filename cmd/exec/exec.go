package exec

import (
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:        "execute",
	Usage:       "tbd",
	Description: "execute command",
	Aliases:     []string{"exec"},
	Subcommands: Commands,
}

var Commands = []cli.Command{
	commandList,
	commandAdd,
	commandDelete,
}

var commandList = cli.Command{
	Name:        "list",
	Usage:       "tbd",
	Description: "listing repository",
	Action:      doList,
}

func doList(c *cli.Context) {}

var commandAdd = cli.Command{
	Name:        "add",
	Usage:       "tbd",
	Description: "adding repository",
	Action:      doAdd,
}

func doAdd(c *cli.Context) {}

var commandDelete = cli.Command{
	Name:        "delete",
	Usage:       "tbd",
	Description: "deleting repository",
	Action:      doDelete,
}

func doDelete(c *cli.Context) {}
