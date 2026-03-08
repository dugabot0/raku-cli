package books

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newBookCmd() *cobra.Command {
	var (
		keyword string
		sort    string
		page    int
		hits    int
	)

	c := &cobra.Command{
		Use:   "book",
		Short: "Search books",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.BooksBookSearch(client.BooksSearchParams{
				Keyword: keyword,
				Sort:    sort,
				Page:    page,
				Hits:    hits,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&keyword, "keyword", "", "Search keyword")
	f.StringVar(&sort, "sort", "", "Sort order")
	f.IntVar(&page, "page", 1, "Page number")
	f.IntVar(&hits, "hits", 30, "Results per page (1-30)")

	return c
}
