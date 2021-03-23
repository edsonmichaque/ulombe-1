package create

import (
	"fmt"
	"strings"
	"text/template"
)

type Decorator struct {
	Provider *Provider
}

func (d *Decorator) Command() string {
	command := make([]string, 0)
	
	command = append(command, "useradd")

	if d.Provider.System {
		command = append(command, "-r")
	}

	if d.Provider.Shell != "" {
		command = append(command, "-s", d.Provider.Shell)
	}

	if d.Provider.Group != "" {
		command = append(command, "-g", d.Provider.Group)
	}

	if len(d.Provider.Groups) > 0 {
		command = append(command, "-G", strings.Join(d.Provider.Groups, ","))
	}

	if d.Provider.Comment != "" {
		command = append(command, "-c", fmt.Sprintf("\"%s\"", d.Provider.Comment))
	}

	if d.Provider.Username != "" {
		command = append(command, d.Provider.Username)
	}

	if d.Provider.Home && (d.Provider.HomeDir == "") {
		d.Provider.HomeDir = fmt.Sprintf("/home/%s", d.Provider.Username)
	}

	if d.Provider.HomeDir != "" {
		command = append(command, "-d", d.Provider.HomeDir)
	}

	return strings.Join(command, " ")
}

func (d *Decorator) Script() {
	tmpl := template.Must(template.ParseFiles("script.tmpl"))
	tmpl.Execute(os.Stdout, d)
}

