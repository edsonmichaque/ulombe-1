package create

import (
)

type HandlerOption func(*Handler)

type Handler struct {
	UID string
	GID string
	Name string
	Password string
	Expire string
	Comment string
	Group string
	Groups []string
	System bool
	Shell string
	Home bool
	HomeDir string
}

func New(name string, options ...HandlerOption) *Handler {
	handler := &Handler{
		Name: name,
	}

	for _, option := range options {
		option(handler)
	} 
}

func WithBash() HandlerOption {
	return func(p *Handler) {
		p.Shell = "/usr/bin/bash"
	}
}

func WithZsh() HandlerOption {
	return func(p *Handler) {
		p.Shell = "/usr/bin/zsh"
	}
}

func System() HandlerOption {
	return func(p *Handler) {
		p.System = true
	}
}

func WithGroups(list []string) HandlerOption {
	return func(p *Handler) {
		p.Groups = list
	}
}

func WithGroup(group string) HandlerOption {
	return func(p *Handler) {
		p.Group = group
	}
}

func (p Handler) validate() error {
}

func (p Handler) command() string {
	command := make([]string, 0)
	
	command = append(command, "useradd")

	if p.System {
		command = append(command, "-r")
	}

	if p.Shell != "" {
		command = append(command, "-s", p.Shell)
	}

	if p.Group != "" {
		command = append(command, "-g", p.Group)
	}

	if len(p.Groups) > 0 {
		command = append(command, "-G", strings.Join(p.Groups, ","))
	}

	if p.Comment != "" {
		command = append(command, "-c", fmt.Sprintf("\"%s\"", p.Comment))
	}

	if p.Name != "" {
		command = append(command, p.Name)
	}

	if p.Home && (p.HomeDir == "") {
		command = append(command, "-m")
		p.HomeDir = fmt.Sprintf("/home/%s", p.Name)
	}

	if p.HomeDir != "" {
		command = append(command, "-m")
		command = append(command, "-d", p.HomeDir)
	}

	return strings.Join(command, " ")
}

func (p *Handler) Script() {
	tmpl := template.Must(template.ParseFiles("script.tmpl"))
	tmpl.Execute(os.Stdout, p)
}