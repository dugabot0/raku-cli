package client

import (
	"encoding/json"
	"fmt"
)

// Ichiba endpoint versions
const (
	ichibaItemSearchStd    = "IchibaItem/Search/20170706"
	ichibaItemSearchNew    = "IchibaItem/Search/20220601"
	ichibaGenreSearchStd   = "IchibaGenre/Search/20170711"
	ichibaGenreSearchNew   = "IchibaGenre/Search/20220601"
	ichibaRankingStd       = "IchibaItem/Ranking/20170628"
	ichibaRankingNew       = "IchibaItem/Ranking/20220601"
)

// --- Ichiba Item Search ---

type IchibaItemSearchParams struct {
	Keyword  string
	GenreID  string
	MinPrice int
	MaxPrice int
	Sort     string
	Page     int
	Hits     int
}

func (c *Client) IchibaItemSearch(p IchibaItemSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "genreId", p.GenreID)
	if p.MinPrice > 0 {
		params.Set("minPrice", fmt.Sprint(p.MinPrice))
	}
	if p.MaxPrice > 0 {
		params.Set("maxPrice", fmt.Sprint(p.MaxPrice))
	}
	setIfNonEmpty(params, "sort", p.Sort)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.ichibaGet(ichibaItemSearchStd, ichibaItemSearchNew, params)
}

// --- Ichiba Genre Search ---

type IchibaGenreSearchParams struct {
	GenreID string
}

func (c *Client) IchibaGenreSearch(p IchibaGenreSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "genreId", p.GenreID)
	return c.ichibaGet(ichibaGenreSearchStd, ichibaGenreSearchNew, params)
}

// --- Ichiba Ranking ---

type IchibaRankingParams struct {
	GenreID string
	Age     string
	Sex     string
	Carrier string
}

func (c *Client) IchibaRanking(p IchibaRankingParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "genreId", p.GenreID)
	setIfNonEmpty(params, "age", p.Age)
	setIfNonEmpty(params, "sex", p.Sex)
	setIfNonEmpty(params, "carrier", p.Carrier)
	return c.ichibaGet(ichibaRankingStd, ichibaRankingNew, params)
}
