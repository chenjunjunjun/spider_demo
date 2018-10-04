package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"spider_demo/crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {
	profile := model.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Gender:     "女",
		Name:       "安静的雪",
		Xingzuo:    "处女座",
		Occupation: "人事",
		Marriage:   "离异",
		House:      "已购房",
		Hokou:      "广西梧州",
		Education:  "大学本科",
		Car:        "未购车",
	}

	id, err := save(profile)

	if err != nil{
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	result, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)
	//json.Unmarshal(result.Source)

}
