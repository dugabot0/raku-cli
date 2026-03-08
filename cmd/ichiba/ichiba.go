package ichiba

import "github.com/spf13/cobra"

// NewCmd returns the ichiba command group.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ichiba",
		Short: "Rakuten Ichiba (Shopping) API commands",
	}
	cmd.AddCommand(newItemsCmd())
	cmd.AddCommand(newGenreCmd())
	cmd.AddCommand(newRankingCmd())
	return cmd
}
