package main

import (
	"fmt"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// Version of cli
var Version = "v0.1.2"
var (
	fileBrowserClient file_browser_client.FileBrowserClient
	isCliDebug        = false
)

// action
// do cli Action before flag.
func action(c *cli.Context) error {
	//err := publicFlagsCheckAndInitFileBrowserClient(c)
	//if err != nil {
	//	return err
	//}
	return nil
}

// publicFlags
// set public flag at here
func publicFlags() []cli.Flag {
	return []cli.Flag{
		// plugin start
		&cli.StringFlag{
			Name:    "config.file_browser_host,file_browser_host",
			Usage:   "must set args, file_browser username",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_HOST"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_username,file_browser_username",
			Usage:   "must set args, file_browser username",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_user_password,file_browser_user_password",
			Usage:   "must set args, file_browser user password",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_USER_PASSWORD"},
		},
		&cli.UintFlag{
			Name:    "config.file_browser_timeout_push_second,file_browser_timeout_push_second",
			Usage:   "file_browser push each file timeout push second",
			Value:   60,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TIMEOUT_PUSH_SECOND"},
		},

		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Aliases: []string{"debug"},
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		// plugin end
	}
}

func publicFlagsCheckAndInitFileBrowserClient(context *cli.Context) error {
	host, err := checkFlagStringIsEmpty(context, "config.file_browser_host")
	if err != nil {
		return err
	}
	userName, err := checkFlagStringIsEmpty(context, "config.file_browser_username")
	if err != nil {
		return err
	}
	passwd, err := checkFlagStringIsEmpty(context, "config.file_browser_user_password")
	if err != nil {
		return err
	}
	timeOutSecond := context.Uint("config.timeout_second")
	timeOutFileSecond := context.Uint("config.file_browser_timeout_push_second")
	fileBrowserClient, err = file_browser_client.NewClient(
		userName,
		passwd,
		host,
		timeOutSecond,
		timeOutFileSecond,
	)
	if err != nil {
		return err
	}
	isCliDebug := context.Bool("config.debug")
	fileBrowserClient.Debug(isCliDebug)
	if isCliDebug {
		log.Println("publicFlagsCheckAndInitFileBrowserClient finish")
	}
	return nil
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

// resourceFlag
// set resource flag at here
func resourceFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config.resource_post_remote_path,resource_post_remote_path",
			Usage:   "resource post remote path",
			EnvVars: []string{"PLUGIN_RESOURCE_POST_REMOTE_PATH"},
		},
		&cli.StringFlag{
			Name:    "config.resource_post_local_file,resource_post_local_file",
			Usage:   "resource post local file",
			EnvVars: []string{"PLUGIN_RESOURCE_POST_LOCAL_FILE"},
		},
		&cli.BoolFlag{
			Name:    "config.resource_post_override,resource_post_override",
			Usage:   "resource post override",
			Value:   false,
			EnvVars: []string{"PLUGIN_RESOURCE_POST_OVERRIDE"},
		},
		// new flag string template if no use, please replace this
		//&cli.StringFlag{
		//	Name:    "config.new_arg,new_arg",
		//	Usage:   "",
		//	EnvVars: []string{"PLUGIN_new_arg"},
		//},
	}
}

func resourceCmd() *cli.Command {
	flags := appendCliFlag(publicFlags(), pluginHideFlag())
	flags = appendCliFlag(flags, resourceFlag())
	return &cli.Command{
		Name:        "resource",
		Description: "file browser resource cli",
		Flags:       flags,
		HideHelp:    false,
		Hidden:      false,
		Before:      resourceBeforeAction,
		Action:      resourceAction,
	}
}

func resourceBeforeAction(context *cli.Context) error {
	err := publicFlagsCheckAndInitFileBrowserClient(context)
	if err != nil {
		return err
	}
	err = fileBrowserClient.Login()
	if err != nil {
		return err
	}
	return nil
}

func resourceAction(context *cli.Context) error {
	resFile := file_browser_client.ResourcePostFile{
		RemoteFilePath: context.String("config.resource_post_remote_path"),
		LocalFilePath:  context.String("config.resource_post_local_file"),
	}
	err := fileBrowserClient.ResourcesPostFile(resFile, context.Bool("config.resource_post_override"))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = "filebrowser client"
	app.Usage = "more info see[ help COMMAND ]"
	app.Description = "client use base web api for https://github.com/filebrowser/filebrowser"
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
	commands := []*cli.Command{
		resourceCmd(),
	}
	//appendCommand(commands, resourceCmd())
	app.Commands = commands
	flags := appendCliFlag(publicFlags(), pluginHideFlag())
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

// appendCliFlag
// append cli.Flag
func appendCommand(target []*cli.Command, elem *cli.Command) []*cli.Command {
	if elem == nil {
		return target
	}

	return append(target, elem)
}

func checkFlagStringIsEmpty(context *cli.Context, key string) (string, error) {
	valStr := context.String(key)
	if valStr == "" {
		return "", fmt.Errorf("flag string not set [ %s ] is empty", key)
	}
	return valStr, nil
}
