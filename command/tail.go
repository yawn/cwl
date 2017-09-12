package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yawn/cwl"
)

var tailFlags = struct {
	json bool
}{}

var tailCmd = &cobra.Command{

	Use:   "tail",
	Short: "Tail a log group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			clients  = cwl.NewClients()
			callback = cwl.CallbackTabs
		)

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

		if tailFlags.json {
			callback = cwl.CallbackJSON
		}

		return client.FindEvents(callback, group)

	},
}

func init() {

	tailCmd.Flags().BoolVarP(&tailFlags.json, "json", "j", false, "Output logs as JSON")

	rootCmd.AddCommand(tailCmd)
}
