package client

import (
	"encoding/json"
	"fmt"
)

const (
	ichibaItemSearchEndpoint  = "IchibaItem/Search/20220601"
	ichibaGenreSearchEndpoint = "IchibaGenre/Search/20220601"
	ichibaRankingEndpoint     = "IchibaItem/Ranking/20220601"
)

// --- Ichiba Item Search ---

type IchibaItemSearchParams struct {
	// Search targets (at least one required)
	Keyword  string
	ShopCode string
	ItemCode string
	GenreID  string
	TagID    string // comma-separated, max 10

	// Pagination
	Page int
	Hits int

	// Price range
	MinPrice int
	MaxPrice int

	// Sort
	Sort string

	// Filtering
	Availability int // 0=all, 1=available only (default 1)
	Field        int // 0=broad, 1=restricted (default 1)
	Carrier      int // 0=PC, 1=mobile, 2=smartphone
	ImageFlag    int // 1=with images only
	OrFlag       int // 0=AND, 1=OR
	NGKeyword    string

	// Shipping
	PurchaseType     int    // 0=normal, 1=periodic, 2=group
	ShipOverseasFlag int    // 1=overseas shippable only
	ShipOverseasArea string // area code
	AsurakuFlag      int    // 1=next-day delivery
	AsurakuArea      int    // area code

	// Promotions & features
	PointRateFlag           int // 1=point multiplier items only
	PointRate               int // minimum multiplier (2-10)
	PostageFlag             int // 1=free shipping only
	CreditCardFlag          int // 1=credit card accepted only
	GiftFlag                int // 1=gift wrapping only
	HasReviewFlag           int // 1=has reviews only
	HasMovieFlag            int // 1=has video only
	PamphletFlag            int // 1=has pamphlet only
	AppointDeliveryDateFlag int // 1=delivery date specifiable only

	// Affiliate rate range
	MinAffiliateRate float64
	MaxAffiliateRate float64

	// Extra info in response
	GenreInformationFlag int // 1=include genre counts
	TagInformationFlag   int // 1=include tag counts
}

func (c *Client) IchibaItemSearch(p IchibaItemSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "shopCode", p.ShopCode)
	setIfNonEmpty(params, "itemCode", p.ItemCode)
	setIfNonEmpty(params, "genreId", p.GenreID)
	setIfNonEmpty(params, "tagId", p.TagID)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	if p.MinPrice > 0 {
		params.Set("minPrice", fmt.Sprint(p.MinPrice))
	}
	if p.MaxPrice > 0 {
		params.Set("maxPrice", fmt.Sprint(p.MaxPrice))
	}
	setIfNonEmpty(params, "sort", p.Sort)
	setIntFlag(params, "availability", p.Availability)
	setIntFlag(params, "field", p.Field)
	setIntFlag(params, "carrier", p.Carrier)
	setIntFlag(params, "imageFlag", p.ImageFlag)
	setIntFlag(params, "orFlag", p.OrFlag)
	setIfNonEmpty(params, "NGKeyword", p.NGKeyword)
	setIntFlag(params, "purchaseType", p.PurchaseType)
	setIntFlag(params, "shipOverseasFlag", p.ShipOverseasFlag)
	setIfNonEmpty(params, "shipOverseasArea", p.ShipOverseasArea)
	setIntFlag(params, "asurakuFlag", p.AsurakuFlag)
	setIntFlag(params, "asurakuArea", p.AsurakuArea)
	setIntFlag(params, "pointRateFlag", p.PointRateFlag)
	if p.PointRate >= 2 {
		params.Set("pointRate", fmt.Sprint(p.PointRate))
	}
	setIntFlag(params, "postageFlag", p.PostageFlag)
	setIntFlag(params, "creditCardFlag", p.CreditCardFlag)
	setIntFlag(params, "giftFlag", p.GiftFlag)
	setIntFlag(params, "hasReviewFlag", p.HasReviewFlag)
	setIntFlag(params, "hasMovieFlag", p.HasMovieFlag)
	setIntFlag(params, "pamphletFlag", p.PamphletFlag)
	setIntFlag(params, "appointDeliveryDateFlag", p.AppointDeliveryDateFlag)
	if p.MinAffiliateRate > 0 {
		params.Set("minAffiliateRate", fmt.Sprintf("%.1f", p.MinAffiliateRate))
	}
	if p.MaxAffiliateRate > 0 {
		params.Set("maxAffiliateRate", fmt.Sprintf("%.1f", p.MaxAffiliateRate))
	}
	setIntFlag(params, "genreInformationFlag", p.GenreInformationFlag)
	setIntFlag(params, "tagInformationFlag", p.TagInformationFlag)
	return c.ichibaGet(ichibaItemSearchEndpoint, params)
}

// setIntFlag sets the param only when val > 0.
func setIntFlag(p interface{ Set(string, string) }, key string, val int) {
	if val > 0 {
		p.Set(key, fmt.Sprint(val))
	}
}

// --- Ichiba Genre Search ---

type IchibaGenreSearchParams struct {
	GenreID string
}

func (c *Client) IchibaGenreSearch(p IchibaGenreSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "genreId", p.GenreID)
	return c.ichibaGet(ichibaGenreSearchEndpoint, params)
}

// --- Ichiba Ranking ---

type IchibaRankingParams struct {
	GenreID string
	Age     string // 10/20/30/40/50 (cannot combine with genreId)
	Sex     int    // 0=male, 1=female
	Carrier int    // 0=PC, 1=mobile
	Page    int    // 1-34
	Period  string // "realtime" for real-time ranking
}

func (c *Client) IchibaRanking(p IchibaRankingParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "genreId", p.GenreID)
	setIfNonEmpty(params, "age", p.Age)
	setIntFlag(params, "sex", p.Sex)
	setIntFlag(params, "carrier", p.Carrier)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	setIfNonEmpty(params, "period", p.Period)
	return c.ichibaGet(ichibaRankingEndpoint, params)
}
