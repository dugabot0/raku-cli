package books

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newSearchCmd() *cobra.Command {
	var p client.BooksSearchParams

	c := &cobra.Command{
		Use:   "search",
		Short: "Search across all Books media types (BooksTotal)",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.BooksTotalSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.Keyword, "keyword", "", "Search keyword (max 128 bytes)")
	f.StringVar(&p.IsbnJan, "isbnjan", "", "13-digit ISBN/JAN code (overrides keyword and genre)")
	f.StringVar(&p.BooksGenreID, "genre-id", "000", "Books genre ID (000=all)")
	f.StringVar(&p.MediaType, "media-type", "", "0=all, 1=book, 2=cd, 3=dvd, 4=foreign, 5=game, 6=magazine")
	f.StringVar(&p.Sort, "sort", "", "Sort: standard/sales/+releaseDate/-releaseDate/+itemPrice/-itemPrice/reviewCount/reviewAverage")
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")
	f.IntVar(&p.Availability, "availability", 0, "0=all, 1=in stock, 2=3-7 days, 3=3-9 days, 4=backorder, 5=pre-order, 6=check needed")
	f.IntVar(&p.OutOfStockFlag, "out-of-stock-flag", 0, "1=include out-of-stock items")
	f.IntVar(&p.ChirayomiFlag, "chirayomi-flag", 0, "1=preview (立ち読み) items only")
	f.IntVar(&p.LimitedFlag, "limited-flag", 0, "1=limited editions only")
	f.IntVar(&p.Field, "field", 0, "0=broad search, 1=narrow search")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC, 1=mobile")
	f.IntVar(&p.OrFlag, "or-flag", 0, "0=AND search, 1=OR search")
	f.StringVar(&p.NGKeyword, "ng-keyword", "", "Keywords to exclude")
	f.IntVar(&p.GenreInformationFlag, "genre-information-flag", 0, "1=include genre item counts in response")

	return c
}
