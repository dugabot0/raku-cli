package ichiba

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newItemsCmd() *cobra.Command {
	var (
		keyword  string
		genreID  string
		minPrice int
		maxPrice int
		sort     string
		page     int
		hits     int
	)

	c := &cobra.Command{
		Use:   "items",
		Short: "Search Ichiba items",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.IchibaItemSearch(client.IchibaItemSearchParams{
				Keyword:  keyword,
				GenreID:  genreID,
				MinPrice: minPrice,
				MaxPrice: maxPrice,
				Sort:     sort,
				Page:     page,
				Hits:     hits,
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
	f.StringVar(&genreID, "genre-id", "", "Genre ID to search within")
	f.IntVar(&minPrice, "min-price", 0, "Minimum price (yen)")
	f.IntVar(&maxPrice, "max-price", 0, "Maximum price (yen)")
	f.StringVar(&sort, "sort", "", "Sort order: standard/+itemPrice/-itemPrice/+updateTimestamp/-updateTimestamp/reviewCount/reviewAverage/affiliateRate/+price/-price")
	f.IntVar(&page, "page", 1, "Page number")
	f.IntVar(&hits, "hits", 30, "Results per page (1-30)")

	return c
}
