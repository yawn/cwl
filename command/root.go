package command

import (
	"github.com/spf13/cobra"
)

const App = "cwl"

var (
	this = Version{}
)

var rootCmd = &cobra.Command{
	Use: App,
}
