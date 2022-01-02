package engine

import "spider/fetcher"

func work(r Request) (ParseResult, error) {
	bytes, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParserFunc(bytes), nil
}
