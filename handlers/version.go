package handlers

import (
		"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

type VersionHandler struct {
	SubCommand *kingpin.CmdClause
}

func NewVersionHandler(app *kingpin.Application) VersionHandler {
	subCommand := app.Command("version", "show aws-profile-utils version")
	return VersionHandler {
		SubCommand: subCommand,
	}
}

var version = "undefined"
func (handler VersionHandler) Handle() (bool, string) {
	fmt.Printf("aws-profile-utils (%s)", version)
	return true, ""
}
