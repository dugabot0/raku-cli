package books

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newSearchCmd() *cobra.Command {
	var (
		keyword   string
		sort      string
		page      int
		hits      int
		mediaType string
	)

	c := &cobra.Command{
		Use:   "search",
		Short: "Search across all Books media types",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.BooksTotalSearch(client.BooksSearchParams{
				Keyword:   keyword,
				Sort:      sort,
				Page:      page,
				Hits:      hits,
				MediaType: mediaType,
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
	f.StringVar(&sort, "sort", "", "Sort order: standard/sales/+releaseDate/-releaseDate/+itemPrice/-itemPrice/reviewCount/reviewAverage")
	f.IntVar(&page, "page", 1, "Page number")
	f.IntVar(&hits, "hits", 30, "Results per page (1-30)")
	f.StringVar(&mediaType, "media-type", "", "Media type: 0=all, 1=book, 2=cd, 3=dvd, 4=foreign, 5=game, 6=magazine, 7=foreign-book")

	return c
}
