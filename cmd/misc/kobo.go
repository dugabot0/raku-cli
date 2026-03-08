package misc

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newKoboCmd() *cobra.Command {
	var (
		keyword string
		title   string
		author  string
		sort    string
		page    int
		hits    int
	)

	c := &cobra.Command{
		Use:   "kobo",
		Short: "Search Kobo ebooks",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.KoboEbookSearch(client.KoboSearchParams{
				Keyword: keyword,
				Title:   title,
				Author:  author,
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
	f.StringVar(&title, "title", "", "Book title")
	f.StringVar(&author, "author", "", "Author name")
	f.StringVar(&sort, "sort", "", "Sort order: standard/sales/+releaseDate/-releaseDate/+itemPrice/-itemPrice/reviewCount/reviewAverage")
	f.IntVar(&page, "page", 1, "Page number")
	f.IntVar(&hits, "hits", 30, "Results per page (1-30)")

	return c
}
