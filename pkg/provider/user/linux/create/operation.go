package create

import (
)

type OperationOption func(*Operation)

type Operation struct {
	UID string
	GID string
	Username string
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

func New(name string, options ...OperationOption) *Operation {
	operation := &Operation{
		Username: name,
	}

	for _, option := range options {
		option(operation)
	} 
}

func WithBash() OperationOption {
	return func(o *Operation) {
		o.Shell = "/usr/bin/bash"
	}
}

func WithZsh() OperationOption {
	return func(o *Operation) {
		o.Shell = "/usr/bin/zsh"
	}
}

func System() OperationOption {
	return func(o *Operation) {
		o.System = true
	}
}

func Groups(list []string) OperationOption {
	return func(o *Operation) {
		o.Groups = list
	}
}

func Group(group string) OperationOption {
	return func(o *Operation) {
		o.Group = group
	}
}
