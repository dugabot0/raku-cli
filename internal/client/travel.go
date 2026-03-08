package client

import (
	"encoding/json"
	"fmt"
)

// --- Travel Hotel Simple Search ---

type TravelHotelSearchParams struct {
	LargeArea    string
	MiddleArea   string
	SmallArea    string
	HotelNo      string
	Latitude     string
	Longitude    string
	SearchRadius string
	Page         int
	Hits         int
}

func (c *Client) TravelSimpleHotelSearch(p TravelHotelSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "largeClassCode", p.LargeArea)
	setIfNonEmpty(params, "middleClassCode", p.MiddleArea)
	setIfNonEmpty(params, "smallClassCode", p.SmallArea)
	setIfNonEmpty(params, "hotelNo", p.HotelNo)
	setIfNonEmpty(params, "latitude", p.Latitude)
	setIfNonEmpty(params, "longitude", p.Longitude)
	setIfNonEmpty(params, "searchRadius", p.SearchRadius)
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
	LargeArea    string
	MiddleArea   string
	SmallArea    string
	HotelNo      string
	CheckinDate  string
	CheckoutDate string
	AdultNum     int
	Page         int
	Hits         int
}

func (c *Client) TravelVacantHotelSearch(p TravelVacantParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "largeClassCode", p.LargeArea)
	setIfNonEmpty(params, "middleClassCode", p.MiddleArea)
	setIfNonEmpty(params, "smallClassCode", p.SmallArea)
	setIfNonEmpty(params, "hotelNo", p.HotelNo)
	setIfNonEmpty(params, "checkinDate", p.CheckinDate)
	setIfNonEmpty(params, "checkoutDate", p.CheckoutDate)
	if p.AdultNum > 0 {
		params.Set("adultNum", fmt.Sprint(p.AdultNum))
	}
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
	return c.get("Travel/GetAreaClass/20131024", params)
}

// --- Travel Hotel Ranking ---

type TravelRankingParams struct {
	Genre string
}

func (c *Client) TravelHotelRanking(p TravelRankingParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "genre", p.Genre)
	return c.get("Travel/HotelRanking/20170426", params)
}
