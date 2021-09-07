package nhentai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Sort string

const (
	Recent          = ""
	PopularAllTime  = "popular"
	PopularThisWeek = "popular-week"
	PopularToday    = "popular-today"
)

// QueryOptions is the option of a query
type QueryOptions struct {
	Page string
	Sort Sort
}

// QueryResult is the result of a Query
type QueryResult struct {
	Result   []*Doujinshi `json:"result"`
	NumPages int          `json:"num_pages"`
	PerPage  int          `json:"per_page"`
}

// RecentDoujinshi is a function that returns the recent doujinshi
func RecentDoujinshi(opts QueryOptions) (qr QueryResult, err error) {
	opts.Sort = ""
	return search("\"\"", opts)
}

// Search is a function that returns the doujinshi searched
func Search(query string, opt QueryOptions) (qr QueryResult, err error) {

	if validateQuerySort(opt.Sort) {
		return QueryResult{}, errors.New("Sort query option not valid")
	}

	// Check if page is setted to all
	if opt.Page == "" || opt.Page == "all" || opt.Page == "All" || opt.Page == "ALL" {
		fmt.Println("ALL")
		return searchAll(query, opt)
	}

	// Check if page option is valid
	if ok, _ := regexp.MatchString(`^[1-9]+$`, opt.Page); !ok {
		return QueryResult{}, errors.New("Page not valid")
	}

	return search(query, opt)
}

// SearchTag is a function that returns the doujinshi related to searched tags
func SearchTag(tagId int, opt QueryOptions) (qr QueryResult, err error) {

	if validateQuerySort(opt.Sort) {
		return QueryResult{}, errors.New("Sort query option not valid")
	}

	// Set template parameters
	tmpVar := struct {
		Option QueryOptions
		TagId  int
	}{
		opt,
		tagId,
	}

	// Resolver template
	queryUrl, err := templateSolver(searchTagIdApi, tmpVar)
	if err != nil {
		return QueryResult{}, err
	}
	queryUrl = baseUrlApi + queryUrl
	if opt.Sort == "" {
		queryUrl = strings.Replace(queryUrl, "&sort=", "", -1)
	}

	// Do request
	resp, err := ClientHttp.Get(queryUrl)
	if err != nil {
		return QueryResult{}, err
	}
	defer resp.Body.Close()

	// Read data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return QueryResult{}, err
	}

	// Parse json response
	err = json.Unmarshal(data, &qr)
	if err != nil {
		return QueryResult{}, err
	}

	return qr, nil
}

// search is utility function
func search(query string, opt QueryOptions) (qr QueryResult, err error) {
	// Set template parameters
	tmpVar := struct {
		Option QueryOptions
		Search string
	}{
		opt,
		query,
	}

	// Resolver template
	queryUrl, err := templateSolver(searchApi, tmpVar)
	if err != nil {
		return QueryResult{}, err
	}
	queryUrl = baseUrlApi + queryUrl

	// Do request
	resp, err := ClientHttp.Get(queryUrl)
	if err != nil {
		return QueryResult{}, err
	}
	defer resp.Body.Close()

	// Read data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return QueryResult{}, err
	}

	// Parse json response
	err = json.Unmarshal(data, &qr)
	if err != nil {
		return QueryResult{}, err
	}

	return qr, nil
}

// searchAll all is utility function
func searchAll(query string, opt QueryOptions) (qr QueryResult, err error) {
	opt.Page = "1"

	// Obtain total number pages and first page results
	queryPag1, err := search(query, opt)
	if err != nil {
		return QueryResult{}, err
	}

	// for each page
	for i := 2; i < queryPag1.NumPages; i++ {
		opt.Page = strconv.Itoa(i)

		// Get query result for single page
		queryResult, err := search(query, opt)
		if err != nil {
			return QueryResult{}, err
		}

		fmt.Println(queryResult.Result)

		// Append results of single page
		queryPag1.Result = append(queryPag1.Result, queryResult.Result...)
	}

	return queryPag1, nil
}
