package handler

import (
	"github.com/edsonmichaque/ulombe/pkg/ulombe"
)

const (
	providerName = "user"
)

type UpdateUserOption func(*UpdateUser)

type UpdateUser struct {
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

func New(name string, options ...UpdateUserOption) *UpdateUser {
	h := &UpdateUser{
		Name: name,
	}

	for _, option := range options {
		option(handler)
	} 
}

func Init(args ...types.ProviderArgument) *Command {
	user := UpdateUser{}

	for _, arg := range args {
		if arg.Name == "name" {
			user.Name = arg.Value.(string)
		}

		if arg.Name == "shell" {
			user.Shell = arg.Value.(string)
		}

		if arg.Name == "password" {
			user.Password = arg.Value.(string)
		}

		if arg.Name == "comment" {
			user.Comment = arg.Value.(string)
		}
	}

	return &provider
}

func WithBash() UpdateUserOption {
	return func(u *UpdateUser) {
		u.Shell = "/usr/bin/bash"
	}
}

func WithZsh() UpdateUserOption {
	return func(u *UpdateUser) {
		u.Shell = "/usr/bin/zsh"
	}
}

func System() UpdateUserOption {
	return func(u *UpdateUser) {
		u.System = true
	}
}

func WithGroups(list []string) UpdateUserOption {
	return func(u *UpdateUser) {
		u.Groups = list
	}
}

func WithGroup(group string) UpdateUserOption {
	return func(u *UpdateUser) {
		u.Group = group
	}
}

func (p UpdateUser) validate() error {
}

func (p UpdateUser) command() string {
	command := make([]string, 0)
	
	command = append(command, "usermod")

	if u.System {
		command = append(command, "-r")
	}

	if u.Shell != "" {
		command = append(command, "-s", u.Shell)
	}

	if u.Group != "" {
		command = append(command, "-g", u.Group)
	}

	if len(u.Groups) > 0 {
		command = append(command, "-G", strings.Join(u.Groups, ","))
	}

	if u.Comment != "" {
		command = append(command, "-c", fmt.Sprintf("\"%s\"", u.Comment))
	}

	if u.Name != "" {
		command = append(command, u.Name)
	}

	if u.Home && (u.HomeDir == "") {
		command = append(command, "-m")
		u.HomeDir = fmt.Sprintf("/home/%s", u.Name)
	}

	if u.HomeDir != "" {
		command = append(command, "-m")
		command = append(command, "-d", u.HomeDir)
	}

	return strings.Join(command, " ")
}

func (u *UpdateUser) Script() {
	tmpl := template.Must(template.ParseFiles("templates/update.tmpl"))
	tmpl.Execute(os.Stdout, p)
}

func (u *UpdateUser) Info() {
	return types.ProviderInfo{
		Name: providerName,
	}
}