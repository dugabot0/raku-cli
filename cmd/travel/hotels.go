package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newHotelsCmd() *cobra.Command {
	var p client.TravelHotelSearchParams

	c := &cobra.Command{
		Use:   "hotels",
		Short: "Search hotels (simple search)",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelSimpleHotelSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.LargeArea, "large-area", "", "Large area code (e.g. japan)")
	f.StringVar(&p.MiddleArea, "middle-area", "", "Middle area code (prefecture)")
	f.StringVar(&p.SmallArea, "small-area", "", "Small area code (city)")
	f.StringVar(&p.DetailArea, "detail-area", "", "Detail area code (station/district)")
	f.StringVar(&p.HotelNo, "hotel-no", "", "Hotel number(s), comma-separated (max 15)")
	f.StringVar(&p.Latitude, "latitude", "", "Latitude (requires --longitude and --search-radius)")
	f.StringVar(&p.Longitude, "longitude", "", "Longitude (requires --latitude and --search-radius)")
	f.StringVar(&p.SearchRadius, "search-radius", "", "Search radius in km (0.1-3.0)")
	f.StringVar(&p.SqueezeCondition, "squeeze", "", "Filters: kinen,internet,daiyoku,onsen (comma-separated)")
	f.StringVar(&p.Sort, "sort", "", "Sort: standard/+roomCharge/-roomCharge")
	f.StringVar(&p.ResponseType, "response-type", "", "Detail level: small/middle/large")
	f.IntVar(&p.HotelThumbnailSize, "thumbnail-size", 0, "Image size: 1=small, 2=medium, 3=large")
	f.IntVar(&p.DatumType, "datum-type", 0, "Coordinate system: 1=WGS84 (degrees), 2=Tokyo Datum (seconds)")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC/smartphone, 1=feature phone")
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")

	return c
}
