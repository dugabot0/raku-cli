package books

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newGenreCmd() *cobra.Command {
	var p client.BooksGenreSearchParams

	c := &cobra.Command{
		Use:   "genre",
		Short: "Search Books genre information",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.BooksGenreSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.GenreID, "genre-id", "001", "Books genre ID (000=root)")
	f.IntVar(&p.GenrePath, "genre-path", 0, "1=include ancestor genres in response")

	return c
}
