package user

import (
	"errors"
	"github.com/edsonmichaque/ulombe/pkg/types"
	linuxProvider "github.com/edsonmichaque/ulombe/pkg/providers/user/linux"
)

type User struct {
}

func New() *User {
	return &User{}
}

func(u User) Name() string {
	return "user"
}

func(u User) Handles() map[]string {
	return []string{
		"linux",
	}
}

func(u User) Execute(levels ...string, args ...types.Argument) (types.Command, error) {
	if len(level) != 2 {
		return nil, errors.New("Must have 2 levels")
	}

	switch levels[0] {
	case "linux":
		switch levels[1] {
		case "create":
			return linuxHandler.Init(args...)
		}
	}

	return nil, errors.New("Cannot find provider")
}
