package persist

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider_demo/crawler/engine"
)

func ItemSaver(index string)  (chan engine.Item, error){
	client, err := elastic.NewClient(
		//Must turn off sniff , because it run in docker
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}


	out := make(chan engine.Item)
	go func() {
		itemConut := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemConut, item)
			itemConut++
			err := save(client, index, item)

			if err != nil {
				log.Printf("Item Saver: error " + "saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func save(client *elastic.Client,index string, item engine.Item) error{

	if item.Type == "" {
		return errors.New("Must Supply Type.")
	}


	 indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	 if item.Id != "" {
	 	indexService.Id(item.Id)
	 }

	_, err := indexService.
		Do(context.Background())

	if err != nil{
		return err
	}


	return nil
}
