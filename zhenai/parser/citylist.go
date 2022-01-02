package parser

import (
	"regexp"
	"spider/engine"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	submatch := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range submatch {
		result.Request = append(
			result.Request,
			engine.Request{
				Url:        string(v[1]),
				ParserFunc: ParseCity,
			})
		//result.Item = append(result.Item, v[2])
	}
	return result
}
