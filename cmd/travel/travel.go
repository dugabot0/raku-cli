package travel

import "github.com/spf13/cobra"

// NewCmd returns the travel command group.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "travel",
		Short: "Rakuten Travel API commands",
	}
	cmd.AddCommand(newHotelsCmd())
	cmd.AddCommand(newHotelCmd())
	cmd.AddCommand(newVacantCmd())
	cmd.AddCommand(newAreaCmd())
	cmd.AddCommand(newRankingCmd())
	return cmd
}
