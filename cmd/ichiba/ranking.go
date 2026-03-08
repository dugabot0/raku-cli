package ichiba

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newRankingCmd() *cobra.Command {
	var (
		genreID string
		age     string
		sex     string
		carrier string
	)

	c := &cobra.Command{
		Use:   "ranking",
		Short: "Get Ichiba item ranking",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.IchibaRanking(client.IchibaRankingParams{
				GenreID: genreID,
				Age:     age,
				Sex:     sex,
				Carrier: carrier,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&genreID, "genre-id", "", "Genre ID to rank within")
	f.StringVar(&age, "age", "", "Age group: 10/20/30/40/50")
	f.StringVar(&sex, "sex", "", "Sex: 0=all, 1=male, 2=female")
	f.StringVar(&carrier, "carrier", "", "Carrier: 0=all, 1=PC, 2=mobile, 3=smartphone")

	return c
}
