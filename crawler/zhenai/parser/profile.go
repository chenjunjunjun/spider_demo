package parser

import (
	"regexp"
	"spider_demo/crawler/engine"
	"spider_demo/crawler/model"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)

var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)

var nameRe = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)

var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)

var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)

var incomeRe =regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)

var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)

var occupationRe =regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)

var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)

var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)

var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)

var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)


func ParseProfile(contents []byte, url string) engine.ParseResult{

	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Name = extractString(contents, nameRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type : "zhenai",
				Id: extractString([]byte(url), idRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}else {
		return ""
	}
}
