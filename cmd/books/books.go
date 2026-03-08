package books

import "github.com/spf13/cobra"

// NewCmd returns the books command group.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "books",
		Short: "Rakuten Books API commands",
	}
	cmd.AddCommand(newSearchCmd())
	cmd.AddCommand(newBookCmd())
	cmd.AddCommand(newCDCmd())
	cmd.AddCommand(newDVDCmd())
	cmd.AddCommand(newMagazineCmd())
	cmd.AddCommand(newGameCmd())
	cmd.AddCommand(newGenreCmd())
	return cmd
}
