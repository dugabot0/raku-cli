package misc

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newKoboCmd() *cobra.Command {
	var p client.KoboSearchParams

	c := &cobra.Command{
		Use:   "kobo",
		Short: "Search Kobo ebooks",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.KoboEbookSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.Keyword, "keyword", "", "Search keyword")
	f.StringVar(&p.Title, "title", "", "Book title")
	f.StringVar(&p.Author, "author", "", "Author name")
	f.StringVar(&p.PublisherName, "publisher", "", "Publisher name")
	f.StringVar(&p.ItemNumber, "item-number", "", "Product number")
	f.StringVar(&p.KoboGenreID, "genre-id", "101", "Kobo genre ID (subcategories of 101)")
	f.StringVar(&p.Language, "language", "", "Language code")
	f.StringVar(&p.Sort, "sort", "", "Sort: standard/+releaseDate/-releaseDate/+itemPrice/-itemPrice/reviewCount/reviewAverage")
	f.StringVar(&p.NGKeyword, "ng-keyword", "", "Keywords to exclude")
	f.IntVar(&p.SalesType, "sales-type", 0, "0=regular, 1=pre-order")
	f.IntVar(&p.Field, "field", 0, "0=broad search, 1=narrow search")
	f.IntVar(&p.OrFlag, "or-flag", 0, "0=AND search, 1=OR search")
	f.IntVar(&p.GenreInformationFlag, "genre-information-flag", 0, "1=include genre item counts in response")
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")

	return c
}
