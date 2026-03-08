package ichiba

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newRankingCmd() *cobra.Command {
	var p client.IchibaRankingParams

	c := &cobra.Command{
		Use:   "ranking",
		Short: "Get Ichiba item ranking",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadIchibaClient()
			result, err := cl.IchibaRanking(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.GenreID, "genre-id", "", "Genre ID (cannot combine with --age or --sex)")
	f.StringVar(&p.Age, "age", "", "Age group: 10/20/30/40/50 (cannot combine with --genre-id)")
	f.IntVar(&p.Sex, "sex", 0, "0=male, 1=female (use with --age)")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC, 1=mobile")
	f.IntVar(&p.Page, "page", 1, "Page number (1-34, 30 items per page)")
	f.StringVar(&p.Period, "period", "", "Set to 'realtime' for real-time ranking")

	return c
}
