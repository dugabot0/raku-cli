package misc

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newGoraCmd() *cobra.Command {
	var p client.GoraSearchParams

	c := &cobra.Command{
		Use:   "gora",
		Short: "Search GORA golf courses",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.GoraGolfCourseSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()
	f.StringVar(&p.Keyword, "keyword", "", "Search keyword")
	f.StringVar(&p.AreaCode, "area-code", "", "Area code (use instead of keyword or coordinates)")
	f.StringVar(&p.Latitude, "latitude", "", "Latitude (takes priority over area-code)")
	f.StringVar(&p.Longitude, "longitude", "", "Longitude (takes priority over area-code)")
	f.IntVar(&p.SearchRadius, "search-radius", 0, "Search radius in km for coordinate search (10-300, default 150)")
	f.StringVar(&p.Sort, "sort", "", "Sort: rating/50on/prefecture/highway/reservation/evaluation/staff/facility/meal/course/costperformance/distance/fairway/friends/entertainment/couple/athlete/beginner/normal/senior/woman")
	f.IntVar(&p.Reservation, "reservation", 0, "0=all courses, 1=reservable on GORA only")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC, 1=mobile")
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")

	return c
}
