package misc

import "github.com/spf13/cobra"

// NewCmd returns the misc command group (Recipe, Kobo, GORA).
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "misc",
		Short: "Miscellaneous Rakuten API commands (Recipe, Kobo, GORA)",
	}
	cmd.AddCommand(newRecipeCmd())
	cmd.AddCommand(newKoboCmd())
	cmd.AddCommand(newGoraCmd())
	return cmd
}
