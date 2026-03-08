package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newHotelCmd() *cobra.Command {
	var hotelNo string

	c := &cobra.Command{
		Use:   "hotel",
		Short: "Get hotel detail",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelHotelDetail(client.TravelHotelDetailParams{
				HotelNo: hotelNo,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&hotelNo, "hotel-no", "", "Hotel number")
	_ = c.MarkFlagRequired("hotel-no")

	return c
}
