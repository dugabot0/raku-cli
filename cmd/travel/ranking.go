package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newRankingCmd() *cobra.Command {
	var genre string

	c := &cobra.Command{
		Use:   "ranking",
		Short: "Get hotel ranking",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelHotelRanking(client.TravelRankingParams{
				Genre: genre,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&genre, "genre", "", "Ranking genre: all/onsen/business/ski/pet")

	return c
}
