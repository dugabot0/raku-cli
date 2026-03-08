package client

import (
	"encoding/json"
	"fmt"
)

// BooksSearchParams is shared across all Books API endpoints.
type BooksSearchParams struct {
	Keyword   string
	Sort      string
	Page      int
	Hits      int
	MediaType string // used by BooksTotal only
}

func (c *Client) bookSearch(endpoint string, p BooksSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "sort", p.Sort)
	setIfNonEmpty(params, "mediaType", p.MediaType)
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	return c.get(endpoint, params)
}

func (c *Client) BooksTotalSearch(p BooksSearchParams) (json.RawMessage, error) {
	return c.bookSearch("BooksTotal/Search/20170404", p)
}

func (c *Client) BooksBookSearch(p BooksSearchParams) (json.RawMessage, error) {
	return c.bookSearch("BooksBook/Search/20170404", p)
}

func (c *Client) BooksCDSearch(p BooksSearchParams) (json.RawMessage, error) {
	return c.bookSearch("BooksCD/Search/20170404", p)
}

func (c *Client) BooksDVDSearch(p BooksSearchParams) (json.RawMessage, error) {
	return c.bookSearch("BooksDVD/Search/20170404", p)
}

func (c *Client) BooksMagazineSearch(p BooksSearchParams) (json.RawMessage, error) {
	return c.bookSearch("BooksMagazine/Search/20170404", p)
}

func (c *Client) BooksGameSearch(p BooksSearchParams) (json.RawMessage, error) {
	return c.bookSearch("BooksGame/Search/20170404", p)
}

// --- Books Genre Search ---

type BooksGenreSearchParams struct {
	GenreID string
}

func (c *Client) BooksGenreSearch(p BooksGenreSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "booksGenreId", p.GenreID)
	return c.get("BooksGenre/Search/20170404", params)
}
