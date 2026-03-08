package travel

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newVacantCmd() *cobra.Command {
	var (
		largeArea    string
		middleArea   string
		smallArea    string
		hotelNo      string
		checkinDate  string
		checkoutDate string
		adultNum     int
		page         int
		hits         int
	)

	c := &cobra.Command{
		Use:   "vacant",
		Short: "Search vacant hotels",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.TravelVacantHotelSearch(client.TravelVacantParams{
				LargeArea:    largeArea,
				MiddleArea:   middleArea,
				SmallArea:    smallArea,
				HotelNo:      hotelNo,
				CheckinDate:  checkinDate,
				CheckoutDate: checkoutDate,
				AdultNum:     adultNum,
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
	f.StringVar(&largeArea, "large-area", "", "Large area code")
	f.StringVar(&middleArea, "middle-area", "", "Middle area code")
	f.StringVar(&smallArea, "small-area", "", "Small area code")
	f.StringVar(&hotelNo, "hotel-no", "", "Hotel number")
	f.StringVar(&checkinDate, "checkin-date", "", "Check-in date (YYYY-MM-DD)")
	f.StringVar(&checkoutDate, "checkout-date", "", "Check-out date (YYYY-MM-DD)")
	f.IntVar(&adultNum, "adult-num", 2, "Number of adults")
	f.IntVar(&page, "page", 1, "Page number")
	f.IntVar(&hits, "hits", 30, "Results per page (1-30)")

	return c
}
