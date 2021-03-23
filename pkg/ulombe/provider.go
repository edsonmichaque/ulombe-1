package ulombe

import (
	userProvider "github.com/edsonmichaque/ulombe/pkg/providers/user"
)

type ProviderList struct {
	providers map[string]*Provider
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
