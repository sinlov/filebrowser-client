package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// Version of cli
var Version = "v0.1.2"

// action
// do cli Action before flag.
func action(c *cli.Context) error {
	return nil
}

// pluginFlag
// set plugin flag at here
func pluginFlag() []cli.Flag {
	return []cli.Flag{
		// plugin start
		// new flag string template if no use, please replace this
		&cli.StringFlag{
			Name:    "config.new_arg,new_arg",
			Usage:   "",
			EnvVars: []string{"PLUGIN_new_arg"},
		},

		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		// plugin end
	}
}

// pluginHideFlag
// set plugin hide flag at here
func pluginHideFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = "filebrowser client"
	app.Usage = "client use base web api for https://github.com/filebrowser/filebrowser"
	year := time.Now().Year()
	app.Copyright = fmt.Sprintf("© 2022-%d sinlov", year)
	author := &cli.Author{
		Name:  "sinlov",
		Email: "sinlovgmppt@gmail.com",
	}
	app.Authors = []*cli.Author{
		author,
	}

	app.Action = action
	flags := appendCliFlag(pluginFlag(), pluginHideFlag())
	app.Flags = flags

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}

// appendCliFlag
// append cli.Flag
func appendCliFlag(target []cli.Flag, elem []cli.Flag) []cli.Flag {
	if len(elem) == 0 {
		return target
	}

	return append(target, elem...)
}
