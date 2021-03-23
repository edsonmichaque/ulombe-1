package handler

import (
	"github.com/edsonmichaque/ulombe/pkg/types"
)

const (
	providerName = "user"
)

type CreateUserOption func(*CreateUser)

type CreateUser struct {
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

func New(name string, options ...CreateUserOption) *CreateUser {
	h := &CreateUser{
		Name: name,
	}

	for _, option := range options {
		option(handler)
	} 
}

func Init(args ...types.ProviderArgument) *Command {
	user := CreateUser{}

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

func WithBash() CreateUserOption {
	return func(u *CreateUser) {
		u.Shell = "/usr/bin/bash"
	}
}

func WithZsh() CreateUserOption {
	return func(u *CreateUser) {
		u.Shell = "/usr/bin/zsh"
	}
}

func System() CreateUserOption {
	return func(u *CreateUser) {
		u.System = true
	}
}

func WithGroups(list []string) CreateUserOption {
	return func(u *CreateUser) {
		u.Groups = list
	}
}

func WithGroup(group string) CreateUserOption {
	return func(u *CreateUser) {
		u.Group = group
	}
}

func (p CreateUser) validate() error {
}

func (p CreateUser) command() string {
	command := make([]string, 0)
	
	command = append(command, "useradd")

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

func (u *CreateUser) Script() {
	tmpl := template.Must(template.ParseFiles("templates/create.tmpl"))
	tmpl.Execute(os.Stdout, p)
}

func (u *CreateUser) Info() {
	return types.ProviderInfo{
		Name: providerName,
	}
}