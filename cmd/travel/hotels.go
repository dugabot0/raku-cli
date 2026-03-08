package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newHotelsCmd() *cobra.Command {
	var (
		largeArea    string
		middleArea   string
		smallArea    string
		hotelNo      string
		latitude     string
		longitude    string
		searchRadius string
		page         int
		hits         int
	)

	c := &cobra.Command{
		Use:   "hotels",
		Short: "Search hotels (simple search)",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelSimpleHotelSearch(client.TravelHotelSearchParams{
				LargeArea:    largeArea,
				MiddleArea:   middleArea,
				SmallArea:    smallArea,
				HotelNo:      hotelNo,
				Latitude:     latitude,
				Longitude:    longitude,
				SearchRadius: searchRadius,
				Page:         page,
				Hits:         hits,
			})
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&largeArea, "large-area", "", "Large area code (e.g. hokkaido)")
	f.StringVar(&middleArea, "middle-area", "", "Middle area code")
	f.StringVar(&smallArea, "small-area", "", "Small area code")
	f.StringVar(&hotelNo, "hotel-no", "", "Hotel number")
	f.StringVar(&latitude, "latitude", "", "Latitude for location search")
	f.StringVar(&longitude, "longitude", "", "Longitude for location search")
	f.StringVar(&searchRadius, "search-radius", "", "Search radius (km) for location search")
	f.IntVar(&page, "page", 1, "Page number")
	f.IntVar(&hits, "hits", 30, "Results per page (1-30)")

	return c
}
