package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newVacantCmd() *cobra.Command {
	var p client.TravelVacantParams

	c := &cobra.Command{
		Use:   "vacant",
		Short: "Search vacant hotels",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelVacantHotelSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	// Area
	f.StringVar(&p.LargeArea, "large-area", "", "Large area code")
	f.StringVar(&p.MiddleArea, "middle-area", "", "Middle area code (prefecture)")
	f.StringVar(&p.SmallArea, "small-area", "", "Small area code (city)")
	f.StringVar(&p.DetailArea, "detail-area", "", "Detail area code (station/district)")
	f.StringVar(&p.HotelNo, "hotel-no", "", "Hotel number(s), comma-separated (max 15)")
	f.StringVar(&p.Latitude, "latitude", "", "Latitude")
	f.StringVar(&p.Longitude, "longitude", "", "Longitude")
	f.StringVar(&p.SearchRadius, "search-radius", "", "Search radius in km (0.1-3.0)")

	// Stay
	f.StringVar(&p.CheckinDate, "checkin-date", "", "Check-in date (YYYY-MM-DD)")
	f.StringVar(&p.CheckoutDate, "checkout-date", "", "Check-out date (YYYY-MM-DD)")
	f.IntVar(&p.AdultNum, "adult-num", 2, "Number of adults (1-10)")
	f.IntVar(&p.RoomNum, "room-num", 1, "Number of rooms (1-10)")

	// Children/infants
	f.IntVar(&p.UpClassNum, "upper-child-num", 0, "Upper elementary school children (0-10)")
	f.IntVar(&p.LowClassNum, "lower-child-num", 0, "Lower elementary school children (0-10)")
	f.IntVar(&p.InfantWithMBNum, "infant-mb-num", 0, "Infants with meals and bedding (0-10)")
	f.IntVar(&p.InfantWithMNum, "infant-m-num", 0, "Infants with meals only (0-10)")
	f.IntVar(&p.InfantWithBNum, "infant-b-num", 0, "Infants with bedding only (0-10)")
	f.IntVar(&p.InfantWithoutMBNum, "infant-none-num", 0, "Infants with neither meals nor bedding (0-10)")

	// Price
	f.IntVar(&p.MinCharge, "min-charge", 0, "Minimum price per room per night (yen)")
	f.IntVar(&p.MaxCharge, "max-charge", 0, "Maximum price per room per night (yen)")

	// Filters & display
	f.StringVar(&p.SqueezeCondition, "squeeze", "", "Filters: kinen,internet,daiyoku,onsen,breakfast,dinner (comma-separated)")
	f.IntVar(&p.SearchPattern, "search-pattern", 0, "0=by facility, 1=by room/plan")
	f.StringVar(&p.Sort, "sort", "", "Sort: standard/+roomCharge/-roomCharge")
	f.StringVar(&p.ResponseType, "response-type", "", "Detail level: small/middle/large")
	f.IntVar(&p.HotelThumbnailSize, "thumbnail-size", 0, "Image size: 1=small, 2=medium, 3=large")
	f.IntVar(&p.DatumType, "datum-type", 0, "Coordinate system: 1=WGS84 (degrees), 2=Tokyo Datum (seconds)")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC/smartphone, 1=feature phone")
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")

	return c
}
