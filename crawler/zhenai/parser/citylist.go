package parser

import (
	"regexp"
	"spider_demo/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`


// return city and city's URL exp: 武汉, http://...
func ParseCityList(contents []byte)  engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	// limit 2 city data
	//limit := 2
	for _, m := range matches {
		//result.Items = append(result.Items,"City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
