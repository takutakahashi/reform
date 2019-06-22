package main

import (
	"exec"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"os"
	"repository"
)

var Version string = "0.9.0"

var Commands = []cli.Command{
	repository.Commands,
	Exec.Commands,
}

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "reform"
	app.Usage = "wallpaper switcher"
	app.Version = Version
	app.Author = "takutakahashi"
	app.Email = "taku.takahashi120@gmail.com"
	app.Commands = Commands
	return app
}
