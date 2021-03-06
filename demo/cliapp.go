package main

import (
	"runtime"
	"github.com/gookit/cliapp"
	"github.com/gookit/cliapp/demo/cmd"
	"github.com/gookit/cliapp/builtin"
)

// for test run: go build ./demo/cliapp.go && ./cliapp
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cliapp.NewApp()
	app.Version = "1.0.3"
	app.Description = "this is my cli application"

	app.SetVerbose(cliapp.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(cmd.ExampleCommand())
	app.Add(cmd.EnvInfoCommand())
	app.Add(cmd.GitCommand())
	app.Add(cmd.ColorCommand())
	app.Add(&cliapp.Command{
		Name:        "test",
		Aliases:     []string{"ts"},
		Description: "this is a description <info>message</> for command {$cmd}",
		Fn: func(cmd *cliapp.Command, args []string) int {
			cliapp.Print("hello, in the test command\n")
			return 0
		},
	})

	app.Add(builtin.GenShAutoComplete())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run()
}
