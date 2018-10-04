package parser

import (
	"regexp"
	"spider_demo/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

	nextPageRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "User " +string(m[2]))
		url := string(m[1])

		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url)
			},
		})

	}

	matches = nextPageRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc:ParseCity,
		})
	}

	return result
}
