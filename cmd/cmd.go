package main

import (
	"github.com/takutakahashi/reform/cmd/exec"
	"github.com/takutakahashi/reform/cmd/repository"
	"github.com/urfave/cli"
	"os"
)

var Version string = "0.9.0"

var Commands = []cli.Command{
	repository.Command,
	exec.Command,
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
