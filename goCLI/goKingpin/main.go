// Credit: https://developer.atlassian.com/blog/2016/03/building-helpful-golang-cli-tools/
package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func addSubCommand(app *kingpin.Application, name string, description string) {
	app.Command(name, description).Action(func(c *kingpin.ParseContext) error {
		fmt.Printf("Would have run command %s.\n", name)
		return nil
	})
}

func main() {
	app := kingpin.New("goKingpin", "My pretend app CLI with Bash completion")

	app.Flag("flag-1", "").String()
	app.Flag("flag-2", "").HintOptions("opts1", "opts2").String()

	// Add some additional top level commands
	addSubCommand(app, "ls", "Additional top level command to show command completion")
	addSubCommand(app, "ping", "Additional top level command to show command completion")
	addSubCommand(app, "nmap", "Additional top level command to show command completion")

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
