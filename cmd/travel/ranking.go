package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newRankingCmd() *cobra.Command {
	var p client.TravelRankingParams

	c := &cobra.Command{
		Use:   "ranking",
		Short: "Get hotel ranking",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelHotelRanking(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.Genre, "genre", "all", "Ranking genre: all/onsen/premium (comma-separated for multiple)")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC/smartphone, 1=feature phone")

	return c
}
