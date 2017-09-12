package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yawn/cwl"
)

var tailFlags = struct {
}{}

var tailCmd = &cobra.Command{

	Use:   "tail",
	Short: "Tail a log group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {


		parts := strings.Split(args[0], "@")

		if len(parts) != 2 {
			return fmt.Errorf("malformed log group %q", args[0])
		}

		var (
			group  = parts[0]
			region = parts[1]
		)

		client, ok := clients[region]

		if !ok {
			return fmt.Errorf("no client for region %q", region)
		}

		return client.FindEvents(group)

	},
}

func init() {

	rootCmd.AddCommand(tailCmd)

}
