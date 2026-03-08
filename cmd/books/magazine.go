package books

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newMagazineCmd() *cobra.Command {
	var p client.BooksSearchParams

	c := &cobra.Command{
		Use:   "magazine",
		Short: "Search magazines",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.BooksMagazineSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.Title, "title", "", "Magazine title")
	f.StringVar(&p.PublisherName, "publisher", "", "Publisher name")
	f.StringVar(&p.JAN, "jan", "", "JAN barcode")
	f.StringVar(&p.BooksGenreID, "genre-id", "007", "Books genre ID (007=magazine)")
	f.StringVar(&p.Sort, "sort", "", "Sort: standard/sales/+releaseDate/-releaseDate/+itemPrice/-itemPrice/reviewCount/reviewAverage")
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")
	f.IntVar(&p.Availability, "availability", 0, "0=all, 1=in stock, 2-6=various shipping times")
	f.IntVar(&p.OutOfStockFlag, "out-of-stock-flag", 0, "1=include out-of-stock items")
	f.IntVar(&p.ChirayomiFlag, "chirayomi-flag", 0, "1=preview (立ち読み) items only")
	f.IntVar(&p.LimitedFlag, "limited-flag", 0, "1=limited editions only")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC, 1=mobile")
	f.IntVar(&p.GenreInformationFlag, "genre-information-flag", 0, "1=include genre item counts in response")

	return c
}
