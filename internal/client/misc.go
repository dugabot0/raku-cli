package client

import (
	"encoding/json"
	"fmt"
)

// --- Recipe Category List ---

type RecipeParams struct {
	CategoryType string // large/medium/small
}

func (c *Client) RecipeCategoryList(p RecipeParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "categoryType", p.CategoryType)
	return c.get("Recipe/CategoryList/20170426", params)
}

// --- Kobo Ebook Search ---

type KoboSearchParams struct {
	Keyword       string
	Title         string
	Author        string
	PublisherName string
	ItemNumber    string
	KoboGenreID   string // subcategories of 101
	Language      string
	Sort          string // standard/+releaseDate/-releaseDate/+itemPrice/-itemPrice/reviewCount/reviewAverage
	NGKeyword     string
	SalesType     int // 0=regular, 1=pre-order
	Field         int // 0=broad, 1=narrow
	OrFlag        int // 0=AND, 1=OR
	GenreInformationFlag int // 1=include genre counts
	Page          int
	Hits          int
}

func (c *Client) KoboEbookSearch(p KoboSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "title", p.Title)
	setIfNonEmpty(params, "author", p.Author)
	setIfNonEmpty(params, "publisherName", p.PublisherName)
	setIfNonEmpty(params, "itemNumber", p.ItemNumber)
	setIfNonEmpty(params, "koboGenreId", p.KoboGenreID)
	setIfNonEmpty(params, "language", p.Language)
	setIfNonEmpty(params, "sort", p.Sort)
	setIfNonEmpty(params, "NGKeyword", p.NGKeyword)
	setIntFlag(params, "salesType", p.SalesType)
	setIntFlag(params, "field", p.Field)
	setIntFlag(params, "orFlag", p.OrFlag)
	setIntFlag(params, "genreInformationFlag", p.GenreInformationFlag)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.get("Kobo/EbookSearch/20170426", params)
}

// --- GORA Golf Course Search ---

type GoraSearchParams struct {
	Keyword      string
	AreaCode     string
	Latitude     string
	Longitude    string
	SearchRadius int // 10-300 km, default 150
	Sort         string
	Reservation  int // 0=all, 1=reservable only (default 1)
	Carrier      int // 0=PC, 1=mobile
	Page         int
	Hits         int
}

func (c *Client) GoraGolfCourseSearch(p GoraSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "areaCode", p.AreaCode)
	setIfNonEmpty(params, "latitude", p.Latitude)
	setIfNonEmpty(params, "longitude", p.Longitude)
	if p.SearchRadius > 0 {
		params.Set("searchRadius", fmt.Sprint(p.SearchRadius))
	}
	setIfNonEmpty(params, "sort", p.Sort)
	setIntFlag(params, "reservation", p.Reservation)
	setIntFlag(params, "carrier", p.Carrier)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.get("Gora/GoraGolfCourseSearch/20170623", params)
}
