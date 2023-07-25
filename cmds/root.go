package cmds

import (
	"github.com/alphauslabs/bluectl/pkg/logger"
	"github.com/spf13/cobra"
)

func ListOpsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Subcommand for CUR import/export operations",
		Long:  `Subcommand for CUR import/export operations.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("see -h for more information")
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
