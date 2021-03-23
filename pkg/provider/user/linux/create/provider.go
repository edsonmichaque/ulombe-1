package create

import (
)

type ProviderOption func(*Provider)

type Provider struct {
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

func New(name string, options ...ProviderOption) *Provider {
	provider := &Provider{
		Username: name,
	}

	for _, option := range options {
		option(provider)
	} 
}

func WithBash() ProviderOption {
	return func(p *Provider) {
		o.Shell = "/usr/bin/bash"
	}
}

func WithZsh() ProviderOption {
	return func(p *Provider) {
		o.Shell = "/usr/bin/zsh"
	}
}

func System() ProviderOption {
	return func(p *Provider) {
		o.System = true
	}
}

func WithGroups(list []string) ProviderOption {
	return func(p *Provider) {
		o.Groups = list
	}
}

func WithGroup(group string) ProviderOption {
	return func(p *Provider) {
		o.Group = group
	}
}
