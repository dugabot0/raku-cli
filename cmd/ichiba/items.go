package ichiba

import (
	"github.com/spf13/cobra"
	"github.com/lumiere/raku-cli/cmd"
	"github.com/lumiere/raku-cli/internal/client"
	"github.com/lumiere/raku-cli/internal/output"
)

func newItemsCmd() *cobra.Command {
	var p client.IchibaItemSearchParams

	c := &cobra.Command{
		Use:   "items",
		Short: "Search Ichiba items",
		RunE: func(c *cobra.Command, args []string) error {
			cl := cmd.LoadRakutenClient()
			result, err := cl.IchibaItemSearch(p)
			if err != nil {
				cmd.HandleError(err)
			}
			output.Print(result, cmd.Pretty())
			return nil
		},
	}

	f := c.Flags()

	// Search targets
	f.StringVar(&p.Keyword, "keyword", "", "Search keyword (max 128 bytes)")
	f.StringVar(&p.ShopCode, "shop-code", "", "Shop code")
	f.StringVar(&p.ItemCode, "item-code", "", "Item code (format: shop:1234)")
	f.StringVar(&p.GenreID, "genre-id", "", "Genre ID")
	f.StringVar(&p.TagID, "tag-id", "", "Tag ID(s), comma-separated (max 10)")

	// Pagination
	f.IntVar(&p.Page, "page", 1, "Page number (1-100)")
	f.IntVar(&p.Hits, "hits", 30, "Results per page (1-30)")

	// Price
	f.IntVar(&p.MinPrice, "min-price", 0, "Minimum price (yen)")
	f.IntVar(&p.MaxPrice, "max-price", 0, "Maximum price (yen)")

	// Sort
	f.StringVar(&p.Sort, "sort", "", "Sort: standard/+itemPrice/-itemPrice/+updateTimestamp/-updateTimestamp/+reviewCount/-reviewCount/+reviewAverage/-reviewAverage/+affiliateRate/-affiliateRate")

	// Search options
	f.IntVar(&p.Availability, "availability", 0, "0=all products, 1=available only")
	f.IntVar(&p.Field, "field", 0, "0=broad search, 1=restricted to item name/caption")
	f.IntVar(&p.Carrier, "carrier", 0, "0=PC, 1=mobile, 2=smartphone")
	f.IntVar(&p.ImageFlag, "image-flag", 0, "1=products with images only")
	f.IntVar(&p.OrFlag, "or-flag", 0, "0=AND search, 1=OR search")
	f.StringVar(&p.NGKeyword, "ng-keyword", "", "Keywords to exclude")

	// Shipping
	f.IntVar(&p.PurchaseType, "purchase-type", 0, "0=normal, 1=periodic purchase, 2=distribution group")
	f.IntVar(&p.ShipOverseasFlag, "ship-overseas-flag", 0, "1=overseas shippable only")
	f.StringVar(&p.ShipOverseasArea, "ship-overseas-area", "", "Overseas delivery area code")
	f.IntVar(&p.AsurakuFlag, "asuraku-flag", 0, "1=next-day delivery (あす楽) only")
	f.IntVar(&p.AsurakuArea, "asuraku-area", 0, "Next-day delivery area code")

	// Promotions & features
	f.IntVar(&p.PointRateFlag, "point-rate-flag", 0, "1=point multiplier items only")
	f.IntVar(&p.PointRate, "point-rate", 0, "Minimum point multiplier (2-10)")
	f.IntVar(&p.PostageFlag, "postage-flag", 0, "1=free shipping only")
	f.IntVar(&p.CreditCardFlag, "credit-card-flag", 0, "1=credit card accepted only")
	f.IntVar(&p.GiftFlag, "gift-flag", 0, "1=gift wrapping available only")
	f.IntVar(&p.HasReviewFlag, "has-review-flag", 0, "1=has reviews only")
	f.IntVar(&p.HasMovieFlag, "has-movie-flag", 0, "1=has video only")
	f.IntVar(&p.PamphletFlag, "pamphlet-flag", 0, "1=has pamphlet only")
	f.IntVar(&p.AppointDeliveryDateFlag, "appoint-delivery-date-flag", 0, "1=delivery date specifiable only")

	// Affiliate rate
	f.Float64Var(&p.MinAffiliateRate, "min-affiliate-rate", 0, "Minimum affiliate rate (1.0-99.9)")
	f.Float64Var(&p.MaxAffiliateRate, "max-affiliate-rate", 0, "Maximum affiliate rate (1.0-99.9)")

	// Response extras
	f.IntVar(&p.GenreInformationFlag, "genre-information-flag", 0, "1=include genre item counts in response")
	f.IntVar(&p.TagInformationFlag, "tag-information-flag", 0, "1=include tag item counts in response")

	return c
}
