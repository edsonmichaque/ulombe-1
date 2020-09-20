package user

import (
	fmt
	ulombe "gitlab.com/ulombe/sdk"
	provider "gitlab.com/ulombe/sdk/provider"
)

Linux := ulombe.NewProvider("linux")

const (
	opCreate = "create"
	opUpdate = "update"
	opDelete = "delete"
)

const (
	attrUsername = "username"
	attrComment = "comment"
	attrSystem = "system"
	attrShell = "shell"
	attrGroup = "group"
	attrUid = "uid"
	attrGid = "gid"
	attrGroups = "groups"
)

var aliases = map[string][]string{
	OpCreate: []string{"new"},
	OpUpdate: []string{"modify"},
	Opdelete: []string{"remove"}
}

Linux.AddScriptGenerator(func(p *Provider) func(op string, data map[string]interface{}) string {
	return func(op string, data map[string]interface{}) string {

		if p.GetOperation(op) == opCreate {


		} else if p.GetOperation(op) == opUpdate {

			baseCommand = "usermod"

		} else if p.GetOperation(op) == opDelete {

			baseCommand = "userdel"

		} else {
			// Return Error
		}
	}
});

Linux.AddOperation(opCreate, aliases[opCreate], []Attribute{
	provider.NewAttribute(attrComment, provider.String),
	provider.NewRequiredAttribute(attrUsername, provider.String),
	provider.NewAttribute(attrPassword, provider.String),
	provider.NewAttribute(attrUid, provider.String),
	provider.NewAttribute(attrGid, provider.String),
	provider.NewAttribute(attrGroups, provider.ListOfString),
	provider.NewAttribute(attrShell, provider.String),
	provider.NewAttribute(attrSystem, provider.Boolean),
})

