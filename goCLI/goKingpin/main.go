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

	configureNetCatCommand(app)

	app.Flag("flag-1", "").String()
	app.Flag("flag-2", "").HintOptions("opts1", "opts2").String()

	// Add some additional top level commands
	addSubCommand(app, "ls", "Additional top level command to show command completion")
	addSubCommand(app, "ping", "Additional top level command to show command completion")
	addSubCommand(app, "nmap", "Additional top level command to show command completion")

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

// NetcatCommand ...
type NetcatCommand struct {
	hostName string
	port     int
	format   string
}

func (n *NetcatCommand) run(c *kingpin.ParseContext) error {
	fmt.Printf("Would have run netcat to hostname %v, port %d, and output format %v\n", n.hostName, n.port, n.format)
	return nil
}

func configureNetCatCommand(app *kingpin.Application) {
	c := &NetcatCommand{}
	nc := app.Command("nc", "Connect to a host").Action(c.run)
	nc.Flag("nop-flag", "Example of a flag with no options").Bool()

	// You can provide hint options statically
	nc.Flag("port", "Provide a port to connect to").
		Required().
		HintOptions("80", "443", "8080").
		IntVar(&c.port)

	// Enum/EnumVar options will be turned into completion options automatically
	nc.Flag("format", "Define the output format").
		Default("raw").
		EnumVar(&c.format, "raw", "json")

	nc.Flag("host", "Provide a hostnanme to nc").
		Required().
		HintAction(listHosts).
		StringVar(&c.hostName)
}

func listHosts() []string {
	return []string{"sshost.exmapkle", "webhost.Example", "ftphost.Example"}
}
