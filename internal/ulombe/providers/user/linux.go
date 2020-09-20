package user

import (
	"fmt"
	"strings"
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
	attrExpireDate = "expire_date"
	attrHome = "home"
)

var aliases = map[string][]string{
	OpCreate: []string{
		"new",
	},
	OpUpdate: []string{
		"modify",
	},
	Opdelete: []string{
		"remove",
	},
}

testUserExists := `
if getent passwd <user> > /dev/null; then
    <command>
else
    printf "User %s doens't exist\n", <user>
fi
`

testUserDoenstExist := `
if getent passwd  <user> > /dev/null; then
    printf "User %s already exists\n", <user> 
else
    <command>
fi
`

Linux.AddScriptGenerator(func(p *Provider) func(op string, data map[string]interface{}) string {
	return func(op string, data map[string]interface{}) string {
		username, _ := data[attrUsername]

		if p.GetOperation(op) == opCreate {
			command := "useradd"

		} else if operation := p.GetOperation(op);  operation.Name == opUpdate {
			command := "usermod"

			command = fmt.Sprintf("%s %s", command, username.(string))

			if uid, ok := data[attrUid]; ok {
				command = fmt.Sprintf("%s -u %s", command, uid.(string))
			}

			if gid, ok := data[attrGid]; ok {
				command = fmt.Sprintf("%s -g %s", command, gid.(string))
			}

			if system, ok := data[attrSystem]; ok && system.(bool) {
				command = fmt.Sprintf("%s -r", command)
			}

			if shell, ok := data[attrShell]; ok {
				command = fmt.Sprintf("%s -s %s", command, shell.(string))
			}

			r := strings.NewReplacer(
				"<user>", username, "<command>", command
			)

			return r.Replace(testUserExist)

		} else if p.GetOperation(op) == opDelete {
			command = "userdel"

		} else {

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


