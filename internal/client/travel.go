package client

import (
	"encoding/json"
	"fmt"
)

// --- Travel Hotel Simple Search ---

type TravelHotelSearchParams struct {
	// Area (one of area codes, hotelNo, or lat/lon required)
	LargeArea       string
	MiddleArea      string
	SmallArea       string
	DetailArea      string
	HotelNo         string // up to 15, comma-separated
	Latitude        string
	Longitude       string
	SearchRadius    string // 0.1-3.0 km
	SqueezeCondition string // kinen,internet,daiyoku,onsen (comma-separated)

	// Display
	Sort              string // standard/+roomCharge/-roomCharge
	ResponseType      string // small/middle/large
	HotelThumbnailSize int   // 1=small, 2=medium, 3=large
	DatumType         int   // 1=WGS84, 2=Tokyo Datum (default 2)
	Carrier           int   // 0=PC, 1=mobile

	// Pagination
	Page int
	Hits int
}

func (c *Client) TravelSimpleHotelSearch(p TravelHotelSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "largeClassCode", p.LargeArea)
	setIfNonEmpty(params, "middleClassCode", p.MiddleArea)
	setIfNonEmpty(params, "smallClassCode", p.SmallArea)
	setIfNonEmpty(params, "detailClassCode", p.DetailArea)
	setIfNonEmpty(params, "hotelNo", p.HotelNo)
	setIfNonEmpty(params, "latitude", p.Latitude)
	setIfNonEmpty(params, "longitude", p.Longitude)
	setIfNonEmpty(params, "searchRadius", p.SearchRadius)
	setIfNonEmpty(params, "squeezeCondition", p.SqueezeCondition)
	setIfNonEmpty(params, "sort", p.Sort)
	setIfNonEmpty(params, "responseType", p.ResponseType)
	if p.HotelThumbnailSize > 0 {
		params.Set("hotelThumbnailSize", fmt.Sprint(p.HotelThumbnailSize))
	}
	setIntFlag(params, "datumType", p.DatumType)
	setIntFlag(params, "carrier", p.Carrier)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.get("Travel/SimpleHotelSearch/20170426", params)
}

// --- Travel Hotel Detail Search ---

type TravelHotelDetailParams struct {
	HotelNo string
}

func (c *Client) TravelHotelDetail(p TravelHotelDetailParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "hotelNo", p.HotelNo)
	return c.get("Travel/HotelDetailSearch/20170426", params)
}

// --- Travel Vacant Hotel Search ---

type TravelVacantParams struct {
	// Area
	LargeArea   string
	MiddleArea  string
	SmallArea   string
	DetailArea  string
	HotelNo     string
	Latitude    string
	Longitude   string
	SearchRadius string

	// Stay details
	CheckinDate  string
	CheckoutDate string
	AdultNum     int
	RoomNum      int // 1-10, default 1

	// Children/infants
	UpClassNum         int // upper elementary
	LowClassNum        int // lower elementary
	InfantWithMBNum    int // infants with meals+bedding
	InfantWithMNum     int // infants with meals only
	InfantWithBNum     int // infants with bedding only
	InfantWithoutMBNum int // infants with neither

	// Price range
	MinCharge int
	MaxCharge int

	// Filters
	SqueezeCondition string // kinen,internet,daiyoku,onsen,breakfast,dinner
	SearchPattern    int   // 0=by facility, 1=by room/plan

	// Display
	Sort               string // standard/+roomCharge/-roomCharge
	ResponseType       string // small/middle/large
	HotelThumbnailSize int   // 1-3
	DatumType          int   // 1=WGS84, 2=Tokyo Datum
	Carrier            int

	// Pagination
	Page int
	Hits int
}

func (c *Client) TravelVacantHotelSearch(p TravelVacantParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "largeClassCode", p.LargeArea)
	setIfNonEmpty(params, "middleClassCode", p.MiddleArea)
	setIfNonEmpty(params, "smallClassCode", p.SmallArea)
	setIfNonEmpty(params, "detailClassCode", p.DetailArea)
	setIfNonEmpty(params, "hotelNo", p.HotelNo)
	setIfNonEmpty(params, "latitude", p.Latitude)
	setIfNonEmpty(params, "longitude", p.Longitude)
	setIfNonEmpty(params, "searchRadius", p.SearchRadius)
	setIfNonEmpty(params, "checkinDate", p.CheckinDate)
	setIfNonEmpty(params, "checkoutDate", p.CheckoutDate)
	if p.AdultNum > 0 {
		params.Set("adultNum", fmt.Sprint(p.AdultNum))
	}
	if p.RoomNum > 0 {
		params.Set("roomNum", fmt.Sprint(p.RoomNum))
	}
	setIntFlag(params, "upClassNum", p.UpClassNum)
	setIntFlag(params, "lowClassNum", p.LowClassNum)
	setIntFlag(params, "infantWithMBNum", p.InfantWithMBNum)
	setIntFlag(params, "infantWithMNum", p.InfantWithMNum)
	setIntFlag(params, "infantWithBNum", p.InfantWithBNum)
	setIntFlag(params, "infantWithoutMBNum", p.InfantWithoutMBNum)
	if p.MinCharge > 0 {
		params.Set("minCharge", fmt.Sprint(p.MinCharge))
	}
	if p.MaxCharge > 0 {
		params.Set("maxCharge", fmt.Sprint(p.MaxCharge))
	}
	setIfNonEmpty(params, "squeezeCondition", p.SqueezeCondition)
	setIntFlag(params, "searchPattern", p.SearchPattern)
	setIfNonEmpty(params, "sort", p.Sort)
	setIfNonEmpty(params, "responseType", p.ResponseType)
	if p.HotelThumbnailSize > 0 {
		params.Set("hotelThumbnailSize", fmt.Sprint(p.HotelThumbnailSize))
	}
	setIntFlag(params, "datumType", p.DatumType)
	setIntFlag(params, "carrier", p.Carrier)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.get("Travel/VacantHotelSearch/20170426", params)
}

// --- Travel Area Class ---

func (c *Client) TravelGetAreaClass() (json.RawMessage, error) {
	params := c.baseParams()
	return c.get("Travel/GetAreaClass/20140210", params)
}

// --- Travel Hotel Ranking ---

type TravelRankingParams struct {
	Genre   string // all/onsen/premium (CSV supported)
	Carrier int    // 0=PC, 1=mobile
}

func (c *Client) TravelHotelRanking(p TravelRankingParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "genre", p.Genre)
	setIntFlag(params, "carrier", p.Carrier)
	return c.get("Travel/HotelRanking/20170426", params)
}
