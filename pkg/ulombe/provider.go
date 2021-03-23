package ulombe

import (
	userProvider "github.com/edsonmichaque/ulombe/pkg/providers/user"
)

type ProviderList struct {
	providers map[string]*Provider
}

type Provider interface {
	Name() string
	Handles() map[string]interface{}
	Execute(...string, ...Argument) (*Command, error)
}

type Command interface {
	Script() ([]byte, error)
}

type Argument struct {
	Name string
	Value interface{}
}

func(p *ProviderList) Add(newProvider *Provider) {
	p.providers[newProvider.Name()] = newProvider
}

func New(providers ...Provider) *ProviderList {
	list := new(ProviderList)
	list.providers = make([map[string]*provider, 0)

	for _, provider := range providers {
		list.Add(provider)
	}

	return list
}
