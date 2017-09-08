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
	RunE: func(cmd *cobra.Command, args []string) error {

		clients := cwl.NewClients()

		if len(args) != 1 {
			return errors.New("missing log group")
		}

		parts := strings.Split(args[0], "@")

		if len(parts) != 2 {
			return errors.New("malformed log group")
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
