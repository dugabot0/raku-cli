package client

import (
	"encoding/json"
	"fmt"
)

// BooksSearchParams covers all fields across Books API endpoints.
// Not every field applies to every endpoint — unused fields are ignored.
type BooksSearchParams struct {
	// Common
	Keyword  string
	Sort     string
	Page     int
	Hits     int

	// BooksTotal only
	MediaType string // 0=all, 1=book, 2=cd, 3=dvd, 4=foreign, 5=game, 6=magazine
	IsbnJan   string // 13-digit ISBN/JAN (overrides keyword + genre)

	// Book-specific
	Title         string
	Author        string
	PublisherName string
	Size          int    // 0=all, 1=book, 2=paperback, 3=new, 4=set, 5=dict, 6=illus, 7=picture, 8=cassette, 9=comic, 10=magazine
	ISBN          string

	// CD/DVD-specific
	ArtistName string
	Label      string
	JAN        string // barcode

	// Game-specific
	Hardware  string
	MakerCode string

	// Genre filter (all endpoints)
	BooksGenreID string

	// Common flags
	Availability         int // 0=all, 1=in stock, 2-6=various
	OutOfStockFlag       int // 1=include out-of-stock
	ChirayomiFlag        int // 1=preview items only
	LimitedFlag          int // 1=limited editions only
	Field                int // 0=broad, 1=narrow
	Carrier              int // 0=PC, 1=mobile
	OrFlag               int // 0=AND, 1=OR
	NGKeyword            string
	GenreInformationFlag int // 1=include genre counts
}

func (c *Client) bookSearch(endpoint string, p BooksSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "keyword", p.Keyword)
	setIfNonEmpty(params, "title", p.Title)
	setIfNonEmpty(params, "author", p.Author)
	setIfNonEmpty(params, "publisherName", p.PublisherName)
	setIfNonEmpty(params, "artistName", p.ArtistName)
	setIfNonEmpty(params, "label", p.Label)
	setIfNonEmpty(params, "jan", p.JAN)
	setIfNonEmpty(params, "isbn", p.ISBN)
	setIfNonEmpty(params, "isbnjan", p.IsbnJan)
	setIfNonEmpty(params, "hardware", p.Hardware)
	setIfNonEmpty(params, "makerCode", p.MakerCode)
	setIfNonEmpty(params, "booksGenreId", p.BooksGenreID)
	setIfNonEmpty(params, "mediaType", p.MediaType)
	setIfNonEmpty(params, "sort", p.Sort)
	setIfNonEmpty(params, "NGKeyword", p.NGKeyword)
	if p.Size > 0 {
		params.Set("size", fmt.Sprint(p.Size))
	}
	if p.Page > 0 {
		params.Set("page", fmt.Sprint(p.Page))
	}
	if p.Hits > 0 {
		params.Set("hits", fmt.Sprint(p.Hits))
	}
	setIntFlag(params, "availability", p.Availability)
	setIntFlag(params, "outOfStockFlag", p.OutOfStockFlag)
	setIntFlag(params, "chirayomiFlag", p.ChirayomiFlag)
	setIntFlag(params, "limitedFlag", p.LimitedFlag)
	setIntFlag(params, "field", p.Field)
	setIntFlag(params, "carrier", p.Carrier)
	setIntFlag(params, "orFlag", p.OrFlag)
	setIntFlag(params, "genreInformationFlag", p.GenreInformationFlag)
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
	GenreID   string
	GenrePath int // 1=include ancestor genres
}

func (c *Client) BooksGenreSearch(p BooksGenreSearchParams) (json.RawMessage, error) {
	params := c.baseParams()
	setIfNonEmpty(params, "booksGenreId", p.GenreID)
	setIntFlag(params, "genrePath", p.GenrePath)
	return c.get("BooksGenre/Search/20170404", params)
}
