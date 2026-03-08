package ichiba

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newGenreCmd() *cobra.Command {
	var genreID string

	c := &cobra.Command{
		Use:   "genre",
		Short: "Search Ichiba genre information",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadIchibaClient()
			result, err := cl.IchibaGenreSearch(client.IchibaGenreSearchParams{
				GenreID: genreID,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&genreID, "genre-id", "0", "Genre ID (0 = root)")

	return c
}
