package create

import (
	"fmt"
	"strings"
	"text/template"
)

type Decorator struct {
	Operation *Operation
}

func (d *Decorator) Command() string {
	
	command := make([]string, 0)
	
	command = append(command, "useradd")

	if d.Operation.System {
		command = append(command, "-r")
	}

	if d.Operation.Shell != "" {
		command = append(command, "-s", d.Operation.Shell)
	}

	if d.Operation.Group != "" {
		command = append(command, "-g", d.Operation.Group)
	}

	if len(d.Operation.Groups) > 0 {
		command = append(command, "-G", strings.Join(d.Operation.Groups, ","))
	}

	if d.Operation.Comment != "" {
		command = append(command, "-c", fmt.Sprintf("\"%s\"", d.Operation.Comment))
	}

	if d.Operation.Username != "" {
		command = append(command, d.Operation.Username)
	}

	if d.Operation.Home && (d.Operation.HomeDir == "") {
		d.Operation.HomeDir = fmt.Sprintf("/home/%s", d.Operation.Username)
	}

	if d.Operation.HomeDir != "" {
		command = append(command, "-d", d.Operation.HomeDir)
	}


	return strings.Join(command, " ")
}

func (d *Decorator) Script() {
	tmpl := template.Must(template.ParseFiles("script.tmpl"))
	tmpl.Execute(os.Stdout, d)
}

