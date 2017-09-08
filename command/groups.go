package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yawn/cwl"
)

var groupsFlags = struct {
	prefix string
}{}

var groupsCmd = &cobra.Command{

	Use:   "groups",
	Short: "Show all log groups",
	RunE: func(cmd *cobra.Command, args []string) error {

		clients := cwl.NewClients()

		result, err := clients.FindGroups(groupsFlags.prefix)

		if err != nil {
			return err
		}

		for region, groups := range result {

			for _, group := range groups {
				fmt.Println(strings.Join([]string{group, region}, "@"))
			}

		}

		return nil

	},
}

func init() {

	groupsCmd.Flags().StringVarP(&groupsFlags.prefix, "prefix", "p", "", "Prefix for filtering log groups")

	rootCmd.AddCommand(groupsCmd)

}
