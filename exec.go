package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteTask(execCmd string) {
	execParts := strings.Split(execCmd, " ")
	execName := execParts[0]
	execParams := strings.Join(execParts[1:], " ")
	cmd := exec.Command(execName, execParams)

	if err := cmd.Run(); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		fmt.Println(err.Error())
	}
}
