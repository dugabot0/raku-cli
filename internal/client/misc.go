package client

import (
	"encoding/json"
	"fmt"
)

// --- Recipe Category List ---

type RecipeParams struct {
	CategoryType string
	CategoryID   string
}

func (c *Client) RecipeCategoryList(p RecipeParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "categoryType", p.CategoryType)
	setIfNonEmpty(params, "categoryId", p.CategoryID)
	return c.get("Recipe/CategoryList/20170426", params)
}

// --- Kobo Ebook Search ---

type KoboSearchParams struct {
	Keyword  string
	Title    string
	Author   string
	Sort     string
	Page     int
	Hits     int
}

func (c *Client) KoboEbookSearch(p KoboSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "title", p.Title)
	setIfNonEmpty(params, "author", p.Author)
	setIfNonEmpty(params, "sort", p.Sort)
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
	Keyword  string
	AreaCode string
	Sort     string
	Page     int
	Hits     int
}

func (c *Client) GoraGolfCourseSearch(p GoraSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "areaCode", p.AreaCode)
	setIfNonEmpty(params, "sort", p.Sort)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.get("Gora/GoraGolfCourseSearch/20170623", params)
}
