package user

import (
	"fmt"
	"strings"
	compositor "github.com/edsonmichaque/artista/internal/compositor"
)

func Create(t compositor.Task) string {
	baseCommand := "useradd"
	var options string

	if name, ok := t.Data["name"]; ok {
		options = name.(string)
	}

	if shell, ok := t.Data["shell"]; ok {
		options = fmt.Sprintf("%s -s %s", options, shell.(string))
	}

	if home, ok := t.Data["home"]; ok {
		options = fmt.Sprintf("%s -d %s", options, home.(string))
	}

	if system, ok := t.Data["system"]; ok && system.(bool){
		options = fmt.Sprintf("%s -r", options)
	}

	if active, ok := t.Data["active"]; ok && !active.(bool) {
		options = fmt.Sprintf("%s -f", options)
	}

	if expire, ok := t.Data["expire"]; ok {
		options = fmt.Sprintf("%s -e \"%s\"", options, expire.(string))
	}

	if comment, ok := t.Data["comment"]; ok {
		options = fmt.Sprintf("%s -c \"%s\"", options, comment.(string))
	}

	if groups, ok := t.Data["groups"]; ok {
		groupList := make([]string,0)
		for _, group := range groups.([]interface{}) {
			groupList = append(groupList, group.(string))
		}

		options = fmt.Sprintf("%s -G %s", options, strings.Join(groupList, " "))
	}

	if uid, ok := t.Data["uid"]; ok {
		options = fmt.Sprintf("%s -u %s", options, uid)
	}

        if gid, ok := t.Data["gid"]; ok {
                options = fmt.Sprintf("%s -g %s", options, gid)
        }


	command := fmt.Sprintf("%s %s", baseCommand, options)
	command = fmt.Sprintf("echo \"%s\" && %s", t.Description, command)

	return command
}
