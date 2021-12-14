package commands

import (
	"github.com/abiosoft/ishell"
)

func Xk(shell *ishell.Shell) (f func(c *ishell.Context)) {
	return func(c *ishell.Context) {
		errString, ok := syntaxCheck(c, 1)
		path := buildPath(c.Args)
		if !ok {
			shell.Println(errString)
			return
		}
		if err := getFromEntry(shell, path, "password"); err != nil {
			shell.Printf("could not copy password: %s", err)
			return
		}
	}
}
