package command

import (
	"fmt"
	"os"
)

func Execute(version, build, t string) {

	this.Build = build
	this.Time = t
	this.Version = version

	if _, err := rootCmd.ExecuteC(); err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)

	}

}
